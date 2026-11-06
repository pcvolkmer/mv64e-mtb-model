package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbEpisodeOfCare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbEpisodeOfCare{}

// MtbEpisodeOfCare struct for MtbEpisodeOfCare
type MtbEpisodeOfCare struct {
	Id        string      `json:"id"`
	Patient   Reference   `json:"patient"`
	Period    PeriodDate  `json:"period"`
	Diagnoses []Reference `json:"diagnoses,omitempty"`
}

type _MtbEpisodeOfCare MtbEpisodeOfCare

// NewMtbEpisodeOfCare instantiates a new MtbEpisodeOfCare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbEpisodeOfCare(id string, patient Reference, period PeriodDate) *MtbEpisodeOfCare {
	this := MtbEpisodeOfCare{}
	this.Id = id
	this.Patient = patient
	this.Period = period
	return &this
}

// NewMtbEpisodeOfCareWithDefaults instantiates a new MtbEpisodeOfCare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbEpisodeOfCareWithDefaults() *MtbEpisodeOfCare {
	this := MtbEpisodeOfCare{}
	return &this
}

// GetId returns the Id field value
func (o *MtbEpisodeOfCare) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbEpisodeOfCare) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbEpisodeOfCare) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *MtbEpisodeOfCare) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbEpisodeOfCare) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbEpisodeOfCare) SetPatient(v Reference) {
	o.Patient = v
}

// GetPeriod returns the Period field value
func (o *MtbEpisodeOfCare) GetPeriod() PeriodDate {
	if o == nil {
		var ret PeriodDate
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *MtbEpisodeOfCare) GetPeriodOk() (*PeriodDate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *MtbEpisodeOfCare) SetPeriod(v PeriodDate) {
	o.Period = v
}

// GetDiagnoses returns the Diagnoses field value if set, zero value otherwise.
func (o *MtbEpisodeOfCare) GetDiagnoses() []Reference {
	if o == nil || IsNil(o.Diagnoses) {
		var ret []Reference
		return ret
	}
	return o.Diagnoses
}

// GetDiagnosesOk returns a tuple with the Diagnoses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbEpisodeOfCare) GetDiagnosesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Diagnoses) {
		return nil, false
	}
	return o.Diagnoses, true
}

// HasDiagnoses returns a boolean if a field has been set.
func (o *MtbEpisodeOfCare) HasDiagnoses() bool {
	if o != nil && !IsNil(o.Diagnoses) {
		return true
	}

	return false
}

// SetDiagnoses gets a reference to the given []Reference and assigns it to the Diagnoses field.
func (o *MtbEpisodeOfCare) SetDiagnoses(v []Reference) {
	o.Diagnoses = v
}

func (o MtbEpisodeOfCare) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbEpisodeOfCare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["period"] = o.Period
	if !IsNil(o.Diagnoses) {
		toSerialize["diagnoses"] = o.Diagnoses
	}
	return toSerialize, nil
}

func (o *MtbEpisodeOfCare) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"period",
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

	varMtbEpisodeOfCare := _MtbEpisodeOfCare{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbEpisodeOfCare)

	if err != nil {
		return err
	}

	*o = MtbEpisodeOfCare(varMtbEpisodeOfCare)

	return err
}

type NullableMtbEpisodeOfCare struct {
	value *MtbEpisodeOfCare
	isSet bool
}

func (v NullableMtbEpisodeOfCare) Get() *MtbEpisodeOfCare {
	return v.value
}

func (v *NullableMtbEpisodeOfCare) Set(val *MtbEpisodeOfCare) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbEpisodeOfCare) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbEpisodeOfCare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbEpisodeOfCare(val *MtbEpisodeOfCare) *NullableMtbEpisodeOfCare {
	return &NullableMtbEpisodeOfCare{value: val, isSet: true}
}

func (v NullableMtbEpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbEpisodeOfCare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
