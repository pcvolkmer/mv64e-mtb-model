package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PatientRecordStudyInclusionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientRecordStudyInclusionsInner{}

// PatientRecordStudyInclusionsInner struct for PatientRecordStudyInclusionsInner
type PatientRecordStudyInclusionsInner struct {
	History []StudyInclusion `json:"history"`
}

type _PatientRecordStudyInclusionsInner PatientRecordStudyInclusionsInner

// NewPatientRecordStudyInclusionsInner instantiates a new PatientRecordStudyInclusionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientRecordStudyInclusionsInner(history []StudyInclusion) *PatientRecordStudyInclusionsInner {
	this := PatientRecordStudyInclusionsInner{}
	this.History = history
	return &this
}

// NewPatientRecordStudyInclusionsInnerWithDefaults instantiates a new PatientRecordStudyInclusionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientRecordStudyInclusionsInnerWithDefaults() *PatientRecordStudyInclusionsInner {
	this := PatientRecordStudyInclusionsInner{}
	return &this
}

// GetHistory returns the History field value
func (o *PatientRecordStudyInclusionsInner) GetHistory() []StudyInclusion {
	if o == nil {
		var ret []StudyInclusion
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *PatientRecordStudyInclusionsInner) GetHistoryOk() ([]StudyInclusion, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *PatientRecordStudyInclusionsInner) SetHistory(v []StudyInclusion) {
	o.History = v
}

func (o PatientRecordStudyInclusionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientRecordStudyInclusionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["history"] = o.History
	return toSerialize, nil
}

func (o *PatientRecordStudyInclusionsInner) UnmarshalJSON(data []byte) (err error) {
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

	varPatientRecordStudyInclusionsInner := _PatientRecordStudyInclusionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientRecordStudyInclusionsInner)

	if err != nil {
		return err
	}

	*o = PatientRecordStudyInclusionsInner(varPatientRecordStudyInclusionsInner)

	return err
}

type NullablePatientRecordStudyInclusionsInner struct {
	value *PatientRecordStudyInclusionsInner
	isSet bool
}

func (v NullablePatientRecordStudyInclusionsInner) Get() *PatientRecordStudyInclusionsInner {
	return v.value
}

func (v *NullablePatientRecordStudyInclusionsInner) Set(val *PatientRecordStudyInclusionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientRecordStudyInclusionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientRecordStudyInclusionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientRecordStudyInclusionsInner(val *PatientRecordStudyInclusionsInner) *NullablePatientRecordStudyInclusionsInner {
	return &NullablePatientRecordStudyInclusionsInner{value: val, isSet: true}
}

func (v NullablePatientRecordStudyInclusionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientRecordStudyInclusionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
