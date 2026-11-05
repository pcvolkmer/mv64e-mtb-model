package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PatientRecordSystemicTherapiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientRecordSystemicTherapiesInner{}

// PatientRecordSystemicTherapiesInner struct for PatientRecordSystemicTherapiesInner
type PatientRecordSystemicTherapiesInner struct {
	History []MtbSystemicTherapy `json:"history"`
}

type _PatientRecordSystemicTherapiesInner PatientRecordSystemicTherapiesInner

// NewPatientRecordSystemicTherapiesInner instantiates a new PatientRecordSystemicTherapiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientRecordSystemicTherapiesInner(history []MtbSystemicTherapy) *PatientRecordSystemicTherapiesInner {
	this := PatientRecordSystemicTherapiesInner{}
	this.History = history
	return &this
}

// NewPatientRecordSystemicTherapiesInnerWithDefaults instantiates a new PatientRecordSystemicTherapiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientRecordSystemicTherapiesInnerWithDefaults() *PatientRecordSystemicTherapiesInner {
	this := PatientRecordSystemicTherapiesInner{}
	return &this
}

// GetHistory returns the History field value
func (o *PatientRecordSystemicTherapiesInner) GetHistory() []MtbSystemicTherapy {
	if o == nil {
		var ret []MtbSystemicTherapy
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *PatientRecordSystemicTherapiesInner) GetHistoryOk() ([]MtbSystemicTherapy, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *PatientRecordSystemicTherapiesInner) SetHistory(v []MtbSystemicTherapy) {
	o.History = v
}

func (o PatientRecordSystemicTherapiesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientRecordSystemicTherapiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["history"] = o.History
	return toSerialize, nil
}

func (o *PatientRecordSystemicTherapiesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"history",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPatientRecordSystemicTherapiesInner := _PatientRecordSystemicTherapiesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientRecordSystemicTherapiesInner)

	if err != nil {
		return err
	}

	*o = PatientRecordSystemicTherapiesInner(varPatientRecordSystemicTherapiesInner)

	return err
}

type NullablePatientRecordSystemicTherapiesInner struct {
	value *PatientRecordSystemicTherapiesInner
	isSet bool
}

func (v NullablePatientRecordSystemicTherapiesInner) Get() *PatientRecordSystemicTherapiesInner {
	return v.value
}

func (v *NullablePatientRecordSystemicTherapiesInner) Set(val *PatientRecordSystemicTherapiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientRecordSystemicTherapiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientRecordSystemicTherapiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientRecordSystemicTherapiesInner(val *PatientRecordSystemicTherapiesInner) *NullablePatientRecordSystemicTherapiesInner {
	return &NullablePatientRecordSystemicTherapiesInner{value: val, isSet: true}
}

func (v NullablePatientRecordSystemicTherapiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientRecordSystemicTherapiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
