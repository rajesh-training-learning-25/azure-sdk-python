//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azpki

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CertificateAttributes.
func (c CertificateAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "extendedKeyUsage", c.ExtendedKeyUsage)
	populate(objectMap, "issuer", c.Issuer)
	populate(objectMap, "keyUsage", c.KeyUsage)
	populate(objectMap, "serialNumber", c.SerialNumber)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "subjectAlternativeNames", c.SubjectAlternativeNames)
	populateTimeRFC3339(objectMap, "validityNotAfter", c.ValidityNotAfter)
	populateTimeRFC3339(objectMap, "validityNotBefore", c.ValidityNotBefore)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateAttributes.
func (c *CertificateAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "extendedKeyUsage":
				err = unpopulate(val, "ExtendedKeyUsage", &c.ExtendedKeyUsage)
				delete(rawMsg, key)
		case "issuer":
				err = unpopulate(val, "Issuer", &c.Issuer)
				delete(rawMsg, key)
		case "keyUsage":
				err = unpopulate(val, "KeyUsage", &c.KeyUsage)
				delete(rawMsg, key)
		case "serialNumber":
				err = unpopulate(val, "SerialNumber", &c.SerialNumber)
				delete(rawMsg, key)
		case "subject":
				err = unpopulate(val, "Subject", &c.Subject)
				delete(rawMsg, key)
		case "subjectAlternativeNames":
				err = unpopulate(val, "SubjectAlternativeNames", &c.SubjectAlternativeNames)
				delete(rawMsg, key)
		case "validityNotAfter":
				err = unpopulateTimeRFC3339(val, "ValidityNotAfter", &c.ValidityNotAfter)
				delete(rawMsg, key)
		case "validityNotBefore":
				err = unpopulateTimeRFC3339(val, "ValidityNotBefore", &c.ValidityNotBefore)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CertificateDescriptionResponse.
func (c CertificateDescriptionResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "CAName", c.CAName)
	populate(objectMap, "enrollmentPolicyName", c.EnrollmentPolicyName)
	populate(objectMap, "enrollmentPolicyVersion", c.EnrollmentPolicyVersion)
	populateTimeRFC3339(objectMap, "issuedAt", c.IssuedAt)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "revocationReason", c.RevocationReason)
	populateTimeRFC3339(objectMap, "revocationTime", c.RevocationTime)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateDescriptionResponse.
func (c *CertificateDescriptionResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "CAName":
				err = unpopulate(val, "CAName", &c.CAName)
				delete(rawMsg, key)
		case "enrollmentPolicyName":
				err = unpopulate(val, "EnrollmentPolicyName", &c.EnrollmentPolicyName)
				delete(rawMsg, key)
		case "enrollmentPolicyVersion":
				err = unpopulate(val, "EnrollmentPolicyVersion", &c.EnrollmentPolicyVersion)
				delete(rawMsg, key)
		case "issuedAt":
				err = unpopulateTimeRFC3339(val, "IssuedAt", &c.IssuedAt)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &c.Name)
				delete(rawMsg, key)
		case "revocationReason":
				err = unpopulate(val, "RevocationReason", &c.RevocationReason)
				delete(rawMsg, key)
		case "revocationTime":
				err = unpopulateTimeRFC3339(val, "RevocationTime", &c.RevocationTime)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &c.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CertificateValue.
func (c CertificateValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "extendedKeyUsage", c.ExtendedKeyUsage)
	populate(objectMap, "issuer", c.Issuer)
	populate(objectMap, "keyUsage", c.KeyUsage)
	populate(objectMap, "PEM", c.PEM)
	populate(objectMap, "PEMChain", c.PEMChain)
	populate(objectMap, "PKCS7", c.PKCS7)
	populate(objectMap, "serialNumber", c.SerialNumber)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "subjectAlternativeNames", c.SubjectAlternativeNames)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populateTimeRFC3339(objectMap, "validityNotAfter", c.ValidityNotAfter)
	populateTimeRFC3339(objectMap, "validityNotBefore", c.ValidityNotBefore)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateValue.
func (c *CertificateValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "extendedKeyUsage":
				err = unpopulate(val, "ExtendedKeyUsage", &c.ExtendedKeyUsage)
				delete(rawMsg, key)
		case "issuer":
				err = unpopulate(val, "Issuer", &c.Issuer)
				delete(rawMsg, key)
		case "keyUsage":
				err = unpopulate(val, "KeyUsage", &c.KeyUsage)
				delete(rawMsg, key)
		case "PEM":
				err = unpopulate(val, "PEM", &c.PEM)
				delete(rawMsg, key)
		case "PEMChain":
				err = unpopulate(val, "PEMChain", &c.PEMChain)
				delete(rawMsg, key)
		case "PKCS7":
				err = unpopulate(val, "PKCS7", &c.PKCS7)
				delete(rawMsg, key)
		case "serialNumber":
				err = unpopulate(val, "SerialNumber", &c.SerialNumber)
				delete(rawMsg, key)
		case "subject":
				err = unpopulate(val, "Subject", &c.Subject)
				delete(rawMsg, key)
		case "subjectAlternativeNames":
				err = unpopulate(val, "SubjectAlternativeNames", &c.SubjectAlternativeNames)
				delete(rawMsg, key)
		case "thumbprint":
				err = unpopulate(val, "Thumbprint", &c.Thumbprint)
				delete(rawMsg, key)
		case "validityNotAfter":
				err = unpopulateTimeRFC3339(val, "ValidityNotAfter", &c.ValidityNotAfter)
				delete(rawMsg, key)
		case "validityNotBefore":
				err = unpopulateTimeRFC3339(val, "ValidityNotBefore", &c.ValidityNotBefore)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EnrollRequest.
func (e EnrollRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "CSR", e.CSR)
	populate(objectMap, "substitutes", e.Substitutes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnrollRequest.
func (e *EnrollRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "CSR":
				err = unpopulate(val, "CSR", &e.CSR)
				delete(rawMsg, key)
		case "substitutes":
				err = unpopulate(val, "Substitutes", &e.Substitutes)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EnrollResponse.
func (e EnrollResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "CAName", e.CAName)
	populate(objectMap, "enrollmentPolicyName", e.EnrollmentPolicyName)
	populate(objectMap, "enrollmentPolicyVersion", e.EnrollmentPolicyVersion)
	populateTimeRFC3339(objectMap, "issuedAt", e.IssuedAt)
	populate(objectMap, "log", e.Log)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "revocationReason", e.RevocationReason)
	populateTimeRFC3339(objectMap, "revocationTime", e.RevocationTime)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnrollResponse.
func (e *EnrollResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "CAName":
				err = unpopulate(val, "CAName", &e.CAName)
				delete(rawMsg, key)
		case "enrollmentPolicyName":
				err = unpopulate(val, "EnrollmentPolicyName", &e.EnrollmentPolicyName)
				delete(rawMsg, key)
		case "enrollmentPolicyVersion":
				err = unpopulate(val, "EnrollmentPolicyVersion", &e.EnrollmentPolicyVersion)
				delete(rawMsg, key)
		case "issuedAt":
				err = unpopulateTimeRFC3339(val, "IssuedAt", &e.IssuedAt)
				delete(rawMsg, key)
		case "log":
				err = unpopulate(val, "Log", &e.Log)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &e.Name)
				delete(rawMsg, key)
		case "revocationReason":
				err = unpopulate(val, "RevocationReason", &e.RevocationReason)
				delete(rawMsg, key)
		case "revocationTime":
				err = unpopulateTimeRFC3339(val, "RevocationTime", &e.RevocationTime)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &e.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}



// MarshalJSON implements the json.Marshaller interface for type LogEntry.
func (l LogEntry) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", l.Description)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogEntry.
func (l *LogEntry) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
				err = unpopulate(val, "Description", &l.Description)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &l.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RevokeRequest.
func (r RevokeRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "revocationReason", r.RevocationReason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RevokeRequest.
func (r *RevokeRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "revocationReason":
				err = unpopulate(val, "RevocationReason", &r.RevocationReason)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RevokeResponse.
func (r RevokeResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "CAName", r.CAName)
	populate(objectMap, "log", r.Log)
	populate(objectMap, "name", r.Name)
	populateTimeRFC3339(objectMap, "revokedAt", r.RevokedAt)
	populate(objectMap, "serialNumber", r.SerialNumber)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RevokeResponse.
func (r *RevokeResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "CAName":
				err = unpopulate(val, "CAName", &r.CAName)
				delete(rawMsg, key)
		case "log":
				err = unpopulate(val, "Log", &r.Log)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &r.Name)
				delete(rawMsg, key)
		case "revokedAt":
				err = unpopulateTimeRFC3339(val, "RevokedAt", &r.RevokedAt)
				delete(rawMsg, key)
		case "serialNumber":
				err = unpopulate(val, "SerialNumber", &r.SerialNumber)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Substitutes.
func (s Substitutes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "distinguishedName", s.DistinguishedName)
	populate(objectMap, "extendedKeyUsage", s.ExtendedKeyUsage)
	populate(objectMap, "keyUsage", s.KeyUsage)
	populate(objectMap, "subjectAlternativeNames", s.SubjectAlternativeNames)
	populate(objectMap, "subjectName", s.SubjectName)
	populateTimeRFC3339(objectMap, "validityNotAfter", s.ValidityNotAfter)
	populateTimeRFC3339(objectMap, "validityNotBefore", s.ValidityNotBefore)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Substitutes.
func (s *Substitutes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "distinguishedName":
				err = unpopulate(val, "DistinguishedName", &s.DistinguishedName)
				delete(rawMsg, key)
		case "extendedKeyUsage":
				err = unpopulate(val, "ExtendedKeyUsage", &s.ExtendedKeyUsage)
				delete(rawMsg, key)
		case "keyUsage":
				err = unpopulate(val, "KeyUsage", &s.KeyUsage)
				delete(rawMsg, key)
		case "subjectAlternativeNames":
				err = unpopulate(val, "SubjectAlternativeNames", &s.SubjectAlternativeNames)
				delete(rawMsg, key)
		case "subjectName":
				err = unpopulate(val, "SubjectName", &s.SubjectName)
				delete(rawMsg, key)
		case "validityNotAfter":
				err = unpopulateTimeRFC3339(val, "ValidityNotAfter", &s.ValidityNotAfter)
				delete(rawMsg, key)
		case "validityNotBefore":
				err = unpopulateTimeRFC3339(val, "ValidityNotBefore", &s.ValidityNotBefore)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

