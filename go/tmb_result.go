package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TmbResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TmbResult{}

// TmbResult struct for TmbResult
type TmbResult struct {
	Value                float32 `json:"value"`
	Unit                 *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TmbResult TmbResult

// NewTmbResult instantiates a new TmbResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmbResult(value float32) *TmbResult {
	this := TmbResult{}
	this.Value = value
	return &this
}

// NewTmbResultWithDefaults instantiates a new TmbResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmbResultWithDefaults() *TmbResult {
	this := TmbResult{}
	return &this
}

// GetValue returns the Value field value
func (o *TmbResult) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TmbResult) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TmbResult) SetValue(v float32) {
	o.Value = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TmbResult) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmbResult) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TmbResult) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *TmbResult) SetUnit(v string) {
	o.Unit = &v
}

func (o TmbResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TmbResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TmbResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varTmbResult := _TmbResult{}

	err = json.Unmarshal(data, &varTmbResult)

	if err != nil {
		return err
	}

	*o = TmbResult(varTmbResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTmbResult struct {
	value *TmbResult
	isSet bool
}

func (v NullableTmbResult) Get() *TmbResult {
	return v.value
}

func (v *NullableTmbResult) Set(val *TmbResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTmbResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTmbResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmbResult(val *TmbResult) *NullableTmbResult {
	return &NullableTmbResult{value: val, isSet: true}
}

func (v NullableTmbResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmbResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
