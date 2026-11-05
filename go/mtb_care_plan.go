package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MtbCarePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbCarePlan{}

// MtbCarePlan struct for MtbCarePlan
type MtbCarePlan struct {
	BoardType                       *CarePlanBoardTypeCoding                        `json:"boardType,omitempty"`
	RecommendationsMissingReason    *MtbCarePlanRecommendationsMissingReasonCoding  `json:"recommendationsMissingReason,omitempty"`
	StudyEnrollmentRecommendations  []MtbStudyEnrollmentRecommendation              `json:"studyEnrollmentRecommendations,omitempty"`
	Reason                          *Reference                                      `json:"reason,omitempty"`
	GeneticCounselingRecommendation *GeneticCounselingRecommendation                `json:"geneticCounselingRecommendation,omitempty"`
	Id                              string                                          `json:"id"`
	MedicationRecommendations       []MtbMedicationRecommendation                   `json:"medicationRecommendations,omitempty"`
	RebiopsyRequests                []MtbCarePlanRebiopsyRequestsInner              `json:"rebiopsyRequests,omitempty"`
	HistologyReevaluationRequests   []MtbCarePlanHistologyReevaluationRequestsInner `json:"histologyReevaluationRequests,omitempty"`
	Notes                           []string                                        `json:"notes,omitempty"`
	ProcedureRecommendations        []MtbCarePlanProcedureRecommendationsInner      `json:"procedureRecommendations,omitempty"`
	IssuedOn                        string                                          `json:"issuedOn"`
	NoSequencingPerformedReason     *CarePlanNoSequencingPerformedReasonCoding      `json:"noSequencingPerformedReason,omitempty"`
	Patient                         Reference                                       `json:"patient"`
	AdditionalProperties            map[string]interface{}
}

type _MtbCarePlan MtbCarePlan

// NewMtbCarePlan instantiates a new MtbCarePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbCarePlan(id string, issuedOn string, patient Reference) *MtbCarePlan {
	this := MtbCarePlan{}
	this.Id = id
	this.IssuedOn = issuedOn
	this.Patient = patient
	return &this
}

// NewMtbCarePlanWithDefaults instantiates a new MtbCarePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbCarePlanWithDefaults() *MtbCarePlan {
	this := MtbCarePlan{}
	return &this
}

// GetBoardType returns the BoardType field value if set, zero value otherwise.
func (o *MtbCarePlan) GetBoardType() CarePlanBoardTypeCoding {
	if o == nil || IsNil(o.BoardType) {
		var ret CarePlanBoardTypeCoding
		return ret
	}
	return *o.BoardType
}

// GetBoardTypeOk returns a tuple with the BoardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetBoardTypeOk() (*CarePlanBoardTypeCoding, bool) {
	if o == nil || IsNil(o.BoardType) {
		return nil, false
	}
	return o.BoardType, true
}

// HasBoardType returns a boolean if a field has been set.
func (o *MtbCarePlan) HasBoardType() bool {
	if o != nil && !IsNil(o.BoardType) {
		return true
	}

	return false
}

// SetBoardType gets a reference to the given CarePlanBoardTypeCoding and assigns it to the BoardType field.
func (o *MtbCarePlan) SetBoardType(v CarePlanBoardTypeCoding) {
	o.BoardType = &v
}

// GetRecommendationsMissingReason returns the RecommendationsMissingReason field value if set, zero value otherwise.
func (o *MtbCarePlan) GetRecommendationsMissingReason() MtbCarePlanRecommendationsMissingReasonCoding {
	if o == nil || IsNil(o.RecommendationsMissingReason) {
		var ret MtbCarePlanRecommendationsMissingReasonCoding
		return ret
	}
	return *o.RecommendationsMissingReason
}

// GetRecommendationsMissingReasonOk returns a tuple with the RecommendationsMissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetRecommendationsMissingReasonOk() (*MtbCarePlanRecommendationsMissingReasonCoding, bool) {
	if o == nil || IsNil(o.RecommendationsMissingReason) {
		return nil, false
	}
	return o.RecommendationsMissingReason, true
}

// HasRecommendationsMissingReason returns a boolean if a field has been set.
func (o *MtbCarePlan) HasRecommendationsMissingReason() bool {
	if o != nil && !IsNil(o.RecommendationsMissingReason) {
		return true
	}

	return false
}

// SetRecommendationsMissingReason gets a reference to the given MtbCarePlanRecommendationsMissingReasonCoding and assigns it to the RecommendationsMissingReason field.
func (o *MtbCarePlan) SetRecommendationsMissingReason(v MtbCarePlanRecommendationsMissingReasonCoding) {
	o.RecommendationsMissingReason = &v
}

// GetStudyEnrollmentRecommendations returns the StudyEnrollmentRecommendations field value if set, zero value otherwise.
func (o *MtbCarePlan) GetStudyEnrollmentRecommendations() []MtbStudyEnrollmentRecommendation {
	if o == nil || IsNil(o.StudyEnrollmentRecommendations) {
		var ret []MtbStudyEnrollmentRecommendation
		return ret
	}
	return o.StudyEnrollmentRecommendations
}

// GetStudyEnrollmentRecommendationsOk returns a tuple with the StudyEnrollmentRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetStudyEnrollmentRecommendationsOk() ([]MtbStudyEnrollmentRecommendation, bool) {
	if o == nil || IsNil(o.StudyEnrollmentRecommendations) {
		return nil, false
	}
	return o.StudyEnrollmentRecommendations, true
}

// HasStudyEnrollmentRecommendations returns a boolean if a field has been set.
func (o *MtbCarePlan) HasStudyEnrollmentRecommendations() bool {
	if o != nil && !IsNil(o.StudyEnrollmentRecommendations) {
		return true
	}

	return false
}

// SetStudyEnrollmentRecommendations gets a reference to the given []MtbStudyEnrollmentRecommendation and assigns it to the StudyEnrollmentRecommendations field.
func (o *MtbCarePlan) SetStudyEnrollmentRecommendations(v []MtbStudyEnrollmentRecommendation) {
	o.StudyEnrollmentRecommendations = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MtbCarePlan) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MtbCarePlan) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *MtbCarePlan) SetReason(v Reference) {
	o.Reason = &v
}

// GetGeneticCounselingRecommendation returns the GeneticCounselingRecommendation field value if set, zero value otherwise.
func (o *MtbCarePlan) GetGeneticCounselingRecommendation() GeneticCounselingRecommendation {
	if o == nil || IsNil(o.GeneticCounselingRecommendation) {
		var ret GeneticCounselingRecommendation
		return ret
	}
	return *o.GeneticCounselingRecommendation
}

// GetGeneticCounselingRecommendationOk returns a tuple with the GeneticCounselingRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetGeneticCounselingRecommendationOk() (*GeneticCounselingRecommendation, bool) {
	if o == nil || IsNil(o.GeneticCounselingRecommendation) {
		return nil, false
	}
	return o.GeneticCounselingRecommendation, true
}

// HasGeneticCounselingRecommendation returns a boolean if a field has been set.
func (o *MtbCarePlan) HasGeneticCounselingRecommendation() bool {
	if o != nil && !IsNil(o.GeneticCounselingRecommendation) {
		return true
	}

	return false
}

// SetGeneticCounselingRecommendation gets a reference to the given GeneticCounselingRecommendation and assigns it to the GeneticCounselingRecommendation field.
func (o *MtbCarePlan) SetGeneticCounselingRecommendation(v GeneticCounselingRecommendation) {
	o.GeneticCounselingRecommendation = &v
}

// GetId returns the Id field value
func (o *MtbCarePlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbCarePlan) SetId(v string) {
	o.Id = v
}

// GetMedicationRecommendations returns the MedicationRecommendations field value if set, zero value otherwise.
func (o *MtbCarePlan) GetMedicationRecommendations() []MtbMedicationRecommendation {
	if o == nil || IsNil(o.MedicationRecommendations) {
		var ret []MtbMedicationRecommendation
		return ret
	}
	return o.MedicationRecommendations
}

// GetMedicationRecommendationsOk returns a tuple with the MedicationRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetMedicationRecommendationsOk() ([]MtbMedicationRecommendation, bool) {
	if o == nil || IsNil(o.MedicationRecommendations) {
		return nil, false
	}
	return o.MedicationRecommendations, true
}

// HasMedicationRecommendations returns a boolean if a field has been set.
func (o *MtbCarePlan) HasMedicationRecommendations() bool {
	if o != nil && !IsNil(o.MedicationRecommendations) {
		return true
	}

	return false
}

// SetMedicationRecommendations gets a reference to the given []MtbMedicationRecommendation and assigns it to the MedicationRecommendations field.
func (o *MtbCarePlan) SetMedicationRecommendations(v []MtbMedicationRecommendation) {
	o.MedicationRecommendations = v
}

// GetRebiopsyRequests returns the RebiopsyRequests field value if set, zero value otherwise.
func (o *MtbCarePlan) GetRebiopsyRequests() []MtbCarePlanRebiopsyRequestsInner {
	if o == nil || IsNil(o.RebiopsyRequests) {
		var ret []MtbCarePlanRebiopsyRequestsInner
		return ret
	}
	return o.RebiopsyRequests
}

// GetRebiopsyRequestsOk returns a tuple with the RebiopsyRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetRebiopsyRequestsOk() ([]MtbCarePlanRebiopsyRequestsInner, bool) {
	if o == nil || IsNil(o.RebiopsyRequests) {
		return nil, false
	}
	return o.RebiopsyRequests, true
}

// HasRebiopsyRequests returns a boolean if a field has been set.
func (o *MtbCarePlan) HasRebiopsyRequests() bool {
	if o != nil && !IsNil(o.RebiopsyRequests) {
		return true
	}

	return false
}

// SetRebiopsyRequests gets a reference to the given []MtbCarePlanRebiopsyRequestsInner and assigns it to the RebiopsyRequests field.
func (o *MtbCarePlan) SetRebiopsyRequests(v []MtbCarePlanRebiopsyRequestsInner) {
	o.RebiopsyRequests = v
}

// GetHistologyReevaluationRequests returns the HistologyReevaluationRequests field value if set, zero value otherwise.
func (o *MtbCarePlan) GetHistologyReevaluationRequests() []MtbCarePlanHistologyReevaluationRequestsInner {
	if o == nil || IsNil(o.HistologyReevaluationRequests) {
		var ret []MtbCarePlanHistologyReevaluationRequestsInner
		return ret
	}
	return o.HistologyReevaluationRequests
}

// GetHistologyReevaluationRequestsOk returns a tuple with the HistologyReevaluationRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetHistologyReevaluationRequestsOk() ([]MtbCarePlanHistologyReevaluationRequestsInner, bool) {
	if o == nil || IsNil(o.HistologyReevaluationRequests) {
		return nil, false
	}
	return o.HistologyReevaluationRequests, true
}

// HasHistologyReevaluationRequests returns a boolean if a field has been set.
func (o *MtbCarePlan) HasHistologyReevaluationRequests() bool {
	if o != nil && !IsNil(o.HistologyReevaluationRequests) {
		return true
	}

	return false
}

// SetHistologyReevaluationRequests gets a reference to the given []MtbCarePlanHistologyReevaluationRequestsInner and assigns it to the HistologyReevaluationRequests field.
func (o *MtbCarePlan) SetHistologyReevaluationRequests(v []MtbCarePlanHistologyReevaluationRequestsInner) {
	o.HistologyReevaluationRequests = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *MtbCarePlan) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MtbCarePlan) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *MtbCarePlan) SetNotes(v []string) {
	o.Notes = v
}

// GetProcedureRecommendations returns the ProcedureRecommendations field value if set, zero value otherwise.
func (o *MtbCarePlan) GetProcedureRecommendations() []MtbCarePlanProcedureRecommendationsInner {
	if o == nil || IsNil(o.ProcedureRecommendations) {
		var ret []MtbCarePlanProcedureRecommendationsInner
		return ret
	}
	return o.ProcedureRecommendations
}

// GetProcedureRecommendationsOk returns a tuple with the ProcedureRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetProcedureRecommendationsOk() ([]MtbCarePlanProcedureRecommendationsInner, bool) {
	if o == nil || IsNil(o.ProcedureRecommendations) {
		return nil, false
	}
	return o.ProcedureRecommendations, true
}

// HasProcedureRecommendations returns a boolean if a field has been set.
func (o *MtbCarePlan) HasProcedureRecommendations() bool {
	if o != nil && !IsNil(o.ProcedureRecommendations) {
		return true
	}

	return false
}

// SetProcedureRecommendations gets a reference to the given []MtbCarePlanProcedureRecommendationsInner and assigns it to the ProcedureRecommendations field.
func (o *MtbCarePlan) SetProcedureRecommendations(v []MtbCarePlanProcedureRecommendationsInner) {
	o.ProcedureRecommendations = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbCarePlan) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbCarePlan) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetNoSequencingPerformedReason returns the NoSequencingPerformedReason field value if set, zero value otherwise.
func (o *MtbCarePlan) GetNoSequencingPerformedReason() CarePlanNoSequencingPerformedReasonCoding {
	if o == nil || IsNil(o.NoSequencingPerformedReason) {
		var ret CarePlanNoSequencingPerformedReasonCoding
		return ret
	}
	return *o.NoSequencingPerformedReason
}

// GetNoSequencingPerformedReasonOk returns a tuple with the NoSequencingPerformedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetNoSequencingPerformedReasonOk() (*CarePlanNoSequencingPerformedReasonCoding, bool) {
	if o == nil || IsNil(o.NoSequencingPerformedReason) {
		return nil, false
	}
	return o.NoSequencingPerformedReason, true
}

// HasNoSequencingPerformedReason returns a boolean if a field has been set.
func (o *MtbCarePlan) HasNoSequencingPerformedReason() bool {
	if o != nil && !IsNil(o.NoSequencingPerformedReason) {
		return true
	}

	return false
}

// SetNoSequencingPerformedReason gets a reference to the given CarePlanNoSequencingPerformedReasonCoding and assigns it to the NoSequencingPerformedReason field.
func (o *MtbCarePlan) SetNoSequencingPerformedReason(v CarePlanNoSequencingPerformedReasonCoding) {
	o.NoSequencingPerformedReason = &v
}

// GetPatient returns the Patient field value
func (o *MtbCarePlan) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlan) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbCarePlan) SetPatient(v Reference) {
	o.Patient = v
}

func (o MtbCarePlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbCarePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoardType) {
		toSerialize["boardType"] = o.BoardType
	}
	if !IsNil(o.RecommendationsMissingReason) {
		toSerialize["recommendationsMissingReason"] = o.RecommendationsMissingReason
	}
	if !IsNil(o.StudyEnrollmentRecommendations) {
		toSerialize["studyEnrollmentRecommendations"] = o.StudyEnrollmentRecommendations
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.GeneticCounselingRecommendation) {
		toSerialize["geneticCounselingRecommendation"] = o.GeneticCounselingRecommendation
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.MedicationRecommendations) {
		toSerialize["medicationRecommendations"] = o.MedicationRecommendations
	}
	if !IsNil(o.RebiopsyRequests) {
		toSerialize["rebiopsyRequests"] = o.RebiopsyRequests
	}
	if !IsNil(o.HistologyReevaluationRequests) {
		toSerialize["histologyReevaluationRequests"] = o.HistologyReevaluationRequests
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ProcedureRecommendations) {
		toSerialize["procedureRecommendations"] = o.ProcedureRecommendations
	}
	toSerialize["issuedOn"] = o.IssuedOn
	if !IsNil(o.NoSequencingPerformedReason) {
		toSerialize["noSequencingPerformedReason"] = o.NoSequencingPerformedReason
	}
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MtbCarePlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"issuedOn",
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

	varMtbCarePlan := _MtbCarePlan{}

	err = json.Unmarshal(data, &varMtbCarePlan)

	if err != nil {
		return err
	}

	*o = MtbCarePlan(varMtbCarePlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "boardType")
		delete(additionalProperties, "recommendationsMissingReason")
		delete(additionalProperties, "studyEnrollmentRecommendations")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "geneticCounselingRecommendation")
		delete(additionalProperties, "id")
		delete(additionalProperties, "medicationRecommendations")
		delete(additionalProperties, "rebiopsyRequests")
		delete(additionalProperties, "histologyReevaluationRequests")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "procedureRecommendations")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "noSequencingPerformedReason")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMtbCarePlan struct {
	value *MtbCarePlan
	isSet bool
}

func (v NullableMtbCarePlan) Get() *MtbCarePlan {
	return v.value
}

func (v *NullableMtbCarePlan) Set(val *MtbCarePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbCarePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbCarePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbCarePlan(val *MtbCarePlan) *NullableMtbCarePlan {
	return &NullableMtbCarePlan{value: val, isSet: true}
}

func (v NullableMtbCarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbCarePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
