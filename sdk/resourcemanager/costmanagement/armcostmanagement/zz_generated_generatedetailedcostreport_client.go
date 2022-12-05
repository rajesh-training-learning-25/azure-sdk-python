//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// GenerateDetailedCostReportClient contains the methods for the GenerateDetailedCostReport group.
// Don't use this type directly, use NewGenerateDetailedCostReportClient() instead.
type GenerateDetailedCostReportClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGenerateDetailedCostReportClient creates a new instance of GenerateDetailedCostReportClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGenerateDetailedCostReportClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateDetailedCostReportClient, error) {
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
	client := &GenerateDetailedCostReportClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOperation - Generates the detailed cost report for provided date range, billing period(Only enterprise customers)
// or Invoice Id asynchronously at a certain scope. Call returns a 202 with header
// Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status
// and if the operation is completed the blob file where generated detailed
// cost report is being stored.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with usage details operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for
// Billing Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount
// scope. Also, Modern Commerce Account scopes are '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for billingAccount
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile
// scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// parameters - Parameters supplied to the Create detailed cost report operation.
// options - GenerateDetailedCostReportClientBeginCreateOperationOptions contains the optional parameters for the GenerateDetailedCostReportClient.BeginCreateOperation
// method.
func (client *GenerateDetailedCostReportClient) BeginCreateOperation(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateDetailedCostReportClientCreateOperationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOperation(ctx, scope, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GenerateDetailedCostReportClientCreateOperationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GenerateDetailedCostReportClientCreateOperationResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOperation - Generates the detailed cost report for provided date range, billing period(Only enterprise customers)
// or Invoice Id asynchronously at a certain scope. Call returns a 202 with header
// Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status
// and if the operation is completed the blob file where generated detailed
// cost report is being stored.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
func (client *GenerateDetailedCostReportClient) createOperation(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*http.Response, error) {
	req, err := client.createOperationCreateRequest(ctx, scope, parameters, options)
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

// createOperationCreateRequest creates the CreateOperation request.
func (client *GenerateDetailedCostReportClient) createOperationCreateRequest(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/generateDetailedCostReport"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}