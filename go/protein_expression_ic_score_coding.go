package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the ProteinExpressionIcScoreCoding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProteinExpressionIcScoreCoding{}

// ProteinExpressionIcScoreCoding struct for ProteinExpressionIcScoreCoding
type ProteinExpressionIcScoreCoding struct {
	Code                 string  `json:"code"`
	Display              *string `json:"display,omitempty"`
	System               *string `json:"system,omitempty"`
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProteinExpressionIcScoreCoding ProteinExpressionIcScoreCoding

// NewProteinExpressionIcScoreCoding instantiates a new ProteinExpressionIcScoreCoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProteinExpressionIcScoreCoding(code string) *ProteinExpressionIcScoreCoding {
	this := ProteinExpressionIcScoreCoding{}
	this.Code = code
	return &this
}

// NewProteinExpressionIcScoreCodingWithDefaults instantiates a new ProteinExpressionIcScoreCoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProteinExpressionIcScoreCodingWithDefaults() *ProteinExpressionIcScoreCoding {
	this := ProteinExpressionIcScoreCoding{}
	return &this
}

// GetCode returns the Code field value
func (o *ProteinExpressionIcScoreCoding) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ProteinExpressionIcScoreCoding) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ProteinExpressionIcScoreCoding) SetCode(v string) {
	o.Code = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ProteinExpressionIcScoreCoding) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpressionIcScoreCoding) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ProteinExpressionIcScoreCoding) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ProteinExpressionIcScoreCoding) SetDisplay(v string) {
	o.Display = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ProteinExpressionIcScoreCoding) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpressionIcScoreCoding) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ProteinExpressionIcScoreCoding) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *ProteinExpressionIcScoreCoding) SetSystem(v string) {
	o.System = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ProteinExpressionIcScoreCoding) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpressionIcScoreCoding) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ProteinExpressionIcScoreCoding) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ProteinExpressionIcScoreCoding) SetVersion(v string) {
	o.Version = &v
}

func (o ProteinExpressionIcScoreCoding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProteinExpressionIcScoreCoding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProteinExpressionIcScoreCoding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varProteinExpressionIcScoreCoding := _ProteinExpressionIcScoreCoding{}

	err = json.Unmarshal(data, &varProteinExpressionIcScoreCoding)

	if err != nil {
		return err
	}

	*o = ProteinExpressionIcScoreCoding(varProteinExpressionIcScoreCoding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "display")
		delete(additionalProperties, "system")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProteinExpressionIcScoreCoding struct {
	value *ProteinExpressionIcScoreCoding
	isSet bool
}

func (v NullableProteinExpressionIcScoreCoding) Get() *ProteinExpressionIcScoreCoding {
	return v.value
}

func (v *NullableProteinExpressionIcScoreCoding) Set(val *ProteinExpressionIcScoreCoding) {
	v.value = val
	v.isSet = true
}

func (v NullableProteinExpressionIcScoreCoding) IsSet() bool {
	return v.isSet
}

func (v *NullableProteinExpressionIcScoreCoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProteinExpressionIcScoreCoding(val *ProteinExpressionIcScoreCoding) *NullableProteinExpressionIcScoreCoding {
	return &NullableProteinExpressionIcScoreCoding{value: val, isSet: true}
}

func (v NullableProteinExpressionIcScoreCoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProteinExpressionIcScoreCoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
