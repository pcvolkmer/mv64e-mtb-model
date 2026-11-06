package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// RnaFusionStrand the model 'RnaFusionStrand'
type RnaFusionStrand string

// List of RnaFusionStrand
const (
	PLUS  RnaFusionStrand = "+"
	MINUS RnaFusionStrand = "-"
)

// All allowed values of RnaFusionStrand enum
var AllowedRnaFusionStrandEnumValues = []RnaFusionStrand{
	"+",
	"-",
}

func (v *RnaFusionStrand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RnaFusionStrand(value)
	for _, existing := range AllowedRnaFusionStrandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RnaFusionStrand", value)
}

// NewRnaFusionStrandFromValue returns a pointer to a valid RnaFusionStrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRnaFusionStrandFromValue(v string) (*RnaFusionStrand, error) {
	ev := RnaFusionStrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RnaFusionStrand: valid values are %v", v, AllowedRnaFusionStrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RnaFusionStrand) IsValid() bool {
	for _, existing := range AllowedRnaFusionStrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RnaFusionStrand value
func (v RnaFusionStrand) Ptr() *RnaFusionStrand {
	return &v
}

type NullableRnaFusionStrand struct {
	value *RnaFusionStrand
	isSet bool
}

func (v NullableRnaFusionStrand) Get() *RnaFusionStrand {
	return v.value
}

func (v *NullableRnaFusionStrand) Set(val *RnaFusionStrand) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaFusionStrand) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaFusionStrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaFusionStrand(val *RnaFusionStrand) *NullableRnaFusionStrand {
	return &NullableRnaFusionStrand{value: val, isSet: true}
}

func (v NullableRnaFusionStrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaFusionStrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
