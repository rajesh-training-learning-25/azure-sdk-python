//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAPIClient creates a new instance of APIClient.
func (c *ClientFactory) NewAPIClient() *APIClient {
	return &APIClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIDiagnosticClient creates a new instance of APIDiagnosticClient.
func (c *ClientFactory) NewAPIDiagnosticClient() *APIDiagnosticClient {
	return &APIDiagnosticClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIExportClient creates a new instance of APIExportClient.
func (c *ClientFactory) NewAPIExportClient() *APIExportClient {
	return &APIExportClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIIssueAttachmentClient creates a new instance of APIIssueAttachmentClient.
func (c *ClientFactory) NewAPIIssueAttachmentClient() *APIIssueAttachmentClient {
	return &APIIssueAttachmentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIIssueClient creates a new instance of APIIssueClient.
func (c *ClientFactory) NewAPIIssueClient() *APIIssueClient {
	return &APIIssueClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIIssueCommentClient creates a new instance of APIIssueCommentClient.
func (c *ClientFactory) NewAPIIssueCommentClient() *APIIssueCommentClient {
	return &APIIssueCommentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIOperationClient creates a new instance of APIOperationClient.
func (c *ClientFactory) NewAPIOperationClient() *APIOperationClient {
	return &APIOperationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIOperationPolicyClient creates a new instance of APIOperationPolicyClient.
func (c *ClientFactory) NewAPIOperationPolicyClient() *APIOperationPolicyClient {
	return &APIOperationPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIPolicyClient creates a new instance of APIPolicyClient.
func (c *ClientFactory) NewAPIPolicyClient() *APIPolicyClient {
	return &APIPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIProductClient creates a new instance of APIProductClient.
func (c *ClientFactory) NewAPIProductClient() *APIProductClient {
	return &APIProductClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIReleaseClient creates a new instance of APIReleaseClient.
func (c *ClientFactory) NewAPIReleaseClient() *APIReleaseClient {
	return &APIReleaseClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIRevisionClient creates a new instance of APIRevisionClient.
func (c *ClientFactory) NewAPIRevisionClient() *APIRevisionClient {
	return &APIRevisionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPISchemaClient creates a new instance of APISchemaClient.
func (c *ClientFactory) NewAPISchemaClient() *APISchemaClient {
	return &APISchemaClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPITagDescriptionClient creates a new instance of APITagDescriptionClient.
func (c *ClientFactory) NewAPITagDescriptionClient() *APITagDescriptionClient {
	return &APITagDescriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIVersionSetClient creates a new instance of APIVersionSetClient.
func (c *ClientFactory) NewAPIVersionSetClient() *APIVersionSetClient {
	return &APIVersionSetClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIWikiClient creates a new instance of APIWikiClient.
func (c *ClientFactory) NewAPIWikiClient() *APIWikiClient {
	return &APIWikiClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIWikisClient creates a new instance of APIWikisClient.
func (c *ClientFactory) NewAPIWikisClient() *APIWikisClient {
	return &APIWikisClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAuthorizationAccessPolicyClient creates a new instance of AuthorizationAccessPolicyClient.
func (c *ClientFactory) NewAuthorizationAccessPolicyClient() *AuthorizationAccessPolicyClient {
	return &AuthorizationAccessPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAuthorizationClient creates a new instance of AuthorizationClient.
func (c *ClientFactory) NewAuthorizationClient() *AuthorizationClient {
	return &AuthorizationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAuthorizationLoginLinksClient creates a new instance of AuthorizationLoginLinksClient.
func (c *ClientFactory) NewAuthorizationLoginLinksClient() *AuthorizationLoginLinksClient {
	return &AuthorizationLoginLinksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAuthorizationProviderClient creates a new instance of AuthorizationProviderClient.
func (c *ClientFactory) NewAuthorizationProviderClient() *AuthorizationProviderClient {
	return &AuthorizationProviderClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAuthorizationServerClient creates a new instance of AuthorizationServerClient.
func (c *ClientFactory) NewAuthorizationServerClient() *AuthorizationServerClient {
	return &AuthorizationServerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackendClient creates a new instance of BackendClient.
func (c *ClientFactory) NewBackendClient() *BackendClient {
	return &BackendClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCacheClient creates a new instance of CacheClient.
func (c *ClientFactory) NewCacheClient() *CacheClient {
	return &CacheClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCertificateClient creates a new instance of CertificateClient.
func (c *ClientFactory) NewCertificateClient() *CertificateClient {
	return &CertificateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContentItemClient creates a new instance of ContentItemClient.
func (c *ClientFactory) NewContentItemClient() *ContentItemClient {
	return &ContentItemClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContentTypeClient creates a new instance of ContentTypeClient.
func (c *ClientFactory) NewContentTypeClient() *ContentTypeClient {
	return &ContentTypeClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDelegationSettingsClient creates a new instance of DelegationSettingsClient.
func (c *ClientFactory) NewDelegationSettingsClient() *DelegationSettingsClient {
	return &DelegationSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedServicesClient creates a new instance of DeletedServicesClient.
func (c *ClientFactory) NewDeletedServicesClient() *DeletedServicesClient {
	return &DeletedServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiagnosticClient creates a new instance of DiagnosticClient.
func (c *ClientFactory) NewDiagnosticClient() *DiagnosticClient {
	return &DiagnosticClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDocumentationClient creates a new instance of DocumentationClient.
func (c *ClientFactory) NewDocumentationClient() *DocumentationClient {
	return &DocumentationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEmailTemplateClient creates a new instance of EmailTemplateClient.
func (c *ClientFactory) NewEmailTemplateClient() *EmailTemplateClient {
	return &EmailTemplateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGatewayAPIClient creates a new instance of GatewayAPIClient.
func (c *ClientFactory) NewGatewayAPIClient() *GatewayAPIClient {
	return &GatewayAPIClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGatewayCertificateAuthorityClient creates a new instance of GatewayCertificateAuthorityClient.
func (c *ClientFactory) NewGatewayCertificateAuthorityClient() *GatewayCertificateAuthorityClient {
	return &GatewayCertificateAuthorityClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGatewayClient creates a new instance of GatewayClient.
func (c *ClientFactory) NewGatewayClient() *GatewayClient {
	return &GatewayClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGatewayHostnameConfigurationClient creates a new instance of GatewayHostnameConfigurationClient.
func (c *ClientFactory) NewGatewayHostnameConfigurationClient() *GatewayHostnameConfigurationClient {
	return &GatewayHostnameConfigurationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGlobalSchemaClient creates a new instance of GlobalSchemaClient.
func (c *ClientFactory) NewGlobalSchemaClient() *GlobalSchemaClient {
	return &GlobalSchemaClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGraphQLAPIResolverClient creates a new instance of GraphQLAPIResolverClient.
func (c *ClientFactory) NewGraphQLAPIResolverClient() *GraphQLAPIResolverClient {
	return &GraphQLAPIResolverClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGraphQLAPIResolverPolicyClient creates a new instance of GraphQLAPIResolverPolicyClient.
func (c *ClientFactory) NewGraphQLAPIResolverPolicyClient() *GraphQLAPIResolverPolicyClient {
	return &GraphQLAPIResolverPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGroupClient creates a new instance of GroupClient.
func (c *ClientFactory) NewGroupClient() *GroupClient {
	return &GroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGroupUserClient creates a new instance of GroupUserClient.
func (c *ClientFactory) NewGroupUserClient() *GroupUserClient {
	return &GroupUserClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIdentityProviderClient creates a new instance of IdentityProviderClient.
func (c *ClientFactory) NewIdentityProviderClient() *IdentityProviderClient {
	return &IdentityProviderClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIssueClient creates a new instance of IssueClient.
func (c *ClientFactory) NewIssueClient() *IssueClient {
	return &IssueClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLoggerClient creates a new instance of LoggerClient.
func (c *ClientFactory) NewLoggerClient() *LoggerClient {
	return &LoggerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNamedValueClient creates a new instance of NamedValueClient.
func (c *ClientFactory) NewNamedValueClient() *NamedValueClient {
	return &NamedValueClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkStatusClient creates a new instance of NetworkStatusClient.
func (c *ClientFactory) NewNetworkStatusClient() *NetworkStatusClient {
	return &NetworkStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNotificationClient creates a new instance of NotificationClient.
func (c *ClientFactory) NewNotificationClient() *NotificationClient {
	return &NotificationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNotificationRecipientEmailClient creates a new instance of NotificationRecipientEmailClient.
func (c *ClientFactory) NewNotificationRecipientEmailClient() *NotificationRecipientEmailClient {
	return &NotificationRecipientEmailClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNotificationRecipientUserClient creates a new instance of NotificationRecipientUserClient.
func (c *ClientFactory) NewNotificationRecipientUserClient() *NotificationRecipientUserClient {
	return &NotificationRecipientUserClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOpenIDConnectProviderClient creates a new instance of OpenIDConnectProviderClient.
func (c *ClientFactory) NewOpenIDConnectProviderClient() *OpenIDConnectProviderClient {
	return &OpenIDConnectProviderClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationClient creates a new instance of OperationClient.
func (c *ClientFactory) NewOperationClient() *OperationClient {
	return &OperationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewOutboundNetworkDependenciesEndpointsClient creates a new instance of OutboundNetworkDependenciesEndpointsClient.
func (c *ClientFactory) NewOutboundNetworkDependenciesEndpointsClient() *OutboundNetworkDependenciesEndpointsClient {
	return &OutboundNetworkDependenciesEndpointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPolicyClient creates a new instance of PolicyClient.
func (c *ClientFactory) NewPolicyClient() *PolicyClient {
	return &PolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPolicyDescriptionClient creates a new instance of PolicyDescriptionClient.
func (c *ClientFactory) NewPolicyDescriptionClient() *PolicyDescriptionClient {
	return &PolicyDescriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPolicyFragmentClient creates a new instance of PolicyFragmentClient.
func (c *ClientFactory) NewPolicyFragmentClient() *PolicyFragmentClient {
	return &PolicyFragmentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPortalConfigClient creates a new instance of PortalConfigClient.
func (c *ClientFactory) NewPortalConfigClient() *PortalConfigClient {
	return &PortalConfigClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPortalRevisionClient creates a new instance of PortalRevisionClient.
func (c *ClientFactory) NewPortalRevisionClient() *PortalRevisionClient {
	return &PortalRevisionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPortalSettingsClient creates a new instance of PortalSettingsClient.
func (c *ClientFactory) NewPortalSettingsClient() *PortalSettingsClient {
	return &PortalSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient.
func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	return &PrivateEndpointConnectionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductAPIClient creates a new instance of ProductAPIClient.
func (c *ClientFactory) NewProductAPIClient() *ProductAPIClient {
	return &ProductAPIClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductClient creates a new instance of ProductClient.
func (c *ClientFactory) NewProductClient() *ProductClient {
	return &ProductClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductGroupClient creates a new instance of ProductGroupClient.
func (c *ClientFactory) NewProductGroupClient() *ProductGroupClient {
	return &ProductGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductPolicyClient creates a new instance of ProductPolicyClient.
func (c *ClientFactory) NewProductPolicyClient() *ProductPolicyClient {
	return &ProductPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductSubscriptionsClient creates a new instance of ProductSubscriptionsClient.
func (c *ClientFactory) NewProductSubscriptionsClient() *ProductSubscriptionsClient {
	return &ProductSubscriptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductWikiClient creates a new instance of ProductWikiClient.
func (c *ClientFactory) NewProductWikiClient() *ProductWikiClient {
	return &ProductWikiClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProductWikisClient creates a new instance of ProductWikisClient.
func (c *ClientFactory) NewProductWikisClient() *ProductWikisClient {
	return &ProductWikisClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewQuotaByCounterKeysClient creates a new instance of QuotaByCounterKeysClient.
func (c *ClientFactory) NewQuotaByCounterKeysClient() *QuotaByCounterKeysClient {
	return &QuotaByCounterKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewQuotaByPeriodKeysClient creates a new instance of QuotaByPeriodKeysClient.
func (c *ClientFactory) NewQuotaByPeriodKeysClient() *QuotaByPeriodKeysClient {
	return &QuotaByPeriodKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRegionClient creates a new instance of RegionClient.
func (c *ClientFactory) NewRegionClient() *RegionClient {
	return &RegionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewReportsClient creates a new instance of ReportsClient.
func (c *ClientFactory) NewReportsClient() *ReportsClient {
	return &ReportsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSKUsClient creates a new instance of SKUsClient.
func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	return &SKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServiceClient creates a new instance of ServiceClient.
func (c *ClientFactory) NewServiceClient() *ServiceClient {
	return &ServiceClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServiceSKUsClient creates a new instance of ServiceSKUsClient.
func (c *ClientFactory) NewServiceSKUsClient() *ServiceSKUsClient {
	return &ServiceSKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSignInSettingsClient creates a new instance of SignInSettingsClient.
func (c *ClientFactory) NewSignInSettingsClient() *SignInSettingsClient {
	return &SignInSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSignUpSettingsClient creates a new instance of SignUpSettingsClient.
func (c *ClientFactory) NewSignUpSettingsClient() *SignUpSettingsClient {
	return &SignUpSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSubscriptionClient creates a new instance of SubscriptionClient.
func (c *ClientFactory) NewSubscriptionClient() *SubscriptionClient {
	return &SubscriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTagClient creates a new instance of TagClient.
func (c *ClientFactory) NewTagClient() *TagClient {
	return &TagClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTagResourceClient creates a new instance of TagResourceClient.
func (c *ClientFactory) NewTagResourceClient() *TagResourceClient {
	return &TagResourceClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTenantAccessClient creates a new instance of TenantAccessClient.
func (c *ClientFactory) NewTenantAccessClient() *TenantAccessClient {
	return &TenantAccessClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTenantAccessGitClient creates a new instance of TenantAccessGitClient.
func (c *ClientFactory) NewTenantAccessGitClient() *TenantAccessGitClient {
	return &TenantAccessGitClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTenantConfigurationClient creates a new instance of TenantConfigurationClient.
func (c *ClientFactory) NewTenantConfigurationClient() *TenantConfigurationClient {
	return &TenantConfigurationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTenantSettingsClient creates a new instance of TenantSettingsClient.
func (c *ClientFactory) NewTenantSettingsClient() *TenantSettingsClient {
	return &TenantSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserClient creates a new instance of UserClient.
func (c *ClientFactory) NewUserClient() *UserClient {
	return &UserClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserConfirmationPasswordClient creates a new instance of UserConfirmationPasswordClient.
func (c *ClientFactory) NewUserConfirmationPasswordClient() *UserConfirmationPasswordClient {
	return &UserConfirmationPasswordClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserGroupClient creates a new instance of UserGroupClient.
func (c *ClientFactory) NewUserGroupClient() *UserGroupClient {
	return &UserGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserIdentitiesClient creates a new instance of UserIdentitiesClient.
func (c *ClientFactory) NewUserIdentitiesClient() *UserIdentitiesClient {
	return &UserIdentitiesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserSubscriptionClient creates a new instance of UserSubscriptionClient.
func (c *ClientFactory) NewUserSubscriptionClient() *UserSubscriptionClient {
	return &UserSubscriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
