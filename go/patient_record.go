package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PatientRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientRecord{}

// PatientRecord struct for PatientRecord
type PatientRecord struct {
	Diagnoses              []MtbDiagnosis                        `json:"diagnoses"`
	StudyInclusions        []PatientRecordStudyInclusionsInner   `json:"studyInclusions,omitempty"`
	NgsReports             []SomaticNgsReport                    `json:"ngsReports,omitempty"`
	PerformanceStatus      []PerformanceStatus                   `json:"performanceStatus,omitempty"`
	GuidelineProcedures    []OncoProcedure                       `json:"guidelineProcedures,omitempty"`
	FamilyMemberHistories  []FamilyMemberHistory                 `json:"familyMemberHistories,omitempty"`
	ClaimResponses         []ClaimResponse                       `json:"claimResponses,omitempty"`
	MsiFindings            []Msi                                 `json:"msiFindings,omitempty"`
	CarePlans              []MtbCarePlan                         `json:"carePlans,omitempty"`
	Specimens              []TumorSpecimen                       `json:"specimens,omitempty"`
	IhcReports             []IhcReport                           `json:"ihcReports,omitempty"`
	HistologyReports       []HistologyReport                     `json:"histologyReports,omitempty"`
	GuidelineTherapies     []MtbSystemicTherapy                  `json:"guidelineTherapies,omitempty"`
	EpisodesOfCare         []MtbEpisodeOfCare                    `json:"episodesOfCare"`
	PriorDiagnosticReports []MolecularDiagnosticReport           `json:"priorDiagnosticReports,omitempty"`
	SystemicTherapies      []PatientRecordSystemicTherapiesInner `json:"systemicTherapies,omitempty"`
	FollowUps              []FollowUp                            `json:"followUps,omitempty"`
	Claims                 []Claim                               `json:"claims,omitempty"`
	Metadata               *MvhMetadata                          `json:"metadata,omitempty"`
	Responses              []Response                            `json:"responses,omitempty"`
	Patient                Patient                               `json:"patient"`
}

type _PatientRecord PatientRecord

// NewPatientRecord instantiates a new PatientRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientRecord(diagnoses []MtbDiagnosis, episodesOfCare []MtbEpisodeOfCare, patient Patient) *PatientRecord {
	this := PatientRecord{}
	this.Diagnoses = diagnoses
	this.EpisodesOfCare = episodesOfCare
	this.Patient = patient
	return &this
}

// NewPatientRecordWithDefaults instantiates a new PatientRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientRecordWithDefaults() *PatientRecord {
	this := PatientRecord{}
	return &this
}

// GetDiagnoses returns the Diagnoses field value
func (o *PatientRecord) GetDiagnoses() []MtbDiagnosis {
	if o == nil {
		var ret []MtbDiagnosis
		return ret
	}

	return o.Diagnoses
}

// GetDiagnosesOk returns a tuple with the Diagnoses field value
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetDiagnosesOk() ([]MtbDiagnosis, bool) {
	if o == nil {
		return nil, false
	}
	return o.Diagnoses, true
}

// SetDiagnoses sets field value
func (o *PatientRecord) SetDiagnoses(v []MtbDiagnosis) {
	o.Diagnoses = v
}

// GetStudyInclusions returns the StudyInclusions field value if set, zero value otherwise.
func (o *PatientRecord) GetStudyInclusions() []PatientRecordStudyInclusionsInner {
	if o == nil || IsNil(o.StudyInclusions) {
		var ret []PatientRecordStudyInclusionsInner
		return ret
	}
	return o.StudyInclusions
}

// GetStudyInclusionsOk returns a tuple with the StudyInclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetStudyInclusionsOk() ([]PatientRecordStudyInclusionsInner, bool) {
	if o == nil || IsNil(o.StudyInclusions) {
		return nil, false
	}
	return o.StudyInclusions, true
}

// HasStudyInclusions returns a boolean if a field has been set.
func (o *PatientRecord) HasStudyInclusions() bool {
	if o != nil && !IsNil(o.StudyInclusions) {
		return true
	}

	return false
}

// SetStudyInclusions gets a reference to the given []PatientRecordStudyInclusionsInner and assigns it to the StudyInclusions field.
func (o *PatientRecord) SetStudyInclusions(v []PatientRecordStudyInclusionsInner) {
	o.StudyInclusions = v
}

// GetNgsReports returns the NgsReports field value if set, zero value otherwise.
func (o *PatientRecord) GetNgsReports() []SomaticNgsReport {
	if o == nil || IsNil(o.NgsReports) {
		var ret []SomaticNgsReport
		return ret
	}
	return o.NgsReports
}

// GetNgsReportsOk returns a tuple with the NgsReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetNgsReportsOk() ([]SomaticNgsReport, bool) {
	if o == nil || IsNil(o.NgsReports) {
		return nil, false
	}
	return o.NgsReports, true
}

// HasNgsReports returns a boolean if a field has been set.
func (o *PatientRecord) HasNgsReports() bool {
	if o != nil && !IsNil(o.NgsReports) {
		return true
	}

	return false
}

// SetNgsReports gets a reference to the given []SomaticNgsReport and assigns it to the NgsReports field.
func (o *PatientRecord) SetNgsReports(v []SomaticNgsReport) {
	o.NgsReports = v
}

// GetPerformanceStatus returns the PerformanceStatus field value if set, zero value otherwise.
func (o *PatientRecord) GetPerformanceStatus() []PerformanceStatus {
	if o == nil || IsNil(o.PerformanceStatus) {
		var ret []PerformanceStatus
		return ret
	}
	return o.PerformanceStatus
}

// GetPerformanceStatusOk returns a tuple with the PerformanceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetPerformanceStatusOk() ([]PerformanceStatus, bool) {
	if o == nil || IsNil(o.PerformanceStatus) {
		return nil, false
	}
	return o.PerformanceStatus, true
}

// HasPerformanceStatus returns a boolean if a field has been set.
func (o *PatientRecord) HasPerformanceStatus() bool {
	if o != nil && !IsNil(o.PerformanceStatus) {
		return true
	}

	return false
}

// SetPerformanceStatus gets a reference to the given []PerformanceStatus and assigns it to the PerformanceStatus field.
func (o *PatientRecord) SetPerformanceStatus(v []PerformanceStatus) {
	o.PerformanceStatus = v
}

// GetGuidelineProcedures returns the GuidelineProcedures field value if set, zero value otherwise.
func (o *PatientRecord) GetGuidelineProcedures() []OncoProcedure {
	if o == nil || IsNil(o.GuidelineProcedures) {
		var ret []OncoProcedure
		return ret
	}
	return o.GuidelineProcedures
}

// GetGuidelineProceduresOk returns a tuple with the GuidelineProcedures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetGuidelineProceduresOk() ([]OncoProcedure, bool) {
	if o == nil || IsNil(o.GuidelineProcedures) {
		return nil, false
	}
	return o.GuidelineProcedures, true
}

// HasGuidelineProcedures returns a boolean if a field has been set.
func (o *PatientRecord) HasGuidelineProcedures() bool {
	if o != nil && !IsNil(o.GuidelineProcedures) {
		return true
	}

	return false
}

// SetGuidelineProcedures gets a reference to the given []OncoProcedure and assigns it to the GuidelineProcedures field.
func (o *PatientRecord) SetGuidelineProcedures(v []OncoProcedure) {
	o.GuidelineProcedures = v
}

// GetFamilyMemberHistories returns the FamilyMemberHistories field value if set, zero value otherwise.
func (o *PatientRecord) GetFamilyMemberHistories() []FamilyMemberHistory {
	if o == nil || IsNil(o.FamilyMemberHistories) {
		var ret []FamilyMemberHistory
		return ret
	}
	return o.FamilyMemberHistories
}

// GetFamilyMemberHistoriesOk returns a tuple with the FamilyMemberHistories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetFamilyMemberHistoriesOk() ([]FamilyMemberHistory, bool) {
	if o == nil || IsNil(o.FamilyMemberHistories) {
		return nil, false
	}
	return o.FamilyMemberHistories, true
}

// HasFamilyMemberHistories returns a boolean if a field has been set.
func (o *PatientRecord) HasFamilyMemberHistories() bool {
	if o != nil && !IsNil(o.FamilyMemberHistories) {
		return true
	}

	return false
}

// SetFamilyMemberHistories gets a reference to the given []FamilyMemberHistory and assigns it to the FamilyMemberHistories field.
func (o *PatientRecord) SetFamilyMemberHistories(v []FamilyMemberHistory) {
	o.FamilyMemberHistories = v
}

// GetClaimResponses returns the ClaimResponses field value if set, zero value otherwise.
func (o *PatientRecord) GetClaimResponses() []ClaimResponse {
	if o == nil || IsNil(o.ClaimResponses) {
		var ret []ClaimResponse
		return ret
	}
	return o.ClaimResponses
}

// GetClaimResponsesOk returns a tuple with the ClaimResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetClaimResponsesOk() ([]ClaimResponse, bool) {
	if o == nil || IsNil(o.ClaimResponses) {
		return nil, false
	}
	return o.ClaimResponses, true
}

// HasClaimResponses returns a boolean if a field has been set.
func (o *PatientRecord) HasClaimResponses() bool {
	if o != nil && !IsNil(o.ClaimResponses) {
		return true
	}

	return false
}

// SetClaimResponses gets a reference to the given []ClaimResponse and assigns it to the ClaimResponses field.
func (o *PatientRecord) SetClaimResponses(v []ClaimResponse) {
	o.ClaimResponses = v
}

// GetMsiFindings returns the MsiFindings field value if set, zero value otherwise.
func (o *PatientRecord) GetMsiFindings() []Msi {
	if o == nil || IsNil(o.MsiFindings) {
		var ret []Msi
		return ret
	}
	return o.MsiFindings
}

// GetMsiFindingsOk returns a tuple with the MsiFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetMsiFindingsOk() ([]Msi, bool) {
	if o == nil || IsNil(o.MsiFindings) {
		return nil, false
	}
	return o.MsiFindings, true
}

// HasMsiFindings returns a boolean if a field has been set.
func (o *PatientRecord) HasMsiFindings() bool {
	if o != nil && !IsNil(o.MsiFindings) {
		return true
	}

	return false
}

// SetMsiFindings gets a reference to the given []Msi and assigns it to the MsiFindings field.
func (o *PatientRecord) SetMsiFindings(v []Msi) {
	o.MsiFindings = v
}

// GetCarePlans returns the CarePlans field value if set, zero value otherwise.
func (o *PatientRecord) GetCarePlans() []MtbCarePlan {
	if o == nil || IsNil(o.CarePlans) {
		var ret []MtbCarePlan
		return ret
	}
	return o.CarePlans
}

// GetCarePlansOk returns a tuple with the CarePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetCarePlansOk() ([]MtbCarePlan, bool) {
	if o == nil || IsNil(o.CarePlans) {
		return nil, false
	}
	return o.CarePlans, true
}

// HasCarePlans returns a boolean if a field has been set.
func (o *PatientRecord) HasCarePlans() bool {
	if o != nil && !IsNil(o.CarePlans) {
		return true
	}

	return false
}

// SetCarePlans gets a reference to the given []MtbCarePlan and assigns it to the CarePlans field.
func (o *PatientRecord) SetCarePlans(v []MtbCarePlan) {
	o.CarePlans = v
}

// GetSpecimens returns the Specimens field value if set, zero value otherwise.
func (o *PatientRecord) GetSpecimens() []TumorSpecimen {
	if o == nil || IsNil(o.Specimens) {
		var ret []TumorSpecimen
		return ret
	}
	return o.Specimens
}

// GetSpecimensOk returns a tuple with the Specimens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetSpecimensOk() ([]TumorSpecimen, bool) {
	if o == nil || IsNil(o.Specimens) {
		return nil, false
	}
	return o.Specimens, true
}

// HasSpecimens returns a boolean if a field has been set.
func (o *PatientRecord) HasSpecimens() bool {
	if o != nil && !IsNil(o.Specimens) {
		return true
	}

	return false
}

// SetSpecimens gets a reference to the given []TumorSpecimen and assigns it to the Specimens field.
func (o *PatientRecord) SetSpecimens(v []TumorSpecimen) {
	o.Specimens = v
}

// GetIhcReports returns the IhcReports field value if set, zero value otherwise.
func (o *PatientRecord) GetIhcReports() []IhcReport {
	if o == nil || IsNil(o.IhcReports) {
		var ret []IhcReport
		return ret
	}
	return o.IhcReports
}

// GetIhcReportsOk returns a tuple with the IhcReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetIhcReportsOk() ([]IhcReport, bool) {
	if o == nil || IsNil(o.IhcReports) {
		return nil, false
	}
	return o.IhcReports, true
}

// HasIhcReports returns a boolean if a field has been set.
func (o *PatientRecord) HasIhcReports() bool {
	if o != nil && !IsNil(o.IhcReports) {
		return true
	}

	return false
}

// SetIhcReports gets a reference to the given []IhcReport and assigns it to the IhcReports field.
func (o *PatientRecord) SetIhcReports(v []IhcReport) {
	o.IhcReports = v
}

// GetHistologyReports returns the HistologyReports field value if set, zero value otherwise.
func (o *PatientRecord) GetHistologyReports() []HistologyReport {
	if o == nil || IsNil(o.HistologyReports) {
		var ret []HistologyReport
		return ret
	}
	return o.HistologyReports
}

// GetHistologyReportsOk returns a tuple with the HistologyReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetHistologyReportsOk() ([]HistologyReport, bool) {
	if o == nil || IsNil(o.HistologyReports) {
		return nil, false
	}
	return o.HistologyReports, true
}

// HasHistologyReports returns a boolean if a field has been set.
func (o *PatientRecord) HasHistologyReports() bool {
	if o != nil && !IsNil(o.HistologyReports) {
		return true
	}

	return false
}

// SetHistologyReports gets a reference to the given []HistologyReport and assigns it to the HistologyReports field.
func (o *PatientRecord) SetHistologyReports(v []HistologyReport) {
	o.HistologyReports = v
}

// GetGuidelineTherapies returns the GuidelineTherapies field value if set, zero value otherwise.
func (o *PatientRecord) GetGuidelineTherapies() []MtbSystemicTherapy {
	if o == nil || IsNil(o.GuidelineTherapies) {
		var ret []MtbSystemicTherapy
		return ret
	}
	return o.GuidelineTherapies
}

// GetGuidelineTherapiesOk returns a tuple with the GuidelineTherapies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetGuidelineTherapiesOk() ([]MtbSystemicTherapy, bool) {
	if o == nil || IsNil(o.GuidelineTherapies) {
		return nil, false
	}
	return o.GuidelineTherapies, true
}

// HasGuidelineTherapies returns a boolean if a field has been set.
func (o *PatientRecord) HasGuidelineTherapies() bool {
	if o != nil && !IsNil(o.GuidelineTherapies) {
		return true
	}

	return false
}

// SetGuidelineTherapies gets a reference to the given []MtbSystemicTherapy and assigns it to the GuidelineTherapies field.
func (o *PatientRecord) SetGuidelineTherapies(v []MtbSystemicTherapy) {
	o.GuidelineTherapies = v
}

// GetEpisodesOfCare returns the EpisodesOfCare field value
func (o *PatientRecord) GetEpisodesOfCare() []MtbEpisodeOfCare {
	if o == nil {
		var ret []MtbEpisodeOfCare
		return ret
	}

	return o.EpisodesOfCare
}

// GetEpisodesOfCareOk returns a tuple with the EpisodesOfCare field value
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetEpisodesOfCareOk() ([]MtbEpisodeOfCare, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpisodesOfCare, true
}

// SetEpisodesOfCare sets field value
func (o *PatientRecord) SetEpisodesOfCare(v []MtbEpisodeOfCare) {
	o.EpisodesOfCare = v
}

// GetPriorDiagnosticReports returns the PriorDiagnosticReports field value if set, zero value otherwise.
func (o *PatientRecord) GetPriorDiagnosticReports() []MolecularDiagnosticReport {
	if o == nil || IsNil(o.PriorDiagnosticReports) {
		var ret []MolecularDiagnosticReport
		return ret
	}
	return o.PriorDiagnosticReports
}

// GetPriorDiagnosticReportsOk returns a tuple with the PriorDiagnosticReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetPriorDiagnosticReportsOk() ([]MolecularDiagnosticReport, bool) {
	if o == nil || IsNil(o.PriorDiagnosticReports) {
		return nil, false
	}
	return o.PriorDiagnosticReports, true
}

// HasPriorDiagnosticReports returns a boolean if a field has been set.
func (o *PatientRecord) HasPriorDiagnosticReports() bool {
	if o != nil && !IsNil(o.PriorDiagnosticReports) {
		return true
	}

	return false
}

// SetPriorDiagnosticReports gets a reference to the given []MolecularDiagnosticReport and assigns it to the PriorDiagnosticReports field.
func (o *PatientRecord) SetPriorDiagnosticReports(v []MolecularDiagnosticReport) {
	o.PriorDiagnosticReports = v
}

// GetSystemicTherapies returns the SystemicTherapies field value if set, zero value otherwise.
func (o *PatientRecord) GetSystemicTherapies() []PatientRecordSystemicTherapiesInner {
	if o == nil || IsNil(o.SystemicTherapies) {
		var ret []PatientRecordSystemicTherapiesInner
		return ret
	}
	return o.SystemicTherapies
}

// GetSystemicTherapiesOk returns a tuple with the SystemicTherapies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetSystemicTherapiesOk() ([]PatientRecordSystemicTherapiesInner, bool) {
	if o == nil || IsNil(o.SystemicTherapies) {
		return nil, false
	}
	return o.SystemicTherapies, true
}

// HasSystemicTherapies returns a boolean if a field has been set.
func (o *PatientRecord) HasSystemicTherapies() bool {
	if o != nil && !IsNil(o.SystemicTherapies) {
		return true
	}

	return false
}

// SetSystemicTherapies gets a reference to the given []PatientRecordSystemicTherapiesInner and assigns it to the SystemicTherapies field.
func (o *PatientRecord) SetSystemicTherapies(v []PatientRecordSystemicTherapiesInner) {
	o.SystemicTherapies = v
}

// GetFollowUps returns the FollowUps field value if set, zero value otherwise.
func (o *PatientRecord) GetFollowUps() []FollowUp {
	if o == nil || IsNil(o.FollowUps) {
		var ret []FollowUp
		return ret
	}
	return o.FollowUps
}

// GetFollowUpsOk returns a tuple with the FollowUps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetFollowUpsOk() ([]FollowUp, bool) {
	if o == nil || IsNil(o.FollowUps) {
		return nil, false
	}
	return o.FollowUps, true
}

// HasFollowUps returns a boolean if a field has been set.
func (o *PatientRecord) HasFollowUps() bool {
	if o != nil && !IsNil(o.FollowUps) {
		return true
	}

	return false
}

// SetFollowUps gets a reference to the given []FollowUp and assigns it to the FollowUps field.
func (o *PatientRecord) SetFollowUps(v []FollowUp) {
	o.FollowUps = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *PatientRecord) GetClaims() []Claim {
	if o == nil || IsNil(o.Claims) {
		var ret []Claim
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetClaimsOk() ([]Claim, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *PatientRecord) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []Claim and assigns it to the Claims field.
func (o *PatientRecord) SetClaims(v []Claim) {
	o.Claims = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatientRecord) GetMetadata() MvhMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret MvhMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetMetadataOk() (*MvhMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatientRecord) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MvhMetadata and assigns it to the Metadata field.
func (o *PatientRecord) SetMetadata(v MvhMetadata) {
	o.Metadata = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *PatientRecord) GetResponses() []Response {
	if o == nil || IsNil(o.Responses) {
		var ret []Response
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetResponsesOk() ([]Response, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *PatientRecord) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []Response and assigns it to the Responses field.
func (o *PatientRecord) SetResponses(v []Response) {
	o.Responses = v
}

// GetPatient returns the Patient field value
func (o *PatientRecord) GetPatient() Patient {
	if o == nil {
		var ret Patient
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *PatientRecord) GetPatientOk() (*Patient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *PatientRecord) SetPatient(v Patient) {
	o.Patient = v
}

func (o PatientRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["diagnoses"] = o.Diagnoses
	if !IsNil(o.StudyInclusions) {
		toSerialize["studyInclusions"] = o.StudyInclusions
	}
	if !IsNil(o.NgsReports) {
		toSerialize["ngsReports"] = o.NgsReports
	}
	if !IsNil(o.PerformanceStatus) {
		toSerialize["performanceStatus"] = o.PerformanceStatus
	}
	if !IsNil(o.GuidelineProcedures) {
		toSerialize["guidelineProcedures"] = o.GuidelineProcedures
	}
	if !IsNil(o.FamilyMemberHistories) {
		toSerialize["familyMemberHistories"] = o.FamilyMemberHistories
	}
	if !IsNil(o.ClaimResponses) {
		toSerialize["claimResponses"] = o.ClaimResponses
	}
	if !IsNil(o.MsiFindings) {
		toSerialize["msiFindings"] = o.MsiFindings
	}
	if !IsNil(o.CarePlans) {
		toSerialize["carePlans"] = o.CarePlans
	}
	if !IsNil(o.Specimens) {
		toSerialize["specimens"] = o.Specimens
	}
	if !IsNil(o.IhcReports) {
		toSerialize["ihcReports"] = o.IhcReports
	}
	if !IsNil(o.HistologyReports) {
		toSerialize["histologyReports"] = o.HistologyReports
	}
	if !IsNil(o.GuidelineTherapies) {
		toSerialize["guidelineTherapies"] = o.GuidelineTherapies
	}
	toSerialize["episodesOfCare"] = o.EpisodesOfCare
	if !IsNil(o.PriorDiagnosticReports) {
		toSerialize["priorDiagnosticReports"] = o.PriorDiagnosticReports
	}
	if !IsNil(o.SystemicTherapies) {
		toSerialize["systemicTherapies"] = o.SystemicTherapies
	}
	if !IsNil(o.FollowUps) {
		toSerialize["followUps"] = o.FollowUps
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *PatientRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diagnoses",
		"episodesOfCare",
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

	varPatientRecord := _PatientRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientRecord)

	if err != nil {
		return err
	}

	*o = PatientRecord(varPatientRecord)

	return err
}

type NullablePatientRecord struct {
	value *PatientRecord
	isSet bool
}

func (v NullablePatientRecord) Get() *PatientRecord {
	return v.value
}

func (v *NullablePatientRecord) Set(val *PatientRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientRecord(val *PatientRecord) *NullablePatientRecord {
	return &NullablePatientRecord{value: val, isSet: true}
}

func (v NullablePatientRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
