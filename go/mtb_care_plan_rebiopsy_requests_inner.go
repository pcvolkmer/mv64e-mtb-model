package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbCarePlanRebiopsyRequestsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbCarePlanRebiopsyRequestsInner{}

// MtbCarePlanRebiopsyRequestsInner struct for MtbCarePlanRebiopsyRequestsInner
type MtbCarePlanRebiopsyRequestsInner struct {
	Id          string    `json:"id"`
	Patient     Reference `json:"patient"`
	TumorEntity Reference `json:"tumorEntity"`
	IssuedOn    string    `json:"issuedOn"`
}

type _MtbCarePlanRebiopsyRequestsInner MtbCarePlanRebiopsyRequestsInner

// NewMtbCarePlanRebiopsyRequestsInner instantiates a new MtbCarePlanRebiopsyRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbCarePlanRebiopsyRequestsInner(id string, patient Reference, tumorEntity Reference, issuedOn string) *MtbCarePlanRebiopsyRequestsInner {
	this := MtbCarePlanRebiopsyRequestsInner{}
	this.Id = id
	this.Patient = patient
	this.TumorEntity = tumorEntity
	this.IssuedOn = issuedOn
	return &this
}

// NewMtbCarePlanRebiopsyRequestsInnerWithDefaults instantiates a new MtbCarePlanRebiopsyRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbCarePlanRebiopsyRequestsInnerWithDefaults() *MtbCarePlanRebiopsyRequestsInner {
	this := MtbCarePlanRebiopsyRequestsInner{}
	return &this
}

// GetId returns the Id field value
func (o *MtbCarePlanRebiopsyRequestsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanRebiopsyRequestsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbCarePlanRebiopsyRequestsInner) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *MtbCarePlanRebiopsyRequestsInner) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanRebiopsyRequestsInner) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbCarePlanRebiopsyRequestsInner) SetPatient(v Reference) {
	o.Patient = v
}

// GetTumorEntity returns the TumorEntity field value
func (o *MtbCarePlanRebiopsyRequestsInner) GetTumorEntity() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.TumorEntity
}

// GetTumorEntityOk returns a tuple with the TumorEntity field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanRebiopsyRequestsInner) GetTumorEntityOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TumorEntity, true
}

// SetTumorEntity sets field value
func (o *MtbCarePlanRebiopsyRequestsInner) SetTumorEntity(v Reference) {
	o.TumorEntity = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbCarePlanRebiopsyRequestsInner) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanRebiopsyRequestsInner) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbCarePlanRebiopsyRequestsInner) SetIssuedOn(v string) {
	o.IssuedOn = v
}

func (o MtbCarePlanRebiopsyRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbCarePlanRebiopsyRequestsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["tumorEntity"] = o.TumorEntity
	toSerialize["issuedOn"] = o.IssuedOn
	return toSerialize, nil
}

func (o *MtbCarePlanRebiopsyRequestsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"tumorEntity",
		"issuedOn",
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

	varMtbCarePlanRebiopsyRequestsInner := _MtbCarePlanRebiopsyRequestsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbCarePlanRebiopsyRequestsInner)

	if err != nil {
		return err
	}

	*o = MtbCarePlanRebiopsyRequestsInner(varMtbCarePlanRebiopsyRequestsInner)

	return err
}

type NullableMtbCarePlanRebiopsyRequestsInner struct {
	value *MtbCarePlanRebiopsyRequestsInner
	isSet bool
}

func (v NullableMtbCarePlanRebiopsyRequestsInner) Get() *MtbCarePlanRebiopsyRequestsInner {
	return v.value
}

func (v *NullableMtbCarePlanRebiopsyRequestsInner) Set(val *MtbCarePlanRebiopsyRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbCarePlanRebiopsyRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbCarePlanRebiopsyRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbCarePlanRebiopsyRequestsInner(val *MtbCarePlanRebiopsyRequestsInner) *NullableMtbCarePlanRebiopsyRequestsInner {
	return &NullableMtbCarePlanRebiopsyRequestsInner{value: val, isSet: true}
}

func (v NullableMtbCarePlanRebiopsyRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbCarePlanRebiopsyRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
