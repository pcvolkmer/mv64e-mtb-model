package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbDiagnosisTypeHistoryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbDiagnosisTypeHistoryInner{}

// MtbDiagnosisTypeHistoryInner struct for MtbDiagnosisTypeHistoryInner
type MtbDiagnosisTypeHistoryInner struct {
	Value MtbDiagnosisTypeCoding `json:"value"`
	Date  string                 `json:"date"`
}

type _MtbDiagnosisTypeHistoryInner MtbDiagnosisTypeHistoryInner

// NewMtbDiagnosisTypeHistoryInner instantiates a new MtbDiagnosisTypeHistoryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbDiagnosisTypeHistoryInner(value MtbDiagnosisTypeCoding, date string) *MtbDiagnosisTypeHistoryInner {
	this := MtbDiagnosisTypeHistoryInner{}
	this.Value = value
	this.Date = date
	return &this
}

// NewMtbDiagnosisTypeHistoryInnerWithDefaults instantiates a new MtbDiagnosisTypeHistoryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbDiagnosisTypeHistoryInnerWithDefaults() *MtbDiagnosisTypeHistoryInner {
	this := MtbDiagnosisTypeHistoryInner{}
	return &this
}

// GetValue returns the Value field value
func (o *MtbDiagnosisTypeHistoryInner) GetValue() MtbDiagnosisTypeCoding {
	if o == nil {
		var ret MtbDiagnosisTypeCoding
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosisTypeHistoryInner) GetValueOk() (*MtbDiagnosisTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MtbDiagnosisTypeHistoryInner) SetValue(v MtbDiagnosisTypeCoding) {
	o.Value = v
}

// GetDate returns the Date field value
func (o *MtbDiagnosisTypeHistoryInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosisTypeHistoryInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MtbDiagnosisTypeHistoryInner) SetDate(v string) {
	o.Date = v
}

func (o MtbDiagnosisTypeHistoryInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbDiagnosisTypeHistoryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

func (o *MtbDiagnosisTypeHistoryInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"date",
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

	varMtbDiagnosisTypeHistoryInner := _MtbDiagnosisTypeHistoryInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbDiagnosisTypeHistoryInner)

	if err != nil {
		return err
	}

	*o = MtbDiagnosisTypeHistoryInner(varMtbDiagnosisTypeHistoryInner)

	return err
}

type NullableMtbDiagnosisTypeHistoryInner struct {
	value *MtbDiagnosisTypeHistoryInner
	isSet bool
}

func (v NullableMtbDiagnosisTypeHistoryInner) Get() *MtbDiagnosisTypeHistoryInner {
	return v.value
}

func (v *NullableMtbDiagnosisTypeHistoryInner) Set(val *MtbDiagnosisTypeHistoryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbDiagnosisTypeHistoryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbDiagnosisTypeHistoryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbDiagnosisTypeHistoryInner(val *MtbDiagnosisTypeHistoryInner) *NullableMtbDiagnosisTypeHistoryInner {
	return &NullableMtbDiagnosisTypeHistoryInner{value: val, isSet: true}
}

func (v NullableMtbDiagnosisTypeHistoryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbDiagnosisTypeHistoryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
