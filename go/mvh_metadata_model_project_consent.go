package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MvhMetadataModelProjectConsent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MvhMetadataModelProjectConsent{}

// MvhMetadataModelProjectConsent struct for MvhMetadataModelProjectConsent
type MvhMetadataModelProjectConsent struct {
	Version    string                                          `json:"version"`
	Date       *string                                         `json:"date,omitempty"`
	Provisions []MvhMetadataModelProjectConsentProvisionsInner `json:"provisions"`
}

type _MvhMetadataModelProjectConsent MvhMetadataModelProjectConsent

// NewMvhMetadataModelProjectConsent instantiates a new MvhMetadataModelProjectConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMvhMetadataModelProjectConsent(version string, provisions []MvhMetadataModelProjectConsentProvisionsInner) *MvhMetadataModelProjectConsent {
	this := MvhMetadataModelProjectConsent{}
	this.Version = version
	this.Provisions = provisions
	return &this
}

// NewMvhMetadataModelProjectConsentWithDefaults instantiates a new MvhMetadataModelProjectConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMvhMetadataModelProjectConsentWithDefaults() *MvhMetadataModelProjectConsent {
	this := MvhMetadataModelProjectConsent{}
	return &this
}

// GetVersion returns the Version field value
func (o *MvhMetadataModelProjectConsent) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsent) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *MvhMetadataModelProjectConsent) SetVersion(v string) {
	o.Version = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MvhMetadataModelProjectConsent) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsent) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MvhMetadataModelProjectConsent) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *MvhMetadataModelProjectConsent) SetDate(v string) {
	o.Date = &v
}

// GetProvisions returns the Provisions field value
func (o *MvhMetadataModelProjectConsent) GetProvisions() []MvhMetadataModelProjectConsentProvisionsInner {
	if o == nil {
		var ret []MvhMetadataModelProjectConsentProvisionsInner
		return ret
	}

	return o.Provisions
}

// GetProvisionsOk returns a tuple with the Provisions field value
// and a boolean to check if the value has been set.
func (o *MvhMetadataModelProjectConsent) GetProvisionsOk() ([]MvhMetadataModelProjectConsentProvisionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provisions, true
}

// SetProvisions sets field value
func (o *MvhMetadataModelProjectConsent) SetProvisions(v []MvhMetadataModelProjectConsentProvisionsInner) {
	o.Provisions = v
}

func (o MvhMetadataModelProjectConsent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MvhMetadataModelProjectConsent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	toSerialize["provisions"] = o.Provisions
	return toSerialize, nil
}

func (o *MvhMetadataModelProjectConsent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"provisions",
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

	varMvhMetadataModelProjectConsent := _MvhMetadataModelProjectConsent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMvhMetadataModelProjectConsent)

	if err != nil {
		return err
	}

	*o = MvhMetadataModelProjectConsent(varMvhMetadataModelProjectConsent)

	return err
}

type NullableMvhMetadataModelProjectConsent struct {
	value *MvhMetadataModelProjectConsent
	isSet bool
}

func (v NullableMvhMetadataModelProjectConsent) Get() *MvhMetadataModelProjectConsent {
	return v.value
}

func (v *NullableMvhMetadataModelProjectConsent) Set(val *MvhMetadataModelProjectConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableMvhMetadataModelProjectConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableMvhMetadataModelProjectConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMvhMetadataModelProjectConsent(val *MvhMetadataModelProjectConsent) *NullableMvhMetadataModelProjectConsent {
	return &NullableMvhMetadataModelProjectConsent{value: val, isSet: true}
}

func (v NullableMvhMetadataModelProjectConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMvhMetadataModelProjectConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
