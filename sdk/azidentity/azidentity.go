// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const (
	defaultAuthorityHost = "https://login.microsoftonline.com/"
)

var (
	successStatusCodes = [2]int{
		http.StatusOK,      // 200
		http.StatusCreated, // 201
	}
)

var (
	defaultAuthorityHostURL    *url.URL
	defaultTokenCredentialOpts *TokenCredentialOptions
)

func init() {
	// The error check is handled in azidentity_test.go
	defaultAuthorityHostURL, _ = url.Parse(defaultAuthorityHost)
	defaultTokenCredentialOpts = &TokenCredentialOptions{AuthorityHost: defaultAuthorityHostURL}
}

type internalAccessToken struct {
	Token        string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    json.Number `json:"expires_in"`
	ExpiresOn    string      `json:"expires_on"`
	NotBefore    string      `json:"not_before"`
	Resource     string      `json:"resource"`
	TokenType    string      `json:"token_type"`
}

// AuthenticationResponseError is a struct used to marshal responses when authentication has failed
type AuthenticationResponseError struct {
	Message       string `json:"error"`
	Description   string `json:"error_description"`
	Timestamp     string `json:"timestamp"`
	TraceID       string `json:"trace_id"`
	CorrelationID string `json:"correlation_id"`
	URI           string `json:"error_uri"`
	Response      *azcore.Response
}

func (e *AuthenticationResponseError) Error() string {
	msg := e.Message
	if len(e.Description) > 0 {
		msg += " " + e.Description
	}
	return msg
}

// AuthenticationFailedError is a struct used to marshal responses when authentication has failed
type AuthenticationFailedError struct {
	AuthError error
}

// IsNotRetriable allows retry policy to stop execution in case it receives a AuthenticationFailedError
func (e *AuthenticationFailedError) IsNotRetriable() bool {
	return true
}

func (e *AuthenticationFailedError) Error() string {
	return e.AuthError.Error()
}

func As(err error, target interface{}) bool {
	var authFailedErr *AuthenticationFailedError
	if reflect.TypeOf(err) == reflect.TypeOf(authFailedErr) {
		err = err.(*AuthenticationFailedError).AuthError
	}

	if target == nil {
		panic("errors: target cannot be nil")
	}
	typ := reflect.TypeOf(target)
	if typ.Kind() != reflect.Ptr {
		panic("errors: target must be a pointer")
	}
	targetType := typ.Elem()
	for {
		if reflect.TypeOf(err) == targetType {
			reflect.ValueOf(target).Elem().Set(reflect.ValueOf(err))
			return true
		}
		if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
			return true
		}
		if err = errors.Unwrap(err); err == nil {
			return false
		}
	}
}

func newAuthenticationResponseError(resp *azcore.Response) error {
	authFailed := &AuthenticationResponseError{}
	err := json.Unmarshal(resp.Payload, authFailed)
	if err != nil {
		authFailed.Message = resp.Status
		authFailed.Description = "Failed to unmarshal response: " + err.Error()
	}
	authFailed.Response = resp
	return authFailed
}

// CredentialUnavailableError an error type returned when the conditions required to create a credential do not exist
type CredentialUnavailableError struct {
	CredentialType string
	Message        string
}

func (e *CredentialUnavailableError) Error() string {
	return e.CredentialType + ": " + e.Message
}

// ChainedCredentialError an error specific to ChainedTokenCredential and DefaultTokenCredential
// this error type will return a list of Credential Unavailable errors
type ChainedCredentialError struct {
	ErrorList []*CredentialUnavailableError
}

// IsNotRetriable allows retry policy to stop execution in case it receives a CredentialUnavailableError
func (e *CredentialUnavailableError) IsNotRetriable() bool {
	return true
}

func (e *ChainedCredentialError) Error() string {
	if len(e.ErrorList) > 0 {
		msg := ""
		for _, err := range e.ErrorList {
			msg += err.Error() + "\n"
		}
		return msg
	}
	return "Chained Token Credential: An unexpected error has occurred"
}

// TokenCredentialOptions to configure requests made to Azure Identity Services
type TokenCredentialOptions struct {
	// The host of the Azure Active Directory authority. The default is https://login.microsoft.com
	AuthorityHost *url.URL

	// HTTPClient sets the transport for making HTTP requests.
	// Leave this as nil to use the default HTTP transport.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// NewIdentityClientOptions initializes an instance of IdentityClientOptions with default settings
func (c *TokenCredentialOptions) setDefaultValues() *TokenCredentialOptions {
	if c == nil {
		c = defaultTokenCredentialOpts
	}

	if c.AuthorityHost == nil {
		c.AuthorityHost = defaultTokenCredentialOpts.AuthorityHost
	}
	if len(c.AuthorityHost.Path) == 0 || c.AuthorityHost.Path[len(c.AuthorityHost.Path)-1:] != "/" {
		c.AuthorityHost.Path = c.AuthorityHost.Path + "/"
	}

	return c
}

// NewDefaultPipeline creates a Pipeline using the specified pipeline options
func newDefaultPipeline(o TokenCredentialOptions) azcore.Pipeline {
	if o.HTTPClient == nil {
		o.HTTPClient = azcore.DefaultHTTPClientTransport()
	}

	return azcore.NewPipeline(
		o.HTTPClient,
		azcore.NewTelemetryPolicy(o.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(o.Retry),
		azcore.NewRequestLogPolicy(o.LogOptions))
}

// NewDefaultMSIPipeline creates a Pipeline using the specified pipeline options needed
// for a Managed Identity, such as a MSI specific retry policy
func newDefaultMSIPipeline(o ManagedIdentityCredentialOptions) azcore.Pipeline {
	if o.HTTPClient == nil {
		o.HTTPClient = azcore.DefaultHTTPClientTransport()
	}
	var statusCodes []int
	// retry policy for MSI is not end-user configurable
	retryOpts := azcore.RetryOptions{
		MaxTries:   5,
		RetryDelay: 2 * time.Second,
		StatusCodes: append(statusCodes,
			http.StatusRequestTimeout,      // 408
			http.StatusTooManyRequests,     // 429
			http.StatusInternalServerError, // 500
			http.StatusBadGateway,          // 502
			http.StatusGatewayTimeout,      // 504
			http.StatusNotFound,
			http.StatusGone,
			// all remaining 5xx
			http.StatusNotImplemented,
			http.StatusHTTPVersionNotSupported,
			http.StatusVariantAlsoNegotiates,
			http.StatusInsufficientStorage,
			http.StatusLoopDetected,
			http.StatusNotExtended,
			http.StatusNetworkAuthenticationRequired),
	}

	return azcore.NewPipeline(
		o.HTTPClient,
		azcore.NewTelemetryPolicy(o.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(retryOpts),
		azcore.NewRequestLogPolicy(o.LogOptions))
}

/*********************************************************************
func newDeviceCodeInfo(dc DeviceCodeResult) DeviceCodeInfo {
	return DeviceCodeInfo{UserCode: dc.UserCode, DeviceCode: dc.DeviceCode, VerificationUrl: dc.VerificationUri,
		ExpiresOn: dc.ExpiresOn, Internal: dc.Interval, Message: dc.Message, ClientId: dc.ClientId, Scopes: dc.Scopes}
}

// Details of the device code to present to a user to allow them to authenticate through the device code authentication flow.
type DeviceCodeInfo struct {
	// JMR: Make all these private and add public getter methods?
	UserCode        string        // User code returned by the service
	DeviceCode      string        // Device code returned by the service
	VerificationUrl string        // Verification URL where the user must navigate to authenticate using the device code and credentials. JMR: URL?
	ExpiresOn       time.Duration // Time when the device code will expire.
	Interval        int64         // Polling interval time to check for completion of authentication flow.
	Message         string        // User friendly text response that can be used for display purpose.
	ClientId        string        // Identifier of the client requesting device code.
	Scopes          []string      // List of the scopes that would be held by token. JMR: Should be readonly
}
**************************************************************/

/*******************************************************
const defaultSuffix = "/.default"

func scopesToResource(scope string) string {
	if strings.EndsWith(scope, defaultSuffix) {
		return scope[:len(scope)-len(defaultSuffix)]
	}
	return scope
}

func resourceToScope(resource string) string {
	resource += defaultSuffix
	return resource
}
**********************************************************/
