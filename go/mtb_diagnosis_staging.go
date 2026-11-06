package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbDiagnosisStaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbDiagnosisStaging{}

// MtbDiagnosisStaging struct for MtbDiagnosisStaging
type MtbDiagnosisStaging struct {
	History []TumorStaging `json:"history"`
}

type _MtbDiagnosisStaging MtbDiagnosisStaging

// NewMtbDiagnosisStaging instantiates a new MtbDiagnosisStaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbDiagnosisStaging(history []TumorStaging) *MtbDiagnosisStaging {
	this := MtbDiagnosisStaging{}
	this.History = history
	return &this
}

// NewMtbDiagnosisStagingWithDefaults instantiates a new MtbDiagnosisStaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbDiagnosisStagingWithDefaults() *MtbDiagnosisStaging {
	this := MtbDiagnosisStaging{}
	return &this
}

// GetHistory returns the History field value
func (o *MtbDiagnosisStaging) GetHistory() []TumorStaging {
	if o == nil {
		var ret []TumorStaging
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosisStaging) GetHistoryOk() ([]TumorStaging, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *MtbDiagnosisStaging) SetHistory(v []TumorStaging) {
	o.History = v
}

func (o MtbDiagnosisStaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbDiagnosisStaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["history"] = o.History
	return toSerialize, nil
}

func (o *MtbDiagnosisStaging) UnmarshalJSON(data []byte) (err error) {
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

	varMtbDiagnosisStaging := _MtbDiagnosisStaging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbDiagnosisStaging)

	if err != nil {
		return err
	}

	*o = MtbDiagnosisStaging(varMtbDiagnosisStaging)

	return err
}

type NullableMtbDiagnosisStaging struct {
	value *MtbDiagnosisStaging
	isSet bool
}

func (v NullableMtbDiagnosisStaging) Get() *MtbDiagnosisStaging {
	return v.value
}

func (v *NullableMtbDiagnosisStaging) Set(val *MtbDiagnosisStaging) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbDiagnosisStaging) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbDiagnosisStaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbDiagnosisStaging(val *MtbDiagnosisStaging) *NullableMtbDiagnosisStaging {
	return &NullableMtbDiagnosisStaging{value: val, isSet: true}
}

func (v NullableMtbDiagnosisStaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbDiagnosisStaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
