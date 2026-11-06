package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MtbStudyEnrollmentRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbStudyEnrollmentRecommendation{}

// MtbStudyEnrollmentRecommendation struct for MtbStudyEnrollmentRecommendation
type MtbStudyEnrollmentRecommendation struct {
	Priority             MtbRecommendationPriorityCoding   `json:"priority"`
	Medication           []AtcUnregisteredMedicationCoding `json:"medication,omitempty"`
	SupportingVariants   []GeneAlterationReference         `json:"supportingVariants,omitempty"`
	Reason               *Reference                        `json:"reason,omitempty"`
	Patient              Reference                         `json:"patient"`
	SupportingFindings   []Reference                       `json:"supportingFindings,omitempty"`
	LevelOfEvidence      *LevelOfEvidence                  `json:"levelOfEvidence,omitempty"`
	Id                   string                            `json:"id"`
	IssuedOn             string                            `json:"issuedOn"`
	Study                []StudyReference                  `json:"study"`
	AdditionalProperties map[string]interface{}
}

type _MtbStudyEnrollmentRecommendation MtbStudyEnrollmentRecommendation

// NewMtbStudyEnrollmentRecommendation instantiates a new MtbStudyEnrollmentRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbStudyEnrollmentRecommendation(priority MtbRecommendationPriorityCoding, patient Reference, id string, issuedOn string, study []StudyReference) *MtbStudyEnrollmentRecommendation {
	this := MtbStudyEnrollmentRecommendation{}
	this.Priority = priority
	this.Patient = patient
	this.Id = id
	this.IssuedOn = issuedOn
	this.Study = study
	return &this
}

// NewMtbStudyEnrollmentRecommendationWithDefaults instantiates a new MtbStudyEnrollmentRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbStudyEnrollmentRecommendationWithDefaults() *MtbStudyEnrollmentRecommendation {
	this := MtbStudyEnrollmentRecommendation{}
	return &this
}

// GetPriority returns the Priority field value
func (o *MtbStudyEnrollmentRecommendation) GetPriority() MtbRecommendationPriorityCoding {
	if o == nil {
		var ret MtbRecommendationPriorityCoding
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetPriorityOk() (*MtbRecommendationPriorityCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *MtbStudyEnrollmentRecommendation) SetPriority(v MtbRecommendationPriorityCoding) {
	o.Priority = v
}

// GetMedication returns the Medication field value if set, zero value otherwise.
func (o *MtbStudyEnrollmentRecommendation) GetMedication() []AtcUnregisteredMedicationCoding {
	if o == nil || IsNil(o.Medication) {
		var ret []AtcUnregisteredMedicationCoding
		return ret
	}
	return o.Medication
}

// GetMedicationOk returns a tuple with the Medication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetMedicationOk() ([]AtcUnregisteredMedicationCoding, bool) {
	if o == nil || IsNil(o.Medication) {
		return nil, false
	}
	return o.Medication, true
}

// HasMedication returns a boolean if a field has been set.
func (o *MtbStudyEnrollmentRecommendation) HasMedication() bool {
	if o != nil && !IsNil(o.Medication) {
		return true
	}

	return false
}

// SetMedication gets a reference to the given []AtcUnregisteredMedicationCoding and assigns it to the Medication field.
func (o *MtbStudyEnrollmentRecommendation) SetMedication(v []AtcUnregisteredMedicationCoding) {
	o.Medication = v
}

// GetSupportingVariants returns the SupportingVariants field value if set, zero value otherwise.
func (o *MtbStudyEnrollmentRecommendation) GetSupportingVariants() []GeneAlterationReference {
	if o == nil || IsNil(o.SupportingVariants) {
		var ret []GeneAlterationReference
		return ret
	}
	return o.SupportingVariants
}

// GetSupportingVariantsOk returns a tuple with the SupportingVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetSupportingVariantsOk() ([]GeneAlterationReference, bool) {
	if o == nil || IsNil(o.SupportingVariants) {
		return nil, false
	}
	return o.SupportingVariants, true
}

// HasSupportingVariants returns a boolean if a field has been set.
func (o *MtbStudyEnrollmentRecommendation) HasSupportingVariants() bool {
	if o != nil && !IsNil(o.SupportingVariants) {
		return true
	}

	return false
}

// SetSupportingVariants gets a reference to the given []GeneAlterationReference and assigns it to the SupportingVariants field.
func (o *MtbStudyEnrollmentRecommendation) SetSupportingVariants(v []GeneAlterationReference) {
	o.SupportingVariants = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MtbStudyEnrollmentRecommendation) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MtbStudyEnrollmentRecommendation) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *MtbStudyEnrollmentRecommendation) SetReason(v Reference) {
	o.Reason = &v
}

// GetPatient returns the Patient field value
func (o *MtbStudyEnrollmentRecommendation) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbStudyEnrollmentRecommendation) SetPatient(v Reference) {
	o.Patient = v
}

// GetSupportingFindings returns the SupportingFindings field value if set, zero value otherwise.
func (o *MtbStudyEnrollmentRecommendation) GetSupportingFindings() []Reference {
	if o == nil || IsNil(o.SupportingFindings) {
		var ret []Reference
		return ret
	}
	return o.SupportingFindings
}

// GetSupportingFindingsOk returns a tuple with the SupportingFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetSupportingFindingsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.SupportingFindings) {
		return nil, false
	}
	return o.SupportingFindings, true
}

// HasSupportingFindings returns a boolean if a field has been set.
func (o *MtbStudyEnrollmentRecommendation) HasSupportingFindings() bool {
	if o != nil && !IsNil(o.SupportingFindings) {
		return true
	}

	return false
}

// SetSupportingFindings gets a reference to the given []Reference and assigns it to the SupportingFindings field.
func (o *MtbStudyEnrollmentRecommendation) SetSupportingFindings(v []Reference) {
	o.SupportingFindings = v
}

// GetLevelOfEvidence returns the LevelOfEvidence field value if set, zero value otherwise.
func (o *MtbStudyEnrollmentRecommendation) GetLevelOfEvidence() LevelOfEvidence {
	if o == nil || IsNil(o.LevelOfEvidence) {
		var ret LevelOfEvidence
		return ret
	}
	return *o.LevelOfEvidence
}

// GetLevelOfEvidenceOk returns a tuple with the LevelOfEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetLevelOfEvidenceOk() (*LevelOfEvidence, bool) {
	if o == nil || IsNil(o.LevelOfEvidence) {
		return nil, false
	}
	return o.LevelOfEvidence, true
}

// HasLevelOfEvidence returns a boolean if a field has been set.
func (o *MtbStudyEnrollmentRecommendation) HasLevelOfEvidence() bool {
	if o != nil && !IsNil(o.LevelOfEvidence) {
		return true
	}

	return false
}

// SetLevelOfEvidence gets a reference to the given LevelOfEvidence and assigns it to the LevelOfEvidence field.
func (o *MtbStudyEnrollmentRecommendation) SetLevelOfEvidence(v LevelOfEvidence) {
	o.LevelOfEvidence = &v
}

// GetId returns the Id field value
func (o *MtbStudyEnrollmentRecommendation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbStudyEnrollmentRecommendation) SetId(v string) {
	o.Id = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbStudyEnrollmentRecommendation) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbStudyEnrollmentRecommendation) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetStudy returns the Study field value
func (o *MtbStudyEnrollmentRecommendation) GetStudy() []StudyReference {
	if o == nil {
		var ret []StudyReference
		return ret
	}

	return o.Study
}

// GetStudyOk returns a tuple with the Study field value
// and a boolean to check if the value has been set.
func (o *MtbStudyEnrollmentRecommendation) GetStudyOk() ([]StudyReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Study, true
}

// SetStudy sets field value
func (o *MtbStudyEnrollmentRecommendation) SetStudy(v []StudyReference) {
	o.Study = v
}

func (o MtbStudyEnrollmentRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbStudyEnrollmentRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.Medication) {
		toSerialize["medication"] = o.Medication
	}
	if !IsNil(o.SupportingVariants) {
		toSerialize["supportingVariants"] = o.SupportingVariants
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["patient"] = o.Patient
	if !IsNil(o.SupportingFindings) {
		toSerialize["supportingFindings"] = o.SupportingFindings
	}
	if !IsNil(o.LevelOfEvidence) {
		toSerialize["levelOfEvidence"] = o.LevelOfEvidence
	}
	toSerialize["id"] = o.Id
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["study"] = o.Study

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MtbStudyEnrollmentRecommendation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"priority",
		"patient",
		"id",
		"issuedOn",
		"study",
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

	varMtbStudyEnrollmentRecommendation := _MtbStudyEnrollmentRecommendation{}

	err = json.Unmarshal(data, &varMtbStudyEnrollmentRecommendation)

	if err != nil {
		return err
	}

	*o = MtbStudyEnrollmentRecommendation(varMtbStudyEnrollmentRecommendation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "priority")
		delete(additionalProperties, "medication")
		delete(additionalProperties, "supportingVariants")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "supportingFindings")
		delete(additionalProperties, "levelOfEvidence")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "study")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMtbStudyEnrollmentRecommendation struct {
	value *MtbStudyEnrollmentRecommendation
	isSet bool
}

func (v NullableMtbStudyEnrollmentRecommendation) Get() *MtbStudyEnrollmentRecommendation {
	return v.value
}

func (v *NullableMtbStudyEnrollmentRecommendation) Set(val *MtbStudyEnrollmentRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbStudyEnrollmentRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbStudyEnrollmentRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbStudyEnrollmentRecommendation(val *MtbStudyEnrollmentRecommendation) *NullableMtbStudyEnrollmentRecommendation {
	return &NullableMtbStudyEnrollmentRecommendation{value: val, isSet: true}
}

func (v NullableMtbStudyEnrollmentRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbStudyEnrollmentRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
