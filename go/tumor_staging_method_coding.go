package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorStagingMethodCoding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorStagingMethodCoding{}

// TumorStagingMethodCoding struct for TumorStagingMethodCoding
type TumorStagingMethodCoding struct {
	Code                 string  `json:"code"`
	Display              *string `json:"display,omitempty"`
	System               *string `json:"system,omitempty"`
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TumorStagingMethodCoding TumorStagingMethodCoding

// NewTumorStagingMethodCoding instantiates a new TumorStagingMethodCoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorStagingMethodCoding(code string) *TumorStagingMethodCoding {
	this := TumorStagingMethodCoding{}
	this.Code = code
	return &this
}

// NewTumorStagingMethodCodingWithDefaults instantiates a new TumorStagingMethodCoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorStagingMethodCodingWithDefaults() *TumorStagingMethodCoding {
	this := TumorStagingMethodCoding{}
	return &this
}

// GetCode returns the Code field value
func (o *TumorStagingMethodCoding) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TumorStagingMethodCoding) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TumorStagingMethodCoding) SetCode(v string) {
	o.Code = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *TumorStagingMethodCoding) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorStagingMethodCoding) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *TumorStagingMethodCoding) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *TumorStagingMethodCoding) SetDisplay(v string) {
	o.Display = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *TumorStagingMethodCoding) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorStagingMethodCoding) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *TumorStagingMethodCoding) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *TumorStagingMethodCoding) SetSystem(v string) {
	o.System = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TumorStagingMethodCoding) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorStagingMethodCoding) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TumorStagingMethodCoding) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TumorStagingMethodCoding) SetVersion(v string) {
	o.Version = &v
}

func (o TumorStagingMethodCoding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorStagingMethodCoding) ToMap() (map[string]interface{}, error) {
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

func (o *TumorStagingMethodCoding) UnmarshalJSON(data []byte) (err error) {
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

	varTumorStagingMethodCoding := _TumorStagingMethodCoding{}

	err = json.Unmarshal(data, &varTumorStagingMethodCoding)

	if err != nil {
		return err
	}

	*o = TumorStagingMethodCoding(varTumorStagingMethodCoding)

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

type NullableTumorStagingMethodCoding struct {
	value *TumorStagingMethodCoding
	isSet bool
}

func (v NullableTumorStagingMethodCoding) Get() *TumorStagingMethodCoding {
	return v.value
}

func (v *NullableTumorStagingMethodCoding) Set(val *TumorStagingMethodCoding) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorStagingMethodCoding) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorStagingMethodCoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorStagingMethodCoding(val *TumorStagingMethodCoding) *NullableTumorStagingMethodCoding {
	return &NullableTumorStagingMethodCoding{value: val, isSet: true}
}

func (v NullableTumorStagingMethodCoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorStagingMethodCoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
