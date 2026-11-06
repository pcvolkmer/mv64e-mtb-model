package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// ModelProjectConsentPurpose the model 'ModelProjectConsentPurpose'
type ModelProjectConsentPurpose string

// List of ModelProjectConsentPurpose
const (
	CASE_IDENTIFICATION ModelProjectConsentPurpose = "case-identification"
	REIDENTIFICATION    ModelProjectConsentPurpose = "reidentification"
	SEQUENCING          ModelProjectConsentPurpose = "sequencing"
)

// All allowed values of ModelProjectConsentPurpose enum
var AllowedModelProjectConsentPurposeEnumValues = []ModelProjectConsentPurpose{
	"case-identification",
	"reidentification",
	"sequencing",
}

func (v *ModelProjectConsentPurpose) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelProjectConsentPurpose(value)
	for _, existing := range AllowedModelProjectConsentPurposeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelProjectConsentPurpose", value)
}

// NewModelProjectConsentPurposeFromValue returns a pointer to a valid ModelProjectConsentPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelProjectConsentPurposeFromValue(v string) (*ModelProjectConsentPurpose, error) {
	ev := ModelProjectConsentPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelProjectConsentPurpose: valid values are %v", v, AllowedModelProjectConsentPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelProjectConsentPurpose) IsValid() bool {
	for _, existing := range AllowedModelProjectConsentPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelProjectConsentPurpose value
func (v ModelProjectConsentPurpose) Ptr() *ModelProjectConsentPurpose {
	return &v
}

type NullableModelProjectConsentPurpose struct {
	value *ModelProjectConsentPurpose
	isSet bool
}

func (v NullableModelProjectConsentPurpose) Get() *ModelProjectConsentPurpose {
	return v.value
}

func (v *NullableModelProjectConsentPurpose) Set(val *ModelProjectConsentPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableModelProjectConsentPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableModelProjectConsentPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelProjectConsentPurpose(val *ModelProjectConsentPurpose) *NullableModelProjectConsentPurpose {
	return &NullableModelProjectConsentPurpose{value: val, isSet: true}
}

func (v NullableModelProjectConsentPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelProjectConsentPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
