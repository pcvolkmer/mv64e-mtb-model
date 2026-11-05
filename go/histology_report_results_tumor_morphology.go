package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HistologyReportResultsTumorMorphology type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistologyReportResultsTumorMorphology{}

// HistologyReportResultsTumorMorphology struct for HistologyReportResultsTumorMorphology
type HistologyReportResultsTumorMorphology struct {
	Id       string    `json:"id"`
	Note     *string   `json:"note,omitempty"`
	Specimen Reference `json:"specimen"`
	Value    Coding    `json:"value"`
	Patient  Reference `json:"patient"`
}

type _HistologyReportResultsTumorMorphology HistologyReportResultsTumorMorphology

// NewHistologyReportResultsTumorMorphology instantiates a new HistologyReportResultsTumorMorphology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistologyReportResultsTumorMorphology(id string, specimen Reference, value Coding, patient Reference) *HistologyReportResultsTumorMorphology {
	this := HistologyReportResultsTumorMorphology{}
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewHistologyReportResultsTumorMorphologyWithDefaults instantiates a new HistologyReportResultsTumorMorphology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistologyReportResultsTumorMorphologyWithDefaults() *HistologyReportResultsTumorMorphology {
	this := HistologyReportResultsTumorMorphology{}
	return &this
}

// GetId returns the Id field value
func (o *HistologyReportResultsTumorMorphology) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistologyReportResultsTumorMorphology) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistologyReportResultsTumorMorphology) SetId(v string) {
	o.Id = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *HistologyReportResultsTumorMorphology) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistologyReportResultsTumorMorphology) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *HistologyReportResultsTumorMorphology) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *HistologyReportResultsTumorMorphology) SetNote(v string) {
	o.Note = &v
}

// GetSpecimen returns the Specimen field value
func (o *HistologyReportResultsTumorMorphology) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *HistologyReportResultsTumorMorphology) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *HistologyReportResultsTumorMorphology) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *HistologyReportResultsTumorMorphology) GetValue() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HistologyReportResultsTumorMorphology) GetValueOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HistologyReportResultsTumorMorphology) SetValue(v Coding) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *HistologyReportResultsTumorMorphology) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *HistologyReportResultsTumorMorphology) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *HistologyReportResultsTumorMorphology) SetPatient(v Reference) {
	o.Patient = v
}

func (o HistologyReportResultsTumorMorphology) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistologyReportResultsTumorMorphology) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *HistologyReportResultsTumorMorphology) UnmarshalJSON(data []byte) (err error) {
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

	varHistologyReportResultsTumorMorphology := _HistologyReportResultsTumorMorphology{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHistologyReportResultsTumorMorphology)

	if err != nil {
		return err
	}

	*o = HistologyReportResultsTumorMorphology(varHistologyReportResultsTumorMorphology)

	return err
}

type NullableHistologyReportResultsTumorMorphology struct {
	value *HistologyReportResultsTumorMorphology
	isSet bool
}

func (v NullableHistologyReportResultsTumorMorphology) Get() *HistologyReportResultsTumorMorphology {
	return v.value
}

func (v *NullableHistologyReportResultsTumorMorphology) Set(val *HistologyReportResultsTumorMorphology) {
	v.value = val
	v.isSet = true
}

func (v NullableHistologyReportResultsTumorMorphology) IsSet() bool {
	return v.isSet
}

func (v *NullableHistologyReportResultsTumorMorphology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistologyReportResultsTumorMorphology(val *HistologyReportResultsTumorMorphology) *NullableHistologyReportResultsTumorMorphology {
	return &NullableHistologyReportResultsTumorMorphology{value: val, isSet: true}
}

func (v NullableHistologyReportResultsTumorMorphology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistologyReportResultsTumorMorphology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
