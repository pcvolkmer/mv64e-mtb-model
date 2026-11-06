package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CnvEndRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CnvEndRange{}

// CnvEndRange struct for CnvEndRange
type CnvEndRange struct {
	Start float32  `json:"start"`
	End   *float32 `json:"end,omitempty"`
}

type _CnvEndRange CnvEndRange

// NewCnvEndRange instantiates a new CnvEndRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnvEndRange(start float32) *CnvEndRange {
	this := CnvEndRange{}
	this.Start = start
	return &this
}

// NewCnvEndRangeWithDefaults instantiates a new CnvEndRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnvEndRangeWithDefaults() *CnvEndRange {
	this := CnvEndRange{}
	return &this
}

// GetStart returns the Start field value
func (o *CnvEndRange) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CnvEndRange) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CnvEndRange) SetStart(v float32) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CnvEndRange) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnvEndRange) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CnvEndRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *CnvEndRange) SetEnd(v float32) {
	o.End = &v
}

func (o CnvEndRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CnvEndRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

func (o *CnvEndRange) UnmarshalJSON(data []byte) (err error) {
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

	varCnvEndRange := _CnvEndRange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCnvEndRange)

	if err != nil {
		return err
	}

	*o = CnvEndRange(varCnvEndRange)

	return err
}

type NullableCnvEndRange struct {
	value *CnvEndRange
	isSet bool
}

func (v NullableCnvEndRange) Get() *CnvEndRange {
	return v.value
}

func (v *NullableCnvEndRange) Set(val *CnvEndRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCnvEndRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCnvEndRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnvEndRange(val *CnvEndRange) *NullableCnvEndRange {
	return &NullableCnvEndRange{value: val, isSet: true}
}

func (v NullableCnvEndRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnvEndRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
