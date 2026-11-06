package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the EcogCoding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcogCoding{}

// EcogCoding struct for EcogCoding
type EcogCoding struct {
	Code                 string  `json:"code"`
	Display              *string `json:"display,omitempty"`
	System               *string `json:"system,omitempty"`
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EcogCoding EcogCoding

// NewEcogCoding instantiates a new EcogCoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcogCoding(code string) *EcogCoding {
	this := EcogCoding{}
	this.Code = code
	return &this
}

// NewEcogCodingWithDefaults instantiates a new EcogCoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcogCodingWithDefaults() *EcogCoding {
	this := EcogCoding{}
	return &this
}

// GetCode returns the Code field value
func (o *EcogCoding) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EcogCoding) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EcogCoding) SetCode(v string) {
	o.Code = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *EcogCoding) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcogCoding) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *EcogCoding) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *EcogCoding) SetDisplay(v string) {
	o.Display = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *EcogCoding) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcogCoding) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *EcogCoding) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *EcogCoding) SetSystem(v string) {
	o.System = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EcogCoding) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcogCoding) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EcogCoding) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EcogCoding) SetVersion(v string) {
	o.Version = &v
}

func (o EcogCoding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcogCoding) ToMap() (map[string]interface{}, error) {
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

func (o *EcogCoding) UnmarshalJSON(data []byte) (err error) {
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

	varEcogCoding := _EcogCoding{}

	err = json.Unmarshal(data, &varEcogCoding)

	if err != nil {
		return err
	}

	*o = EcogCoding(varEcogCoding)

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

type NullableEcogCoding struct {
	value *EcogCoding
	isSet bool
}

func (v NullableEcogCoding) Get() *EcogCoding {
	return v.value
}

func (v *NullableEcogCoding) Set(val *EcogCoding) {
	v.value = val
	v.isSet = true
}

func (v NullableEcogCoding) IsSet() bool {
	return v.isSet
}

func (v *NullableEcogCoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcogCoding(val *EcogCoding) *NullableEcogCoding {
	return &NullableEcogCoding{value: val, isSet: true}
}

func (v NullableEcogCoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcogCoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
