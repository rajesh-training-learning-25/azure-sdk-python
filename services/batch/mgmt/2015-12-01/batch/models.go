package batch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccountKeyType enumerates the values for account key type.
type AccountKeyType string

const (
	// Primary ...
	Primary AccountKeyType = "Primary"
	// Secondary ...
	Secondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns an array of possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() [2]AccountKeyType {
	return [2]AccountKeyType{Primary, Secondary}
}

// PackageState enumerates the values for package state.
type PackageState string

const (
	// Active ...
	Active PackageState = "active"
	// Pending ...
	Pending PackageState = "pending"
	// Unmapped ...
	Unmapped PackageState = "unmapped"
)

// PossiblePackageStateValues returns an array of possible values for the PackageState const type.
func PossiblePackageStateValues() [3]PackageState {
	return [3]PackageState{Active, Pending, Unmapped}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Invalid ...
	Invalid ProvisioningState = "Invalid"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() [6]ProvisioningState {
	return [6]ProvisioningState{Cancelled, Creating, Deleting, Failed, Invalid, Succeeded}
}

// Account contains information about an Azure Batch account.
type Account struct {
	autorest.Response `json:"-"`
	// AccountProperties - The properties associated with the account.
	*AccountProperties `json:"properties,omitempty"`
	// ID - The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
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
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
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
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
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
		}
	}

	return nil
}

// AccountBaseProperties the properties of a Batch account.
type AccountBaseProperties struct {
	// AutoStorage - The properties related to auto storage account.
	AutoStorage *AutoStorageBaseProperties `json:"autoStorage,omitempty"`
}

// AccountCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AccountCreateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future AccountCreateFuture) Result(client AccountClient) (a Account, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return a, azure.NewAsyncOpIncompleteError("batch.AccountCreateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		a, err = client.CreateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", resp, "Failure sending request")
		return
	}
	a, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// AccountCreateParameters parameters supplied to the Create operation.
type AccountCreateParameters struct {
	// Location - The region in which to create the account.
	Location *string `json:"location,omitempty"`
	// Tags - The user specified tags associated with the account.
	Tags map[string]*string `json:"tags"`
	// AccountBaseProperties - The properties of the account.
	*AccountBaseProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	if acp.AccountBaseProperties != nil {
		objectMap["properties"] = acp.AccountBaseProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountCreateParameters struct.
func (acp *AccountCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				acp.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				acp.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountBaseProperties AccountBaseProperties
				err = json.Unmarshal(*v, &accountBaseProperties)
				if err != nil {
					return err
				}
				acp.AccountBaseProperties = &accountBaseProperties
			}
		}
	}

	return nil
}

// AccountDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AccountDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future AccountDeleteFuture) Result(client AccountClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return ar, azure.NewAsyncOpIncompleteError("batch.AccountDeleteFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.AccountDeleteFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountDeleteFuture", "Result", resp, "Failure sending request")
		return
	}
	ar, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountDeleteFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// AccountKeys a set of Azure Batch account keys.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Primary - The primary key associated with the account.
	Primary *string `json:"primary,omitempty"`
	// Secondary - The secondary key associated with the account.
	Secondary *string `json:"secondary,omitempty"`
}

// AccountListResult values returned by the List operation.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of returned Batch accounts.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - The continuation token.
	NextLink *string `json:"nextLink,omitempty"`
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListResultIterator) Response() AccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListResultIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer() (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) Next() error {
	next, err := page.fn(page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListResultPage) Values() []Account {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// AccountProperties account specific properties.
type AccountProperties struct {
	// AccountEndpoint - The endpoint used by this account to interact with the Batch services.
	AccountEndpoint *string `json:"accountEndpoint,omitempty"`
	// ProvisioningState - The provisioned state of the resource. Possible values include: 'Invalid', 'Creating', 'Deleting', 'Succeeded', 'Failed', 'Cancelled'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AutoStorage - The properties and status of any auto storage account associated with the account.
	AutoStorage *AutoStorageProperties `json:"autoStorage,omitempty"`
	// CoreQuota - The core quota for this Batch account.
	CoreQuota *int32 `json:"coreQuota,omitempty"`
	// PoolQuota - The pool quota for this Batch account.
	PoolQuota *int32 `json:"poolQuota,omitempty"`
	// ActiveJobAndJobScheduleQuota - The active job and job schedule quota for this Batch account.
	ActiveJobAndJobScheduleQuota *int32 `json:"activeJobAndJobScheduleQuota,omitempty"`
}

// AccountRegenerateKeyParameters parameters supplied to the RegenerateKey operation.
type AccountRegenerateKeyParameters struct {
	// KeyName - The type of account key to regenerate. Possible values include: 'Primary', 'Secondary'
	KeyName AccountKeyType `json:"keyName,omitempty"`
}

// AccountUpdateParameters parameters supplied to the Update operation.
type AccountUpdateParameters struct {
	// Tags - The user specified tags associated with the account.
	Tags map[string]*string `json:"tags"`
	// AccountBaseProperties - The properties of the account.
	*AccountBaseProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.AccountBaseProperties != nil {
		objectMap["properties"] = aup.AccountBaseProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				aup.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountBaseProperties AccountBaseProperties
				err = json.Unmarshal(*v, &accountBaseProperties)
				if err != nil {
					return err
				}
				aup.AccountBaseProperties = &accountBaseProperties
			}
		}
	}

	return nil
}

// ActivateApplicationPackageParameters parameters for an ApplicationOperations.ActivateApplicationPackage request.
type ActivateApplicationPackageParameters struct {
	// Format - The format of the application package binary file.
	Format *string `json:"format,omitempty"`
}

// AddApplicationParameters parameters for an ApplicationOperations.AddApplication request.
type AddApplicationParameters struct {
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
}

// Application contains information about an application in a Batch account.
type Application struct {
	autorest.Response `json:"-"`
	// ID - A string that uniquely identifies the application within the account.
	ID *string `json:"id,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
	// Packages - The list of packages under this application.
	Packages *[]ApplicationPackage `json:"packages,omitempty"`
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DefaultVersion - The package to use if a client requests the application but does not specify a version.
	DefaultVersion *string `json:"defaultVersion,omitempty"`
}

// ApplicationPackage an application package which represents a particular version of an application.
type ApplicationPackage struct {
	autorest.Response `json:"-"`
	// ID - The ID of the application.
	ID *string `json:"id,omitempty"`
	// Version - The version of the application package.
	Version *string `json:"version,omitempty"`
	// State - The current state of the application package. Possible values include: 'Pending', 'Active', 'Unmapped'
	State PackageState `json:"state,omitempty"`
	// Format - The format of the application package, if the package is active.
	Format *string `json:"format,omitempty"`
	// StorageURL - The storage URL at which the application package is stored.
	StorageURL *string `json:"storageUrl,omitempty"`
	// StorageURLExpiry - The UTC time at which the storage URL will expire.
	StorageURLExpiry *date.Time `json:"storageUrlExpiry,omitempty"`
	// LastActivationTime - The time at which the package was last activated, if the package is active.
	LastActivationTime *date.Time `json:"lastActivationTime,omitempty"`
}

// AutoStorageBaseProperties the properties related to auto storage account.
type AutoStorageBaseProperties struct {
	// StorageAccountID - The resource ID of the storage account to be used for auto storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
}

// AutoStorageProperties contains information about the auto storage account associated with a Batch account.
type AutoStorageProperties struct {
	// StorageAccountID - The resource ID of the storage account to be used for auto storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
	// LastKeySync - The UTC time at which storage keys were last synchronized with the Batch account.
	LastKeySync *date.Time `json:"lastKeySync,omitempty"`
}

// CloudError an error response from the Batch service.
type CloudError struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudError `json:"details,omitempty"`
}

// ListApplicationsResult response to an ApplicationOperations.ListApplications request.
type ListApplicationsResult struct {
	autorest.Response `json:"-"`
	// Value - The list of applications.
	Value *[]Application `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListApplicationsResultIterator provides access to a complete listing of Application values.
type ListApplicationsResultIterator struct {
	i    int
	page ListApplicationsResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListApplicationsResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListApplicationsResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListApplicationsResultIterator) Response() ListApplicationsResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListApplicationsResultIterator) Value() Application {
	if !iter.page.NotDone() {
		return Application{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lar ListApplicationsResult) IsEmpty() bool {
	return lar.Value == nil || len(*lar.Value) == 0
}

// listApplicationsResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lar ListApplicationsResult) listApplicationsResultPreparer() (*http.Request, error) {
	if lar.NextLink == nil || len(to.String(lar.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lar.NextLink)))
}

// ListApplicationsResultPage contains a page of Application values.
type ListApplicationsResultPage struct {
	fn  func(ListApplicationsResult) (ListApplicationsResult, error)
	lar ListApplicationsResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListApplicationsResultPage) Next() error {
	next, err := page.fn(page.lar)
	if err != nil {
		return err
	}
	page.lar = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListApplicationsResultPage) NotDone() bool {
	return !page.lar.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListApplicationsResultPage) Response() ListApplicationsResult {
	return page.lar
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListApplicationsResultPage) Values() []Application {
	if page.lar.IsEmpty() {
		return nil
	}
	return *page.lar.Value
}

// LocationQuota quotas associated with a Batch region for a particular subscription.
type LocationQuota struct {
	autorest.Response `json:"-"`
	// AccountQuota - The number of Batch accounts that may be created under the subscription in the specified region.
	AccountQuota *int32 `json:"accountQuota,omitempty"`
}

// Resource a definition of an Azure resource.
type Resource struct {
	// ID - The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// UpdateApplicationParameters parameters for an ApplicationOperations.UpdateApplication request.
type UpdateApplicationParameters struct {
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DefaultVersion - The package to use if a client requests the application but does not specify a version.
	DefaultVersion *string `json:"defaultVersion,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
}
