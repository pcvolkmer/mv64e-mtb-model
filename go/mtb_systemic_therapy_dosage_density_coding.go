package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MtbSystemicTherapyDosageDensityCoding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbSystemicTherapyDosageDensityCoding{}

// MtbSystemicTherapyDosageDensityCoding struct for MtbSystemicTherapyDosageDensityCoding
type MtbSystemicTherapyDosageDensityCoding struct {
	Code                 string  `json:"code"`
	Display              *string `json:"display,omitempty"`
	System               *string `json:"system,omitempty"`
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MtbSystemicTherapyDosageDensityCoding MtbSystemicTherapyDosageDensityCoding

// NewMtbSystemicTherapyDosageDensityCoding instantiates a new MtbSystemicTherapyDosageDensityCoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbSystemicTherapyDosageDensityCoding(code string) *MtbSystemicTherapyDosageDensityCoding {
	this := MtbSystemicTherapyDosageDensityCoding{}
	this.Code = code
	return &this
}

// NewMtbSystemicTherapyDosageDensityCodingWithDefaults instantiates a new MtbSystemicTherapyDosageDensityCoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbSystemicTherapyDosageDensityCodingWithDefaults() *MtbSystemicTherapyDosageDensityCoding {
	this := MtbSystemicTherapyDosageDensityCoding{}
	return &this
}

// GetCode returns the Code field value
func (o *MtbSystemicTherapyDosageDensityCoding) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *MtbSystemicTherapyDosageDensityCoding) SetCode(v string) {
	o.Code = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *MtbSystemicTherapyDosageDensityCoding) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *MtbSystemicTherapyDosageDensityCoding) SetDisplay(v string) {
	o.Display = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *MtbSystemicTherapyDosageDensityCoding) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *MtbSystemicTherapyDosageDensityCoding) SetSystem(v string) {
	o.System = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MtbSystemicTherapyDosageDensityCoding) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MtbSystemicTherapyDosageDensityCoding) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MtbSystemicTherapyDosageDensityCoding) SetVersion(v string) {
	o.Version = &v
}

func (o MtbSystemicTherapyDosageDensityCoding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbSystemicTherapyDosageDensityCoding) ToMap() (map[string]interface{}, error) {
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

func (o *MtbSystemicTherapyDosageDensityCoding) UnmarshalJSON(data []byte) (err error) {
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

	varMtbSystemicTherapyDosageDensityCoding := _MtbSystemicTherapyDosageDensityCoding{}

	err = json.Unmarshal(data, &varMtbSystemicTherapyDosageDensityCoding)

	if err != nil {
		return err
	}

	*o = MtbSystemicTherapyDosageDensityCoding(varMtbSystemicTherapyDosageDensityCoding)

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

type NullableMtbSystemicTherapyDosageDensityCoding struct {
	value *MtbSystemicTherapyDosageDensityCoding
	isSet bool
}

func (v NullableMtbSystemicTherapyDosageDensityCoding) Get() *MtbSystemicTherapyDosageDensityCoding {
	return v.value
}

func (v *NullableMtbSystemicTherapyDosageDensityCoding) Set(val *MtbSystemicTherapyDosageDensityCoding) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbSystemicTherapyDosageDensityCoding) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbSystemicTherapyDosageDensityCoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbSystemicTherapyDosageDensityCoding(val *MtbSystemicTherapyDosageDensityCoding) *NullableMtbSystemicTherapyDosageDensityCoding {
	return &NullableMtbSystemicTherapyDosageDensityCoding{value: val, isSet: true}
}

func (v NullableMtbSystemicTherapyDosageDensityCoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbSystemicTherapyDosageDensityCoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
