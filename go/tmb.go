package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the Tmb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tmb{}

// Tmb struct for Tmb
type Tmb struct {
	Interpretation       *TmbInterpretationCoding `json:"interpretation,omitempty"`
	Id                   string                   `json:"id"`
	Specimen             Reference                `json:"specimen"`
	Value                TmbResult                `json:"value"`
	Patient              Reference                `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _Tmb Tmb

// NewTmb instantiates a new Tmb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmb(id string, specimen Reference, value TmbResult, patient Reference) *Tmb {
	this := Tmb{}
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewTmbWithDefaults instantiates a new Tmb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmbWithDefaults() *Tmb {
	this := Tmb{}
	return &this
}

// GetInterpretation returns the Interpretation field value if set, zero value otherwise.
func (o *Tmb) GetInterpretation() TmbInterpretationCoding {
	if o == nil || IsNil(o.Interpretation) {
		var ret TmbInterpretationCoding
		return ret
	}
	return *o.Interpretation
}

// GetInterpretationOk returns a tuple with the Interpretation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tmb) GetInterpretationOk() (*TmbInterpretationCoding, bool) {
	if o == nil || IsNil(o.Interpretation) {
		return nil, false
	}
	return o.Interpretation, true
}

// HasInterpretation returns a boolean if a field has been set.
func (o *Tmb) HasInterpretation() bool {
	if o != nil && !IsNil(o.Interpretation) {
		return true
	}

	return false
}

// SetInterpretation gets a reference to the given TmbInterpretationCoding and assigns it to the Interpretation field.
func (o *Tmb) SetInterpretation(v TmbInterpretationCoding) {
	o.Interpretation = &v
}

// GetId returns the Id field value
func (o *Tmb) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tmb) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tmb) SetId(v string) {
	o.Id = v
}

// GetSpecimen returns the Specimen field value
func (o *Tmb) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *Tmb) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *Tmb) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *Tmb) GetValue() TmbResult {
	if o == nil {
		var ret TmbResult
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Tmb) GetValueOk() (*TmbResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Tmb) SetValue(v TmbResult) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *Tmb) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Tmb) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Tmb) SetPatient(v Reference) {
	o.Patient = v
}

func (o Tmb) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tmb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interpretation) {
		toSerialize["interpretation"] = o.Interpretation
	}
	toSerialize["id"] = o.Id
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tmb) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varTmb := _Tmb{}

	err = json.Unmarshal(data, &varTmb)

	if err != nil {
		return err
	}

	*o = Tmb(varTmb)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interpretation")
		delete(additionalProperties, "id")
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "value")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTmb struct {
	value *Tmb
	isSet bool
}

func (v NullableTmb) Get() *Tmb {
	return v.value
}

func (v *NullableTmb) Set(val *Tmb) {
	v.value = val
	v.isSet = true
}

func (v NullableTmb) IsSet() bool {
	return v.isSet
}

func (v *NullableTmb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmb(val *Tmb) *NullableTmb {
	return &NullableTmb{value: val, isSet: true}
}

func (v NullableTmb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
