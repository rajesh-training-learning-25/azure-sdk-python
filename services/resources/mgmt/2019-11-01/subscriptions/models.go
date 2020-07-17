package subscriptions

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
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"

// RegionCategory enumerates the values for region category.
type RegionCategory string

const (
	// Other ...
	Other RegionCategory = "Other"
	// Recommended ...
	Recommended RegionCategory = "Recommended"
)

// PossibleRegionCategoryValues returns an array of possible values for the RegionCategory const type.
func PossibleRegionCategoryValues() []RegionCategory {
	return []RegionCategory{Other, Recommended}
}

// RegionType enumerates the values for region type.
type RegionType string

const (
	// Logical ...
	Logical RegionType = "Logical"
	// Physical ...
	Physical RegionType = "Physical"
)

// PossibleRegionTypeValues returns an array of possible values for the RegionType const type.
func PossibleRegionTypeValues() []RegionType {
	return []RegionType{Logical, Physical}
}

// ResourceNameStatus enumerates the values for resource name status.
type ResourceNameStatus string

const (
	// Allowed ...
	Allowed ResourceNameStatus = "Allowed"
	// Reserved ...
	Reserved ResourceNameStatus = "Reserved"
)

// PossibleResourceNameStatusValues returns an array of possible values for the ResourceNameStatus const type.
func PossibleResourceNameStatusValues() []ResourceNameStatus {
	return []ResourceNameStatus{Allowed, Reserved}
}

// SpendingLimit enumerates the values for spending limit.
type SpendingLimit string

const (
	// CurrentPeriodOff ...
	CurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
	// Off ...
	Off SpendingLimit = "Off"
	// On ...
	On SpendingLimit = "On"
)

// PossibleSpendingLimitValues returns an array of possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{CurrentPeriodOff, Off, On}
}

// State enumerates the values for state.
type State string

const (
	// Deleted ...
	Deleted State = "Deleted"
	// Disabled ...
	Disabled State = "Disabled"
	// Enabled ...
	Enabled State = "Enabled"
	// PastDue ...
	PastDue State = "PastDue"
	// Warned ...
	Warned State = "Warned"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Deleted, Disabled, Enabled, PastDue, Warned}
}

// TenantCategory enumerates the values for tenant category.
type TenantCategory string

const (
	// Home ...
	Home TenantCategory = "Home"
	// ManagedBy ...
	ManagedBy TenantCategory = "ManagedBy"
	// ProjectedBy ...
	ProjectedBy TenantCategory = "ProjectedBy"
)

// PossibleTenantCategoryValues returns an array of possible values for the TenantCategory const type.
func PossibleTenantCategoryValues() []TenantCategory {
	return []TenantCategory{Home, ManagedBy, ProjectedBy}
}

// CheckResourceNameResult resource Name valid if not a reserved word, does not contain a reserved word and
// does not start with a reserved word
type CheckResourceNameResult struct {
	autorest.Response `json:"-"`
	// Name - Name of Resource
	Name *string `json:"name,omitempty"`
	// Type - Type of Resource
	Type *string `json:"type,omitempty"`
	// Status - Is the resource name Allowed or Reserved. Possible values include: 'Allowed', 'Reserved'
	Status ResourceNameStatus `json:"status,omitempty"`
}

// ErrorDefinition error description and code explaining why resource name is invalid.
type ErrorDefinition struct {
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
	// Code - Code of the error.
	Code *string `json:"code,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// ListResult subscription list operation response.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of subscriptions.
	Value *[]Subscription `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of Subscription values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
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
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
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
func (iter ListResultIterator) Value() Subscription {
	if !iter.page.NotDone() {
		return Subscription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of Subscription values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.lr)
	if err != nil {
		return err
	}
	page.lr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
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
func (page ListResultPage) Values() []Subscription {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{fn: getNextPage}
}

// Location location information.
type Location struct {
	// ID - READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// Name - READ-ONLY; The location name.
	Name *string `json:"name,omitempty"`
	// DisplayName - READ-ONLY; The display name of the location.
	DisplayName *string `json:"displayName,omitempty"`
	// RegionalDisplayName - READ-ONLY; The display name of the location and its region.
	RegionalDisplayName *string `json:"regionalDisplayName,omitempty"`
	// Metadata - Metadata of the location, such as lat/long, paired region, and others.
	Metadata *LocationMetadata `json:"metadata,omitempty"`
}

// LocationListResult location list operation response.
type LocationListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of locations.
	Value *[]Location `json:"value,omitempty"`
}

// LocationMetadata location metadata information
type LocationMetadata struct {
	// RegionType - READ-ONLY; The type of the region. Possible values include: 'Physical', 'Logical'
	RegionType RegionType `json:"regionType,omitempty"`
	// RegionCategory - READ-ONLY; The category of the region. Possible values include: 'Recommended', 'Other'
	RegionCategory RegionCategory `json:"regionCategory,omitempty"`
	// GeographyGroup - READ-ONLY; The geography group of the location.
	GeographyGroup *string `json:"geographyGroup,omitempty"`
	// Longitude - READ-ONLY; The longitude of the location.
	Longitude *string `json:"longitude,omitempty"`
	// Latitude - READ-ONLY; The latitude of the location.
	Latitude *string `json:"latitude,omitempty"`
	// PhysicalLocation - READ-ONLY; The physical location of the Azure location.
	PhysicalLocation *string `json:"physicalLocation,omitempty"`
	// PairedRegion - The regions paired to this region.
	PairedRegion *[]PairedRegion `json:"pairedRegion,omitempty"`
}

// ManagedByTenant information about a tenant managing the subscription.
type ManagedByTenant struct {
	// TenantID - READ-ONLY; The tenant ID of the managing tenant. This is a GUID.
	TenantID *string `json:"tenantId,omitempty"`
}

// Operation microsoft.Resources operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Resources
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Microsoft.Resources operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Microsoft.Resources operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
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

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
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
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
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
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// PairedRegion information regarding paired region.
type PairedRegion struct {
	// Name - READ-ONLY; The name of the paired region.
	Name *string `json:"name,omitempty"`
	// ID - READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// Policies subscription policies.
type Policies struct {
	// LocationPlacementID - READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example, a subscription with a location placement Id of Public_2014-09-01 has access to Azure public regions.
	LocationPlacementID *string `json:"locationPlacementId,omitempty"`
	// QuotaID - READ-ONLY; The subscription quota ID.
	QuotaID *string `json:"quotaId,omitempty"`
	// SpendingLimit - READ-ONLY; The subscription spending limit. Possible values include: 'On', 'Off', 'CurrentPeriodOff'
	SpendingLimit SpendingLimit `json:"spendingLimit,omitempty"`
}

// ResourceName name and Type of the Resource
type ResourceName struct {
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource
	Type *string `json:"type,omitempty"`
}

// Subscription subscription information.
type Subscription struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// DisplayName - READ-ONLY; The subscription display name.
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID - READ-ONLY; The subscription tenant ID.
	TenantID *string `json:"tenantId,omitempty"`
	// State - READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted. Possible values include: 'Enabled', 'Warned', 'PastDue', 'Disabled', 'Deleted'
	State State `json:"state,omitempty"`
	// SubscriptionPolicies - The subscription policies.
	SubscriptionPolicies *Policies `json:"subscriptionPolicies,omitempty"`
	// AuthorizationSource - The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string `json:"authorizationSource,omitempty"`
	// ManagedByTenants - An array containing the tenants managing the subscription.
	ManagedByTenants *[]ManagedByTenant `json:"managedByTenants,omitempty"`
	// Tags - The tags attached to the subscription.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Subscription.
func (s Subscription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.SubscriptionPolicies != nil {
		objectMap["subscriptionPolicies"] = s.SubscriptionPolicies
	}
	if s.AuthorizationSource != nil {
		objectMap["authorizationSource"] = s.AuthorizationSource
	}
	if s.ManagedByTenants != nil {
		objectMap["managedByTenants"] = s.ManagedByTenants
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	return json.Marshal(objectMap)
}

// TenantIDDescription tenant Id information.
type TenantIDDescription struct {
	// ID - READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// TenantID - READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string `json:"tenantId,omitempty"`
	// TenantCategory - READ-ONLY; Category of the tenant. Possible values include: 'Home', 'ProjectedBy', 'ManagedBy'
	TenantCategory TenantCategory `json:"tenantCategory,omitempty"`
	// Country - READ-ONLY; Country/region name of the address for the tenant.
	Country *string `json:"country,omitempty"`
	// CountryCode - READ-ONLY; Country/region abbreviation for the tenant.
	CountryCode *string `json:"countryCode,omitempty"`
	// DisplayName - READ-ONLY; The display name of the tenant.
	DisplayName *string `json:"displayName,omitempty"`
	// Domains - READ-ONLY; The list of domains for the tenant.
	Domains *[]string `json:"domains,omitempty"`
}

// TenantListResult tenant Ids information.
type TenantListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of tenants.
	Value *[]TenantIDDescription `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// TenantListResultIterator provides access to a complete listing of TenantIDDescription values.
type TenantListResultIterator struct {
	i    int
	page TenantListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *TenantListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultIterator.NextWithContext")
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
func (iter *TenantListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter TenantListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter TenantListResultIterator) Response() TenantListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter TenantListResultIterator) Value() TenantIDDescription {
	if !iter.page.NotDone() {
		return TenantIDDescription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the TenantListResultIterator type.
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return TenantListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (tlr TenantListResult) IsEmpty() bool {
	return tlr.Value == nil || len(*tlr.Value) == 0
}

// tenantListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (tlr TenantListResult) tenantListResultPreparer(ctx context.Context) (*http.Request, error) {
	if tlr.NextLink == nil || len(to.String(tlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(tlr.NextLink)))
}

// TenantListResultPage contains a page of TenantIDDescription values.
type TenantListResultPage struct {
	fn  func(context.Context, TenantListResult) (TenantListResult, error)
	tlr TenantListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *TenantListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.tlr)
	if err != nil {
		return err
	}
	page.tlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *TenantListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page TenantListResultPage) NotDone() bool {
	return !page.tlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page TenantListResultPage) Response() TenantListResult {
	return page.tlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page TenantListResultPage) Values() []TenantIDDescription {
	if page.tlr.IsEmpty() {
		return nil
	}
	return *page.tlr.Value
}

// Creates a new instance of the TenantListResultPage type.
func NewTenantListResultPage(getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
	return TenantListResultPage{fn: getNextPage}
}
