//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GenerateReservationDetailsReportClient contains the methods for the GenerateReservationDetailsReport group.
// Don't use this type directly, use NewGenerateReservationDetailsReportClient() instead.
type GenerateReservationDetailsReportClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGenerateReservationDetailsReportClient creates a new instance of GenerateReservationDetailsReportClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGenerateReservationDetailsReportClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateReservationDetailsReportClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &GenerateReservationDetailsReportClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginByBillingAccountID - Generates the reservations details report for provided date range asynchronously based on enrollment
// id. The Reservation usage details can be viewed only by certain enterprise roles. For more details
// on the roles see, https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/understand-ea-roles#usage-and-costs-access-by-role
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// billingAccountID - Enrollment ID (Legacy BillingAccount ID)
// startDate - Start Date
// endDate - End Date
// options - GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions contains the optional parameters for the
// GenerateReservationDetailsReportClient.BeginByBillingAccountID method.
func (client *GenerateReservationDetailsReportClient) BeginByBillingAccountID(ctx context.Context, billingAccountID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions) (*runtime.Poller[GenerateReservationDetailsReportClientByBillingAccountIDResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.byBillingAccountID(ctx, billingAccountID, startDate, endDate, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GenerateReservationDetailsReportClientByBillingAccountIDResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GenerateReservationDetailsReportClientByBillingAccountIDResponse](options.ResumeToken, client.pl, nil)
	}
}

// ByBillingAccountID - Generates the reservations details report for provided date range asynchronously based on enrollment
// id. The Reservation usage details can be viewed only by certain enterprise roles. For more details
// on the roles see, https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/understand-ea-roles#usage-and-costs-access-by-role
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
func (client *GenerateReservationDetailsReportClient) byBillingAccountID(ctx context.Context, billingAccountID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions) (*http.Response, error) {
	req, err := client.byBillingAccountIDCreateRequest(ctx, billingAccountID, startDate, endDate, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// byBillingAccountIDCreateRequest creates the ByBillingAccountID request.
func (client *GenerateReservationDetailsReportClient) byBillingAccountIDCreateRequest(ctx context.Context, billingAccountID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/generateReservationDetailsReport"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginByBillingProfileID - Generates the reservations details report for provided date range asynchronously by billing profile.
// The Reservation usage details can be viewed by only certain enterprise roles by default. For more
// details on the roles see, https://docs.microsoft.com/en-us/azure/cost-management-billing/reservations/reservation-utilization#view-utilization-in-the-azure-portal-with-azure-rbac-access
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// billingAccountID - BillingAccount ID
// billingProfileID - BillingProfile ID
// startDate - Start Date
// endDate - End Date
// options - GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions contains the optional parameters for the
// GenerateReservationDetailsReportClient.BeginByBillingProfileID method.
func (client *GenerateReservationDetailsReportClient) BeginByBillingProfileID(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions) (*runtime.Poller[GenerateReservationDetailsReportClientByBillingProfileIDResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.byBillingProfileID(ctx, billingAccountID, billingProfileID, startDate, endDate, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GenerateReservationDetailsReportClientByBillingProfileIDResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GenerateReservationDetailsReportClientByBillingProfileIDResponse](options.ResumeToken, client.pl, nil)
	}
}

// ByBillingProfileID - Generates the reservations details report for provided date range asynchronously by billing profile.
// The Reservation usage details can be viewed by only certain enterprise roles by default. For more
// details on the roles see, https://docs.microsoft.com/en-us/azure/cost-management-billing/reservations/reservation-utilization#view-utilization-in-the-azure-portal-with-azure-rbac-access
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
func (client *GenerateReservationDetailsReportClient) byBillingProfileID(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions) (*http.Response, error) {
	req, err := client.byBillingProfileIDCreateRequest(ctx, billingAccountID, billingProfileID, startDate, endDate, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// byBillingProfileIDCreateRequest creates the ByBillingProfileID request.
func (client *GenerateReservationDetailsReportClient) byBillingProfileIDCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.CostManagement/generateReservationDetailsReport"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}