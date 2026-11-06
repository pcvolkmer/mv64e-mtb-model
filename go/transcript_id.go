package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TranscriptId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscriptId{}

// TranscriptId struct for TranscriptId
type TranscriptId struct {
	Value                string `json:"value"`
	System               string `json:"system"`
	AdditionalProperties map[string]interface{}
}

type _TranscriptId TranscriptId

// NewTranscriptId instantiates a new TranscriptId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscriptId(value string, system string) *TranscriptId {
	this := TranscriptId{}
	this.Value = value
	this.System = system
	return &this
}

// NewTranscriptIdWithDefaults instantiates a new TranscriptId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscriptIdWithDefaults() *TranscriptId {
	this := TranscriptId{}
	return &this
}

// GetValue returns the Value field value
func (o *TranscriptId) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TranscriptId) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TranscriptId) SetValue(v string) {
	o.Value = v
}

// GetSystem returns the System field value
func (o *TranscriptId) GetSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *TranscriptId) GetSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *TranscriptId) SetSystem(v string) {
	o.System = v
}

func (o TranscriptId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscriptId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["system"] = o.System

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TranscriptId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"system",
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

	varTranscriptId := _TranscriptId{}

	err = json.Unmarshal(data, &varTranscriptId)

	if err != nil {
		return err
	}

	*o = TranscriptId(varTranscriptId)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "system")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTranscriptId struct {
	value *TranscriptId
	isSet bool
}

func (v NullableTranscriptId) Get() *TranscriptId {
	return v.value
}

func (v *NullableTranscriptId) Set(val *TranscriptId) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscriptId) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscriptId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscriptId(val *TranscriptId) *NullableTranscriptId {
	return &NullableTranscriptId{value: val, isSet: true}
}

func (v NullableTranscriptId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscriptId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
