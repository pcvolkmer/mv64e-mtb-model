package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbCarePlanHistologyReevaluationRequestsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbCarePlanHistologyReevaluationRequestsInner{}

// MtbCarePlanHistologyReevaluationRequestsInner struct for MtbCarePlanHistologyReevaluationRequestsInner
type MtbCarePlanHistologyReevaluationRequestsInner struct {
	Id       string    `json:"id"`
	Patient  Reference `json:"patient"`
	Specimen Reference `json:"specimen"`
	IssuedOn string    `json:"issuedOn"`
}

type _MtbCarePlanHistologyReevaluationRequestsInner MtbCarePlanHistologyReevaluationRequestsInner

// NewMtbCarePlanHistologyReevaluationRequestsInner instantiates a new MtbCarePlanHistologyReevaluationRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbCarePlanHistologyReevaluationRequestsInner(id string, patient Reference, specimen Reference, issuedOn string) *MtbCarePlanHistologyReevaluationRequestsInner {
	this := MtbCarePlanHistologyReevaluationRequestsInner{}
	this.Id = id
	this.Patient = patient
	this.Specimen = specimen
	this.IssuedOn = issuedOn
	return &this
}

// NewMtbCarePlanHistologyReevaluationRequestsInnerWithDefaults instantiates a new MtbCarePlanHistologyReevaluationRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbCarePlanHistologyReevaluationRequestsInnerWithDefaults() *MtbCarePlanHistologyReevaluationRequestsInner {
	this := MtbCarePlanHistologyReevaluationRequestsInner{}
	return &this
}

// GetId returns the Id field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) SetPatient(v Reference) {
	o.Patient = v
}

// GetSpecimen returns the Specimen field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanHistologyReevaluationRequestsInner) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbCarePlanHistologyReevaluationRequestsInner) SetIssuedOn(v string) {
	o.IssuedOn = v
}

func (o MtbCarePlanHistologyReevaluationRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbCarePlanHistologyReevaluationRequestsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["specimen"] = o.Specimen
	toSerialize["issuedOn"] = o.IssuedOn
	return toSerialize, nil
}

func (o *MtbCarePlanHistologyReevaluationRequestsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"specimen",
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

	varMtbCarePlanHistologyReevaluationRequestsInner := _MtbCarePlanHistologyReevaluationRequestsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbCarePlanHistologyReevaluationRequestsInner)

	if err != nil {
		return err
	}

	*o = MtbCarePlanHistologyReevaluationRequestsInner(varMtbCarePlanHistologyReevaluationRequestsInner)

	return err
}

type NullableMtbCarePlanHistologyReevaluationRequestsInner struct {
	value *MtbCarePlanHistologyReevaluationRequestsInner
	isSet bool
}

func (v NullableMtbCarePlanHistologyReevaluationRequestsInner) Get() *MtbCarePlanHistologyReevaluationRequestsInner {
	return v.value
}

func (v *NullableMtbCarePlanHistologyReevaluationRequestsInner) Set(val *MtbCarePlanHistologyReevaluationRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbCarePlanHistologyReevaluationRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbCarePlanHistologyReevaluationRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbCarePlanHistologyReevaluationRequestsInner(val *MtbCarePlanHistologyReevaluationRequestsInner) *NullableMtbCarePlanHistologyReevaluationRequestsInner {
	return &NullableMtbCarePlanHistologyReevaluationRequestsInner{value: val, isSet: true}
}

func (v NullableMtbCarePlanHistologyReevaluationRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbCarePlanHistologyReevaluationRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
