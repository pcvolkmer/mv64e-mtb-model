package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the StudyReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyReference{}

// StudyReference struct for StudyReference
type StudyReference struct {
	Id                   string  `json:"id"`
	System               string  `json:"system"`
	Display              *string `json:"display,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyReference StudyReference

// NewStudyReference instantiates a new StudyReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyReference(id string, system string) *StudyReference {
	this := StudyReference{}
	this.Id = id
	this.System = system
	return &this
}

// NewStudyReferenceWithDefaults instantiates a new StudyReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyReferenceWithDefaults() *StudyReference {
	this := StudyReference{}
	return &this
}

// GetId returns the Id field value
func (o *StudyReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StudyReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StudyReference) SetId(v string) {
	o.Id = v
}

// GetSystem returns the System field value
func (o *StudyReference) GetSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *StudyReference) GetSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *StudyReference) SetSystem(v string) {
	o.System = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *StudyReference) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyReference) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *StudyReference) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *StudyReference) SetDisplay(v string) {
	o.Display = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StudyReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StudyReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StudyReference) SetType(v string) {
	o.Type = &v
}

func (o StudyReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["system"] = o.System
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varStudyReference := _StudyReference{}

	err = json.Unmarshal(data, &varStudyReference)

	if err != nil {
		return err
	}

	*o = StudyReference(varStudyReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "system")
		delete(additionalProperties, "display")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyReference struct {
	value *StudyReference
	isSet bool
}

func (v NullableStudyReference) Get() *StudyReference {
	return v.value
}

func (v *NullableStudyReference) Set(val *StudyReference) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyReference) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyReference(val *StudyReference) *NullableStudyReference {
	return &NullableStudyReference{value: val, isSet: true}
}

func (v NullableStudyReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
