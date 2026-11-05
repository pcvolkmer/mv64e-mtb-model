package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PerformanceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceStatus{}

// PerformanceStatus struct for PerformanceStatus
type PerformanceStatus struct {
	Id            string     `json:"id"`
	Patient       Reference  `json:"patient"`
	EffectiveDate string     `json:"effectiveDate"`
	Value         EcogCoding `json:"value"`
}

type _PerformanceStatus PerformanceStatus

// NewPerformanceStatus instantiates a new PerformanceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceStatus(id string, patient Reference, effectiveDate string, value EcogCoding) *PerformanceStatus {
	this := PerformanceStatus{}
	this.Id = id
	this.Patient = patient
	this.EffectiveDate = effectiveDate
	this.Value = value
	return &this
}

// NewPerformanceStatusWithDefaults instantiates a new PerformanceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceStatusWithDefaults() *PerformanceStatus {
	this := PerformanceStatus{}
	return &this
}

// GetId returns the Id field value
func (o *PerformanceStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PerformanceStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PerformanceStatus) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *PerformanceStatus) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *PerformanceStatus) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *PerformanceStatus) SetPatient(v Reference) {
	o.Patient = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PerformanceStatus) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PerformanceStatus) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PerformanceStatus) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetValue returns the Value field value
func (o *PerformanceStatus) GetValue() EcogCoding {
	if o == nil {
		var ret EcogCoding
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PerformanceStatus) GetValueOk() (*EcogCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PerformanceStatus) SetValue(v EcogCoding) {
	o.Value = v
}

func (o PerformanceStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["effectiveDate"] = o.EffectiveDate
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *PerformanceStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"effectiveDate",
		"value",
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

	varPerformanceStatus := _PerformanceStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPerformanceStatus)

	if err != nil {
		return err
	}

	*o = PerformanceStatus(varPerformanceStatus)

	return err
}

type NullablePerformanceStatus struct {
	value *PerformanceStatus
	isSet bool
}

func (v NullablePerformanceStatus) Get() *PerformanceStatus {
	return v.value
}

func (v *NullablePerformanceStatus) Set(val *PerformanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceStatus(val *PerformanceStatus) *NullablePerformanceStatus {
	return &NullablePerformanceStatus{value: val, isSet: true}
}

func (v NullablePerformanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
