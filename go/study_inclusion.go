package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StudyInclusion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyInclusion{}

// StudyInclusion struct for StudyInclusion
type StudyInclusion struct {
	StatusReason *MtbTherapyStatusReasonCoding     `json:"statusReason,omitempty"`
	Medication   []AtcUnregisteredMedicationCoding `json:"medication,omitempty"`
	Reason       *Reference                        `json:"reason,omitempty"`
	Id           string                            `json:"id"`
	RecordedOn   string                            `json:"recordedOn"`
	Status       TherapyStatusCoding               `json:"status"`
	BasedOn      *Reference                        `json:"basedOn,omitempty"`
	Study        []StudyReference                  `json:"study"`
	Period       *PeriodDate                       `json:"period,omitempty"`
	Patient      Reference                         `json:"patient"`
}

type _StudyInclusion StudyInclusion

// NewStudyInclusion instantiates a new StudyInclusion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyInclusion(id string, recordedOn string, status TherapyStatusCoding, study []StudyReference, patient Reference) *StudyInclusion {
	this := StudyInclusion{}
	this.Id = id
	this.RecordedOn = recordedOn
	this.Status = status
	this.Study = study
	this.Patient = patient
	return &this
}

// NewStudyInclusionWithDefaults instantiates a new StudyInclusion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyInclusionWithDefaults() *StudyInclusion {
	this := StudyInclusion{}
	return &this
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *StudyInclusion) GetStatusReason() MtbTherapyStatusReasonCoding {
	if o == nil || IsNil(o.StatusReason) {
		var ret MtbTherapyStatusReasonCoding
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetStatusReasonOk() (*MtbTherapyStatusReasonCoding, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *StudyInclusion) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given MtbTherapyStatusReasonCoding and assigns it to the StatusReason field.
func (o *StudyInclusion) SetStatusReason(v MtbTherapyStatusReasonCoding) {
	o.StatusReason = &v
}

// GetMedication returns the Medication field value if set, zero value otherwise.
func (o *StudyInclusion) GetMedication() []AtcUnregisteredMedicationCoding {
	if o == nil || IsNil(o.Medication) {
		var ret []AtcUnregisteredMedicationCoding
		return ret
	}
	return o.Medication
}

// GetMedicationOk returns a tuple with the Medication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetMedicationOk() ([]AtcUnregisteredMedicationCoding, bool) {
	if o == nil || IsNil(o.Medication) {
		return nil, false
	}
	return o.Medication, true
}

// HasMedication returns a boolean if a field has been set.
func (o *StudyInclusion) HasMedication() bool {
	if o != nil && !IsNil(o.Medication) {
		return true
	}

	return false
}

// SetMedication gets a reference to the given []AtcUnregisteredMedicationCoding and assigns it to the Medication field.
func (o *StudyInclusion) SetMedication(v []AtcUnregisteredMedicationCoding) {
	o.Medication = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *StudyInclusion) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *StudyInclusion) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *StudyInclusion) SetReason(v Reference) {
	o.Reason = &v
}

// GetId returns the Id field value
func (o *StudyInclusion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StudyInclusion) SetId(v string) {
	o.Id = v
}

// GetRecordedOn returns the RecordedOn field value
func (o *StudyInclusion) GetRecordedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordedOn
}

// GetRecordedOnOk returns a tuple with the RecordedOn field value
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetRecordedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedOn, true
}

// SetRecordedOn sets field value
func (o *StudyInclusion) SetRecordedOn(v string) {
	o.RecordedOn = v
}

// GetStatus returns the Status field value
func (o *StudyInclusion) GetStatus() TherapyStatusCoding {
	if o == nil {
		var ret TherapyStatusCoding
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetStatusOk() (*TherapyStatusCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StudyInclusion) SetStatus(v TherapyStatusCoding) {
	o.Status = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *StudyInclusion) GetBasedOn() Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret Reference
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetBasedOnOk() (*Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *StudyInclusion) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given Reference and assigns it to the BasedOn field.
func (o *StudyInclusion) SetBasedOn(v Reference) {
	o.BasedOn = &v
}

// GetStudy returns the Study field value
func (o *StudyInclusion) GetStudy() []StudyReference {
	if o == nil {
		var ret []StudyReference
		return ret
	}

	return o.Study
}

// GetStudyOk returns a tuple with the Study field value
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetStudyOk() ([]StudyReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Study, true
}

// SetStudy sets field value
func (o *StudyInclusion) SetStudy(v []StudyReference) {
	o.Study = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *StudyInclusion) GetPeriod() PeriodDate {
	if o == nil || IsNil(o.Period) {
		var ret PeriodDate
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetPeriodOk() (*PeriodDate, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *StudyInclusion) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given PeriodDate and assigns it to the Period field.
func (o *StudyInclusion) SetPeriod(v PeriodDate) {
	o.Period = &v
}

// GetPatient returns the Patient field value
func (o *StudyInclusion) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *StudyInclusion) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *StudyInclusion) SetPatient(v Reference) {
	o.Patient = v
}

func (o StudyInclusion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyInclusion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	if !IsNil(o.Medication) {
		toSerialize["medication"] = o.Medication
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["id"] = o.Id
	toSerialize["recordedOn"] = o.RecordedOn
	toSerialize["status"] = o.Status
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	toSerialize["study"] = o.Study
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *StudyInclusion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"recordedOn",
		"status",
		"study",
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

	varStudyInclusion := _StudyInclusion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStudyInclusion)

	if err != nil {
		return err
	}

	*o = StudyInclusion(varStudyInclusion)

	return err
}

type NullableStudyInclusion struct {
	value *StudyInclusion
	isSet bool
}

func (v NullableStudyInclusion) Get() *StudyInclusion {
	return v.value
}

func (v *NullableStudyInclusion) Set(val *StudyInclusion) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyInclusion) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyInclusion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyInclusion(val *StudyInclusion) *NullableStudyInclusion {
	return &NullableStudyInclusion{value: val, isSet: true}
}

func (v NullableStudyInclusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyInclusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
