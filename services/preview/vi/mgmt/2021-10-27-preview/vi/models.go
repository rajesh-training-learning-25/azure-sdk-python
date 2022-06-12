package vi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/vi/mgmt/2021-10-27-preview/vi"

// AccessToken azure Video Analyzer for Media access token.
type AccessToken struct {
	autorest.Response `json:"-"`
	// AccessToken - READ-ONLY; The access token.
	AccessToken *string `json:"accessToken,omitempty"`
}

// MarshalJSON is the custom marshaler for AccessToken.
func (at AccessToken) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Account an Azure Video Analyzer for Media account.
type Account struct {
	autorest.Response `json:"-"`
	// AccountPropertiesForPutRequest - List of account properties
	*AccountPropertiesForPutRequest `json:"properties,omitempty"`
	Identity                        *ManagedServiceIdentity `json:"identity,omitempty"`
	// SystemData - READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AccountPropertiesForPutRequest != nil {
		objectMap["properties"] = a.AccountPropertiesForPutRequest
	}
	if a.Identity != nil {
		objectMap["identity"] = a.Identity
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountPropertiesForPutRequest AccountPropertiesForPutRequest
				err = json.Unmarshal(*v, &accountPropertiesForPutRequest)
				if err != nil {
					return err
				}
				a.AccountPropertiesForPutRequest = &accountPropertiesForPutRequest
			}
		case "identity":
			if v != nil {
				var identity ManagedServiceIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				a.Identity = &identity
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				a.SystemData = &systemData
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		}
	}

	return nil
}

// AccountCheckNameAvailabilityParameters the parameters used to check the availability of the Video
// Indexer account name.
type AccountCheckNameAvailabilityParameters struct {
	// Name - The VideoIndexer account name.
	Name *string `json:"name,omitempty"`
	// Type - The type of resource, Microsoft.VideoIndexer/accounts
	Type *string `json:"type,omitempty"`
}

// AccountList the list operation response, that contains the data pools and their properties.
type AccountList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountList.
func (al AccountList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if al.NextLink != nil {
		objectMap["nextLink"] = al.NextLink
	}
	return json.Marshal(objectMap)
}

// AccountListIterator provides access to a complete listing of Account values.
type AccountListIterator struct {
	i    int
	page AccountListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AccountListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListIterator) Response() AccountList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AccountListIterator type.
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return AccountListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (al AccountList) IsEmpty() bool {
	return al.Value == nil || len(*al.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (al AccountList) hasNextLink() bool {
	return al.NextLink != nil && len(*al.NextLink) != 0
}

// accountListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (al AccountList) accountListPreparer(ctx context.Context) (*http.Request, error) {
	if !al.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(al.NextLink)))
}

// AccountListPage contains a page of Account values.
type AccountListPage struct {
	fn func(context.Context, AccountList) (AccountList, error)
	al AccountList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.al)
		if err != nil {
			return err
		}
		page.al = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AccountListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListPage) NotDone() bool {
	return !page.al.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListPage) Response() AccountList {
	return page.al
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListPage) Values() []Account {
	if page.al.IsEmpty() {
		return nil
	}
	return *page.al.Value
}

// Creates a new instance of the AccountListPage type.
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return AccountListPage{
		fn: getNextPage,
		al: cur,
	}
}

// AccountPatch azure Video Analyzer for Media account
type AccountPatch struct {
	// AccountPropertiesForPatchRequest - List of account properties
	*AccountPropertiesForPatchRequest `json:"properties,omitempty"`
	Identity                          *ManagedServiceIdentity `json:"identity,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountPatch.
func (ap AccountPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.AccountPropertiesForPatchRequest != nil {
		objectMap["properties"] = ap.AccountPropertiesForPatchRequest
	}
	if ap.Identity != nil {
		objectMap["identity"] = ap.Identity
	}
	if ap.Tags != nil {
		objectMap["tags"] = ap.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountPatch struct.
func (ap *AccountPatch) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountPropertiesForPatchRequest AccountPropertiesForPatchRequest
				err = json.Unmarshal(*v, &accountPropertiesForPatchRequest)
				if err != nil {
					return err
				}
				ap.AccountPropertiesForPatchRequest = &accountPropertiesForPatchRequest
			}
		case "identity":
			if v != nil {
				var identity ManagedServiceIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				ap.Identity = &identity
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ap.Tags = tags
			}
		}
	}

	return nil
}

// AccountPropertiesForPatchRequest azure Video Analyzer for Media account properties
type AccountPropertiesForPatchRequest struct {
	// TenantID - READ-ONLY; The account's tenant id
	TenantID *string `json:"tenantId,omitempty"`
	// AccountID - READ-ONLY; The account's data-plane ID
	AccountID *string `json:"accountId,omitempty"`
	// MediaServices - The media services details
	MediaServices *MediaServicesForPatchRequest `json:"mediaServices,omitempty"`
	// ProvisioningState - READ-ONLY; Gets the status of the account at the time the operation was called. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Provisioning', 'Deleting'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountPropertiesForPatchRequest.
func (apfpr AccountPropertiesForPatchRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if apfpr.MediaServices != nil {
		objectMap["mediaServices"] = apfpr.MediaServices
	}
	return json.Marshal(objectMap)
}

// AccountPropertiesForPutRequest azure Video Analyzer for Media account properties
type AccountPropertiesForPutRequest struct {
	// TenantID - READ-ONLY; The account's tenant id
	TenantID *string `json:"tenantId,omitempty"`
	// AccountID - The account's data-plane ID. This can be set only when connecting an existing classic account
	AccountID *string `json:"accountId,omitempty"`
	// AccountName - READ-ONLY; The account's name
	AccountName *string `json:"accountName,omitempty"`
	// MediaServices - The media services details
	MediaServices *MediaServicesForPutRequest `json:"mediaServices,omitempty"`
	// ProvisioningState - READ-ONLY; Gets the status of the account at the time the operation was called. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Provisioning', 'Deleting'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountPropertiesForPutRequest.
func (apfpr AccountPropertiesForPutRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if apfpr.AccountID != nil {
		objectMap["accountId"] = apfpr.AccountID
	}
	if apfpr.MediaServices != nil {
		objectMap["mediaServices"] = apfpr.MediaServices
	}
	return json.Marshal(objectMap)
}

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for AzureEntityResource.
func (aer AzureEntityResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; Gets a boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; Gets the reason that a Video Indexer account name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'AlreadyExists'
	Reason Reason `json:"reason,omitempty"`
	// Message - READ-ONLY; Gets an error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckNameAvailabilityResult.
func (cnar CheckNameAvailabilityResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ClassicAccount an Azure Video Analyzer for Media classic account.
type ClassicAccount struct {
	autorest.Response `json:"-"`
	// Name - READ-ONLY; The account's name
	Name *string `json:"name,omitempty"`
	// Location - The account's location
	Location *string `json:"location,omitempty"`
	// ClassicAccountMediaServices - List of classic account properties
	*ClassicAccountMediaServices `json:"mediaServices,omitempty"`
}

// MarshalJSON is the custom marshaler for ClassicAccount.
func (ca ClassicAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ca.Location != nil {
		objectMap["location"] = ca.Location
	}
	if ca.ClassicAccountMediaServices != nil {
		objectMap["mediaServices"] = ca.ClassicAccountMediaServices
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ClassicAccount struct.
func (ca *ClassicAccount) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ca.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ca.Location = &location
			}
		case "mediaServices":
			if v != nil {
				var classicAccountMediaServices ClassicAccountMediaServices
				err = json.Unmarshal(*v, &classicAccountMediaServices)
				if err != nil {
					return err
				}
				ca.ClassicAccountMediaServices = &classicAccountMediaServices
			}
		}
	}

	return nil
}

// ClassicAccountMediaServices azure Video Analyzer for Media classic account properties
type ClassicAccountMediaServices struct {
	// AadApplicationID - The aad application id
	AadApplicationID *string `json:"aadApplicationId,omitempty"`
	// AadTenantID - The aad tenant id
	AadTenantID *string `json:"aadTenantId,omitempty"`
	// Connected - Represents wether the media services is connected or not
	Connected *bool `json:"connected,omitempty"`
	// EventGridProviderRegistered - Represents if the media services event grid is connected or not
	EventGridProviderRegistered *bool `json:"eventGridProviderRegistered,omitempty"`
	// Name - The media services name
	Name *string `json:"name,omitempty"`
	// ResourceGroup - The resource group that the media services belong to
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// StreamingEndpointStarted - Represents wether the media services streaming endpoint has started
	StreamingEndpointStarted *bool `json:"streamingEndpointStarted,omitempty"`
	// SubscriptionID - The media services subscriptionId
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// ClassicAccountSlim an Azure Video Analyzer for Media classic account.
type ClassicAccountSlim struct {
	// Name - READ-ONLY; The account's name
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; The account's location
	Location *string `json:"location,omitempty"`
}

// MarshalJSON is the custom marshaler for ClassicAccountSlim.
func (cas ClassicAccountSlim) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDefinition.
func (ed ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// GenerateAccessTokenParameters access token generation request's parameters
type GenerateAccessTokenParameters struct {
	// PermissionType - The requested permission. Possible values include: 'Contributor', 'Reader'
	PermissionType PermissionType `json:"permissionType,omitempty"`
	// Scope - The requested media type. Possible values include: 'ScopeVideo', 'ScopeAccount', 'ScopeProject'
	Scope Scope `json:"scope,omitempty"`
	// VideoID - The video ID
	VideoID *string `json:"videoId,omitempty"`
	// ProjectID - The project ID
	ProjectID *string `json:"projectId,omitempty"`
}

// ManagedServiceIdentity managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// PrincipalID - READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// Type - Possible values include: 'None', 'SystemAssigned', 'UserAssigned', 'SystemAssignedUserAssigned'
	Type                   ManagedServiceIdentityType       `json:"type,omitempty"`
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities"`
}

// MarshalJSON is the custom marshaler for ManagedServiceIdentity.
func (msi ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if msi.Type != "" {
		objectMap["type"] = msi.Type
	}
	if msi.UserAssignedIdentities != nil {
		objectMap["userAssignedIdentities"] = msi.UserAssignedIdentities
	}
	return json.Marshal(objectMap)
}

// MediaServicesForPatchRequest the media services details
type MediaServicesForPatchRequest struct {
	// UserAssignedIdentity - The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// MediaServicesForPutRequest the media services details
type MediaServicesForPutRequest struct {
	// ResourceID - The media services resource id
	ResourceID *string `json:"resourceId,omitempty"`
	// UserAssignedIdentity - The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Operation operation detail payload
type Operation struct {
	// Name - READ-ONLY; Name of the operation
	Name *string `json:"name,omitempty"`
	// IsDataAction - READ-ONLY; Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`
	// ActionType - READ-ONLY; Indicates the action type.
	ActionType *string `json:"actionType,omitempty"`
	// Display - READ-ONLY; Display of the operation
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - READ-ONLY; Origin of the operation
	Origin *string `json:"origin,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationDisplay operation display payload
type OperationDisplay struct {
	// Provider - READ-ONLY; Resource provider of the operation
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource of the operation
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Localized friendly name for the operation
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; Localized friendly description for the operation
	Description *string `json:"description,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDisplay.
func (od OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationListResult available operations of the service.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of operations supported by the Resource Provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// Tags resource tags
type Tags struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Tags.
func (t Tags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if t.Tags != nil {
		objectMap["tags"] = t.Tags
	}
	return json.Marshal(objectMap)
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// UserAssignedIdentity user assigned identity properties
type UserAssignedIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// ClientID - READ-ONLY; The client ID of the assigned identity.
	ClientID *uuid.UUID `json:"clientId,omitempty"`
}

// MarshalJSON is the custom marshaler for UserAssignedIdentity.
func (uai UserAssignedIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// UserClassicAccountList the list of user classic accounts.
type UserClassicAccountList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of classic account names and their location.
	Value *[]ClassicAccountSlim `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for UserClassicAccountList.
func (ucal UserClassicAccountList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ucal.NextLink != nil {
		objectMap["nextLink"] = ucal.NextLink
	}
	return json.Marshal(objectMap)
}

// UserClassicAccountListIterator provides access to a complete listing of ClassicAccountSlim values.
type UserClassicAccountListIterator struct {
	i    int
	page UserClassicAccountListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UserClassicAccountListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserClassicAccountListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *UserClassicAccountListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UserClassicAccountListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UserClassicAccountListIterator) Response() UserClassicAccountList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UserClassicAccountListIterator) Value() ClassicAccountSlim {
	if !iter.page.NotDone() {
		return ClassicAccountSlim{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the UserClassicAccountListIterator type.
func NewUserClassicAccountListIterator(page UserClassicAccountListPage) UserClassicAccountListIterator {
	return UserClassicAccountListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ucal UserClassicAccountList) IsEmpty() bool {
	return ucal.Value == nil || len(*ucal.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ucal UserClassicAccountList) hasNextLink() bool {
	return ucal.NextLink != nil && len(*ucal.NextLink) != 0
}

// userClassicAccountListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ucal UserClassicAccountList) userClassicAccountListPreparer(ctx context.Context) (*http.Request, error) {
	if !ucal.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ucal.NextLink)))
}

// UserClassicAccountListPage contains a page of ClassicAccountSlim values.
type UserClassicAccountListPage struct {
	fn   func(context.Context, UserClassicAccountList) (UserClassicAccountList, error)
	ucal UserClassicAccountList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UserClassicAccountListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserClassicAccountListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ucal)
		if err != nil {
			return err
		}
		page.ucal = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *UserClassicAccountListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UserClassicAccountListPage) NotDone() bool {
	return !page.ucal.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UserClassicAccountListPage) Response() UserClassicAccountList {
	return page.ucal
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UserClassicAccountListPage) Values() []ClassicAccountSlim {
	if page.ucal.IsEmpty() {
		return nil
	}
	return *page.ucal.Value
}

// Creates a new instance of the UserClassicAccountListPage type.
func NewUserClassicAccountListPage(cur UserClassicAccountList, getNextPage func(context.Context, UserClassicAccountList) (UserClassicAccountList, error)) UserClassicAccountListPage {
	return UserClassicAccountListPage{
		fn:   getNextPage,
		ucal: cur,
	}
}
