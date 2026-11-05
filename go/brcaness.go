package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the Brcaness type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Brcaness{}

// Brcaness struct for Brcaness
type Brcaness struct {
	ConfidenceRange      BrcanessConfidenceRange `json:"confidenceRange"`
	Id                   string                  `json:"id"`
	Specimen             Reference               `json:"specimen"`
	Value                float32                 `json:"value"`
	Patient              Reference               `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _Brcaness Brcaness

// NewBrcaness instantiates a new Brcaness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrcaness(confidenceRange BrcanessConfidenceRange, id string, specimen Reference, value float32, patient Reference) *Brcaness {
	this := Brcaness{}
	this.ConfidenceRange = confidenceRange
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewBrcanessWithDefaults instantiates a new Brcaness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrcanessWithDefaults() *Brcaness {
	this := Brcaness{}
	return &this
}

// GetConfidenceRange returns the ConfidenceRange field value
func (o *Brcaness) GetConfidenceRange() BrcanessConfidenceRange {
	if o == nil {
		var ret BrcanessConfidenceRange
		return ret
	}

	return o.ConfidenceRange
}

// GetConfidenceRangeOk returns a tuple with the ConfidenceRange field value
// and a boolean to check if the value has been set.
func (o *Brcaness) GetConfidenceRangeOk() (*BrcanessConfidenceRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceRange, true
}

// SetConfidenceRange sets field value
func (o *Brcaness) SetConfidenceRange(v BrcanessConfidenceRange) {
	o.ConfidenceRange = v
}

// GetId returns the Id field value
func (o *Brcaness) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Brcaness) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Brcaness) SetId(v string) {
	o.Id = v
}

// GetSpecimen returns the Specimen field value
func (o *Brcaness) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *Brcaness) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *Brcaness) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *Brcaness) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Brcaness) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Brcaness) SetValue(v float32) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *Brcaness) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Brcaness) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Brcaness) SetPatient(v Reference) {
	o.Patient = v
}

func (o Brcaness) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Brcaness) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["confidenceRange"] = o.ConfidenceRange
	toSerialize["id"] = o.Id
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Brcaness) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"confidenceRange",
		"id",
		"specimen",
		"value",
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

	varBrcaness := _Brcaness{}

	err = json.Unmarshal(data, &varBrcaness)

	if err != nil {
		return err
	}

	*o = Brcaness(varBrcaness)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "confidenceRange")
		delete(additionalProperties, "id")
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "value")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrcaness struct {
	value *Brcaness
	isSet bool
}

func (v NullableBrcaness) Get() *Brcaness {
	return v.value
}

func (v *NullableBrcaness) Set(val *Brcaness) {
	v.value = val
	v.isSet = true
}

func (v NullableBrcaness) IsSet() bool {
	return v.isSet
}

func (v *NullableBrcaness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrcaness(val *Brcaness) *NullableBrcaness {
	return &NullableBrcaness{value: val, isSet: true}
}

func (v NullableBrcaness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrcaness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
