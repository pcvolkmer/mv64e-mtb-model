package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Msi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Msi{}

// Msi struct for Msi
type Msi struct {
	Method         MsiMethodCoding         `json:"method"`
	Interpretation MsiInterpretationCoding `json:"interpretation"`
	Id             string                  `json:"id"`
	Specimen       Reference               `json:"specimen"`
	Value          float32                 `json:"value"`
	Patient        Reference               `json:"patient"`
}

type _Msi Msi

// NewMsi instantiates a new Msi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsi(method MsiMethodCoding, interpretation MsiInterpretationCoding, id string, specimen Reference, value float32, patient Reference) *Msi {
	this := Msi{}
	this.Method = method
	this.Interpretation = interpretation
	this.Id = id
	this.Specimen = specimen
	this.Value = value
	this.Patient = patient
	return &this
}

// NewMsiWithDefaults instantiates a new Msi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsiWithDefaults() *Msi {
	this := Msi{}
	return &this
}

// GetMethod returns the Method field value
func (o *Msi) GetMethod() MsiMethodCoding {
	if o == nil {
		var ret MsiMethodCoding
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Msi) GetMethodOk() (*MsiMethodCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *Msi) SetMethod(v MsiMethodCoding) {
	o.Method = v
}

// GetInterpretation returns the Interpretation field value
func (o *Msi) GetInterpretation() MsiInterpretationCoding {
	if o == nil {
		var ret MsiInterpretationCoding
		return ret
	}

	return o.Interpretation
}

// GetInterpretationOk returns a tuple with the Interpretation field value
// and a boolean to check if the value has been set.
func (o *Msi) GetInterpretationOk() (*MsiInterpretationCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interpretation, true
}

// SetInterpretation sets field value
func (o *Msi) SetInterpretation(v MsiInterpretationCoding) {
	o.Interpretation = v
}

// GetId returns the Id field value
func (o *Msi) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Msi) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Msi) SetId(v string) {
	o.Id = v
}

// GetSpecimen returns the Specimen field value
func (o *Msi) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *Msi) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *Msi) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetValue returns the Value field value
func (o *Msi) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Msi) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Msi) SetValue(v float32) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *Msi) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Msi) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Msi) SetPatient(v Reference) {
	o.Patient = v
}

func (o Msi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Msi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["interpretation"] = o.Interpretation
	toSerialize["id"] = o.Id
	toSerialize["specimen"] = o.Specimen
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *Msi) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"interpretation",
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

	varMsi := _Msi{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMsi)

	if err != nil {
		return err
	}

	*o = Msi(varMsi)

	return err
}

type NullableMsi struct {
	value *Msi
	isSet bool
}

func (v NullableMsi) Get() *Msi {
	return v.value
}

func (v *NullableMsi) Set(val *Msi) {
	v.value = val
	v.isSet = true
}

func (v NullableMsi) IsSet() bool {
	return v.isSet
}

func (v *NullableMsi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsi(val *Msi) *NullableMsi {
	return &NullableMsi{value: val, isSet: true}
}

func (v NullableMsi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
