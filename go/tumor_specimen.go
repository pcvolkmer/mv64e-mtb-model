package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorSpecimen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorSpecimen{}

// TumorSpecimen struct for TumorSpecimen
type TumorSpecimen struct {
	Diagnosis            Reference                `json:"diagnosis"`
	Collection           *TumorSpecimenCollection `json:"collection,omitempty"`
	Id                   string                   `json:"id"`
	Type                 TumorSpecimenTypeCoding  `json:"type"`
	Patient              Reference                `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _TumorSpecimen TumorSpecimen

// NewTumorSpecimen instantiates a new TumorSpecimen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorSpecimen(diagnosis Reference, id string, type_ TumorSpecimenTypeCoding, patient Reference) *TumorSpecimen {
	this := TumorSpecimen{}
	this.Diagnosis = diagnosis
	this.Id = id
	this.Type = type_
	this.Patient = patient
	return &this
}

// NewTumorSpecimenWithDefaults instantiates a new TumorSpecimen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorSpecimenWithDefaults() *TumorSpecimen {
	this := TumorSpecimen{}
	return &this
}

// GetDiagnosis returns the Diagnosis field value
func (o *TumorSpecimen) GetDiagnosis() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Diagnosis
}

// GetDiagnosisOk returns a tuple with the Diagnosis field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimen) GetDiagnosisOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diagnosis, true
}

// SetDiagnosis sets field value
func (o *TumorSpecimen) SetDiagnosis(v Reference) {
	o.Diagnosis = v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *TumorSpecimen) GetCollection() TumorSpecimenCollection {
	if o == nil || IsNil(o.Collection) {
		var ret TumorSpecimenCollection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorSpecimen) GetCollectionOk() (*TumorSpecimenCollection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *TumorSpecimen) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given TumorSpecimenCollection and assigns it to the Collection field.
func (o *TumorSpecimen) SetCollection(v TumorSpecimenCollection) {
	o.Collection = &v
}

// GetId returns the Id field value
func (o *TumorSpecimen) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimen) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TumorSpecimen) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TumorSpecimen) GetType() TumorSpecimenTypeCoding {
	if o == nil {
		var ret TumorSpecimenTypeCoding
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimen) GetTypeOk() (*TumorSpecimenTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TumorSpecimen) SetType(v TumorSpecimenTypeCoding) {
	o.Type = v
}

// GetPatient returns the Patient field value
func (o *TumorSpecimen) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimen) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *TumorSpecimen) SetPatient(v Reference) {
	o.Patient = v
}

func (o TumorSpecimen) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorSpecimen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["diagnosis"] = o.Diagnosis
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TumorSpecimen) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diagnosis",
		"id",
		"type",
		"patient",
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

	varTumorSpecimen := _TumorSpecimen{}

	err = json.Unmarshal(data, &varTumorSpecimen)

	if err != nil {
		return err
	}

	*o = TumorSpecimen(varTumorSpecimen)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "diagnosis")
		delete(additionalProperties, "collection")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTumorSpecimen struct {
	value *TumorSpecimen
	isSet bool
}

func (v NullableTumorSpecimen) Get() *TumorSpecimen {
	return v.value
}

func (v *NullableTumorSpecimen) Set(val *TumorSpecimen) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorSpecimen) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorSpecimen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorSpecimen(val *TumorSpecimen) *NullableTumorSpecimen {
	return &NullableTumorSpecimen{value: val, isSet: true}
}

func (v NullableTumorSpecimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorSpecimen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
