//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azappconfig

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-sdk-for-go/sdk/appconfiguration/azappconfig/internal/generated"
)

const timeFormat = time.RFC3339Nano

// Client is the struct for interacting with an Azure App Configuration instance.
type Client struct {
	appConfigClient *generated.AzureAppConfigurationClient
	syncTokenPolicy *syncTokenPolicy
}

// ClientOptions are the configurable options on a Client.
type ClientOptions struct {
	azcore.ClientOptions
}

func (c *ClientOptions) toConnectionOptions() *policy.ClientOptions {
	if c == nil {
		return nil
	}

	return &policy.ClientOptions{
		Logging:          c.Logging,
		Retry:            c.Retry,
		Telemetry:        c.Telemetry,
		Transport:        c.Transport,
		PerCallPolicies:  c.PerCallPolicies,
		PerRetryPolicies: c.PerRetryPolicies,
	}
}
func getDefaultScope(endpoint string) (string, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return "", errors.New("error parsing endpoint url")
	}

	return url.Scheme + "://" + url.Host + "/.default", nil
}

// NewClient returns a pointer to a Client object affinitized to an endpointUrl.
func NewClient(endpointUrl string, cred azcore.TokenCredential, options *ClientOptions) (Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	genOptions := options.toConnectionOptions()

	tokenScope, err := getDefaultScope(endpointUrl)
	if err != nil {
		return Client{}, err
	}

	syncTokenPolicy := newSyncTokenPolicy()
	genOptions.PerRetryPolicies = append(
		genOptions.PerRetryPolicies,
		runtime.NewBearerTokenPolicy(cred, []string{tokenScope}, nil),
		syncTokenPolicy,
	)

	pl := runtime.NewPipeline(generated.ModuleName, generated.ModuleVersion, runtime.PipelineOptions{}, genOptions)
	return Client{
		appConfigClient: generated.NewAzureAppConfigurationClient(endpointUrl, nil, pl),
		syncTokenPolicy: syncTokenPolicy,
	}, nil
}

// NewClientFromConnectionString parses the connection string and returns a pointer to a Client object.
func NewClientFromConnectionString(connectionString string, options *ClientOptions) (Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	genOptions := options.toConnectionOptions()

	endpoint, credential, secret, err := parseConnectionString(connectionString)
	if err != nil {
		return Client{}, err
	}

	syncTokenPolicy := newSyncTokenPolicy()
	genOptions.PerRetryPolicies = append(
		genOptions.PerRetryPolicies,
		newHmacAuthenticationPolicy(credential, secret),
		syncTokenPolicy,
	)

	pl := runtime.NewPipeline(generated.ModuleName, generated.ModuleVersion, runtime.PipelineOptions{}, genOptions)
	return Client{
		appConfigClient: generated.NewAzureAppConfigurationClient(endpoint, nil, pl),
		syncTokenPolicy: syncTokenPolicy,
	}, nil
}

// UpdateSyncToken sets an external synchronization token to ensure service requests receive up-to-date values.
func (c *Client) UpdateSyncToken(token string) {
	c.syncTokenPolicy.addToken(token)
}

func (cs Setting) toGeneratedPutOptions(ifMatch *azcore.ETag, ifNoneMatch *azcore.ETag) *generated.AzureAppConfigurationClientPutKeyValueOptions {
	return &generated.AzureAppConfigurationClientPutKeyValueOptions{
		Entity:      cs.toGenerated(),
		IfMatch:     (*string)(ifMatch),
		IfNoneMatch: (*string)(ifNoneMatch),
		Label:       cs.Label,
	}
}

// AddSettingResult contains the response from AddSetting method.
type AddSettingResult struct {
	Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string
}

func fromGeneratedAdd(g generated.AzureAppConfigurationClientPutKeyValueResponse) AddSettingResult {
	return AddSettingResult{
		Setting:   settingFromGenerated(g.KeyValue),
		SyncToken: g.SyncToken,
	}
}

// AddSettingOptions contains the optional parameters for the AddSetting method.
type AddSettingOptions struct {
}

// AddSetting creates a configuration setting only if the setting does not already exist in the configuration store.
func (c *Client) AddSetting(ctx context.Context, setting Setting, options *AddSettingOptions) (AddSettingResult, error) {
	etagAny := azcore.ETagAny
	resp, err := c.appConfigClient.PutKeyValue(ctx, *setting.Key, setting.toGeneratedPutOptions(nil, &etagAny))
	if err != nil {
		return AddSettingResult{}, err
	}

	return (AddSettingResult)(fromGeneratedAdd(resp)), nil
}

// DeleteSettingResult contains the response from DeleteSetting method.
type DeleteSettingResult struct {
	Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string
}

func fromGeneratedDelete(g generated.AzureAppConfigurationClientDeleteKeyValueResponse) DeleteSettingResult {
	return DeleteSettingResult{
		Setting:   settingFromGenerated(g.KeyValue),
		SyncToken: g.SyncToken,
	}
}

// DeleteSettingOptions contains the optional parameters for the DeleteSetting method.
type DeleteSettingOptions struct {
	// If set to true and the configuration setting exists in the configuration store,
	// delete the setting if the passed-in configuration setting is the same version as the one in the configuration store.
	// The setting versions are the same if their ETag fields match.
	OnlyIfUnchanged bool
}

func (cs Setting) toGeneratedDeleteOptions(ifMatch *azcore.ETag) *generated.AzureAppConfigurationClientDeleteKeyValueOptions {
	return &generated.AzureAppConfigurationClientDeleteKeyValueOptions{
		IfMatch: (*string)(ifMatch),
		Label:   cs.Label,
	}
}

// DeleteSetting deletes a configuration setting from the configuration store.
func (c *Client) DeleteSetting(ctx context.Context, setting Setting, options *DeleteSettingOptions) (DeleteSettingResult, error) {
	var ifMatch *azcore.ETag
	if options != nil && options.OnlyIfUnchanged {
		ifMatch = setting.ETag
	}

	resp, err := c.appConfigClient.DeleteKeyValue(ctx, *setting.Key, setting.toGeneratedDeleteOptions(ifMatch))
	if err != nil {
		return DeleteSettingResult{}, err
	}

	return fromGeneratedDelete(resp), nil
}

// GetSettingResult contains the configuration setting retrieved by GetSetting method.
type GetSettingResult struct {
	Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string

	// Contains the timestamp of when the configuration setting was last modified.
	LastModified *time.Time
}

func fromGeneratedGet(g generated.AzureAppConfigurationClientGetKeyValueResponse) GetSettingResult {
	var t *time.Time
	if g.LastModified != nil {
		if tt, err := time.Parse(timeFormat, *g.LastModified); err == nil {
			t = &tt
		}
	}

	return GetSettingResult{
		Setting:      settingFromGenerated(g.KeyValue),
		SyncToken:    g.SyncToken,
		LastModified: t,
	}
}

// GetSettingOptions contains the optional parameters for the GetSetting method.
type GetSettingOptions struct {
	// If set to true, only retrieve the setting from the configuration store if it has changed since the client last retrieved it.
	// It is determined to have changed if the ETag field on the passed-in configuration setting is different from the ETag
	// of the setting in the configuration store.
	OnlyIfChanged bool

	// The setting will be retrieved exactly as it existed at the provided time.
	AcceptDateTime *time.Time
}

func (cs Setting) toGeneratedGetOptions(ifNoneMatch *azcore.ETag, acceptDateTime *time.Time) *generated.AzureAppConfigurationClientGetKeyValueOptions {
	var dt *string
	if acceptDateTime != nil {
		str := acceptDateTime.Format(timeFormat)
		dt = &str
	}

	return &generated.AzureAppConfigurationClientGetKeyValueOptions{
		AcceptDatetime: dt,
		IfNoneMatch:    (*string)(ifNoneMatch),
		Label:          cs.Label,
	}
}

// GetSetting retrieves an existing configuration setting from the configuration store.
func (c *Client) GetSetting(ctx context.Context, setting Setting, options *GetSettingOptions) (GetSettingResult, error) {
	var ifNoneMatch *azcore.ETag
	var acceptDateTime *time.Time
	if options != nil {
		if options.OnlyIfChanged {
			ifNoneMatch = setting.ETag
		}

		acceptDateTime = options.AcceptDateTime
	}

	resp, err := c.appConfigClient.GetKeyValue(ctx, *setting.Key, setting.toGeneratedGetOptions(ifNoneMatch, acceptDateTime))
	if err != nil {
		return GetSettingResult{}, err
	}

	return fromGeneratedGet(resp), nil
}

// SetReadOnlyResult contains the response from SetReadOnly method.
type SetReadOnlyResult struct {
	Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string
}

func fromGeneratedPutLock(g generated.AzureAppConfigurationClientPutLockResponse) SetReadOnlyResult {
	return SetReadOnlyResult{
		Setting:   settingFromGenerated(g.KeyValue),
		SyncToken: g.SyncToken,
	}
}

func fromGeneratedDeleteLock(g generated.AzureAppConfigurationClientDeleteLockResponse) SetReadOnlyResult {
	return SetReadOnlyResult{
		Setting:   settingFromGenerated(g.KeyValue),
		SyncToken: g.SyncToken,
	}
}

// SetReadOnlyOptions contains the optional parameters for the SetReadOnly method.
type SetReadOnlyOptions struct {
	// If set to true and the configuration setting exists in the configuration store, update the setting
	// if the passed-in configuration setting is the same version as the one in the configuration store.
	// The setting versions are the same if their ETag fields match.
	OnlyIfUnchanged bool
}

func (cs Setting) toGeneratedPutLockOptions(ifMatch *azcore.ETag) *generated.AzureAppConfigurationClientPutLockOptions {
	return &generated.AzureAppConfigurationClientPutLockOptions{
		IfMatch: (*string)(ifMatch),
		Label:   cs.Label,
	}
}

func (cs Setting) toGeneratedDeleteLockOptions(ifMatch *azcore.ETag) *generated.AzureAppConfigurationClientDeleteLockOptions {
	return &generated.AzureAppConfigurationClientDeleteLockOptions{
		IfMatch: (*string)(ifMatch),
		Label:   cs.Label,
	}
}

// SetReadOnly sets an existing configuration setting to read only or read write state in the configuration store.
func (c *Client) SetReadOnly(ctx context.Context, setting Setting, isReadOnly bool, options *SetReadOnlyOptions) (SetReadOnlyResult, error) {
	var ifMatch *azcore.ETag
	if options != nil && options.OnlyIfUnchanged {
		ifMatch = setting.ETag
	}

	var err error
	if isReadOnly {
		var resp generated.AzureAppConfigurationClientPutLockResponse
		resp, err = c.appConfigClient.PutLock(ctx, *setting.Key, setting.toGeneratedPutLockOptions(ifMatch))
		if err == nil {
			return fromGeneratedPutLock(resp), nil
		}
	} else {
		var resp generated.AzureAppConfigurationClientDeleteLockResponse
		resp, err = c.appConfigClient.DeleteLock(ctx, *setting.Key, setting.toGeneratedDeleteLockOptions(ifMatch))
		if err == nil {
			return fromGeneratedDeleteLock(resp), nil
		}
	}

	return SetReadOnlyResult{}, err
}

// SetSettingResult contains the response from SetSetting method.
type SetSettingResult struct {
	Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string
}

func fromGeneratedSet(g generated.AzureAppConfigurationClientPutKeyValueResponse) SetSettingResult {
	return SetSettingResult{
		Setting:   settingFromGenerated(g.KeyValue),
		SyncToken: g.SyncToken,
	}
}

// SetSettingOptions contains the optional parameters for the SetSetting method.
type SetSettingOptions struct {
	// If set to true and the configuration setting exists in the configuration store, overwrite the setting
	// if the passed-in configuration setting is the same version as the one in the configuration store.
	// The setting versions are the same if their ETag fields match.
	OnlyIfUnchanged bool
}

// SetSetting creates a configuration setting if it doesn't exist or overwrites the existing setting in the configuration store.
func (c *Client) SetSetting(ctx context.Context, setting Setting, options *SetSettingOptions) (SetSettingResult, error) {
	var ifMatch *azcore.ETag
	if options != nil && options.OnlyIfUnchanged {
		ifMatch = setting.ETag
	}

	resp, err := c.appConfigClient.PutKeyValue(ctx, *setting.Key, setting.toGeneratedPutOptions(ifMatch, nil))
	if err != nil {
		return SetSettingResult{}, err
	}

	return (SetSettingResult)(fromGeneratedSet(resp)), nil
}

// ListRevisionsPage contains the configuration settings returned by ListRevisionsPager.
type ListRevisionsPage struct {
	// Contains the configuration settings returned that match the setting selector provided.
	Settings []Setting

	// Sync token for the Azure App Configuration client, corresponding to the current state of the client.
	SyncToken *string
}

func fromGeneratedGetRevisionsPage(g generated.AzureAppConfigurationClientGetRevisionsResponse) ListRevisionsPage {
	var css []Setting
	for _, cs := range g.Items {
		if cs != nil {
			css = append(css, settingFromGenerated(*cs))
		}
	}

	return ListRevisionsPage{
		Settings:  css,
		SyncToken: g.SyncToken,
	}
}

// ListRevisionsPager is a Pager for revision list operations.
//
// NextPage should be called first. It fetches the next available page of results from the service.
// If the fetched page contains results, the return value is true, else false.
// Results fetched from the service can be evaluated by calling PageResponse on this Pager.
// If the result is false, the value of Err() will indicate if an error occurred.
type ListRevisionsPager interface {
	// PageResponse returns the current ListRevisionsPage.
	PageResponse() ListRevisionsPage

	// Err returns an error if there was an error on the last request.
	Err() error

	// NextPage returns true if there is another page of data available, false if not.
	NextPage(context.Context) bool
}

type listRevisionsPager struct {
	genPager *generated.AzureAppConfigurationClientGetRevisionsPager
}

func (p listRevisionsPager) PageResponse() ListRevisionsPage {
	return fromGeneratedGetRevisionsPage(p.genPager.PageResponse())
}

func (p listRevisionsPager) Err() error {
	return p.genPager.Err()
}

func (p listRevisionsPager) NextPage(ctx context.Context) bool {
	return p.genPager.NextPage(ctx)
}

// ListRevisionsOptions contains the optional parameters for the ListRevisions method.
type ListRevisionsOptions struct {
}

// NewListRevisionsPager retrieves the revisions of one or more configuration setting entities that match the specified setting selector.
func (c *Client) NewListRevisionsPager(selector SettingSelector, options *ListRevisionsOptions) ListRevisionsPager {
	_ = options
	return listRevisionsPager{genPager: c.appConfigClient.GetRevisions(selector.toGenerated())}
}
