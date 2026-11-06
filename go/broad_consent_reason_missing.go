package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// BroadConsentReasonMissing the model 'BroadConsentReasonMissing'
type BroadConsentReasonMissing string

// List of BroadConsentReasonMissing
const (
	ORGANIZATIONAL_ISSUES BroadConsentReasonMissing = "organizational-issues"
	TECHNICAL_ISSUES      BroadConsentReasonMissing = "technical-issues"
	PATIENT_INABILITY     BroadConsentReasonMissing = "patient-inability"
	OTHER_PATIENT_REASON  BroadConsentReasonMissing = "other-patient-reason"
	CONSENT_NOT_RETURNED  BroadConsentReasonMissing = "consent-not-returned"
	PATIENT_REFUSAL       BroadConsentReasonMissing = "patient-refusal"
)

// All allowed values of BroadConsentReasonMissing enum
var AllowedBroadConsentReasonMissingEnumValues = []BroadConsentReasonMissing{
	"organizational-issues",
	"technical-issues",
	"patient-inability",
	"other-patient-reason",
	"consent-not-returned",
	"patient-refusal",
}

func (v *BroadConsentReasonMissing) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BroadConsentReasonMissing(value)
	for _, existing := range AllowedBroadConsentReasonMissingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BroadConsentReasonMissing", value)
}

// NewBroadConsentReasonMissingFromValue returns a pointer to a valid BroadConsentReasonMissing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBroadConsentReasonMissingFromValue(v string) (*BroadConsentReasonMissing, error) {
	ev := BroadConsentReasonMissing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BroadConsentReasonMissing: valid values are %v", v, AllowedBroadConsentReasonMissingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BroadConsentReasonMissing) IsValid() bool {
	for _, existing := range AllowedBroadConsentReasonMissingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BroadConsentReasonMissing value
func (v BroadConsentReasonMissing) Ptr() *BroadConsentReasonMissing {
	return &v
}

type NullableBroadConsentReasonMissing struct {
	value *BroadConsentReasonMissing
	isSet bool
}

func (v NullableBroadConsentReasonMissing) Get() *BroadConsentReasonMissing {
	return v.value
}

func (v *NullableBroadConsentReasonMissing) Set(val *BroadConsentReasonMissing) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadConsentReasonMissing) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadConsentReasonMissing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadConsentReasonMissing(val *BroadConsentReasonMissing) *NullableBroadConsentReasonMissing {
	return &NullableBroadConsentReasonMissing{value: val, isSet: true}
}

func (v NullableBroadConsentReasonMissing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadConsentReasonMissing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
