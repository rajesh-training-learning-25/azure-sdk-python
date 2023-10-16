//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import "time"

// ChatTranscriptDetails - Object that represents a Chat Transcript resource.
type ChatTranscriptDetails struct {
	// Properties of the resource.
	Properties *ChatTranscriptDetailsProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ChatTranscriptDetailsProperties - Describes the properties of a Chat Transcript Details resource.
type ChatTranscriptDetailsProperties struct {
	// List of chat transcript communication resources.
	Messages []*MessageProperties

	// READ-ONLY; Time in UTC (ISO 8601 format) when the chat began.
	StartTime *time.Time
}

// ChatTranscriptsListResult - Collection of Chat Transcripts resources.
type ChatTranscriptsListResult struct {
	// The URI to fetch the next page of Chat Transcripts resources.
	NextLink *string

	// List of Chat Transcripts resources.
	Value []*ChatTranscriptDetails
}

// CheckNameAvailabilityInput - Input of CheckNameAvailability API.
type CheckNameAvailabilityInput struct {
	// REQUIRED; The resource name to validate.
	Name *string

	// REQUIRED; The type of resource.
	Type *Type
}

// CheckNameAvailabilityOutput - Output of check name availability API.
type CheckNameAvailabilityOutput struct {
	// READ-ONLY; The detailed error message describing why the name is not available.
	Message *string

	// READ-ONLY; Indicates whether the name is available.
	NameAvailable *bool

	// READ-ONLY; The reason why the name is not available.
	Reason *string
}

// CommunicationDetails - Object that represents a Communication resource.
type CommunicationDetails struct {
	// Properties of the resource.
	Properties *CommunicationDetailsProperties

	// READ-ONLY; Id of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of the resource 'Microsoft.Support/communications'.
	Type *string
}

// CommunicationDetailsProperties - Describes the properties of a communication resource.
type CommunicationDetailsProperties struct {
	// REQUIRED; Body of the communication.
	Body *string

	// REQUIRED; Subject of the communication.
	Subject *string

	// Email address of the sender. This property is required if called by a service principal.
	Sender *string

	// READ-ONLY; Direction of communication.
	CommunicationDirection *CommunicationDirection

	// READ-ONLY; Communication type.
	CommunicationType *CommunicationType

	// READ-ONLY; Time in UTC (ISO 8601 format) when the communication was created.
	CreatedDate *time.Time
}

// CommunicationsListResult - Collection of Communication resources.
type CommunicationsListResult struct {
	// The URI to fetch the next page of Communication resources.
	NextLink *string

	// List of Communication resources.
	Value []*CommunicationDetails
}

// ContactProfile - Contact information associated with the support ticket.
type ContactProfile struct {
	// REQUIRED; Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string

	// REQUIRED; First name.
	FirstName *string

	// REQUIRED; Last name.
	LastName *string

	// REQUIRED; Preferred contact method.
	PreferredContactMethod *PreferredContactMethod

	// REQUIRED; Preferred language of support from Azure. Support languages vary based on the severity you choose for your support
	// ticket. Learn more at Azure Severity and responsiveness
	// [https://azure.microsoft.com/support/plans/response]. Use the standard language-country code. Valid values are 'en-us'
	// for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French,
	// 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for
	// Chinese and 'de-de' for German.
	PreferredSupportLanguage *string

	// REQUIRED; Time zone of the user. This is the name of the time zone from Microsoft Time Zone Index Values [https://support.microsoft.com/help/973627/microsoft-time-zone-index-values].
	PreferredTimeZone *string

	// REQUIRED; Primary email address.
	PrimaryEmailAddress *string

	// Additional email addresses listed will be copied on any correspondence about the support ticket.
	AdditionalEmailAddresses []*string

	// Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string
}

// Engineer - Support engineer information.
type Engineer struct {
	// READ-ONLY; Email address of the Azure Support engineer assigned to the support ticket.
	EmailAddress *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// FileDetails - Object that represents File Details resource
type FileDetails struct {
	// Properties of the resource
	Properties *FileDetailsProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FileDetailsProperties - Describes the properties of a file.
type FileDetailsProperties struct {
	// Size of each chunk
	ChunkSize *float32

	// Size of the file to be uploaded
	FileSize *float32

	// Number of chunks to be uploaded
	NumberOfChunks *float32

	// READ-ONLY; Time in UTC (ISO 8601 format) when file workspace was created.
	CreatedOn *time.Time
}

// FileWorkspaceDetails - Object that represents FileWorkspaceDetails resource
type FileWorkspaceDetails struct {
	// Properties of the resource
	Properties *FileWorkspaceDetailsProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FileWorkspaceDetailsProperties - Describes the properties of a file workspace.
type FileWorkspaceDetailsProperties struct {
	// READ-ONLY; Time in UTC (ISO 8601 format) when file workspace was created.
	CreatedOn *time.Time

	// READ-ONLY; Time in UTC (ISO 8601 format) when file workspace is going to expire.
	ExpirationTime *time.Time
}

// FilesListResult - Object that represents a collection of File resources.
type FilesListResult struct {
	// The URI to fetch the next page of File resources.
	NextLink *string

	// List of File resources.
	Value []*FileDetails
}

// MessageProperties - Describes the properties of a Message Details resource.
type MessageProperties struct {
	// REQUIRED; Body of the communication.
	Body *string

	// Name of the sender.
	Sender *string

	// READ-ONLY; Direction of communication.
	CommunicationDirection *CommunicationDirection

	// READ-ONLY; Content type.
	ContentType *TranscriptContentType

	// READ-ONLY; Time in UTC (ISO 8601 format) when the communication was created.
	CreatedDate *time.Time
}

// Operation - The operation supported by Microsoft Support resource provider.
type Operation struct {
	// The object that describes the operation.
	Display *OperationDisplay

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string
}

// OperationDisplay - The object that describes the operation.
type OperationDisplay struct {
	// READ-ONLY; The description of the operation.
	Description *string

	// READ-ONLY; The action that users can perform, based on their permission level.
	Operation *string

	// READ-ONLY; Service provider: Microsoft Support.
	Provider *string

	// READ-ONLY; Resource on which the operation is performed.
	Resource *string
}

// OperationsListResult - The list of operations supported by Microsoft Support resource provider.
type OperationsListResult struct {
	// The list of operations supported by Microsoft Support resource provider.
	Value []*Operation
}

// ProblemClassification resource object.
type ProblemClassification struct {
	// Properties of the resource.
	Properties *ProblemClassificationProperties

	// READ-ONLY; Id of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of the resource 'Microsoft.Support/problemClassification'.
	Type *string
}

// ProblemClassificationProperties - Details about a problem classification available for an Azure service.
type ProblemClassificationProperties struct {
	// Localized name of problem classification.
	DisplayName *string

	// This property indicates whether secondary consent is present for problem classification
	SecondaryConsentEnabled []*SecondaryConsentEnabled
}

// ProblemClassificationsListResult - Collection of ProblemClassification resources.
type ProblemClassificationsListResult struct {
	// List of ProblemClassification resources.
	Value []*ProblemClassification
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// QuotaChangeRequest - This property is required for providing the region and new quota limits.
type QuotaChangeRequest struct {
	// Payload of the quota increase request.
	Payload *string

	// Region for which the quota increase request is being made.
	Region *string
}

// QuotaTicketDetails - Additional set of information required for quota increase support ticket for certain quota types,
// e.g.: Virtual machine cores. Get complete details about Quota payload support request along with
// examples at Support quota request [https://aka.ms/supportrpquotarequestpayload].
type QuotaTicketDetails struct {
	// Required for certain quota types when there is a sub type, such as Batch, for which you are requesting a quota increase.
	QuotaChangeRequestSubType *string

	// Quota change request version.
	QuotaChangeRequestVersion *string

	// This property is required for providing the region and new quota limits.
	QuotaChangeRequests []*QuotaChangeRequest
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SecondaryConsent - This property indicates secondary consent for the support ticket.
type SecondaryConsent struct {
	// The service name for which the secondary consent is being provided. The value needs to be retrieved from the Problem Classification
	// API response.
	Type *string

	// User consent value provided
	UserConsent *UserConsent
}

// SecondaryConsentEnabled - This property indicates whether secondary consent is present for problem classification.
type SecondaryConsentEnabled struct {
	// User consent description.
	Description *string

	// The Azure service for which secondary consent is needed for case creation.
	Type *string
}

// Service - Object that represents a Service resource.
type Service struct {
	// Properties of the resource.
	Properties *ServiceProperties

	// READ-ONLY; Id of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of the resource 'Microsoft.Support/services'.
	Type *string
}

// ServiceLevelAgreement - Service Level Agreement details for a support ticket.
type ServiceLevelAgreement struct {
	// READ-ONLY; Time in UTC (ISO 8601 format) when the service level agreement expires.
	ExpirationTime *time.Time

	// READ-ONLY; Service Level Agreement in minutes.
	SLAMinutes *int32

	// READ-ONLY; Time in UTC (ISO 8601 format) when the service level agreement starts.
	StartTime *time.Time
}

// ServiceProperties - Details about an Azure service available for support ticket creation.
type ServiceProperties struct {
	// Localized name of the Azure service.
	DisplayName *string

	// ARM Resource types.
	ResourceTypes []*string
}

// ServicesListResult - Collection of Service resources.
type ServicesListResult struct {
	// List of Service resources.
	Value []*Service
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TechnicalTicketDetails - Additional information for technical support ticket.
type TechnicalTicketDetails struct {
	// This is the resource Id of the Azure service resource (For example: A virtual machine resource or an HDInsight resource)
	// for which the support ticket is created.
	ResourceID *string
}

// TicketDetails - Object that represents SupportTicketDetails resource.
type TicketDetails struct {
	// Properties of the resource.
	Properties *TicketDetailsProperties

	// READ-ONLY; Id of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of the resource 'Microsoft.Support/supportTickets'.
	Type *string
}

// TicketDetailsProperties - Describes the properties of a support ticket.
type TicketDetailsProperties struct {
	// REQUIRED; Contact information of the user requesting to create a support ticket.
	ContactDetails *ContactProfile

	// REQUIRED; Detailed description of the question or issue.
	Description *string

	// REQUIRED; Each Azure service has its own set of issue categories, also known as problem classification. This parameter
	// is the unique Id for the type of problem you are experiencing.
	ProblemClassificationID *string

	// REQUIRED; This is the resource Id of the Azure service resource associated with the support ticket.
	ServiceID *string

	// REQUIRED; A value that indicates the urgency of the case, which in turn determines the response time according to the service
	// level agreement of the technical support plan you have with Azure. Note: 'Highest
	// critical impact', also known as the 'Emergency - Severe impact' level in the Azure portal is reserved only for our Premium
	// customers.
	Severity *SeverityLevel

	// REQUIRED; Title of the support ticket.
	Title *string

	// Advanced diagnostic consent to be updated on the support ticket.
	AdvancedDiagnosticConsent *Consent

	// File workspace name.
	FileWorkspaceName *string

	// Problem scoping questions associated with the support ticket.
	ProblemScopingQuestions *string

	// Time in UTC (ISO 8601 format) when the problem started.
	ProblemStartTime *time.Time

	// Additional ticket details associated with a quota support ticket request.
	QuotaTicketDetails *QuotaTicketDetails

	// Indicates if this requires a 24x7 response from Azure.
	Require24X7Response *bool

	// This property indicates secondary consents for the support ticket
	SecondaryConsent []*SecondaryConsent

	// Service Level Agreement information for this support ticket.
	ServiceLevelAgreement *ServiceLevelAgreement

	// Information about the support engineer working on this support ticket.
	SupportEngineer *Engineer

	// Support plan id associated with the support ticket.
	SupportPlanID *string

	// System generated support ticket Id that is unique.
	SupportTicketID *string

	// Additional ticket details associated with a technical support ticket request.
	TechnicalTicketDetails *TechnicalTicketDetails

	// READ-ONLY; Time in UTC (ISO 8601 format) when the support ticket was created.
	CreatedDate *time.Time

	// READ-ONLY; Enrollment Id associated with the support ticket.
	EnrollmentID *string

	// READ-ONLY; Time in UTC (ISO 8601 format) when the support ticket was last modified.
	ModifiedDate *time.Time

	// READ-ONLY; Localized name of problem classification.
	ProblemClassificationDisplayName *string

	// READ-ONLY; Localized name of the Azure service.
	ServiceDisplayName *string

	// READ-ONLY; Status of the support ticket.
	Status *string

	// READ-ONLY; Support plan type associated with the support ticket.
	SupportPlanDisplayName *string

	// READ-ONLY; Support plan type associated with the support ticket.
	SupportPlanType *string
}

// TicketsListResult - Object that represents a collection of SupportTicket resources.
type TicketsListResult struct {
	// The URI to fetch the next page of SupportTicket resources.
	NextLink *string

	// List of SupportTicket resources.
	Value []*TicketDetails
}

// UpdateContactProfile - Contact information associated with the support ticket.
type UpdateContactProfile struct {
	// Email addresses listed will be copied on any correspondence about the support ticket.
	AdditionalEmailAddresses []*string

	// Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string

	// First name.
	FirstName *string

	// Last name.
	LastName *string

	// Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string

	// Preferred contact method.
	PreferredContactMethod *PreferredContactMethod

	// Preferred language of support from Azure. Support languages vary based on the severity you choose for your support ticket.
	// Learn more at Azure Severity and responsiveness
	// [https://azure.microsoft.com/support/plans/response/]. Use the standard language-country code. Valid values are 'en-us'
	// for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French,
	// 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for
	// Chinese and 'de-de' for German.
	PreferredSupportLanguage *string

	// Time zone of the user. This is the name of the time zone from Microsoft Time Zone Index Values [https://support.microsoft.com/help/973627/microsoft-time-zone-index-values].
	PreferredTimeZone *string

	// Primary email address.
	PrimaryEmailAddress *string
}

// UpdateSupportTicket - Updates severity, ticket status, and contact details in the support ticket.
type UpdateSupportTicket struct {
	// Advanced diagnostic consent to be updated on the support ticket.
	AdvancedDiagnosticConsent *Consent

	// Contact details to be updated on the support ticket.
	ContactDetails *UpdateContactProfile

	// This property indicates secondary consents for the support ticket
	SecondaryConsent []*SecondaryConsent

	// Severity level.
	Severity *SeverityLevel

	// Status to be updated on the ticket.
	Status *Status
}

// UploadFile - File content associated with the file under a workspace.
type UploadFile struct {
	// Index of the uploaded chunk (Index starts at 0)
	ChunkIndex *float32

	// File Content in base64 encoded format
	Content *string
}
