package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// MvhSubmissionType the model 'MvhSubmissionType'
type MvhSubmissionType string

// List of MvhSubmissionType
const (
	INITIAL    MvhSubmissionType = "initial"
	CORRECTION MvhSubmissionType = "correction"
	TEST       MvhSubmissionType = "test"
	ADDITION   MvhSubmissionType = "addition"
	FOLLOWUP   MvhSubmissionType = "followup"
)

// All allowed values of MvhSubmissionType enum
var AllowedMvhSubmissionTypeEnumValues = []MvhSubmissionType{
	"initial",
	"correction",
	"test",
	"addition",
	"followup",
}

func (v *MvhSubmissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MvhSubmissionType(value)
	for _, existing := range AllowedMvhSubmissionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MvhSubmissionType", value)
}

// NewMvhSubmissionTypeFromValue returns a pointer to a valid MvhSubmissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMvhSubmissionTypeFromValue(v string) (*MvhSubmissionType, error) {
	ev := MvhSubmissionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MvhSubmissionType: valid values are %v", v, AllowedMvhSubmissionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MvhSubmissionType) IsValid() bool {
	for _, existing := range AllowedMvhSubmissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MvhSubmissionType value
func (v MvhSubmissionType) Ptr() *MvhSubmissionType {
	return &v
}

type NullableMvhSubmissionType struct {
	value *MvhSubmissionType
	isSet bool
}

func (v NullableMvhSubmissionType) Get() *MvhSubmissionType {
	return v.value
}

func (v *NullableMvhSubmissionType) Set(val *MvhSubmissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMvhSubmissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMvhSubmissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMvhSubmissionType(val *MvhSubmissionType) *NullableMvhSubmissionType {
	return &NullableMvhSubmissionType{value: val, isSet: true}
}

func (v NullableMvhSubmissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMvhSubmissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
