package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the CnvStartRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CnvStartRange{}

// CnvStartRange struct for CnvStartRange
type CnvStartRange struct {
	Start                float32  `json:"start"`
	End                  *float32 `json:"end,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CnvStartRange CnvStartRange

// NewCnvStartRange instantiates a new CnvStartRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnvStartRange(start float32) *CnvStartRange {
	this := CnvStartRange{}
	this.Start = start
	return &this
}

// NewCnvStartRangeWithDefaults instantiates a new CnvStartRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnvStartRangeWithDefaults() *CnvStartRange {
	this := CnvStartRange{}
	return &this
}

// GetStart returns the Start field value
func (o *CnvStartRange) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CnvStartRange) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CnvStartRange) SetStart(v float32) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CnvStartRange) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnvStartRange) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CnvStartRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *CnvStartRange) SetEnd(v float32) {
	o.End = &v
}

func (o CnvStartRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CnvStartRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CnvStartRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
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

	varCnvStartRange := _CnvStartRange{}

	err = json.Unmarshal(data, &varCnvStartRange)

	if err != nil {
		return err
	}

	*o = CnvStartRange(varCnvStartRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCnvStartRange struct {
	value *CnvStartRange
	isSet bool
}

func (v NullableCnvStartRange) Get() *CnvStartRange {
	return v.value
}

func (v *NullableCnvStartRange) Set(val *CnvStartRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCnvStartRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCnvStartRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnvStartRange(val *CnvStartRange) *NullableCnvStartRange {
	return &NullableCnvStartRange{value: val, isSet: true}
}

func (v NullableCnvStartRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnvStartRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
