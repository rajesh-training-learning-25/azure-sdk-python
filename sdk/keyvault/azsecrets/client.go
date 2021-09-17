//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsecrets

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets/internal"
)

type Client struct {
	kvClient *internal.KeyVaultClient
	vaultUrl string
}

type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient policy.Transporter
	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

func (c *ClientOptions) toConnectionOptions() *internal.ConnectionOptions {
	if c == nil {
		return nil
	}

	return &internal.ConnectionOptions{
		HTTPClient:       c.HTTPClient,
		Retry:            c.Retry,
		Telemetry:        c.Telemetry,
		Logging:          c.Logging,
		PerCallPolicies:  c.PerCallPolicies,
		PerRetryPolicies: c.PerRetryPolicies,
	}
}

// NewClient returns a pointer to a Client object affinitized to a vaultUrl.
func NewClient(vaultUrl string, credential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	conn := internal.NewConnection(credential, options.toConnectionOptions())

	return &Client{
		kvClient: &internal.KeyVaultClient{
			Con: conn,
		},
		vaultUrl: vaultUrl,
	}, nil
}

type GetSecretOptions struct {
	Version string
}

func (g *GetSecretOptions) toGenerated() *internal.KeyVaultClientGetSecretOptions {
	if g == nil {
		return &internal.KeyVaultClientGetSecretOptions{}
	}
	return &internal.KeyVaultClientGetSecretOptions{}
}

type GetSecretResponse struct {
	SecretBundle
	RawResponse *http.Response
}

func getSecretResponseFromGenerated(i internal.KeyVaultClientGetSecretResponse) *GetSecretResponse {
	return &GetSecretResponse{
		RawResponse: i.RawResponse,
		SecretBundle: SecretBundle{
			Attributes:  secretAttributesFromGenerated(i.Attributes),
			ContentType: i.ContentType,
			ID:          i.ID,
			Tags:        i.Tags,
			Value:       i.Value,
			Kid:         i.Kid,
			Managed:     i.Managed,
		},
	}
}

func (c *Client) GetSecret(ctx context.Context, name string, options *GetSecretOptions) (GetSecretResponse, error) {
	if options == nil {
		options = &GetSecretOptions{}
	}
	resp, err := c.kvClient.GetSecret(ctx, c.vaultUrl, name, options.Version, options.toGenerated())
	return *getSecretResponseFromGenerated(resp), err
}

type SetSecretOptions struct {
	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// The secret management attributes.
	SecretAttributes *internal.SecretAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

func (s *SetSecretOptions) toGenerated() *internal.KeyVaultClientSetSecretOptions {
	if s == nil {
		return nil
	}
	return &internal.KeyVaultClientSetSecretOptions{}
}

type SetSecretResponse struct {
	RawResponse *http.Response

	// The secret management attributes.
	Attributes *internal.SecretAttributes `json:"attributes,omitempty"`

	// The secret id.
	ID *string `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// The secret value.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV certificate.
	Kid *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

func setSecretResponseFromGenerated(i internal.KeyVaultClientSetSecretResponse) *SetSecretResponse {
	return &SetSecretResponse{
		RawResponse: i.RawResponse,
		Attributes:  i.Attributes,
		ID:          i.ID,
		Tags:        i.Tags,
		Value:       i.Value,
		Kid:         i.Kid,
		Managed:     i.Managed,
	}
}

type SetSecretParameters struct {
}

func (c *Client) SetSecret(ctx context.Context, name string, value string, options *SetSecretOptions) (SetSecretResponse, error) {
	if options == nil {
		options = &SetSecretOptions{}
	}
	resp, err := c.kvClient.SetSecret(ctx, c.vaultUrl, name, internal.SecretSetParameters{
		Value:            &value,
		ContentType:      options.ContentType,
		SecretAttributes: options.SecretAttributes,
		Tags:             options.Tags,
	}, options.toGenerated())
	return *setSecretResponseFromGenerated(resp), err
}

type DeleteSecretResponse struct {
	DeletedSecretBundle
	// RawResponse holds the underlying HTTP response
	RawResponse *http.Response
}

func deleteSecretResponseFromGenerated(i *internal.KeyVaultClientDeleteSecretResponse) *DeleteSecretResponse {
	if i == nil {
		return nil
	}
	return &DeleteSecretResponse{
		DeletedSecretBundle: DeletedSecretBundle{
			SecretBundle: SecretBundle{
				ContentType: i.ContentType,
				ID:          i.ID,
				Tags:        i.Tags,
				Value:       i.Value,
				Kid:         i.Kid,
				Managed:     i.Managed,
				Attributes: &SecretAttributes{
					Attributes:      Attributes(i.Attributes.Attributes),
					RecoverableDays: i.Attributes.RecoverableDays,
					RecoveryLevel:   (*DeletionRecoveryLevel)(i.Attributes.RecoveryLevel),
				},
			},
			RecoveryID:         i.RecoveryID,
			DeletedDate:        i.DeletedDate,
			ScheduledPurgeDate: i.ScheduledPurgeDate,
		},
		RawResponse: i.RawResponse,
	}
}

// BeginDeleteSecretOptions contains the optional parameters for the Client.BeginDeleteSecret method.
type BeginDeleteSecretOptions struct {
	// placeholder for future optional parameters
}

func (b *BeginDeleteSecretOptions) toGenerated() *internal.KeyVaultClientDeleteSecretOptions {
	return &internal.KeyVaultClientDeleteSecretOptions{}
}

type DeleteSecretPoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// ResumeToken returns a value representing the poller that can be used to resume
	// the LRO at a later time. ResumeTokens are unique per service operation
	ResumeToken() string

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (DeleteSecretResponse, error)
}

type startDeletePoller struct {
	secretName     string // This is the secret to Poll for in GetDeletedSecret
	vaultUrl       string
	client         *internal.KeyVaultClient
	deleteResponse internal.KeyVaultClientDeleteSecretResponse
	lastResponse   internal.KeyVaultClientGetDeletedSecretResponse
	RawResponse    *http.Response
}

func (s *startDeletePoller) Done() bool {
	return s.lastResponse.RawResponse != nil
}

func (s *startDeletePoller) ResumeToken() string {
	return s.secretName
}

func (s *startDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := s.client.GetDeletedSecret(ctx, s.vaultUrl, s.secretName, nil)
	if err == nil {
		// Service recognizes DeletedSecret, operation is done
		s.lastResponse = resp
		return resp.RawResponse, nil
	} else if err != nil {
		return s.deleteResponse.RawResponse, nil
	}
	s.lastResponse = resp
	return resp.RawResponse, nil
}

func (s *startDeletePoller) FinalResponse(ctx context.Context) (DeleteSecretResponse, error) {
	return *deleteSecretResponseFromGenerated(&s.deleteResponse), nil
}

// pollUntilDone continually calls the Poll operation until the operation is completed. In between each
// Poll is a wait determined by the t parameter.
func (s *startDeletePoller) pollUntilDone(ctx context.Context, t time.Duration) (DeleteSecretResponse, error) {
	for {
		resp, err := s.Poll(ctx)
		if err != nil {
			return DeleteSecretResponse{}, err
		}
		s.RawResponse = resp
		if s.Done() {
			break
		}
		time.Sleep(t)
	}
	return DeleteSecretResponse{}, nil
}

type DeleteSecretPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (DeleteSecretResponse, error)

	// Poller contains an initialized WidgetPoller
	Poller DeleteSecretPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) BeginDeleteSecret(ctx context.Context, name string, options *BeginDeleteSecretOptions) (DeleteSecretPollerResponse, error) {
	// TODO: this is kvSecretClient.DeleteSecret and a GetDeletedSecret under the hood for the polling version
	if options == nil {
		options = &BeginDeleteSecretOptions{}
	}
	resp, err := c.kvClient.DeleteSecret(ctx, c.vaultUrl, name, options.toGenerated())
	if err != nil {
		return DeleteSecretPollerResponse{}, err
	}

	getResp, err := c.kvClient.GetDeletedSecret(ctx, c.vaultUrl, name, nil)
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		if httpErr.RawResponse().StatusCode != http.StatusNotFound {
			return DeleteSecretPollerResponse{}, err
		}
	}

	s := &startDeletePoller{
		vaultUrl:       c.vaultUrl,
		secretName:     name,
		client:         c.kvClient,
		deleteResponse: resp,
		lastResponse:   getResp,
	}

	return DeleteSecretPollerResponse{
		Poller:        s,
		RawResponse:   resp.RawResponse,
		PollUntilDone: s.pollUntilDone,
	}, nil
}

func (c *Client) CreateBeginDeleteSecretFromResumeToken(ctx context.Context, resumeToken string) DeleteSecretPoller {
	ret := &startDeletePoller{
		secretName: resumeToken,
		client:     c.kvClient,
	}
	return ret
}

type GetDeletedSecretOptions struct{}

func (g *GetDeletedSecretOptions) toGenerated() *internal.KeyVaultClientGetDeletedSecretOptions {
	return &internal.KeyVaultClientGetDeletedSecretOptions{}
}

type GetDeletedSecretResponse struct {
	SecretBundle
	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`

	RawResponse *http.Response
}

func getDeletedSecretResponseFromGenerated(i internal.KeyVaultClientGetDeletedSecretResponse) GetDeletedSecretResponse {
	return GetDeletedSecretResponse{
		RawResponse:        i.RawResponse,
		RecoveryID:         i.RecoveryID,
		DeletedDate:        i.DeletedDate,
		ScheduledPurgeDate: i.ScheduledPurgeDate,
		SecretBundle:       secretBundleFromGenerated(i.SecretBundle),
	}
}

func (c *Client) GetDeletedSecret(ctx context.Context, name string, options *GetDeletedSecretOptions) (GetDeletedSecretResponse, error) {
	if options == nil {
		options = &GetDeletedSecretOptions{}
	}

	resp, err := c.kvClient.GetDeletedSecret(ctx, c.vaultUrl, name, options.toGenerated())

	return getDeletedSecretResponseFromGenerated(resp), err
}

// UpdateSecretPropertiesOptions contains the optional parameters for the Client.UpdateSecretProperties method.
type UpdateSecretPropertiesOptions struct{}

func (u UpdateSecretPropertiesOptions) toGenerated() *internal.KeyVaultClientUpdateSecretOptions {
	return &internal.KeyVaultClientUpdateSecretOptions{}
}

type UpdateSecretPropertiesResponse struct {
	SecretBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func updateSecretPropertiesResponseFromGenerated(i internal.KeyVaultClientUpdateSecretResponse) UpdateSecretPropertiesResponse {
	return UpdateSecretPropertiesResponse{
		RawResponse: i.RawResponse,
		SecretBundle: SecretBundle{
			Attributes:  secretAttributesFromGenerated(i.Attributes),
			ContentType: i.ContentType,
			ID:          i.ID,
			Tags:        i.Tags,
			Value:       i.Value,
			Kid:         i.Kid,
			Managed:     i.Managed,
		},
	}
}

// SecretUpdateParameters - The secret update parameters.
type SecretProperties struct {
	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// The secret management attributes.
	SecretAttributes *SecretAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

func (s SecretProperties) toGenerated() internal.SecretUpdateParameters {
	var secAttribs *internal.SecretAttributes
	if s.SecretAttributes != nil {
		secAttribs = s.SecretAttributes.toGenerated()
	}
	return internal.SecretUpdateParameters{
		ContentType:      s.ContentType,
		Tags:             s.Tags,
		SecretAttributes: secAttribs,
	}
}

func (c *Client) UpdateSecretProperties(ctx context.Context, secretName string, secretVersion string, parameters SecretProperties, options *UpdateSecretPropertiesOptions) (UpdateSecretPropertiesResponse, error) {
	if options == nil {
		options = &UpdateSecretPropertiesOptions{}
	}

	resp, err := c.kvClient.UpdateSecret(ctx, c.vaultUrl, secretName, secretVersion, parameters.toGenerated(), options.toGenerated())
	if err != nil {
		return UpdateSecretPropertiesResponse{}, err
	}

	return updateSecretPropertiesResponseFromGenerated(resp), err
}

type BackupSecretOptions struct{}

type BackupSecretResponse struct{}

func (c *Client) BackupSecret(ctx context.Context, options *BackupSecretOptions) (BackupSecretResponse, error) {
	if options == nil {
		options = &BackupSecretOptions{}
	}

	return BackupSecretResponse{}, errors.New("not implemented")
}

type RestoreSecretBackupOptions struct{}

type RestoreSecretBackupResponse struct{}

func (c *Client) RestoreSecretBackup(ctx context.Context, options *RestoreSecretBackupOptions) (RestoreSecretBackupResponse, error) {
	if options == nil {
		options = &RestoreSecretBackupOptions{}
	}

	return RestoreSecretBackupResponse{}, errors.New("not implemented")
}

type PurgeDeletedSecretOptions struct{}

func (p *PurgeDeletedSecretOptions) toGenerated() *internal.KeyVaultClientPurgeDeletedSecretOptions {
	return &internal.KeyVaultClientPurgeDeletedSecretOptions{}
}

type PurgeDeletedSecretResponse struct {
	RawResponse *http.Response
}

func purgeDeletedSecretResponseFromGenerated(i internal.KeyVaultClientPurgeDeletedSecretResponse) PurgeDeletedSecretResponse {
	return PurgeDeletedSecretResponse{
		RawResponse: i.RawResponse,
	}
}

func (c *Client) PurgeDeletedSecret(ctx context.Context, secretName string, options *PurgeDeletedSecretOptions) (PurgeDeletedSecretResponse, error) {
	if options == nil {
		options = &PurgeDeletedSecretOptions{}
	}
	resp, err := c.kvClient.PurgeDeletedSecret(ctx, c.vaultUrl, secretName, options.toGenerated())
	return purgeDeletedSecretResponseFromGenerated(resp), err
}

type RecoverDeletedSecretPoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// ResumeToken returns a value representing the poller that can be used to resume
	// the LRO at a later time. ResumeTokens are unique per service operation
	ResumeToken() string

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (RecoverDeletedSecretResponse, error)
}

type beginRecoverPoller struct {
	secretName      string
	vaultUrl        string
	client          *internal.KeyVaultClient
	recoverResponse internal.KeyVaultClientRecoverDeletedSecretResponse
	lastResponse    internal.KeyVaultClientGetSecretResponse
	RawResponse     *http.Response
}

func (b *beginRecoverPoller) Done() bool {
	return b.RawResponse.StatusCode == http.StatusOK
}

func (b *beginRecoverPoller) ResumeToken() string {
	return b.secretName
}

func (b *beginRecoverPoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := b.client.GetSecret(ctx, b.vaultUrl, b.secretName, "", nil)
	b.lastResponse = resp
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		return httpErr.RawResponse(), err
	}
	return resp.RawResponse, nil
}

func (b *beginRecoverPoller) FinalResponse(ctx context.Context) (RecoverDeletedSecretResponse, error) {
	return recoverDeletedSecretResponseFromGenerated(b.recoverResponse), nil
}

func (b *beginRecoverPoller) pollUntilDone(ctx context.Context, t time.Duration) (RecoverDeletedSecretResponse, error) {
	for {
		resp, err := b.Poll(ctx)
		if err != nil {
			b.RawResponse = resp
		}
		if b.Done() {
			break
		}
		b.RawResponse = resp
		time.Sleep(t)
	}
	return recoverDeletedSecretResponseFromGenerated(b.recoverResponse), nil
}

type BeginRecoverDeletedSecretOptions struct{}

func (b BeginRecoverDeletedSecretOptions) toGenerated() *internal.KeyVaultClientRecoverDeletedSecretOptions {
	return &internal.KeyVaultClientRecoverDeletedSecretOptions{}
}

type RecoverDeletedSecretResponse struct {
	SecretBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func recoverDeletedSecretResponseFromGenerated(i internal.KeyVaultClientRecoverDeletedSecretResponse) RecoverDeletedSecretResponse {
	var a *SecretAttributes
	if i.Attributes != nil {
		a = &SecretAttributes{
			Attributes: Attributes{
				Enabled:   i.Attributes.Enabled,
				Expires:   i.Attributes.Expires,
				NotBefore: i.Attributes.NotBefore,
				Created:   i.Attributes.Created,
				Updated:   i.Attributes.Updated,
			},
			RecoverableDays: i.Attributes.RecoverableDays,
			RecoveryLevel:   (*DeletionRecoveryLevel)(i.Attributes.RecoveryLevel),
		}
	}
	return RecoverDeletedSecretResponse{
		RawResponse: i.RawResponse,
		SecretBundle: SecretBundle{
			Attributes:  a,
			ContentType: i.ContentType,
			ID:          i.ID,
			Tags:        i.Tags,
			Value:       i.Value,
			Kid:         i.Kid,
			Managed:     i.Managed,
		},
	}
}

type RecoverDeletedSecretPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (RecoverDeletedSecretResponse, error)

	// Poller contains an initialized RecoverDeletedSecretPoller
	Poller RecoverDeletedSecretPoller

	// RawResponse cotains the underlying HTTP response
	RawResponse *http.Response
}

func (c *Client) BeginRecoverDeletedSecret(ctx context.Context, secretName string, options *BeginRecoverDeletedSecretOptions) (RecoverDeletedSecretPollerResponse, error) {
	// This is a poller. Call RecoverDeletedSecret, then GetSecret until it is successful
	if options == nil {
		options = &BeginRecoverDeletedSecretOptions{}
	}
	resp, err := c.kvClient.RecoverDeletedSecret(ctx, c.vaultUrl, secretName, options.toGenerated())
	if err != nil {
		return RecoverDeletedSecretPollerResponse{}, err
	}

	getResp, err := c.kvClient.GetSecret(ctx, c.vaultUrl, secretName, "", nil)
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		if httpErr.RawResponse().StatusCode != http.StatusNotFound {
			return RecoverDeletedSecretPollerResponse{}, err
		}
	}

	b := &beginRecoverPoller{
		lastResponse:    getResp,
		secretName:      secretName,
		client:          c.kvClient,
		vaultUrl:        c.vaultUrl,
		recoverResponse: resp,
		RawResponse:     getResp.RawResponse,
	}

	return RecoverDeletedSecretPollerResponse{
		PollUntilDone: b.pollUntilDone,
		Poller:        b,
		RawResponse:   getResp.RawResponse,
	}, nil
}

type GetSecretVersionsPropertiesOptions struct{}

type GetSecretVersionsPropertiesResponse struct{}

func (c *Client) GetSecretVersionsProperties(ctx context.Context, options *GetSecretVersionsPropertiesOptions) (GetSecretVersionsPropertiesResponse, error) {
	if options == nil {
		options = &GetSecretVersionsPropertiesOptions{}
	}

	return GetSecretVersionsPropertiesResponse{}, errors.New("not implemented")
}

type ListDeletedSecretsPager interface {
	// PageResponse returns the current ListDeletedSecretPage
	PageResponse() ListDeletedSecretsPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

type listDeletedSecretsPager struct {
	genPager *internal.KeyVaultClientGetDeletedSecretsPager
}

func (l *listDeletedSecretsPager) PageResponse() ListDeletedSecretsPage {
	resp := l.genPager.PageResponse()

	var values []*DeletedSecretItem
	for _, d := range resp.Value {
		values = append(values, deletedSecretItemFromGenerated(d))
	}

	return ListDeletedSecretsPage{
		RawResponse: resp.RawResponse,
		NextLink:    resp.NextLink,
		Value:       values,
	}
}

func (l *listDeletedSecretsPager) Err() error {
	return l.genPager.Err()
}

func (l *listDeletedSecretsPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

type ListDeletedSecretsPage struct {
	RawResponse *http.Response
	// READ-ONLY; The URL to get the next set of deleted secrets.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of the deleted secrets in the vault along with a link to the next page of deleted secrets
	Value []*DeletedSecretItem `json:"value,omitempty" azure:"ro"`
}

type ListDeletedSecretsOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	Maxresults *int32
}

func (l *ListDeletedSecretsOptions) toGenerated() *internal.KeyVaultClientGetDeletedSecretsOptions {
	return &internal.KeyVaultClientGetDeletedSecretsOptions{
		Maxresults: l.Maxresults,
	}
}

type ListDeletedSecretsResponse struct{}

func (c *Client) ListDeletedSecrets(options *ListDeletedSecretsOptions) ListDeletedSecretsPager {
	if options == nil {
		options = &ListDeletedSecretsOptions{}
	}

	return &listDeletedSecretsPager{
		genPager: c.kvClient.GetDeletedSecrets(c.vaultUrl, options.toGenerated()),
	}

}

// ListSecretVersionsPager is a Pager for SecretVersion results
type ListSecretVersionsPager interface {
	// PageResponse returns the current ListSecretVersionsPage
	PageResponse() ListSecretVersionsPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

type listSecretVersionsPager struct {
	genPager *internal.KeyVaultClientGetSecretVersionsPager
}

// PageResponse returns the results from the page most recently fetched from the service.
func (l *listSecretVersionsPager) PageResponse() ListSecretVersionsPage {
	return listSecretVersionsPageFromGenerated(l.genPager.PageResponse())
}

// Err returns an error value if the most recent call to NextPage was not successful, else nil.
func (l *listSecretVersionsPager) Err() error {
	return l.genPager.Err()
}

// NextPage fetches the next available page of results from the service. If the fetched page
// contains results, the return value is true, else false. Results fetched from the service
// can be evaluated by calling PageResponse on this Pager.
func (l *listSecretVersionsPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

// ListSecretVersionsOptions contains the options for the ListSecretVersions operations
type ListSecretVersionsOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// convert the public ListSecretVersionsOptions to the generated version
func (l *ListSecretVersionsOptions) toGenerated() *internal.KeyVaultClientGetSecretVersionsOptions {
	if l == nil {
		return &internal.KeyVaultClientGetSecretVersionsOptions{}
	}
	return &internal.KeyVaultClientGetSecretVersionsOptions{
		Maxresults: l.MaxResults,
	}
}

// The secret list result
type ListSecretVersionsPage struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// READ-ONLY; The URL to get the next set of secrets.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of secrets in the key vault along with a link to the next page of secrets.
	Secrets []*SecretItem `json:"value,omitempty" azure:"ro"`
}

// create ListSecretsPage from generated pager
func listSecretVersionsPageFromGenerated(i internal.KeyVaultClientGetSecretVersionsResponse) ListSecretVersionsPage {
	var secrets []*SecretItem
	for _, s := range i.Value {
		secrets = append(secrets, secretItemFromGenerated(s))
	}
	return ListSecretVersionsPage{
		RawResponse: i.RawResponse,
		NextLink:    i.NextLink,
		Secrets:     secrets,
	}
}

// ListSecretVersions lists all versions of the specified secret. The full secret identifer and
// attributes are provided in the response. No values are returned for the secrets. This operation
// requires the secrets/list permission.
func (c *Client) ListSecretVersions(name string, options *ListSecretVersionsOptions) ListSecretVersionsPager {
	if options == nil {
		options = &ListSecretVersionsOptions{}
	}

	return &listSecretVersionsPager{
		genPager: c.kvClient.GetSecretVersions(
			c.vaultUrl,
			name,
			options.toGenerated(),
		),
	}
}

// ListSecretsPager is a Pager for SecretVersion results
type ListSecretsPager interface {
	// PageResponse returns the current ListSecretsPage
	PageResponse() ListSecretsPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

// listSecretsPager implements the ListSecretsPager interface
type listSecretsPager struct {
	genPager *internal.KeyVaultClientGetSecretsPager
}

// PageResponse returns the results from the page most recently fetched from the service
func (l *listSecretsPager) PageResponse() ListSecretsPage {
	return listSecretsPageFromGenerated(l.genPager.PageResponse())
}

// Err returns an error value if the most recent call to NextPage was not successful, else nil.
func (l *listSecretsPager) Err() error {
	return l.genPager.Err()
}

// NextPage fetches the next available page of results from the service. If the fetched page
// contains results, the return value is true, else false. Results fetched from the service
// can be evaluated by calling PageResponse on this Pager.
func (l *listSecretsPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

// ListSecretsOptions contains the options for the ListSecretVersions operations
type ListSecretsOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// create the generated code version of options
func (l *ListSecretsOptions) toGenerated() *internal.KeyVaultClientGetSecretsOptions {
	if l == nil {
		return nil
	}
	return &internal.KeyVaultClientGetSecretsOptions{
		Maxresults: l.MaxResults,
	}
}

// The list secrets result
type ListSecretsPage struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// READ-ONLY; The URL to get the next set of secrets.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of secrets in the key vault along with a link to the next page of secrets.
	Value []*SecretItem `json:"value,omitempty" azure:"ro"`
}

// create a ListSecretsPage from a generated code response
func listSecretsPageFromGenerated(i internal.KeyVaultClientGetSecretsResponse) ListSecretsPage {
	var secrets []*SecretItem
	for _, s := range i.Value {
		secrets = append(secrets, secretItemFromGenerated(s))
	}
	return ListSecretsPage{
		RawResponse: i.RawResponse,
		NextLink:    i.NextLink,
		Value:       secrets,
	}
}

// List secrets in a specified key vault. The ListSecrets operation is applicable to the entire vault.
// However, only the base secret identifier and its attributes are provided in the response. Individual
// secret versions are not listed in the response. This operation requires the secrets/list permission.
func (c *Client) ListSecrets(options *ListSecretsOptions) ListSecretsPager {
	if options == nil {
		options = &ListSecretsOptions{}
	}

	return &listSecretsPager{
		genPager: c.kvClient.GetSecrets(c.vaultUrl, options.toGenerated()),
	}
}
