package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the HrdScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrdScore{}

// HrdScore struct for HrdScore
type HrdScore struct {
	Interpretation       *HrdScoreInterpretationCoding `json:"interpretation,omitempty"`
	Components           HrdScoreComponents            `json:"components"`
	Id                   string                        `json:"id"`
	Specimen             Reference                     `json:"specimen"`
	Value                float32                       `json:"value"`
	Patient              Reference                     `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _HrdScore HrdScore

// NewHrdScore instantiates a new HrdScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrdScore(components HrdScoreComponents, id string, specimen Reference, value float32, patient Reference) *HrdScore {
	this := HrdScore{}
	this.Components = components
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewHrdScoreWithDefaults instantiates a new HrdScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrdScoreWithDefaults() *HrdScore {
	this := HrdScore{}
	return &this
}

// GetInterpretation returns the Interpretation field value if set, zero value otherwise.
func (o *HrdScore) GetInterpretation() HrdScoreInterpretationCoding {
	if o == nil || IsNil(o.Interpretation) {
		var ret HrdScoreInterpretationCoding
		return ret
	}
	return *o.Interpretation
}

// GetInterpretationOk returns a tuple with the Interpretation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrdScore) GetInterpretationOk() (*HrdScoreInterpretationCoding, bool) {
	if o == nil || IsNil(o.Interpretation) {
		return nil, false
	}
	return o.Interpretation, true
}

// HasInterpretation returns a boolean if a field has been set.
func (o *HrdScore) HasInterpretation() bool {
	if o != nil && !IsNil(o.Interpretation) {
		return true
	}

	return false
}

// SetInterpretation gets a reference to the given HrdScoreInterpretationCoding and assigns it to the Interpretation field.
func (o *HrdScore) SetInterpretation(v HrdScoreInterpretationCoding) {
	o.Interpretation = &v
}

// GetComponents returns the Components field value
func (o *HrdScore) GetComponents() HrdScoreComponents {
	if o == nil {
		var ret HrdScoreComponents
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *HrdScore) GetComponentsOk() (*HrdScoreComponents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *HrdScore) SetComponents(v HrdScoreComponents) {
	o.Components = v
}

// GetId returns the Id field value
func (o *HrdScore) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HrdScore) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HrdScore) SetId(v string) {
	o.Id = v
}

// GetSpecimen returns the Specimen field value
func (o *HrdScore) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *HrdScore) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *HrdScore) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *HrdScore) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HrdScore) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HrdScore) SetValue(v float32) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *HrdScore) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *HrdScore) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *HrdScore) SetPatient(v Reference) {
	o.Patient = v
}

func (o HrdScore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrdScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interpretation) {
		toSerialize["interpretation"] = o.Interpretation
	}
	toSerialize["components"] = o.Components
	toSerialize["id"] = o.Id
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HrdScore) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"components",
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

	varHrdScore := _HrdScore{}

	err = json.Unmarshal(data, &varHrdScore)

	if err != nil {
		return err
	}

	*o = HrdScore(varHrdScore)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interpretation")
		delete(additionalProperties, "components")
		delete(additionalProperties, "id")
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "value")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrdScore struct {
	value *HrdScore
	isSet bool
}

func (v NullableHrdScore) Get() *HrdScore {
	return v.value
}

func (v *NullableHrdScore) Set(val *HrdScore) {
	v.value = val
	v.isSet = true
}

func (v NullableHrdScore) IsSet() bool {
	return v.isSet
}

func (v *NullableHrdScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrdScore(val *HrdScore) *NullableHrdScore {
	return &NullableHrdScore{value: val, isSet: true}
}

func (v NullableHrdScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrdScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
