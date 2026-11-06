package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FollowUp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowUp{}

// FollowUp struct for FollowUp
type FollowUp struct {
	Date            string                       `json:"date"`
	Patient         Reference                    `json:"patient"`
	LastContactDate *string                      `json:"lastContactDate,omitempty"`
	PatientStatus   *FollowUpPatientStatusCoding `json:"patientStatus,omitempty"`
}

type _FollowUp FollowUp

// NewFollowUp instantiates a new FollowUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowUp(date string, patient Reference) *FollowUp {
	this := FollowUp{}
	this.Date = date
	this.Patient = patient
	return &this
}

// NewFollowUpWithDefaults instantiates a new FollowUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowUpWithDefaults() *FollowUp {
	this := FollowUp{}
	return &this
}

// GetDate returns the Date field value
func (o *FollowUp) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *FollowUp) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *FollowUp) SetDate(v string) {
	o.Date = v
}

// GetPatient returns the Patient field value
func (o *FollowUp) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *FollowUp) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *FollowUp) SetPatient(v Reference) {
	o.Patient = v
}

// GetLastContactDate returns the LastContactDate field value if set, zero value otherwise.
func (o *FollowUp) GetLastContactDate() string {
	if o == nil || IsNil(o.LastContactDate) {
		var ret string
		return ret
	}
	return *o.LastContactDate
}

// GetLastContactDateOk returns a tuple with the LastContactDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowUp) GetLastContactDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastContactDate) {
		return nil, false
	}
	return o.LastContactDate, true
}

// HasLastContactDate returns a boolean if a field has been set.
func (o *FollowUp) HasLastContactDate() bool {
	if o != nil && !IsNil(o.LastContactDate) {
		return true
	}

	return false
}

// SetLastContactDate gets a reference to the given string and assigns it to the LastContactDate field.
func (o *FollowUp) SetLastContactDate(v string) {
	o.LastContactDate = &v
}

// GetPatientStatus returns the PatientStatus field value if set, zero value otherwise.
func (o *FollowUp) GetPatientStatus() FollowUpPatientStatusCoding {
	if o == nil || IsNil(o.PatientStatus) {
		var ret FollowUpPatientStatusCoding
		return ret
	}
	return *o.PatientStatus
}

// GetPatientStatusOk returns a tuple with the PatientStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowUp) GetPatientStatusOk() (*FollowUpPatientStatusCoding, bool) {
	if o == nil || IsNil(o.PatientStatus) {
		return nil, false
	}
	return o.PatientStatus, true
}

// HasPatientStatus returns a boolean if a field has been set.
func (o *FollowUp) HasPatientStatus() bool {
	if o != nil && !IsNil(o.PatientStatus) {
		return true
	}

	return false
}

// SetPatientStatus gets a reference to the given FollowUpPatientStatusCoding and assigns it to the PatientStatus field.
func (o *FollowUp) SetPatientStatus(v FollowUpPatientStatusCoding) {
	o.PatientStatus = &v
}

func (o FollowUp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowUp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["patient"] = o.Patient
	if !IsNil(o.LastContactDate) {
		toSerialize["lastContactDate"] = o.LastContactDate
	}
	if !IsNil(o.PatientStatus) {
		toSerialize["patientStatus"] = o.PatientStatus
	}
	return toSerialize, nil
}

func (o *FollowUp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
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

	varFollowUp := _FollowUp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowUp)

	if err != nil {
		return err
	}

	*o = FollowUp(varFollowUp)

	return err
}

type NullableFollowUp struct {
	value *FollowUp
	isSet bool
}

func (v NullableFollowUp) Get() *FollowUp {
	return v.value
}

func (v *NullableFollowUp) Set(val *FollowUp) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowUp) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowUp(val *FollowUp) *NullableFollowUp {
	return &NullableFollowUp{value: val, isSet: true}
}

func (v NullableFollowUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
