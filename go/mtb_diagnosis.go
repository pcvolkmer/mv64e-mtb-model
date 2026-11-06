package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MtbDiagnosis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbDiagnosis{}

// MtbDiagnosis struct for MtbDiagnosis
type MtbDiagnosis struct {
	Grading                  *MtbDiagnosisGrading                        `json:"grading,omitempty"`
	GermlineCodes            []Coding                                    `json:"germlineCodes,omitempty"`
	Code                     Coding                                      `json:"code"`
	Id                       string                                      `json:"id"`
	RecordedOn               string                                      `json:"recordedOn"`
	Type                     MtbDiagnosisType                            `json:"type"`
	Patient                  Reference                                   `json:"patient"`
	Topography               Coding                                      `json:"topography"`
	Staging                  *MtbDiagnosisStaging                        `json:"staging,omitempty"`
	Histology                []Reference                                 `json:"histology,omitempty"`
	GuidelineTreatmentStatus *MtbDiagnosisGuidelineTreatmentStatusCoding `json:"guidelineTreatmentStatus,omitempty"`
	Notes                    []string                                    `json:"notes,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _MtbDiagnosis MtbDiagnosis

// NewMtbDiagnosis instantiates a new MtbDiagnosis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbDiagnosis(code Coding, id string, recordedOn string, type_ MtbDiagnosisType, patient Reference, topography Coding) *MtbDiagnosis {
	this := MtbDiagnosis{}
	this.Code = code
	this.Id = id
	this.RecordedOn = recordedOn
	this.Type = type_
	this.Patient = patient
	this.Topography = topography
	return &this
}

// NewMtbDiagnosisWithDefaults instantiates a new MtbDiagnosis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbDiagnosisWithDefaults() *MtbDiagnosis {
	this := MtbDiagnosis{}
	return &this
}

// GetGrading returns the Grading field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetGrading() MtbDiagnosisGrading {
	if o == nil || IsNil(o.Grading) {
		var ret MtbDiagnosisGrading
		return ret
	}
	return *o.Grading
}

// GetGradingOk returns a tuple with the Grading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetGradingOk() (*MtbDiagnosisGrading, bool) {
	if o == nil || IsNil(o.Grading) {
		return nil, false
	}
	return o.Grading, true
}

// HasGrading returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasGrading() bool {
	if o != nil && !IsNil(o.Grading) {
		return true
	}

	return false
}

// SetGrading gets a reference to the given MtbDiagnosisGrading and assigns it to the Grading field.
func (o *MtbDiagnosis) SetGrading(v MtbDiagnosisGrading) {
	o.Grading = &v
}

// GetGermlineCodes returns the GermlineCodes field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetGermlineCodes() []Coding {
	if o == nil || IsNil(o.GermlineCodes) {
		var ret []Coding
		return ret
	}
	return o.GermlineCodes
}

// GetGermlineCodesOk returns a tuple with the GermlineCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetGermlineCodesOk() ([]Coding, bool) {
	if o == nil || IsNil(o.GermlineCodes) {
		return nil, false
	}
	return o.GermlineCodes, true
}

// HasGermlineCodes returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasGermlineCodes() bool {
	if o != nil && !IsNil(o.GermlineCodes) {
		return true
	}

	return false
}

// SetGermlineCodes gets a reference to the given []Coding and assigns it to the GermlineCodes field.
func (o *MtbDiagnosis) SetGermlineCodes(v []Coding) {
	o.GermlineCodes = v
}

// GetCode returns the Code field value
func (o *MtbDiagnosis) GetCode() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetCodeOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *MtbDiagnosis) SetCode(v Coding) {
	o.Code = v
}

// GetId returns the Id field value
func (o *MtbDiagnosis) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbDiagnosis) SetId(v string) {
	o.Id = v
}

// GetRecordedOn returns the RecordedOn field value
func (o *MtbDiagnosis) GetRecordedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordedOn
}

// GetRecordedOnOk returns a tuple with the RecordedOn field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetRecordedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedOn, true
}

// SetRecordedOn sets field value
func (o *MtbDiagnosis) SetRecordedOn(v string) {
	o.RecordedOn = v
}

// GetType returns the Type field value
func (o *MtbDiagnosis) GetType() MtbDiagnosisType {
	if o == nil {
		var ret MtbDiagnosisType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetTypeOk() (*MtbDiagnosisType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MtbDiagnosis) SetType(v MtbDiagnosisType) {
	o.Type = v
}

// GetPatient returns the Patient field value
func (o *MtbDiagnosis) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbDiagnosis) SetPatient(v Reference) {
	o.Patient = v
}

// GetTopography returns the Topography field value
func (o *MtbDiagnosis) GetTopography() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Topography
}

// GetTopographyOk returns a tuple with the Topography field value
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetTopographyOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topography, true
}

// SetTopography sets field value
func (o *MtbDiagnosis) SetTopography(v Coding) {
	o.Topography = v
}

// GetStaging returns the Staging field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetStaging() MtbDiagnosisStaging {
	if o == nil || IsNil(o.Staging) {
		var ret MtbDiagnosisStaging
		return ret
	}
	return *o.Staging
}

// GetStagingOk returns a tuple with the Staging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetStagingOk() (*MtbDiagnosisStaging, bool) {
	if o == nil || IsNil(o.Staging) {
		return nil, false
	}
	return o.Staging, true
}

// HasStaging returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasStaging() bool {
	if o != nil && !IsNil(o.Staging) {
		return true
	}

	return false
}

// SetStaging gets a reference to the given MtbDiagnosisStaging and assigns it to the Staging field.
func (o *MtbDiagnosis) SetStaging(v MtbDiagnosisStaging) {
	o.Staging = &v
}

// GetHistology returns the Histology field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetHistology() []Reference {
	if o == nil || IsNil(o.Histology) {
		var ret []Reference
		return ret
	}
	return o.Histology
}

// GetHistologyOk returns a tuple with the Histology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetHistologyOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Histology) {
		return nil, false
	}
	return o.Histology, true
}

// HasHistology returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasHistology() bool {
	if o != nil && !IsNil(o.Histology) {
		return true
	}

	return false
}

// SetHistology gets a reference to the given []Reference and assigns it to the Histology field.
func (o *MtbDiagnosis) SetHistology(v []Reference) {
	o.Histology = v
}

// GetGuidelineTreatmentStatus returns the GuidelineTreatmentStatus field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetGuidelineTreatmentStatus() MtbDiagnosisGuidelineTreatmentStatusCoding {
	if o == nil || IsNil(o.GuidelineTreatmentStatus) {
		var ret MtbDiagnosisGuidelineTreatmentStatusCoding
		return ret
	}
	return *o.GuidelineTreatmentStatus
}

// GetGuidelineTreatmentStatusOk returns a tuple with the GuidelineTreatmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetGuidelineTreatmentStatusOk() (*MtbDiagnosisGuidelineTreatmentStatusCoding, bool) {
	if o == nil || IsNil(o.GuidelineTreatmentStatus) {
		return nil, false
	}
	return o.GuidelineTreatmentStatus, true
}

// HasGuidelineTreatmentStatus returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasGuidelineTreatmentStatus() bool {
	if o != nil && !IsNil(o.GuidelineTreatmentStatus) {
		return true
	}

	return false
}

// SetGuidelineTreatmentStatus gets a reference to the given MtbDiagnosisGuidelineTreatmentStatusCoding and assigns it to the GuidelineTreatmentStatus field.
func (o *MtbDiagnosis) SetGuidelineTreatmentStatus(v MtbDiagnosisGuidelineTreatmentStatusCoding) {
	o.GuidelineTreatmentStatus = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *MtbDiagnosis) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbDiagnosis) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MtbDiagnosis) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *MtbDiagnosis) SetNotes(v []string) {
	o.Notes = v
}

func (o MtbDiagnosis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbDiagnosis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grading) {
		toSerialize["grading"] = o.Grading
	}
	if !IsNil(o.GermlineCodes) {
		toSerialize["germlineCodes"] = o.GermlineCodes
	}
	toSerialize["code"] = o.Code
	toSerialize["id"] = o.Id
	toSerialize["recordedOn"] = o.RecordedOn
	toSerialize["type"] = o.Type
	toSerialize["patient"] = o.Patient
	toSerialize["topography"] = o.Topography
	if !IsNil(o.Staging) {
		toSerialize["staging"] = o.Staging
	}
	if !IsNil(o.Histology) {
		toSerialize["histology"] = o.Histology
	}
	if !IsNil(o.GuidelineTreatmentStatus) {
		toSerialize["guidelineTreatmentStatus"] = o.GuidelineTreatmentStatus
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MtbDiagnosis) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"id",
		"recordedOn",
		"type",
		"patient",
		"topography",
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

	varMtbDiagnosis := _MtbDiagnosis{}

	err = json.Unmarshal(data, &varMtbDiagnosis)

	if err != nil {
		return err
	}

	*o = MtbDiagnosis(varMtbDiagnosis)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grading")
		delete(additionalProperties, "germlineCodes")
		delete(additionalProperties, "code")
		delete(additionalProperties, "id")
		delete(additionalProperties, "recordedOn")
		delete(additionalProperties, "type")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "topography")
		delete(additionalProperties, "staging")
		delete(additionalProperties, "histology")
		delete(additionalProperties, "guidelineTreatmentStatus")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMtbDiagnosis struct {
	value *MtbDiagnosis
	isSet bool
}

func (v NullableMtbDiagnosis) Get() *MtbDiagnosis {
	return v.value
}

func (v *NullableMtbDiagnosis) Set(val *MtbDiagnosis) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbDiagnosis) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbDiagnosis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbDiagnosis(val *MtbDiagnosis) *NullableMtbDiagnosis {
	return &NullableMtbDiagnosis{value: val, isSet: true}
}

func (v NullableMtbDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbDiagnosis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
