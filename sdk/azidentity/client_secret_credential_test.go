//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
)

const secret = "secret"

func TestClientSecretCredential_InvalidTenantID(t *testing.T) {
	cred, err := NewClientSecretCredential(badTenantID, fakeClientID, secret, nil)
	if err == nil {
		t.Fatal("Expected an error but received none")
	}
	if cred != nil {
		t.Fatalf("Expected a nil credential value. Received: %v", cred)
	}
}

func TestClientSecretCredential_GetTokenSuccess(t *testing.T) {
	cred, err := NewClientSecretCredential(fakeTenantID, fakeClientID, secret, nil)
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	cred.client = fakeConfidentialClient{}
	_, err = cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{liveTestScope}})
	if err != nil {
		t.Fatalf("Expected an empty error but received: %v", err)
	}
}

func TestClientSecretCredential_Live(t *testing.T) {
	opts, stop := initRecording(t)
	defer stop()
	o := ClientSecretCredentialOptions{ClientOptions: opts}
	cred, err := NewClientSecretCredential(liveSP.tenantID, liveSP.clientID, liveSP.secret, &o)
	if err != nil {
		t.Fatalf("failed to construct credential: %v", err)
	}
	testGetTokenSuccess(t, cred)
}

func TestClientSecretCredentialADFS_Live(t *testing.T) {
	if recording.GetRecordMode() == recording.LiveMode {
		if adfsLiveSP.certPass == "" || adfsLiveSP.clientID == "" || adfsLiveSP.pfxPath == "" || adfsLiveSP.scope == "" || adfsLiveSP.tenantID == "" {
			t.Skip("this test requires manual recording and access to ADFS instance, and can't pass live in CI")
		}
	}
	opts, stop := initRecording(t)
	defer stop()
	disableID, err := strconv.ParseBool(disableInstanceDiscovery)
	if err != nil {
		disableID = false
	}
	o := ClientSecretCredentialOptions{ClientOptions: opts, DisableInstanceDiscovery: disableID}
	cred, err := NewClientSecretCredential(adfsLiveSP.tenantID, adfsLiveSP.clientID, adfsLiveSP.secret, &o)
	if err != nil {
		t.Fatalf("failed to construct credential: %v", err)
	}
	var scope []string
	scope = append(scope, adfsLiveSP.scope)
	testGetTokenSuccess(t, cred, scope...)
}

func TestClientSecretCredential_InvalidSecretLive(t *testing.T) {
	opts, stop := initRecording(t)
	defer stop()
	o := ClientSecretCredentialOptions{ClientOptions: opts}
	cred, err := NewClientSecretCredential(liveSP.tenantID, liveSP.clientID, "invalid secret", &o)
	if err != nil {
		t.Fatalf("failed to construct credential: %v", err)
	}
	tk, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{liveTestScope}})
	if !reflect.ValueOf(tk).IsZero() {
		t.Fatal("expected a zero value AccessToken")
	}
	if e, ok := err.(*AuthenticationFailedError); ok {
		if e.RawResponse == nil {
			t.Fatal("expected a non-nil RawResponse")
		}
	} else {
		t.Fatalf("expected AuthenticationFailedError, received %T", err)
	}
	if !strings.HasPrefix(err.Error(), credNameSecret) {
		t.Fatal("missing credential type prefix")
	}
}
