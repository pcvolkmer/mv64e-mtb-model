package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the HrdScoreComponents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrdScoreComponents{}

// HrdScoreComponents struct for HrdScoreComponents
type HrdScoreComponents struct {
	Lst                  float32 `json:"lst"`
	Loh                  float32 `json:"loh"`
	Tai                  float32 `json:"tai"`
	AdditionalProperties map[string]interface{}
}

type _HrdScoreComponents HrdScoreComponents

// NewHrdScoreComponents instantiates a new HrdScoreComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrdScoreComponents(lst float32, loh float32, tai float32) *HrdScoreComponents {
	this := HrdScoreComponents{}
	this.Lst = lst
	this.Loh = loh
	this.Tai = tai
	return &this
}

// NewHrdScoreComponentsWithDefaults instantiates a new HrdScoreComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrdScoreComponentsWithDefaults() *HrdScoreComponents {
	this := HrdScoreComponents{}
	return &this
}

// GetLst returns the Lst field value
func (o *HrdScoreComponents) GetLst() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lst
}

// GetLstOk returns a tuple with the Lst field value
// and a boolean to check if the value has been set.
func (o *HrdScoreComponents) GetLstOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lst, true
}

// SetLst sets field value
func (o *HrdScoreComponents) SetLst(v float32) {
	o.Lst = v
}

// GetLoh returns the Loh field value
func (o *HrdScoreComponents) GetLoh() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Loh
}

// GetLohOk returns a tuple with the Loh field value
// and a boolean to check if the value has been set.
func (o *HrdScoreComponents) GetLohOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loh, true
}

// SetLoh sets field value
func (o *HrdScoreComponents) SetLoh(v float32) {
	o.Loh = v
}

// GetTai returns the Tai field value
func (o *HrdScoreComponents) GetTai() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *HrdScoreComponents) GetTaiOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *HrdScoreComponents) SetTai(v float32) {
	o.Tai = v
}

func (o HrdScoreComponents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrdScoreComponents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lst"] = o.Lst
	toSerialize["loh"] = o.Loh
	toSerialize["tai"] = o.Tai

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HrdScoreComponents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lst",
		"loh",
		"tai",
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

	varHrdScoreComponents := _HrdScoreComponents{}

	err = json.Unmarshal(data, &varHrdScoreComponents)

	if err != nil {
		return err
	}

	*o = HrdScoreComponents(varHrdScoreComponents)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lst")
		delete(additionalProperties, "loh")
		delete(additionalProperties, "tai")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrdScoreComponents struct {
	value *HrdScoreComponents
	isSet bool
}

func (v NullableHrdScoreComponents) Get() *HrdScoreComponents {
	return v.value
}

func (v *NullableHrdScoreComponents) Set(val *HrdScoreComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableHrdScoreComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableHrdScoreComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrdScoreComponents(val *HrdScoreComponents) *NullableHrdScoreComponents {
	return &NullableHrdScoreComponents{value: val, isSet: true}
}

func (v NullableHrdScoreComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrdScoreComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
