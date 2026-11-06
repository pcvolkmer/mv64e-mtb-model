package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbDiagnosisGrading type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbDiagnosisGrading{}

// MtbDiagnosisGrading struct for MtbDiagnosisGrading
type MtbDiagnosisGrading struct {
	History []TumorGrading `json:"history"`
}

type _MtbDiagnosisGrading MtbDiagnosisGrading

// NewMtbDiagnosisGrading instantiates a new MtbDiagnosisGrading object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbDiagnosisGrading(history []TumorGrading) *MtbDiagnosisGrading {
	this := MtbDiagnosisGrading{}
	this.History = history
	return &this
}

// NewMtbDiagnosisGradingWithDefaults instantiates a new MtbDiagnosisGrading object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbDiagnosisGradingWithDefaults() *MtbDiagnosisGrading {
	this := MtbDiagnosisGrading{}
	return &this
}

// GetHistory returns the History field value
func (o *MtbDiagnosisGrading) GetHistory() []TumorGrading {
	if o == nil {
		var ret []TumorGrading
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosisGrading) GetHistoryOk() ([]TumorGrading, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *MtbDiagnosisGrading) SetHistory(v []TumorGrading) {
	o.History = v
}

func (o MtbDiagnosisGrading) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbDiagnosisGrading) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["history"] = o.History
	return toSerialize, nil
}

func (o *MtbDiagnosisGrading) UnmarshalJSON(data []byte) (err error) {
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

	varMtbDiagnosisGrading := _MtbDiagnosisGrading{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbDiagnosisGrading)

	if err != nil {
		return err
	}

	*o = MtbDiagnosisGrading(varMtbDiagnosisGrading)

	return err
}

type NullableMtbDiagnosisGrading struct {
	value *MtbDiagnosisGrading
	isSet bool
}

func (v NullableMtbDiagnosisGrading) Get() *MtbDiagnosisGrading {
	return v.value
}

func (v *NullableMtbDiagnosisGrading) Set(val *MtbDiagnosisGrading) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbDiagnosisGrading) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbDiagnosisGrading) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbDiagnosisGrading(val *MtbDiagnosisGrading) *NullableMtbDiagnosisGrading {
	return &NullableMtbDiagnosisGrading{value: val, isSet: true}
}

func (v NullableMtbDiagnosisGrading) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbDiagnosisGrading) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
