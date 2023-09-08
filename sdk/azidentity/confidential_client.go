//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

type confidentialClientOptions struct {
	azcore.ClientOptions

	AdditionallyAllowedTenants []string
	// Assertion for on-behalf-of authentication
	Assertion                         string
	DisableInstanceDiscovery, SendX5C bool
}

// confidentialClient wraps the MSAL confidential client
type confidentialClient struct {
	cae, noCAE               msalConfidentialClient
	caeMu, noCAEMu, clientMu *sync.Mutex
	clientID, tenantID       string
	cred                     confidential.Credential
	host                     string
	name                     string
	opts                     confidentialClientOptions
	region                   string
}

func newConfidentialClient(tenantID, clientID, name string, cred confidential.Credential, opts confidentialClientOptions) (*confidentialClient, error) {
	if !validTenantID(tenantID) {
		return nil, errInvalidTenantID
	}
	host, err := setAuthorityHost(opts.Cloud)
	if err != nil {
		return nil, err
	}
	region := os.Getenv(azureRegionalAuthorityName)
	opts.AdditionallyAllowedTenants = resolveAdditionalTenants(opts.AdditionallyAllowedTenants)
	return &confidentialClient{
		caeMu:    &sync.Mutex{},
		clientID: clientID,
		clientMu: &sync.Mutex{},
		cred:     cred,
		host:     host,
		name:     name,
		noCAEMu:  &sync.Mutex{},
		opts:     opts,
		region:   region,
		tenantID: tenantID,
	}, nil
}

// GetToken requests an access token from MSAL, checking the cache first.
func (c *confidentialClient) GetToken(ctx context.Context, tro policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if len(tro.Scopes) < 1 {
		return azcore.AccessToken{}, fmt.Errorf("%s.GetToken() requires at least one scope", c.name)
	}
	// we don't resolve the tenant for managed identities because they acquire tokens only from their home tenants
	if c.name != credNameManagedIdentity {
		tenant, err := c.resolveTenant(tro.TenantID)
		if err != nil {
			return azcore.AccessToken{}, err
		}
		tro.TenantID = tenant
	}
	client, mu, err := c.client(ctx, tro)
	if err != nil {
		return azcore.AccessToken{}, err
	}
	mu.Lock()
	defer mu.Unlock()
	var ar confidential.AuthResult
	if c.opts.Assertion != "" {
		ar, err = client.AcquireTokenOnBehalfOf(ctx, c.opts.Assertion, tro.Scopes, confidential.WithClaims(tro.Claims), confidential.WithTenantID(tro.TenantID))
	} else {
		ar, err = client.AcquireTokenSilent(ctx, tro.Scopes, confidential.WithClaims(tro.Claims), confidential.WithTenantID(tro.TenantID))
		if err != nil {
			ar, err = client.AcquireTokenByCredential(ctx, tro.Scopes, confidential.WithClaims(tro.Claims), confidential.WithTenantID(tro.TenantID))
		}
	}
	if err != nil {
		// We could get a credentialUnavailableError from managed identity authentication because in that case the error comes from our code.
		// We return it directly because it affects the behavior of credential chains. Otherwise, we return AuthenticationFailedError.
		var unavailableErr *credentialUnavailableError
		if !errors.As(err, &unavailableErr) {
			res := getResponseFromError(err)
			err = newAuthenticationFailedError(c.name, err.Error(), res, err)
		}
	} else {
		msg := fmt.Sprintf("%s.GetToken() acquired a token for scope %q", c.name, strings.Join(ar.GrantedScopes, ", "))
		log.Write(EventAuthentication, msg)
	}
	return azcore.AccessToken{Token: ar.AccessToken, ExpiresOn: ar.ExpiresOn.UTC()}, err
}

func (c *confidentialClient) client(ctx context.Context, tro policy.TokenRequestOptions) (msalConfidentialClient, *sync.Mutex, error) {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if tro.EnableCAE {
		if c.cae == nil {
			client, err := c.newMSALClient(true)
			if err != nil {
				return nil, nil, err
			}
			c.cae = client
		}
		return c.cae, c.caeMu, nil
	}
	if c.noCAE == nil {
		client, err := c.newMSALClient(false)
		if err != nil {
			return nil, nil, err
		}
		c.noCAE = client
	}
	return c.noCAE, c.noCAEMu, nil
}

func (c *confidentialClient) newMSALClient(enableCAE bool) (msalConfidentialClient, error) {
	authority := runtime.JoinPaths(c.host, c.tenantID)
	o := []confidential.Option{
		confidential.WithAzureRegion(c.region),
		confidential.WithHTTPClient(newPipelineAdapter(&c.opts.ClientOptions)),
	}
	if enableCAE {
		o = append(o, confidential.WithClientCapabilities(cp1))
	}
	if c.opts.SendX5C {
		o = append(o, confidential.WithX5C())
	}
	if c.opts.DisableInstanceDiscovery || strings.ToLower(c.tenantID) == "adfs" {
		o = append(o, confidential.WithInstanceDiscovery(false))
	}
	return confidential.New(authority, c.clientID, c.cred, o...)
}

// resolveTenant returns the correct tenant for a token request given the client's
// configuration, or an error when that configuration doesn't allow the specified tenant
func (c *confidentialClient) resolveTenant(specified string) (string, error) {
	return resolveTenant(c.tenantID, specified, c.name, c.opts.AdditionallyAllowedTenants)
}
