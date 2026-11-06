package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FamilyMemberHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FamilyMemberHistory{}

// FamilyMemberHistory struct for FamilyMemberHistory
type FamilyMemberHistory struct {
	Id           string                                    `json:"id"`
	Patient      Reference                                 `json:"patient"`
	Relationship FamilyMemberHistoryRelationshipTypeCoding `json:"relationship"`
}

type _FamilyMemberHistory FamilyMemberHistory

// NewFamilyMemberHistory instantiates a new FamilyMemberHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFamilyMemberHistory(id string, patient Reference, relationship FamilyMemberHistoryRelationshipTypeCoding) *FamilyMemberHistory {
	this := FamilyMemberHistory{}
	this.Id = id
	this.Patient = patient
	this.Relationship = relationship
	return &this
}

// NewFamilyMemberHistoryWithDefaults instantiates a new FamilyMemberHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFamilyMemberHistoryWithDefaults() *FamilyMemberHistory {
	this := FamilyMemberHistory{}
	return &this
}

// GetId returns the Id field value
func (o *FamilyMemberHistory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FamilyMemberHistory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FamilyMemberHistory) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *FamilyMemberHistory) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *FamilyMemberHistory) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *FamilyMemberHistory) SetPatient(v Reference) {
	o.Patient = v
}

// GetRelationship returns the Relationship field value
func (o *FamilyMemberHistory) GetRelationship() FamilyMemberHistoryRelationshipTypeCoding {
	if o == nil {
		var ret FamilyMemberHistoryRelationshipTypeCoding
		return ret
	}

	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value
// and a boolean to check if the value has been set.
func (o *FamilyMemberHistory) GetRelationshipOk() (*FamilyMemberHistoryRelationshipTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// SetRelationship sets field value
func (o *FamilyMemberHistory) SetRelationship(v FamilyMemberHistoryRelationshipTypeCoding) {
	o.Relationship = v
}

func (o FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FamilyMemberHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["relationship"] = o.Relationship
	return toSerialize, nil
}

func (o *FamilyMemberHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"relationship",
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

	varFamilyMemberHistory := _FamilyMemberHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFamilyMemberHistory)

	if err != nil {
		return err
	}

	*o = FamilyMemberHistory(varFamilyMemberHistory)

	return err
}

type NullableFamilyMemberHistory struct {
	value *FamilyMemberHistory
	isSet bool
}

func (v NullableFamilyMemberHistory) Get() *FamilyMemberHistory {
	return v.value
}

func (v *NullableFamilyMemberHistory) Set(val *FamilyMemberHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableFamilyMemberHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableFamilyMemberHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFamilyMemberHistory(val *FamilyMemberHistory) *NullableFamilyMemberHistory {
	return &NullableFamilyMemberHistory{value: val, isSet: true}
}

func (v NullableFamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFamilyMemberHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
