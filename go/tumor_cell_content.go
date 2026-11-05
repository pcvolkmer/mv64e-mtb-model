package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorCellContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorCellContent{}

// TumorCellContent struct for TumorCellContent
type TumorCellContent struct {
	Method               TumorCellContentMethodCoding `json:"method"`
	Id                   string                       `json:"id"`
	Specimen             Reference                    `json:"specimen"`
	Value                float32                      `json:"value"`
	Patient              Reference                    `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _TumorCellContent TumorCellContent

// NewTumorCellContent instantiates a new TumorCellContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorCellContent(method TumorCellContentMethodCoding, id string, specimen Reference, value float32, patient Reference) *TumorCellContent {
	this := TumorCellContent{}
	this.Method = method
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewTumorCellContentWithDefaults instantiates a new TumorCellContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorCellContentWithDefaults() *TumorCellContent {
	this := TumorCellContent{}
	return &this
}

// GetMethod returns the Method field value
func (o *TumorCellContent) GetMethod() TumorCellContentMethodCoding {
	if o == nil {
		var ret TumorCellContentMethodCoding
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TumorCellContent) GetMethodOk() (*TumorCellContentMethodCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TumorCellContent) SetMethod(v TumorCellContentMethodCoding) {
	o.Method = v
}

// GetId returns the Id field value
func (o *TumorCellContent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TumorCellContent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TumorCellContent) SetId(v string) {
	o.Id = v
}

// GetSpecimen returns the Specimen field value
func (o *TumorCellContent) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *TumorCellContent) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *TumorCellContent) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *TumorCellContent) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TumorCellContent) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TumorCellContent) SetValue(v float32) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *TumorCellContent) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *TumorCellContent) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *TumorCellContent) SetPatient(v Reference) {
	o.Patient = v
}

func (o TumorCellContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorCellContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["id"] = o.Id
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TumorCellContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
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

	varTumorCellContent := _TumorCellContent{}

	err = json.Unmarshal(data, &varTumorCellContent)

	if err != nil {
		return err
	}

	*o = TumorCellContent(varTumorCellContent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "id")
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "value")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTumorCellContent struct {
	value *TumorCellContent
	isSet bool
}

func (v NullableTumorCellContent) Get() *TumorCellContent {
	return v.value
}

func (v *NullableTumorCellContent) Set(val *TumorCellContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorCellContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorCellContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorCellContent(val *TumorCellContent) *NullableTumorCellContent {
	return &NullableTumorCellContent{value: val, isSet: true}
}

func (v NullableTumorCellContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorCellContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
