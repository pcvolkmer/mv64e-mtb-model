package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbSystemicTherapy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbSystemicTherapy{}

// MtbSystemicTherapy struct for MtbSystemicTherapy
type MtbSystemicTherapy struct {
	StatusReason                    *MtbTherapyStatusReasonCoding                            `json:"statusReason,omitempty"`
	Medication                      []AtcUnregisteredMedicationCoding                        `json:"medication,omitempty"`
	Intent                          *MtbTherapyIntentCoding                                  `json:"intent,omitempty"`
	TherapyLine                     *int32                                                   `json:"therapyLine,omitempty"`
	Reason                          *Reference                                               `json:"reason,omitempty"`
	RecommendationFulfillmentStatus *MtbSystemicTherapyRecommendationFulfillmentStatusCoding `json:"recommendationFulfillmentStatus,omitempty"`
	Id                              string                                                   `json:"id"`
	RecordedOn                      string                                                   `json:"recordedOn"`
	Status                          TherapyStatusCoding                                      `json:"status"`
	Notes                           []string                                                 `json:"notes,omitempty"`
	Dosage                          *MtbSystemicTherapyDosageDensityCoding                   `json:"dosage,omitempty"`
	Period                          *PeriodDate                                              `json:"period,omitempty"`
	Patient                         Reference                                                `json:"patient"`
	Category                        *MtbSystemicTherapyCategoryCoding                        `json:"category,omitempty"`
	BasedOn                         *Reference                                               `json:"basedOn,omitempty"`
}

type _MtbSystemicTherapy MtbSystemicTherapy

// NewMtbSystemicTherapy instantiates a new MtbSystemicTherapy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbSystemicTherapy(id string, recordedOn string, status TherapyStatusCoding, patient Reference) *MtbSystemicTherapy {
	this := MtbSystemicTherapy{}
	this.Id = id
	this.RecordedOn = recordedOn
	this.Status = status
	this.Patient = patient
	return &this
}

// NewMtbSystemicTherapyWithDefaults instantiates a new MtbSystemicTherapy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbSystemicTherapyWithDefaults() *MtbSystemicTherapy {
	this := MtbSystemicTherapy{}
	return &this
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetStatusReason() MtbTherapyStatusReasonCoding {
	if o == nil || IsNil(o.StatusReason) {
		var ret MtbTherapyStatusReasonCoding
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetStatusReasonOk() (*MtbTherapyStatusReasonCoding, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given MtbTherapyStatusReasonCoding and assigns it to the StatusReason field.
func (o *MtbSystemicTherapy) SetStatusReason(v MtbTherapyStatusReasonCoding) {
	o.StatusReason = &v
}

// GetMedication returns the Medication field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetMedication() []AtcUnregisteredMedicationCoding {
	if o == nil || IsNil(o.Medication) {
		var ret []AtcUnregisteredMedicationCoding
		return ret
	}
	return o.Medication
}

// GetMedicationOk returns a tuple with the Medication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetMedicationOk() ([]AtcUnregisteredMedicationCoding, bool) {
	if o == nil || IsNil(o.Medication) {
		return nil, false
	}
	return o.Medication, true
}

// HasMedication returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasMedication() bool {
	if o != nil && !IsNil(o.Medication) {
		return true
	}

	return false
}

// SetMedication gets a reference to the given []AtcUnregisteredMedicationCoding and assigns it to the Medication field.
func (o *MtbSystemicTherapy) SetMedication(v []AtcUnregisteredMedicationCoding) {
	o.Medication = v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetIntent() MtbTherapyIntentCoding {
	if o == nil || IsNil(o.Intent) {
		var ret MtbTherapyIntentCoding
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetIntentOk() (*MtbTherapyIntentCoding, bool) {
	if o == nil || IsNil(o.Intent) {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasIntent() bool {
	if o != nil && !IsNil(o.Intent) {
		return true
	}

	return false
}

// SetIntent gets a reference to the given MtbTherapyIntentCoding and assigns it to the Intent field.
func (o *MtbSystemicTherapy) SetIntent(v MtbTherapyIntentCoding) {
	o.Intent = &v
}

// GetTherapyLine returns the TherapyLine field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetTherapyLine() int32 {
	if o == nil || IsNil(o.TherapyLine) {
		var ret int32
		return ret
	}
	return *o.TherapyLine
}

// GetTherapyLineOk returns a tuple with the TherapyLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetTherapyLineOk() (*int32, bool) {
	if o == nil || IsNil(o.TherapyLine) {
		return nil, false
	}
	return o.TherapyLine, true
}

// HasTherapyLine returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasTherapyLine() bool {
	if o != nil && !IsNil(o.TherapyLine) {
		return true
	}

	return false
}

// SetTherapyLine gets a reference to the given int32 and assigns it to the TherapyLine field.
func (o *MtbSystemicTherapy) SetTherapyLine(v int32) {
	o.TherapyLine = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *MtbSystemicTherapy) SetReason(v Reference) {
	o.Reason = &v
}

// GetRecommendationFulfillmentStatus returns the RecommendationFulfillmentStatus field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetRecommendationFulfillmentStatus() MtbSystemicTherapyRecommendationFulfillmentStatusCoding {
	if o == nil || IsNil(o.RecommendationFulfillmentStatus) {
		var ret MtbSystemicTherapyRecommendationFulfillmentStatusCoding
		return ret
	}
	return *o.RecommendationFulfillmentStatus
}

// GetRecommendationFulfillmentStatusOk returns a tuple with the RecommendationFulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetRecommendationFulfillmentStatusOk() (*MtbSystemicTherapyRecommendationFulfillmentStatusCoding, bool) {
	if o == nil || IsNil(o.RecommendationFulfillmentStatus) {
		return nil, false
	}
	return o.RecommendationFulfillmentStatus, true
}

// HasRecommendationFulfillmentStatus returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasRecommendationFulfillmentStatus() bool {
	if o != nil && !IsNil(o.RecommendationFulfillmentStatus) {
		return true
	}

	return false
}

// SetRecommendationFulfillmentStatus gets a reference to the given MtbSystemicTherapyRecommendationFulfillmentStatusCoding and assigns it to the RecommendationFulfillmentStatus field.
func (o *MtbSystemicTherapy) SetRecommendationFulfillmentStatus(v MtbSystemicTherapyRecommendationFulfillmentStatusCoding) {
	o.RecommendationFulfillmentStatus = &v
}

// GetId returns the Id field value
func (o *MtbSystemicTherapy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbSystemicTherapy) SetId(v string) {
	o.Id = v
}

// GetRecordedOn returns the RecordedOn field value
func (o *MtbSystemicTherapy) GetRecordedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordedOn
}

// GetRecordedOnOk returns a tuple with the RecordedOn field value
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetRecordedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedOn, true
}

// SetRecordedOn sets field value
func (o *MtbSystemicTherapy) SetRecordedOn(v string) {
	o.RecordedOn = v
}

// GetStatus returns the Status field value
func (o *MtbSystemicTherapy) GetStatus() TherapyStatusCoding {
	if o == nil {
		var ret TherapyStatusCoding
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetStatusOk() (*TherapyStatusCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MtbSystemicTherapy) SetStatus(v TherapyStatusCoding) {
	o.Status = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *MtbSystemicTherapy) SetNotes(v []string) {
	o.Notes = v
}

// GetDosage returns the Dosage field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetDosage() MtbSystemicTherapyDosageDensityCoding {
	if o == nil || IsNil(o.Dosage) {
		var ret MtbSystemicTherapyDosageDensityCoding
		return ret
	}
	return *o.Dosage
}

// GetDosageOk returns a tuple with the Dosage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetDosageOk() (*MtbSystemicTherapyDosageDensityCoding, bool) {
	if o == nil || IsNil(o.Dosage) {
		return nil, false
	}
	return o.Dosage, true
}

// HasDosage returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasDosage() bool {
	if o != nil && !IsNil(o.Dosage) {
		return true
	}

	return false
}

// SetDosage gets a reference to the given MtbSystemicTherapyDosageDensityCoding and assigns it to the Dosage field.
func (o *MtbSystemicTherapy) SetDosage(v MtbSystemicTherapyDosageDensityCoding) {
	o.Dosage = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetPeriod() PeriodDate {
	if o == nil || IsNil(o.Period) {
		var ret PeriodDate
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetPeriodOk() (*PeriodDate, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given PeriodDate and assigns it to the Period field.
func (o *MtbSystemicTherapy) SetPeriod(v PeriodDate) {
	o.Period = &v
}

// GetPatient returns the Patient field value
func (o *MtbSystemicTherapy) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbSystemicTherapy) SetPatient(v Reference) {
	o.Patient = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetCategory() MtbSystemicTherapyCategoryCoding {
	if o == nil || IsNil(o.Category) {
		var ret MtbSystemicTherapyCategoryCoding
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetCategoryOk() (*MtbSystemicTherapyCategoryCoding, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given MtbSystemicTherapyCategoryCoding and assigns it to the Category field.
func (o *MtbSystemicTherapy) SetCategory(v MtbSystemicTherapyCategoryCoding) {
	o.Category = &v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *MtbSystemicTherapy) GetBasedOn() Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret Reference
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbSystemicTherapy) GetBasedOnOk() (*Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *MtbSystemicTherapy) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given Reference and assigns it to the BasedOn field.
func (o *MtbSystemicTherapy) SetBasedOn(v Reference) {
	o.BasedOn = &v
}

func (o MtbSystemicTherapy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbSystemicTherapy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	if !IsNil(o.Medication) {
		toSerialize["medication"] = o.Medication
	}
	if !IsNil(o.Intent) {
		toSerialize["intent"] = o.Intent
	}
	if !IsNil(o.TherapyLine) {
		toSerialize["therapyLine"] = o.TherapyLine
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.RecommendationFulfillmentStatus) {
		toSerialize["recommendationFulfillmentStatus"] = o.RecommendationFulfillmentStatus
	}
	toSerialize["id"] = o.Id
	toSerialize["recordedOn"] = o.RecordedOn
	toSerialize["status"] = o.Status
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Dosage) {
		toSerialize["dosage"] = o.Dosage
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["patient"] = o.Patient
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	return toSerialize, nil
}

func (o *MtbSystemicTherapy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"recordedOn",
		"status",
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

	varMtbSystemicTherapy := _MtbSystemicTherapy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbSystemicTherapy)

	if err != nil {
		return err
	}

	*o = MtbSystemicTherapy(varMtbSystemicTherapy)

	return err
}

type NullableMtbSystemicTherapy struct {
	value *MtbSystemicTherapy
	isSet bool
}

func (v NullableMtbSystemicTherapy) Get() *MtbSystemicTherapy {
	return v.value
}

func (v *NullableMtbSystemicTherapy) Set(val *MtbSystemicTherapy) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbSystemicTherapy) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbSystemicTherapy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbSystemicTherapy(val *MtbSystemicTherapy) *NullableMtbSystemicTherapy {
	return &NullableMtbSystemicTherapy{value: val, isSet: true}
}

func (v NullableMtbSystemicTherapy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbSystemicTherapy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
