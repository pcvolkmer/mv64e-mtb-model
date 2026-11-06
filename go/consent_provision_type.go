package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// ConsentProvisionType the model 'ConsentProvisionType'
type ConsentProvisionType string

// List of ConsentProvisionType
const (
	DENY   ConsentProvisionType = "deny"
	PERMIT ConsentProvisionType = "permit"
)

// All allowed values of ConsentProvisionType enum
var AllowedConsentProvisionTypeEnumValues = []ConsentProvisionType{
	"deny",
	"permit",
}

func (v *ConsentProvisionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsentProvisionType(value)
	for _, existing := range AllowedConsentProvisionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsentProvisionType", value)
}

// NewConsentProvisionTypeFromValue returns a pointer to a valid ConsentProvisionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsentProvisionTypeFromValue(v string) (*ConsentProvisionType, error) {
	ev := ConsentProvisionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsentProvisionType: valid values are %v", v, AllowedConsentProvisionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsentProvisionType) IsValid() bool {
	for _, existing := range AllowedConsentProvisionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsentProvisionType value
func (v ConsentProvisionType) Ptr() *ConsentProvisionType {
	return &v
}

type NullableConsentProvisionType struct {
	value *ConsentProvisionType
	isSet bool
}

func (v NullableConsentProvisionType) Get() *ConsentProvisionType {
	return v.value
}

func (v *NullableConsentProvisionType) Set(val *ConsentProvisionType) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentProvisionType) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentProvisionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentProvisionType(val *ConsentProvisionType) *NullableConsentProvisionType {
	return &NullableConsentProvisionType{value: val, isSet: true}
}

func (v NullableConsentProvisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentProvisionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
