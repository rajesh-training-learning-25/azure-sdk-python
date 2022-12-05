//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AccountProperties.
func (a AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountStatus", a.AccountStatus)
	populate(objectMap, "accountType", a.AccountType)
	populate(objectMap, "agreementType", a.AgreementType)
	populate(objectMap, "billingProfiles", a.BillingProfiles)
	populate(objectMap, "departments", a.Departments)
	populate(objectMap, "displayName", a.DisplayName)
	populate(objectMap, "enrollmentAccounts", a.EnrollmentAccounts)
	populate(objectMap, "enrollmentDetails", a.EnrollmentDetails)
	populate(objectMap, "hasReadAccess", a.HasReadAccess)
	populate(objectMap, "notificationEmailAddress", a.NotificationEmailAddress)
	populate(objectMap, "soldTo", a.SoldTo)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountUpdateRequest.
func (a AccountUpdateRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AgreementProperties.
func (a AgreementProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "acceptanceMode", a.AcceptanceMode)
	populate(objectMap, "agreementLink", a.AgreementLink)
	populate(objectMap, "billingProfileInfo", a.BillingProfileInfo)
	populate(objectMap, "category", a.Category)
	populateTimeRFC3339(objectMap, "effectiveDate", a.EffectiveDate)
	populateTimeRFC3339(objectMap, "expirationDate", a.ExpirationDate)
	populate(objectMap, "participants", a.Participants)
	populate(objectMap, "status", a.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AgreementProperties.
func (a *AgreementProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "acceptanceMode":
			err = unpopulate(val, "AcceptanceMode", &a.AcceptanceMode)
			delete(rawMsg, key)
		case "agreementLink":
			err = unpopulate(val, "AgreementLink", &a.AgreementLink)
			delete(rawMsg, key)
		case "billingProfileInfo":
			err = unpopulate(val, "BillingProfileInfo", &a.BillingProfileInfo)
			delete(rawMsg, key)
		case "category":
			err = unpopulate(val, "Category", &a.Category)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &a.EffectiveDate)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &a.ExpirationDate)
			delete(rawMsg, key)
		case "participants":
			err = unpopulate(val, "Participants", &a.Participants)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &a.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CustomerProperties.
func (c CustomerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billingProfileDisplayName", c.BillingProfileDisplayName)
	populate(objectMap, "billingProfileId", c.BillingProfileID)
	populate(objectMap, "displayName", c.DisplayName)
	populate(objectMap, "enabledAzurePlans", c.EnabledAzurePlans)
	populate(objectMap, "resellers", c.Resellers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DepartmentProperties.
func (d DepartmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "costCenter", d.CostCenter)
	populate(objectMap, "departmentName", d.DepartmentName)
	populate(objectMap, "enrollmentAccounts", d.EnrollmentAccounts)
	populate(objectMap, "status", d.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DownloadURL.
func (d *DownloadURL) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expiryTime":
			err = unpopulateTimeRFC3339(val, "ExpiryTime", &d.ExpiryTime)
			delete(rawMsg, key)
		case "url":
			err = unpopulate(val, "URL", &d.URL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Enrollment.
func (e Enrollment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billingCycle", e.BillingCycle)
	populate(objectMap, "channel", e.Channel)
	populate(objectMap, "countryCode", e.CountryCode)
	populate(objectMap, "currency", e.Currency)
	populateTimeRFC3339(objectMap, "endDate", e.EndDate)
	populate(objectMap, "language", e.Language)
	populate(objectMap, "policies", e.Policies)
	populateTimeRFC3339(objectMap, "startDate", e.StartDate)
	populate(objectMap, "status", e.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Enrollment.
func (e *Enrollment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "billingCycle":
			err = unpopulate(val, "BillingCycle", &e.BillingCycle)
			delete(rawMsg, key)
		case "channel":
			err = unpopulate(val, "Channel", &e.Channel)
			delete(rawMsg, key)
		case "countryCode":
			err = unpopulate(val, "CountryCode", &e.CountryCode)
			delete(rawMsg, key)
		case "currency":
			err = unpopulate(val, "Currency", &e.Currency)
			delete(rawMsg, key)
		case "endDate":
			err = unpopulateTimeRFC3339(val, "EndDate", &e.EndDate)
			delete(rawMsg, key)
		case "language":
			err = unpopulate(val, "Language", &e.Language)
			delete(rawMsg, key)
		case "policies":
			err = unpopulate(val, "Policies", &e.Policies)
			delete(rawMsg, key)
		case "startDate":
			err = unpopulateTimeRFC3339(val, "StartDate", &e.StartDate)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &e.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnrollmentAccountContext.
func (e *EnrollmentAccountContext) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "costCenter":
			err = unpopulate(val, "CostCenter", &e.CostCenter)
			delete(rawMsg, key)
		case "endDate":
			err = unpopulateTimeRFC3339(val, "EndDate", &e.EndDate)
			delete(rawMsg, key)
		case "enrollmentAccountName":
			err = unpopulate(val, "EnrollmentAccountName", &e.EnrollmentAccountName)
			delete(rawMsg, key)
		case "startDate":
			err = unpopulateTimeRFC3339(val, "StartDate", &e.StartDate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EnrollmentAccountProperties.
func (e EnrollmentAccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountName", e.AccountName)
	populate(objectMap, "accountOwner", e.AccountOwner)
	populate(objectMap, "accountOwnerEmail", e.AccountOwnerEmail)
	populate(objectMap, "costCenter", e.CostCenter)
	populate(objectMap, "department", e.Department)
	populateTimeRFC3339(objectMap, "endDate", e.EndDate)
	populateTimeRFC3339(objectMap, "startDate", e.StartDate)
	populate(objectMap, "status", e.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnrollmentAccountProperties.
func (e *EnrollmentAccountProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accountName":
			err = unpopulate(val, "AccountName", &e.AccountName)
			delete(rawMsg, key)
		case "accountOwner":
			err = unpopulate(val, "AccountOwner", &e.AccountOwner)
			delete(rawMsg, key)
		case "accountOwnerEmail":
			err = unpopulate(val, "AccountOwnerEmail", &e.AccountOwnerEmail)
			delete(rawMsg, key)
		case "costCenter":
			err = unpopulate(val, "CostCenter", &e.CostCenter)
			delete(rawMsg, key)
		case "department":
			err = unpopulate(val, "Department", &e.Department)
			delete(rawMsg, key)
		case "endDate":
			err = unpopulateTimeRFC3339(val, "EndDate", &e.EndDate)
			delete(rawMsg, key)
		case "startDate":
			err = unpopulateTimeRFC3339(val, "StartDate", &e.StartDate)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &e.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InstructionProperties.
func (i InstructionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "amount", i.Amount)
	populateTimeRFC3339(objectMap, "creationDate", i.CreationDate)
	populateTimeRFC3339(objectMap, "endDate", i.EndDate)
	populateTimeRFC3339(objectMap, "startDate", i.StartDate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InstructionProperties.
func (i *InstructionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "amount":
			err = unpopulate(val, "Amount", &i.Amount)
			delete(rawMsg, key)
		case "creationDate":
			err = unpopulateTimeRFC3339(val, "CreationDate", &i.CreationDate)
			delete(rawMsg, key)
		case "endDate":
			err = unpopulateTimeRFC3339(val, "EndDate", &i.EndDate)
			delete(rawMsg, key)
		case "startDate":
			err = unpopulateTimeRFC3339(val, "StartDate", &i.StartDate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InvoiceProperties.
func (i InvoiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "amountDue", i.AmountDue)
	populate(objectMap, "azurePrepaymentApplied", i.AzurePrepaymentApplied)
	populate(objectMap, "billedAmount", i.BilledAmount)
	populate(objectMap, "billedDocumentId", i.BilledDocumentID)
	populate(objectMap, "billingProfileDisplayName", i.BillingProfileDisplayName)
	populate(objectMap, "billingProfileId", i.BillingProfileID)
	populate(objectMap, "creditAmount", i.CreditAmount)
	populate(objectMap, "creditForDocumentId", i.CreditForDocumentID)
	populate(objectMap, "documentType", i.DocumentType)
	populate(objectMap, "documents", i.Documents)
	populateTimeRFC3339(objectMap, "dueDate", i.DueDate)
	populate(objectMap, "freeAzureCreditApplied", i.FreeAzureCreditApplied)
	populateTimeRFC3339(objectMap, "invoiceDate", i.InvoiceDate)
	populateTimeRFC3339(objectMap, "invoicePeriodEndDate", i.InvoicePeriodEndDate)
	populateTimeRFC3339(objectMap, "invoicePeriodStartDate", i.InvoicePeriodStartDate)
	populate(objectMap, "invoiceType", i.InvoiceType)
	populate(objectMap, "isMonthlyInvoice", i.IsMonthlyInvoice)
	populate(objectMap, "payments", i.Payments)
	populate(objectMap, "purchaseOrderNumber", i.PurchaseOrderNumber)
	populate(objectMap, "rebillDetails", i.RebillDetails)
	populate(objectMap, "status", i.Status)
	populate(objectMap, "subTotal", i.SubTotal)
	populate(objectMap, "subscriptionId", i.SubscriptionID)
	populate(objectMap, "taxAmount", i.TaxAmount)
	populate(objectMap, "totalAmount", i.TotalAmount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InvoiceProperties.
func (i *InvoiceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "amountDue":
			err = unpopulate(val, "AmountDue", &i.AmountDue)
			delete(rawMsg, key)
		case "azurePrepaymentApplied":
			err = unpopulate(val, "AzurePrepaymentApplied", &i.AzurePrepaymentApplied)
			delete(rawMsg, key)
		case "billedAmount":
			err = unpopulate(val, "BilledAmount", &i.BilledAmount)
			delete(rawMsg, key)
		case "billedDocumentId":
			err = unpopulate(val, "BilledDocumentID", &i.BilledDocumentID)
			delete(rawMsg, key)
		case "billingProfileDisplayName":
			err = unpopulate(val, "BillingProfileDisplayName", &i.BillingProfileDisplayName)
			delete(rawMsg, key)
		case "billingProfileId":
			err = unpopulate(val, "BillingProfileID", &i.BillingProfileID)
			delete(rawMsg, key)
		case "creditAmount":
			err = unpopulate(val, "CreditAmount", &i.CreditAmount)
			delete(rawMsg, key)
		case "creditForDocumentId":
			err = unpopulate(val, "CreditForDocumentID", &i.CreditForDocumentID)
			delete(rawMsg, key)
		case "documentType":
			err = unpopulate(val, "DocumentType", &i.DocumentType)
			delete(rawMsg, key)
		case "documents":
			err = unpopulate(val, "Documents", &i.Documents)
			delete(rawMsg, key)
		case "dueDate":
			err = unpopulateTimeRFC3339(val, "DueDate", &i.DueDate)
			delete(rawMsg, key)
		case "freeAzureCreditApplied":
			err = unpopulate(val, "FreeAzureCreditApplied", &i.FreeAzureCreditApplied)
			delete(rawMsg, key)
		case "invoiceDate":
			err = unpopulateTimeRFC3339(val, "InvoiceDate", &i.InvoiceDate)
			delete(rawMsg, key)
		case "invoicePeriodEndDate":
			err = unpopulateTimeRFC3339(val, "InvoicePeriodEndDate", &i.InvoicePeriodEndDate)
			delete(rawMsg, key)
		case "invoicePeriodStartDate":
			err = unpopulateTimeRFC3339(val, "InvoicePeriodStartDate", &i.InvoicePeriodStartDate)
			delete(rawMsg, key)
		case "invoiceType":
			err = unpopulate(val, "InvoiceType", &i.InvoiceType)
			delete(rawMsg, key)
		case "isMonthlyInvoice":
			err = unpopulate(val, "IsMonthlyInvoice", &i.IsMonthlyInvoice)
			delete(rawMsg, key)
		case "payments":
			err = unpopulate(val, "Payments", &i.Payments)
			delete(rawMsg, key)
		case "purchaseOrderNumber":
			err = unpopulate(val, "PurchaseOrderNumber", &i.PurchaseOrderNumber)
			delete(rawMsg, key)
		case "rebillDetails":
			err = unpopulate(val, "RebillDetails", &i.RebillDetails)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &i.Status)
			delete(rawMsg, key)
		case "subTotal":
			err = unpopulate(val, "SubTotal", &i.SubTotal)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &i.SubscriptionID)
			delete(rawMsg, key)
		case "taxAmount":
			err = unpopulate(val, "TaxAmount", &i.TaxAmount)
			delete(rawMsg, key)
		case "totalAmount":
			err = unpopulate(val, "TotalAmount", &i.TotalAmount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InvoiceSectionProperties.
func (i InvoiceSectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", i.DisplayName)
	populate(objectMap, "labels", i.Labels)
	populate(objectMap, "state", i.State)
	populate(objectMap, "systemId", i.SystemID)
	populate(objectMap, "tags", i.Tags)
	populate(objectMap, "targetCloud", i.TargetCloud)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InvoiceSectionsOnExpand.
func (i InvoiceSectionsOnExpand) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hasMoreResults", i.HasMoreResults)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Participants.
func (p Participants) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "email", p.Email)
	populate(objectMap, "status", p.Status)
	populateTimeRFC3339(objectMap, "statusDate", p.StatusDate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Participants.
func (p *Participants) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "email":
			err = unpopulate(val, "Email", &p.Email)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		case "statusDate":
			err = unpopulateTimeRFC3339(val, "StatusDate", &p.StatusDate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PaymentProperties.
func (p PaymentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "amount", p.Amount)
	populateTimeRFC3339(objectMap, "date", p.Date)
	populate(objectMap, "paymentMethodFamily", p.PaymentMethodFamily)
	populate(objectMap, "paymentMethodType", p.PaymentMethodType)
	populate(objectMap, "paymentType", p.PaymentType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PaymentProperties.
func (p *PaymentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "amount":
			err = unpopulate(val, "Amount", &p.Amount)
			delete(rawMsg, key)
		case "date":
			err = unpopulateTimeRFC3339(val, "Date", &p.Date)
			delete(rawMsg, key)
		case "paymentMethodFamily":
			err = unpopulate(val, "PaymentMethodFamily", &p.PaymentMethodFamily)
			delete(rawMsg, key)
		case "paymentMethodType":
			err = unpopulate(val, "PaymentMethodType", &p.PaymentMethodType)
			delete(rawMsg, key)
		case "paymentType":
			err = unpopulate(val, "PaymentType", &p.PaymentType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PeriodProperties.
func (p PeriodProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateDateType(objectMap, "billingPeriodEndDate", p.BillingPeriodEndDate)
	populateDateType(objectMap, "billingPeriodStartDate", p.BillingPeriodStartDate)
	populate(objectMap, "invoiceIds", p.InvoiceIDs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PeriodProperties.
func (p *PeriodProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "billingPeriodEndDate":
			err = unpopulateDateType(val, "BillingPeriodEndDate", &p.BillingPeriodEndDate)
			delete(rawMsg, key)
		case "billingPeriodStartDate":
			err = unpopulateDateType(val, "BillingPeriodStartDate", &p.BillingPeriodStartDate)
			delete(rawMsg, key)
		case "invoiceIds":
			err = unpopulate(val, "InvoiceIDs", &p.InvoiceIDs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PermissionsProperties.
func (p PermissionsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", p.Actions)
	populate(objectMap, "notActions", p.NotActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Product.
func (p Product) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProductProperties.
func (p ProductProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoRenew", p.AutoRenew)
	populate(objectMap, "availabilityId", p.AvailabilityID)
	populate(objectMap, "billingFrequency", p.BillingFrequency)
	populate(objectMap, "billingProfileDisplayName", p.BillingProfileDisplayName)
	populate(objectMap, "billingProfileId", p.BillingProfileID)
	populate(objectMap, "customerDisplayName", p.CustomerDisplayName)
	populate(objectMap, "customerId", p.CustomerID)
	populate(objectMap, "displayName", p.DisplayName)
	populateTimeRFC3339(objectMap, "endDate", p.EndDate)
	populate(objectMap, "invoiceSectionDisplayName", p.InvoiceSectionDisplayName)
	populate(objectMap, "invoiceSectionId", p.InvoiceSectionID)
	populate(objectMap, "lastCharge", p.LastCharge)
	populateTimeRFC3339(objectMap, "lastChargeDate", p.LastChargeDate)
	populate(objectMap, "productType", p.ProductType)
	populate(objectMap, "productTypeId", p.ProductTypeID)
	populateTimeRFC3339(objectMap, "purchaseDate", p.PurchaseDate)
	populate(objectMap, "quantity", p.Quantity)
	populate(objectMap, "reseller", p.Reseller)
	populate(objectMap, "skuDescription", p.SKUDescription)
	populate(objectMap, "skuId", p.SKUID)
	populate(objectMap, "status", p.Status)
	populate(objectMap, "tenantId", p.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProductProperties.
func (p *ProductProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "autoRenew":
			err = unpopulate(val, "AutoRenew", &p.AutoRenew)
			delete(rawMsg, key)
		case "availabilityId":
			err = unpopulate(val, "AvailabilityID", &p.AvailabilityID)
			delete(rawMsg, key)
		case "billingFrequency":
			err = unpopulate(val, "BillingFrequency", &p.BillingFrequency)
			delete(rawMsg, key)
		case "billingProfileDisplayName":
			err = unpopulate(val, "BillingProfileDisplayName", &p.BillingProfileDisplayName)
			delete(rawMsg, key)
		case "billingProfileId":
			err = unpopulate(val, "BillingProfileID", &p.BillingProfileID)
			delete(rawMsg, key)
		case "customerDisplayName":
			err = unpopulate(val, "CustomerDisplayName", &p.CustomerDisplayName)
			delete(rawMsg, key)
		case "customerId":
			err = unpopulate(val, "CustomerID", &p.CustomerID)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "endDate":
			err = unpopulateTimeRFC3339(val, "EndDate", &p.EndDate)
			delete(rawMsg, key)
		case "invoiceSectionDisplayName":
			err = unpopulate(val, "InvoiceSectionDisplayName", &p.InvoiceSectionDisplayName)
			delete(rawMsg, key)
		case "invoiceSectionId":
			err = unpopulate(val, "InvoiceSectionID", &p.InvoiceSectionID)
			delete(rawMsg, key)
		case "lastCharge":
			err = unpopulate(val, "LastCharge", &p.LastCharge)
			delete(rawMsg, key)
		case "lastChargeDate":
			err = unpopulateTimeRFC3339(val, "LastChargeDate", &p.LastChargeDate)
			delete(rawMsg, key)
		case "productType":
			err = unpopulate(val, "ProductType", &p.ProductType)
			delete(rawMsg, key)
		case "productTypeId":
			err = unpopulate(val, "ProductTypeID", &p.ProductTypeID)
			delete(rawMsg, key)
		case "purchaseDate":
			err = unpopulateTimeRFC3339(val, "PurchaseDate", &p.PurchaseDate)
			delete(rawMsg, key)
		case "quantity":
			err = unpopulate(val, "Quantity", &p.Quantity)
			delete(rawMsg, key)
		case "reseller":
			err = unpopulate(val, "Reseller", &p.Reseller)
			delete(rawMsg, key)
		case "skuDescription":
			err = unpopulate(val, "SKUDescription", &p.SKUDescription)
			delete(rawMsg, key)
		case "skuId":
			err = unpopulate(val, "SKUID", &p.SKUID)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &p.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProfileProperties.
func (p ProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billTo", p.BillTo)
	populate(objectMap, "billingRelationshipType", p.BillingRelationshipType)
	populate(objectMap, "currency", p.Currency)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "enabledAzurePlans", p.EnabledAzurePlans)
	populate(objectMap, "hasReadAccess", p.HasReadAccess)
	populate(objectMap, "indirectRelationshipInfo", p.IndirectRelationshipInfo)
	populate(objectMap, "invoiceDay", p.InvoiceDay)
	populate(objectMap, "invoiceEmailOptIn", p.InvoiceEmailOptIn)
	populate(objectMap, "invoiceSections", p.InvoiceSections)
	populate(objectMap, "poNumber", p.PoNumber)
	populate(objectMap, "spendingLimit", p.SpendingLimit)
	populate(objectMap, "status", p.Status)
	populate(objectMap, "statusReasonCode", p.StatusReasonCode)
	populate(objectMap, "systemId", p.SystemID)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "targetClouds", p.TargetClouds)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProfilesOnExpand.
func (p ProfilesOnExpand) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hasMoreResults", p.HasMoreResults)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Property.
func (p Property) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RebillDetails.
func (r RebillDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "creditNoteDocumentId", r.CreditNoteDocumentID)
	populate(objectMap, "invoiceDocumentId", r.InvoiceDocumentID)
	populate(objectMap, "rebillDetails", r.RebillDetails)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinitionProperties.
func (r RoleDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", r.Description)
	populate(objectMap, "permissions", r.Permissions)
	populate(objectMap, "roleName", r.RoleName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Subscription.
func (s Subscription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionProperties.
func (s SubscriptionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billingProfileDisplayName", s.BillingProfileDisplayName)
	populate(objectMap, "billingProfileId", s.BillingProfileID)
	populate(objectMap, "costCenter", s.CostCenter)
	populate(objectMap, "customerDisplayName", s.CustomerDisplayName)
	populate(objectMap, "customerId", s.CustomerID)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "invoiceSectionDisplayName", s.InvoiceSectionDisplayName)
	populate(objectMap, "invoiceSectionId", s.InvoiceSectionID)
	populate(objectMap, "lastMonthCharges", s.LastMonthCharges)
	populate(objectMap, "monthToDateCharges", s.MonthToDateCharges)
	populate(objectMap, "reseller", s.Reseller)
	populate(objectMap, "skuDescription", s.SKUDescription)
	populate(objectMap, "skuId", s.SKUID)
	populate(objectMap, "subscriptionBillingStatus", s.SubscriptionBillingStatus)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populate(objectMap, "suspensionReasons", s.SuspensionReasons)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TransactionProperties.
func (t TransactionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "azureCreditApplied", t.AzureCreditApplied)
	populate(objectMap, "azurePlan", t.AzurePlan)
	populate(objectMap, "billingCurrency", t.BillingCurrency)
	populate(objectMap, "billingProfileDisplayName", t.BillingProfileDisplayName)
	populate(objectMap, "billingProfileId", t.BillingProfileID)
	populate(objectMap, "customerDisplayName", t.CustomerDisplayName)
	populate(objectMap, "customerId", t.CustomerID)
	populateTimeRFC3339(objectMap, "date", t.Date)
	populate(objectMap, "discount", t.Discount)
	populate(objectMap, "effectivePrice", t.EffectivePrice)
	populate(objectMap, "exchangeRate", t.ExchangeRate)
	populate(objectMap, "invoice", t.Invoice)
	populate(objectMap, "invoiceId", t.InvoiceID)
	populate(objectMap, "invoiceSectionDisplayName", t.InvoiceSectionDisplayName)
	populate(objectMap, "invoiceSectionId", t.InvoiceSectionID)
	populate(objectMap, "kind", t.Kind)
	populate(objectMap, "marketPrice", t.MarketPrice)
	populate(objectMap, "orderId", t.OrderID)
	populate(objectMap, "orderName", t.OrderName)
	populate(objectMap, "pricingCurrency", t.PricingCurrency)
	populate(objectMap, "productDescription", t.ProductDescription)
	populate(objectMap, "productFamily", t.ProductFamily)
	populate(objectMap, "productType", t.ProductType)
	populate(objectMap, "productTypeId", t.ProductTypeID)
	populate(objectMap, "quantity", t.Quantity)
	populateTimeRFC3339(objectMap, "servicePeriodEndDate", t.ServicePeriodEndDate)
	populateTimeRFC3339(objectMap, "servicePeriodStartDate", t.ServicePeriodStartDate)
	populate(objectMap, "subTotal", t.SubTotal)
	populate(objectMap, "subscriptionId", t.SubscriptionID)
	populate(objectMap, "subscriptionName", t.SubscriptionName)
	populate(objectMap, "tax", t.Tax)
	populate(objectMap, "transactionAmount", t.TransactionAmount)
	populate(objectMap, "transactionType", t.TransactionType)
	populate(objectMap, "unitOfMeasure", t.UnitOfMeasure)
	populate(objectMap, "unitType", t.UnitType)
	populate(objectMap, "units", t.Units)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TransactionProperties.
func (t *TransactionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureCreditApplied":
			err = unpopulate(val, "AzureCreditApplied", &t.AzureCreditApplied)
			delete(rawMsg, key)
		case "azurePlan":
			err = unpopulate(val, "AzurePlan", &t.AzurePlan)
			delete(rawMsg, key)
		case "billingCurrency":
			err = unpopulate(val, "BillingCurrency", &t.BillingCurrency)
			delete(rawMsg, key)
		case "billingProfileDisplayName":
			err = unpopulate(val, "BillingProfileDisplayName", &t.BillingProfileDisplayName)
			delete(rawMsg, key)
		case "billingProfileId":
			err = unpopulate(val, "BillingProfileID", &t.BillingProfileID)
			delete(rawMsg, key)
		case "customerDisplayName":
			err = unpopulate(val, "CustomerDisplayName", &t.CustomerDisplayName)
			delete(rawMsg, key)
		case "customerId":
			err = unpopulate(val, "CustomerID", &t.CustomerID)
			delete(rawMsg, key)
		case "date":
			err = unpopulateTimeRFC3339(val, "Date", &t.Date)
			delete(rawMsg, key)
		case "discount":
			err = unpopulate(val, "Discount", &t.Discount)
			delete(rawMsg, key)
		case "effectivePrice":
			err = unpopulate(val, "EffectivePrice", &t.EffectivePrice)
			delete(rawMsg, key)
		case "exchangeRate":
			err = unpopulate(val, "ExchangeRate", &t.ExchangeRate)
			delete(rawMsg, key)
		case "invoice":
			err = unpopulate(val, "Invoice", &t.Invoice)
			delete(rawMsg, key)
		case "invoiceId":
			err = unpopulate(val, "InvoiceID", &t.InvoiceID)
			delete(rawMsg, key)
		case "invoiceSectionDisplayName":
			err = unpopulate(val, "InvoiceSectionDisplayName", &t.InvoiceSectionDisplayName)
			delete(rawMsg, key)
		case "invoiceSectionId":
			err = unpopulate(val, "InvoiceSectionID", &t.InvoiceSectionID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &t.Kind)
			delete(rawMsg, key)
		case "marketPrice":
			err = unpopulate(val, "MarketPrice", &t.MarketPrice)
			delete(rawMsg, key)
		case "orderId":
			err = unpopulate(val, "OrderID", &t.OrderID)
			delete(rawMsg, key)
		case "orderName":
			err = unpopulate(val, "OrderName", &t.OrderName)
			delete(rawMsg, key)
		case "pricingCurrency":
			err = unpopulate(val, "PricingCurrency", &t.PricingCurrency)
			delete(rawMsg, key)
		case "productDescription":
			err = unpopulate(val, "ProductDescription", &t.ProductDescription)
			delete(rawMsg, key)
		case "productFamily":
			err = unpopulate(val, "ProductFamily", &t.ProductFamily)
			delete(rawMsg, key)
		case "productType":
			err = unpopulate(val, "ProductType", &t.ProductType)
			delete(rawMsg, key)
		case "productTypeId":
			err = unpopulate(val, "ProductTypeID", &t.ProductTypeID)
			delete(rawMsg, key)
		case "quantity":
			err = unpopulate(val, "Quantity", &t.Quantity)
			delete(rawMsg, key)
		case "servicePeriodEndDate":
			err = unpopulateTimeRFC3339(val, "ServicePeriodEndDate", &t.ServicePeriodEndDate)
			delete(rawMsg, key)
		case "servicePeriodStartDate":
			err = unpopulateTimeRFC3339(val, "ServicePeriodStartDate", &t.ServicePeriodStartDate)
			delete(rawMsg, key)
		case "subTotal":
			err = unpopulate(val, "SubTotal", &t.SubTotal)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &t.SubscriptionID)
			delete(rawMsg, key)
		case "subscriptionName":
			err = unpopulate(val, "SubscriptionName", &t.SubscriptionName)
			delete(rawMsg, key)
		case "tax":
			err = unpopulate(val, "Tax", &t.Tax)
			delete(rawMsg, key)
		case "transactionAmount":
			err = unpopulate(val, "TransactionAmount", &t.TransactionAmount)
			delete(rawMsg, key)
		case "transactionType":
			err = unpopulate(val, "TransactionType", &t.TransactionType)
			delete(rawMsg, key)
		case "unitOfMeasure":
			err = unpopulate(val, "UnitOfMeasure", &t.UnitOfMeasure)
			delete(rawMsg, key)
		case "unitType":
			err = unpopulate(val, "UnitType", &t.UnitType)
			delete(rawMsg, key)
		case "units":
			err = unpopulate(val, "Units", &t.Units)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}