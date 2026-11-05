package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the BrcanessConfidenceRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrcanessConfidenceRange{}

// BrcanessConfidenceRange struct for BrcanessConfidenceRange
type BrcanessConfidenceRange struct {
	Min                  float32 `json:"min"`
	Max                  float32 `json:"max"`
	AdditionalProperties map[string]interface{}
}

type _BrcanessConfidenceRange BrcanessConfidenceRange

// NewBrcanessConfidenceRange instantiates a new BrcanessConfidenceRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrcanessConfidenceRange(min float32, max float32) *BrcanessConfidenceRange {
	this := BrcanessConfidenceRange{}
	this.Min = min
	this.Max = max
	return &this
}

// NewBrcanessConfidenceRangeWithDefaults instantiates a new BrcanessConfidenceRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrcanessConfidenceRangeWithDefaults() *BrcanessConfidenceRange {
	this := BrcanessConfidenceRange{}
	return &this
}

// GetMin returns the Min field value
func (o *BrcanessConfidenceRange) GetMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *BrcanessConfidenceRange) GetMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *BrcanessConfidenceRange) SetMin(v float32) {
	o.Min = v
}

// GetMax returns the Max field value
func (o *BrcanessConfidenceRange) GetMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *BrcanessConfidenceRange) GetMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *BrcanessConfidenceRange) SetMax(v float32) {
	o.Max = v
}

func (o BrcanessConfidenceRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrcanessConfidenceRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BrcanessConfidenceRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"min",
		"max",
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

	varBrcanessConfidenceRange := _BrcanessConfidenceRange{}

	err = json.Unmarshal(data, &varBrcanessConfidenceRange)

	if err != nil {
		return err
	}

	*o = BrcanessConfidenceRange(varBrcanessConfidenceRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "min")
		delete(additionalProperties, "max")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrcanessConfidenceRange struct {
	value *BrcanessConfidenceRange
	isSet bool
}

func (v NullableBrcanessConfidenceRange) Get() *BrcanessConfidenceRange {
	return v.value
}

func (v *NullableBrcanessConfidenceRange) Set(val *BrcanessConfidenceRange) {
	v.value = val
	v.isSet = true
}

func (v NullableBrcanessConfidenceRange) IsSet() bool {
	return v.isSet
}

func (v *NullableBrcanessConfidenceRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrcanessConfidenceRange(val *BrcanessConfidenceRange) *NullableBrcanessConfidenceRange {
	return &NullableBrcanessConfidenceRange{value: val, isSet: true}
}

func (v NullableBrcanessConfidenceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrcanessConfidenceRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
