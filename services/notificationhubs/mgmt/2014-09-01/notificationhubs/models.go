package notificationhubs

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

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen ...
	Listen AccessRights = "Listen"
	// Manage ...
	Manage AccessRights = "Manage"
	// Send ...
	Send AccessRights = "Send"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() [3]AccessRights {
	return [3]AccessRights{Listen, Manage, Send}
}

// NamespaceType enumerates the values for namespace type.
type NamespaceType string

const (
	// Messaging ...
	Messaging NamespaceType = "Messaging"
	// NotificationHub ...
	NotificationHub NamespaceType = "NotificationHub"
)

// PossibleNamespaceTypeValues returns an array of possible values for the NamespaceType const type.
func PossibleNamespaceTypeValues() [2]NamespaceType {
	return [2]NamespaceType{Messaging, NotificationHub}
}

// AdmCredential description of a NotificationHub AdmCredential.
type AdmCredential struct {
	// Properties - Gets or sets properties of NotificationHub AdmCredential.
	Properties *AdmCredentialProperties `json:"properties,omitempty"`
}

// AdmCredentialProperties description of a NotificationHub AdmCredential.
type AdmCredentialProperties struct {
	// ClientID - Gets or sets the client identifier.
	ClientID *string `json:"clientId,omitempty"`
	// ClientSecret - Gets or sets the credential secret access key.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// AuthTokenURL - Gets or sets the URL of the authorization token.
	AuthTokenURL *string `json:"authTokenUrl,omitempty"`
}

// ApnsCredential description of a NotificationHub ApnsCredential.
type ApnsCredential struct {
	// Properties - Gets or sets properties of NotificationHub ApnsCredential.
	Properties *ApnsCredentialProperties `json:"properties,omitempty"`
}

// ApnsCredentialProperties description of a NotificationHub ApnsCredential.
type ApnsCredentialProperties struct {
	// ApnsCertificate - Gets or sets the APNS certificate.
	ApnsCertificate *string `json:"apnsCertificate,omitempty"`
	// CertificateKey - Gets or sets the certificate key.
	CertificateKey *string `json:"certificateKey,omitempty"`
	// Endpoint - Gets or sets the endpoint of this credential.
	Endpoint *string `json:"endpoint,omitempty"`
	// Thumbprint - Gets or sets the Apns certificate Thumbprint
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// BaiduCredential description of a NotificationHub BaiduCredential.
type BaiduCredential struct {
	// Properties - Gets or sets properties of NotificationHub BaiduCredential.
	Properties *BaiduCredentialProperties `json:"properties,omitempty"`
}

// BaiduCredentialProperties description of a NotificationHub BaiduCredential.
type BaiduCredentialProperties struct {
	// BaiduAPIKey - Get or Set Baidu Api Key.
	BaiduAPIKey *string `json:"baiduApiKey,omitempty"`
	// BaiduEndPoint - Get or Set Baidu Endpoint.
	BaiduEndPoint *string `json:"baiduEndPoint,omitempty"`
	// BaiduSecretKey - Get or Set Baidu Secret Key
	BaiduSecretKey *string `json:"baiduSecretKey,omitempty"`
}

// CheckAvailabilityParameters parameters supplied to the Check Name Availability for Namespace and
// NotificationHubs.
type CheckAvailabilityParameters struct {
	// Name - Gets or sets name
	Name *string `json:"name,omitempty"`
	// Location - Gets or sets location.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets tags.
	Tags map[string]*string `json:"tags"`
	// IsAvailiable - Gets or sets true if the name is available and can be used to create new Namespace/NotificationHub. Otherwise false.
	IsAvailiable *bool `json:"isAvailiable,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckAvailabilityParameters.
func (capVar CheckAvailabilityParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if capVar.Name != nil {
		objectMap["name"] = capVar.Name
	}
	if capVar.Location != nil {
		objectMap["location"] = capVar.Location
	}
	if capVar.Tags != nil {
		objectMap["tags"] = capVar.Tags
	}
	if capVar.IsAvailiable != nil {
		objectMap["isAvailiable"] = capVar.IsAvailiable
	}
	return json.Marshal(objectMap)
}

// CheckAvailabilityResource description of a CheckAvailibility resource.
type CheckAvailabilityResource struct {
	autorest.Response `json:"-"`
	// ID - Gets or sets the id
	ID *string `json:"id,omitempty"`
	// Location - Gets or sets datacenter location
	Location *string `json:"location,omitempty"`
	// Name - Gets or sets name
	Name *string `json:"name,omitempty"`
	// Type - Gets or sets resource type
	Type *string `json:"type,omitempty"`
	// Tags - Gets or sets tags
	Tags map[string]*string `json:"tags"`
	// IsAvailiable - Gets or sets true if the name is available and can be used to create new Namespace/NotificationHub. Otherwise false.
	IsAvailiable *bool `json:"isAvailiable,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckAvailabilityResource.
func (car CheckAvailabilityResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if car.ID != nil {
		objectMap["id"] = car.ID
	}
	if car.Location != nil {
		objectMap["location"] = car.Location
	}
	if car.Name != nil {
		objectMap["name"] = car.Name
	}
	if car.Type != nil {
		objectMap["type"] = car.Type
	}
	if car.Tags != nil {
		objectMap["tags"] = car.Tags
	}
	if car.IsAvailiable != nil {
		objectMap["isAvailiable"] = car.IsAvailiable
	}
	return json.Marshal(objectMap)
}

// CreateOrUpdateParameters parameters supplied to the CreateOrUpdate NotificationHub operation.
type CreateOrUpdateParameters struct {
	// Location - Gets or sets NotificationHub data center location.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets NotificationHub tags.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets properties of the NotificationHub.
	Properties *Properties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for CreateOrUpdateParameters.
func (coup CreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if coup.Location != nil {
		objectMap["location"] = coup.Location
	}
	if coup.Tags != nil {
		objectMap["tags"] = coup.Tags
	}
	if coup.Properties != nil {
		objectMap["properties"] = coup.Properties
	}
	return json.Marshal(objectMap)
}

// GcmCredential description of a NotificationHub GcmCredential.
type GcmCredential struct {
	// Properties - Gets or sets properties of NotificationHub GcmCredential.
	Properties *GcmCredentialProperties `json:"properties,omitempty"`
}

// GcmCredentialProperties description of a NotificationHub GcmCredential.
type GcmCredentialProperties struct {
	// GcmEndpoint - Gets or sets the GCM endpoint.
	GcmEndpoint *string `json:"gcmEndpoint,omitempty"`
	// GoogleAPIKey - Gets or sets the Google API key.
	GoogleAPIKey *string `json:"googleApiKey,omitempty"`
}

// ListResult the response of the List NotificationHub operation.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets result of the List NotificationHub operation.
	Value *[]ResourceType `json:"value,omitempty"`
	// NextLink - Gets or sets link to the next set of results. Not empty if Value contains incomplete list of NotificationHub
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of ResourceType values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) Next() error {
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
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() ResourceType {
	if !iter.page.NotDone() {
		return ResourceType{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer() (*http.Request, error) {
	if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of ResourceType values.
type ListResultPage struct {
	fn func(ListResult) (ListResult, error)
	lr ListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) Next() error {
	next, err := page.fn(page.lr)
	if err != nil {
		return err
	}
	page.lr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []ResourceType {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// MpnsCredential description of a NotificationHub MpnsCredential.
type MpnsCredential struct {
	// Properties - Gets or sets properties of NotificationHub MpnsCredential.
	Properties *MpnsCredentialProperties `json:"properties,omitempty"`
}

// MpnsCredentialProperties description of a NotificationHub MpnsCredential.
type MpnsCredentialProperties struct {
	// MpnsCertificate - Gets or sets the MPNS certificate.
	MpnsCertificate *string `json:"mpnsCertificate,omitempty"`
	// CertificateKey - Gets or sets the certificate key for this credential.
	CertificateKey *string `json:"certificateKey,omitempty"`
	// Thumbprint - Gets or sets the Mpns certificate Thumbprint
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// NamespaceCreateOrUpdateParameters parameters supplied to the CreateOrUpdate Namespace operation.
type NamespaceCreateOrUpdateParameters struct {
	// Location - Gets or sets Namespace data center location.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets Namespace tags.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets properties of the Namespace.
	Properties *NamespaceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for NamespaceCreateOrUpdateParameters.
func (ncoup NamespaceCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ncoup.Location != nil {
		objectMap["location"] = ncoup.Location
	}
	if ncoup.Tags != nil {
		objectMap["tags"] = ncoup.Tags
	}
	if ncoup.Properties != nil {
		objectMap["properties"] = ncoup.Properties
	}
	return json.Marshal(objectMap)
}

// NamespaceListResult the response of the List Namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets result of the List Namespace operation.
	Value *[]NamespaceResource `json:"value,omitempty"`
	// NextLink - Gets or sets link to the next set of results. Not empty if Value contains incomplete list of Namespaces
	NextLink *string `json:"nextLink,omitempty"`
}

// NamespaceListResultIterator provides access to a complete listing of NamespaceResource values.
type NamespaceListResultIterator struct {
	i    int
	page NamespaceListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *NamespaceListResultIterator) Next() error {
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
func (iter NamespaceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter NamespaceListResultIterator) Response() NamespaceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter NamespaceListResultIterator) Value() NamespaceResource {
	if !iter.page.NotDone() {
		return NamespaceResource{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (nlr NamespaceListResult) IsEmpty() bool {
	return nlr.Value == nil || len(*nlr.Value) == 0
}

// namespaceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (nlr NamespaceListResult) namespaceListResultPreparer() (*http.Request, error) {
	if nlr.NextLink == nil || len(to.String(nlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(nlr.NextLink)))
}

// NamespaceListResultPage contains a page of NamespaceResource values.
type NamespaceListResultPage struct {
	fn  func(NamespaceListResult) (NamespaceListResult, error)
	nlr NamespaceListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *NamespaceListResultPage) Next() error {
	next, err := page.fn(page.nlr)
	if err != nil {
		return err
	}
	page.nlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page NamespaceListResultPage) NotDone() bool {
	return !page.nlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page NamespaceListResultPage) Response() NamespaceListResult {
	return page.nlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page NamespaceListResultPage) Values() []NamespaceResource {
	if page.nlr.IsEmpty() {
		return nil
	}
	return *page.nlr.Value
}

// NamespaceProperties namespace properties.
type NamespaceProperties struct {
	// Name - The name of the namespace.
	Name *string `json:"name,omitempty"`
	// ProvisioningState - Gets or sets provisioning state of the Namespace.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// Region - Specifies the targeted region in which the namespace should be created. It can be any of the following values: Australia EastAustralia SoutheastCentral USEast USEast US 2West USNorth Central USSouth Central USEast AsiaSoutheast AsiaBrazil SouthJapan EastJapan WestNorth EuropeWest Europe
	Region *string `json:"region,omitempty"`
	// Status - Status of the namespace. It can be any of these values:1 = Created/Active2 = Creating3 = Suspended4 = Deleting
	Status *string `json:"status,omitempty"`
	// CreatedAt - The time the namespace was created.
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// ServiceBusEndpoint - Endpoint you can use to perform NotificationHub operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`
	// SubscriptionID - The Id of the Azure subscription associated with the namespace.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ScaleUnit - ScaleUnit where the namespace gets created
	ScaleUnit *string `json:"scaleUnit,omitempty"`
	// Enabled - Whether or not the namespace is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Critical - Whether or not the namespace is set as Critical.
	Critical *bool `json:"critical,omitempty"`
	// NamespaceType - Gets or sets the namespace type. Possible values include: 'Messaging', 'NotificationHub'
	NamespaceType NamespaceType `json:"namespaceType,omitempty"`
}

// NamespaceResource description of a Namespace resource.
type NamespaceResource struct {
	autorest.Response `json:"-"`
	// ID - Gets or sets the id of the created Namespace.
	ID *string `json:"id,omitempty"`
	// Location - Gets or sets datacenter location of the Namespace.
	Location *string `json:"location,omitempty"`
	// Name - Gets or sets name of the Namespace.
	Name *string `json:"name,omitempty"`
	// Type - Gets or sets resource type of the Namespace.
	Type *string `json:"type,omitempty"`
	// Tags - Gets or sets tags of the Namespace.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets properties of the Namespace.
	Properties *NamespaceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for NamespaceResource.
func (nr NamespaceResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if nr.ID != nil {
		objectMap["id"] = nr.ID
	}
	if nr.Location != nil {
		objectMap["location"] = nr.Location
	}
	if nr.Name != nil {
		objectMap["name"] = nr.Name
	}
	if nr.Type != nil {
		objectMap["type"] = nr.Type
	}
	if nr.Tags != nil {
		objectMap["tags"] = nr.Tags
	}
	if nr.Properties != nil {
		objectMap["properties"] = nr.Properties
	}
	return json.Marshal(objectMap)
}

// NamespacesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type NamespacesDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future NamespacesDeleteFuture) Result(client NamespacesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationhubs.NamespacesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return ar, azure.NewAsyncOpIncompleteError("notificationhubs.NamespacesDeleteFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "notificationhubs.NamespacesDeleteFuture", "Result", future.Response(), "Failure responding to request")
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
		err = autorest.NewErrorWithError(err, "notificationhubs.NamespacesDeleteFuture", "Result", resp, "Failure sending request")
		return
	}
	ar, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationhubs.NamespacesDeleteFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// Properties notificationHub properties.
type Properties struct {
	// Name - The NotificationHub name.
	Name *string `json:"name,omitempty"`
	// RegistrationTTL - The RegistrationTtl of the created NotificationHub
	RegistrationTTL *string `json:"registrationTtl,omitempty"`
	// AuthorizationRules - The AuthorizationRules of the created NotificationHub
	AuthorizationRules *[]SharedAccessAuthorizationRuleProperties `json:"authorizationRules,omitempty"`
	// ApnsCredential - The ApnsCredential of the created NotificationHub
	ApnsCredential *ApnsCredential `json:"apnsCredential,omitempty"`
	// WnsCredential - The WnsCredential of the created NotificationHub
	WnsCredential *WnsCredential `json:"wnsCredential,omitempty"`
	// GcmCredential - The GcmCredential of the created NotificationHub
	GcmCredential *GcmCredential `json:"gcmCredential,omitempty"`
	// MpnsCredential - The MpnsCredential of the created NotificationHub
	MpnsCredential *MpnsCredential `json:"mpnsCredential,omitempty"`
	// AdmCredential - The AdmCredential of the created NotificationHub
	AdmCredential *AdmCredential `json:"admCredential,omitempty"`
	// BaiduCredential - The BaiduCredential of the created NotificationHub
	BaiduCredential *BaiduCredential `json:"baiduCredential,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
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

// ResourceListKeys namespace/NotificationHub Connection String
type ResourceListKeys struct {
	autorest.Response `json:"-"`
	// PrimaryConnectionString - Gets or sets the primaryConnectionString of the created Namespace AuthorizationRule.
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty"`
	// SecondaryConnectionString - Gets or sets the secondaryConnectionString of the created Namespace AuthorizationRule
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
}

// ResourceType description of a NotificationHub Resource.
type ResourceType struct {
	autorest.Response `json:"-"`
	// ID - Gets or sets the id of the created NotificationHub.
	ID *string `json:"id,omitempty"`
	// Location - Gets or sets datacenter location of the NotificationHub.
	Location *string `json:"location,omitempty"`
	// Name - Gets or sets name of the NotificationHub.
	Name *string `json:"name,omitempty"`
	// Type - Gets or sets resource type of the NotificationHub.
	Type *string `json:"type,omitempty"`
	// Tags - Gets or sets tags of the NotificationHub.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets properties of the NotificationHub.
	Properties *Properties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceType.
func (rt ResourceType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rt.ID != nil {
		objectMap["id"] = rt.ID
	}
	if rt.Location != nil {
		objectMap["location"] = rt.Location
	}
	if rt.Name != nil {
		objectMap["name"] = rt.Name
	}
	if rt.Type != nil {
		objectMap["type"] = rt.Type
	}
	if rt.Tags != nil {
		objectMap["tags"] = rt.Tags
	}
	if rt.Properties != nil {
		objectMap["properties"] = rt.Properties
	}
	return json.Marshal(objectMap)
}

// SharedAccessAuthorizationRuleCreateOrUpdateParameters parameters supplied to the CreateOrUpdate Namespace
// AuthorizationRules.
type SharedAccessAuthorizationRuleCreateOrUpdateParameters struct {
	// Location - Gets or sets Namespace data center location.
	Location *string `json:"location,omitempty"`
	// Name - Gets or sets Name of the Namespace AuthorizationRule.
	Name *string `json:"name,omitempty"`
	// Properties - Gets or sets properties of the Namespace AuthorizationRules.
	Properties *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// SharedAccessAuthorizationRuleListResult the response of the List Namespace operation.
type SharedAccessAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets result of the List AuthorizationRules operation.
	Value *[]SharedAccessAuthorizationRuleResource `json:"value,omitempty"`
	// NextLink - Gets or sets link to the next set of results. Not empty if Value contains incomplete list of AuthorizationRules
	NextLink *string `json:"nextLink,omitempty"`
}

// SharedAccessAuthorizationRuleListResultIterator provides access to a complete listing of
// SharedAccessAuthorizationRuleResource values.
type SharedAccessAuthorizationRuleListResultIterator struct {
	i    int
	page SharedAccessAuthorizationRuleListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SharedAccessAuthorizationRuleListResultIterator) Next() error {
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
func (iter SharedAccessAuthorizationRuleListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SharedAccessAuthorizationRuleListResultIterator) Response() SharedAccessAuthorizationRuleListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SharedAccessAuthorizationRuleListResultIterator) Value() SharedAccessAuthorizationRuleResource {
	if !iter.page.NotDone() {
		return SharedAccessAuthorizationRuleResource{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (saarlr SharedAccessAuthorizationRuleListResult) IsEmpty() bool {
	return saarlr.Value == nil || len(*saarlr.Value) == 0
}

// sharedAccessAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (saarlr SharedAccessAuthorizationRuleListResult) sharedAccessAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if saarlr.NextLink == nil || len(to.String(saarlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(saarlr.NextLink)))
}

// SharedAccessAuthorizationRuleListResultPage contains a page of SharedAccessAuthorizationRuleResource values.
type SharedAccessAuthorizationRuleListResultPage struct {
	fn     func(SharedAccessAuthorizationRuleListResult) (SharedAccessAuthorizationRuleListResult, error)
	saarlr SharedAccessAuthorizationRuleListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SharedAccessAuthorizationRuleListResultPage) Next() error {
	next, err := page.fn(page.saarlr)
	if err != nil {
		return err
	}
	page.saarlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SharedAccessAuthorizationRuleListResultPage) NotDone() bool {
	return !page.saarlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SharedAccessAuthorizationRuleListResultPage) Response() SharedAccessAuthorizationRuleListResult {
	return page.saarlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SharedAccessAuthorizationRuleListResultPage) Values() []SharedAccessAuthorizationRuleResource {
	if page.saarlr.IsEmpty() {
		return nil
	}
	return *page.saarlr.Value
}

// SharedAccessAuthorizationRuleProperties sharedAccessAuthorizationRule properties.
type SharedAccessAuthorizationRuleProperties struct {
	// PrimaryKey - The primary key that was used.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The secondary key that was used.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
	// KeyName - The name of the key that was used.
	KeyName *string `json:"keyName,omitempty"`
	// ClaimType - The type of the claim.
	ClaimType *string `json:"claimType,omitempty"`
	// ClaimValue - The value of the claim.
	ClaimValue *string `json:"claimValue,omitempty"`
	// Rights - The rights associated with the rule.
	Rights *[]AccessRights `json:"rights,omitempty"`
	// CreatedTime - The time at which the authorization rule was created.
	CreatedTime *date.Time `json:"createdTime,omitempty"`
	// ModifiedTime - The most recent time the rule was updated.
	ModifiedTime *date.Time `json:"modifiedTime,omitempty"`
	// Revision - The revision number for the rule.
	Revision *int32 `json:"revision,omitempty"`
}

// SharedAccessAuthorizationRuleResource description of a Namespace AuthorizationRules.
type SharedAccessAuthorizationRuleResource struct {
	autorest.Response `json:"-"`
	// ID - Gets or sets the id of the created Namespace AuthorizationRules.
	ID *string `json:"id,omitempty"`
	// Location - Gets or sets datacenter location of the Namespace AuthorizationRules.
	Location *string `json:"location,omitempty"`
	// Name - Gets or sets name of the Namespace AuthorizationRules.
	Name *string `json:"name,omitempty"`
	// Type - Gets or sets resource type of the Namespace AuthorizationRules.
	Type *string `json:"type,omitempty"`
	// Tags - Gets or sets tags of the Namespace AuthorizationRules.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets properties of the Namespace.
	Properties *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for SharedAccessAuthorizationRuleResource.
func (saarr SharedAccessAuthorizationRuleResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if saarr.ID != nil {
		objectMap["id"] = saarr.ID
	}
	if saarr.Location != nil {
		objectMap["location"] = saarr.Location
	}
	if saarr.Name != nil {
		objectMap["name"] = saarr.Name
	}
	if saarr.Type != nil {
		objectMap["type"] = saarr.Type
	}
	if saarr.Tags != nil {
		objectMap["tags"] = saarr.Tags
	}
	if saarr.Properties != nil {
		objectMap["properties"] = saarr.Properties
	}
	return json.Marshal(objectMap)
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// WnsCredential description of a NotificationHub WnsCredential.
type WnsCredential struct {
	// Properties - Gets or sets properties of NotificationHub WnsCredential.
	Properties *WnsCredentialProperties `json:"properties,omitempty"`
}

// WnsCredentialProperties description of a NotificationHub WnsCredential.
type WnsCredentialProperties struct {
	// PackageSid - Gets or sets the package ID for this credential.
	PackageSid *string `json:"packageSid,omitempty"`
	// SecretKey - Gets or sets the secret key.
	SecretKey *string `json:"secretKey,omitempty"`
	// WindowsLiveEndpoint - Gets or sets the Windows Live endpoint.
	WindowsLiveEndpoint *string `json:"windowsLiveEndpoint,omitempty"`
}
