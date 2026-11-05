package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PatientAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientAddress{}

// PatientAddress struct for PatientAddress
type PatientAddress struct {
	MunicipalityCode string `json:"municipalityCode" validate:"regexp=^\\\\d{5}$"`
}

type _PatientAddress PatientAddress

// NewPatientAddress instantiates a new PatientAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientAddress(municipalityCode string) *PatientAddress {
	this := PatientAddress{}
	this.MunicipalityCode = municipalityCode
	return &this
}

// NewPatientAddressWithDefaults instantiates a new PatientAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientAddressWithDefaults() *PatientAddress {
	this := PatientAddress{}
	return &this
}

// GetMunicipalityCode returns the MunicipalityCode field value
func (o *PatientAddress) GetMunicipalityCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MunicipalityCode
}

// GetMunicipalityCodeOk returns a tuple with the MunicipalityCode field value
// and a boolean to check if the value has been set.
func (o *PatientAddress) GetMunicipalityCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MunicipalityCode, true
}

// SetMunicipalityCode sets field value
func (o *PatientAddress) SetMunicipalityCode(v string) {
	o.MunicipalityCode = v
}

func (o PatientAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["municipalityCode"] = o.MunicipalityCode
	return toSerialize, nil
}

func (o *PatientAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"municipalityCode",
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

	varPatientAddress := _PatientAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientAddress)

	if err != nil {
		return err
	}

	*o = PatientAddress(varPatientAddress)

	return err
}

type NullablePatientAddress struct {
	value *PatientAddress
	isSet bool
}

func (v NullablePatientAddress) Get() *PatientAddress {
	return v.value
}

func (v *NullablePatientAddress) Set(val *PatientAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientAddress(val *PatientAddress) *NullablePatientAddress {
	return &NullablePatientAddress{value: val, isSet: true}
}

func (v NullablePatientAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
