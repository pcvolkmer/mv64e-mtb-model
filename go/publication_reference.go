package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicationReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicationReference{}

// PublicationReference struct for PublicationReference
type PublicationReference struct {
	Id                   string  `json:"id"`
	System               string  `json:"system"`
	Display              *string `json:"display,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicationReference PublicationReference

// NewPublicationReference instantiates a new PublicationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicationReference(id string, system string) *PublicationReference {
	this := PublicationReference{}
	this.Id = id
	this.System = system
	return &this
}

// NewPublicationReferenceWithDefaults instantiates a new PublicationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicationReferenceWithDefaults() *PublicationReference {
	this := PublicationReference{}
	return &this
}

// GetId returns the Id field value
func (o *PublicationReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicationReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicationReference) SetId(v string) {
	o.Id = v
}

// GetSystem returns the System field value
func (o *PublicationReference) GetSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *PublicationReference) GetSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *PublicationReference) SetSystem(v string) {
	o.System = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PublicationReference) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicationReference) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PublicationReference) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PublicationReference) SetDisplay(v string) {
	o.Display = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PublicationReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicationReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PublicationReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PublicationReference) SetType(v string) {
	o.Type = &v
}

func (o PublicationReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicationReference) ToMap() (map[string]interface{}, error) {
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

func (o *PublicationReference) UnmarshalJSON(data []byte) (err error) {
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

	varPublicationReference := _PublicationReference{}

	err = json.Unmarshal(data, &varPublicationReference)

	if err != nil {
		return err
	}

	*o = PublicationReference(varPublicationReference)

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

type NullablePublicationReference struct {
	value *PublicationReference
	isSet bool
}

func (v NullablePublicationReference) Get() *PublicationReference {
	return v.value
}

func (v *NullablePublicationReference) Set(val *PublicationReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicationReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicationReference(val *PublicationReference) *NullablePublicationReference {
	return &NullablePublicationReference{value: val, isSet: true}
}

func (v NullablePublicationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
