package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MvhMetadataModelProjectConsentProvisionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MvhMetadataModelProjectConsentProvisionsInner{}

// MvhMetadataModelProjectConsentProvisionsInner struct for MvhMetadataModelProjectConsentProvisionsInner
type MvhMetadataModelProjectConsentProvisionsInner struct {
	Date    string                     `json:"date"`
	Purpose ModelProjectConsentPurpose `json:"purpose"`
	Type    ConsentProvisionType       `json:"type"`
}

type _MvhMetadataModelProjectConsentProvisionsInner MvhMetadataModelProjectConsentProvisionsInner

// NewMvhMetadataModelProjectConsentProvisionsInner instantiates a new MvhMetadataModelProjectConsentProvisionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMvhMetadataModelProjectConsentProvisionsInner(date string, purpose ModelProjectConsentPurpose, type_ ConsentProvisionType) *MvhMetadataModelProjectConsentProvisionsInner {
	this := MvhMetadataModelProjectConsentProvisionsInner{}
	this.Date = date
	this.Purpose = purpose
	this.Type = type_
	return &this
}

// NewMvhMetadataModelProjectConsentProvisionsInnerWithDefaults instantiates a new MvhMetadataModelProjectConsentProvisionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMvhMetadataModelProjectConsentProvisionsInnerWithDefaults() *MvhMetadataModelProjectConsentProvisionsInner {
	this := MvhMetadataModelProjectConsentProvisionsInner{}
	return &this
}

// GetDate returns the Date field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) SetDate(v string) {
	o.Date = v
}

// GetPurpose returns the Purpose field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetPurpose() ModelProjectConsentPurpose {
	if o == nil {
		var ret ModelProjectConsentPurpose
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetPurposeOk() (*ModelProjectConsentPurpose, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) SetPurpose(v ModelProjectConsentPurpose) {
	o.Purpose = v
}

// GetType returns the Type field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetType() ConsentProvisionType {
	if o == nil {
		var ret ConsentProvisionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsentProvisionsInner) GetTypeOk() (*ConsentProvisionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MvhMetadataModelProjectConsentProvisionsInner) SetType(v ConsentProvisionType) {
	o.Type = v
}

func (o MvhMetadataModelProjectConsentProvisionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MvhMetadataModelProjectConsentProvisionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["purpose"] = o.Purpose
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *MvhMetadataModelProjectConsentProvisionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"purpose",
		"type",
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

	varMvhMetadataModelProjectConsentProvisionsInner := _MvhMetadataModelProjectConsentProvisionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMvhMetadataModelProjectConsentProvisionsInner)

	if err != nil {
		return err
	}

	*o = MvhMetadataModelProjectConsentProvisionsInner(varMvhMetadataModelProjectConsentProvisionsInner)

	return err
}

type NullableMvhMetadataModelProjectConsentProvisionsInner struct {
	value *MvhMetadataModelProjectConsentProvisionsInner
	isSet bool
}

func (v NullableMvhMetadataModelProjectConsentProvisionsInner) Get() *MvhMetadataModelProjectConsentProvisionsInner {
	return v.value
}

func (v *NullableMvhMetadataModelProjectConsentProvisionsInner) Set(val *MvhMetadataModelProjectConsentProvisionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMvhMetadataModelProjectConsentProvisionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMvhMetadataModelProjectConsentProvisionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMvhMetadataModelProjectConsentProvisionsInner(val *MvhMetadataModelProjectConsentProvisionsInner) *NullableMvhMetadataModelProjectConsentProvisionsInner {
	return &NullableMvhMetadataModelProjectConsentProvisionsInner{value: val, isSet: true}
}

func (v NullableMvhMetadataModelProjectConsentProvisionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMvhMetadataModelProjectConsentProvisionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
