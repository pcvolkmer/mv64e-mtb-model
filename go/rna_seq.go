package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the RnaSeq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RnaSeq{}

// RnaSeq struct for RnaSeq
type RnaSeq struct {
	LibrarySize               *int32                          `json:"librarySize,omitempty"`
	Localization              []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	TranscriptsPerMillion     float32                         `json:"transcriptsPerMillion"`
	CohortRanking             *int32                          `json:"cohortRanking,omitempty"`
	RawCounts                 int32                           `json:"rawCounts"`
	Variant                   Reference                       `json:"variant"`
	TranscriptId              *TranscriptId                   `json:"transcriptId,omitempty"`
	ExternalIds               []VariantExternalId             `json:"externalIds,omitempty"`
	Id                        string                          `json:"id"`
	TissueCorrectedExpression *bool                           `json:"tissueCorrectedExpression,omitempty"`
	Gene                      *Coding                         `json:"gene,omitempty"`
	Patient                   Reference                       `json:"patient"`
	AdditionalProperties      map[string]interface{}
}

type _RnaSeq RnaSeq

// NewRnaSeq instantiates a new RnaSeq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaSeq(transcriptsPerMillion float32, rawCounts int32, variant Reference, id string, patient Reference) *RnaSeq {
	this := RnaSeq{}
	this.TranscriptsPerMillion = transcriptsPerMillion
	this.RawCounts = rawCounts
	this.Variant = variant
	this.Id = id
	this.Patient = patient
	return &this
}

// NewRnaSeqWithDefaults instantiates a new RnaSeq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaSeqWithDefaults() *RnaSeq {
	this := RnaSeq{}
	return &this
}

// GetLibrarySize returns the LibrarySize field value if set, zero value otherwise.
func (o *RnaSeq) GetLibrarySize() int32 {
	if o == nil || IsNil(o.LibrarySize) {
		var ret int32
		return ret
	}
	return *o.LibrarySize
}

// GetLibrarySizeOk returns a tuple with the LibrarySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetLibrarySizeOk() (*int32, bool) {
	if o == nil || IsNil(o.LibrarySize) {
		return nil, false
	}
	return o.LibrarySize, true
}

// HasLibrarySize returns a boolean if a field has been set.
func (o *RnaSeq) HasLibrarySize() bool {
	if o != nil && !IsNil(o.LibrarySize) {
		return true
	}

	return false
}

// SetLibrarySize gets a reference to the given int32 and assigns it to the LibrarySize field.
func (o *RnaSeq) SetLibrarySize(v int32) {
	o.LibrarySize = &v
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *RnaSeq) GetLocalization() []BaseVariantLocalizationCoding {
	if o == nil || IsNil(o.Localization) {
		var ret []BaseVariantLocalizationCoding
		return ret
	}
	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetLocalizationOk() ([]BaseVariantLocalizationCoding, bool) {
	if o == nil || IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *RnaSeq) HasLocalization() bool {
	if o != nil && !IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given []BaseVariantLocalizationCoding and assigns it to the Localization field.
func (o *RnaSeq) SetLocalization(v []BaseVariantLocalizationCoding) {
	o.Localization = v
}

// GetTranscriptsPerMillion returns the TranscriptsPerMillion field value
func (o *RnaSeq) GetTranscriptsPerMillion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TranscriptsPerMillion
}

// GetTranscriptsPerMillionOk returns a tuple with the TranscriptsPerMillion field value
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetTranscriptsPerMillionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranscriptsPerMillion, true
}

// SetTranscriptsPerMillion sets field value
func (o *RnaSeq) SetTranscriptsPerMillion(v float32) {
	o.TranscriptsPerMillion = v
}

// GetCohortRanking returns the CohortRanking field value if set, zero value otherwise.
func (o *RnaSeq) GetCohortRanking() int32 {
	if o == nil || IsNil(o.CohortRanking) {
		var ret int32
		return ret
	}
	return *o.CohortRanking
}

// GetCohortRankingOk returns a tuple with the CohortRanking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetCohortRankingOk() (*int32, bool) {
	if o == nil || IsNil(o.CohortRanking) {
		return nil, false
	}
	return o.CohortRanking, true
}

// HasCohortRanking returns a boolean if a field has been set.
func (o *RnaSeq) HasCohortRanking() bool {
	if o != nil && !IsNil(o.CohortRanking) {
		return true
	}

	return false
}

// SetCohortRanking gets a reference to the given int32 and assigns it to the CohortRanking field.
func (o *RnaSeq) SetCohortRanking(v int32) {
	o.CohortRanking = &v
}

// GetRawCounts returns the RawCounts field value
func (o *RnaSeq) GetRawCounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RawCounts
}

// GetRawCountsOk returns a tuple with the RawCounts field value
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetRawCountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawCounts, true
}

// SetRawCounts sets field value
func (o *RnaSeq) SetRawCounts(v int32) {
	o.RawCounts = v
}

// GetVariant returns the Variant field value
func (o *RnaSeq) GetVariant() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetVariantOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *RnaSeq) SetVariant(v Reference) {
	o.Variant = v
}

// GetTranscriptId returns the TranscriptId field value if set, zero value otherwise.
func (o *RnaSeq) GetTranscriptId() TranscriptId {
	if o == nil || IsNil(o.TranscriptId) {
		var ret TranscriptId
		return ret
	}
	return *o.TranscriptId
}

// GetTranscriptIdOk returns a tuple with the TranscriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetTranscriptIdOk() (*TranscriptId, bool) {
	if o == nil || IsNil(o.TranscriptId) {
		return nil, false
	}
	return o.TranscriptId, true
}

// HasTranscriptId returns a boolean if a field has been set.
func (o *RnaSeq) HasTranscriptId() bool {
	if o != nil && !IsNil(o.TranscriptId) {
		return true
	}

	return false
}

// SetTranscriptId gets a reference to the given TranscriptId and assigns it to the TranscriptId field.
func (o *RnaSeq) SetTranscriptId(v TranscriptId) {
	o.TranscriptId = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *RnaSeq) GetExternalIds() []VariantExternalId {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []VariantExternalId
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetExternalIdsOk() ([]VariantExternalId, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *RnaSeq) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []VariantExternalId and assigns it to the ExternalIds field.
func (o *RnaSeq) SetExternalIds(v []VariantExternalId) {
	o.ExternalIds = v
}

// GetId returns the Id field value
func (o *RnaSeq) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RnaSeq) SetId(v string) {
	o.Id = v
}

// GetTissueCorrectedExpression returns the TissueCorrectedExpression field value if set, zero value otherwise.
func (o *RnaSeq) GetTissueCorrectedExpression() bool {
	if o == nil || IsNil(o.TissueCorrectedExpression) {
		var ret bool
		return ret
	}
	return *o.TissueCorrectedExpression
}

// GetTissueCorrectedExpressionOk returns a tuple with the TissueCorrectedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetTissueCorrectedExpressionOk() (*bool, bool) {
	if o == nil || IsNil(o.TissueCorrectedExpression) {
		return nil, false
	}
	return o.TissueCorrectedExpression, true
}

// HasTissueCorrectedExpression returns a boolean if a field has been set.
func (o *RnaSeq) HasTissueCorrectedExpression() bool {
	if o != nil && !IsNil(o.TissueCorrectedExpression) {
		return true
	}

	return false
}

// SetTissueCorrectedExpression gets a reference to the given bool and assigns it to the TissueCorrectedExpression field.
func (o *RnaSeq) SetTissueCorrectedExpression(v bool) {
	o.TissueCorrectedExpression = &v
}

// GetGene returns the Gene field value if set, zero value otherwise.
func (o *RnaSeq) GetGene() Coding {
	if o == nil || IsNil(o.Gene) {
		var ret Coding
		return ret
	}
	return *o.Gene
}

// GetGeneOk returns a tuple with the Gene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetGeneOk() (*Coding, bool) {
	if o == nil || IsNil(o.Gene) {
		return nil, false
	}
	return o.Gene, true
}

// HasGene returns a boolean if a field has been set.
func (o *RnaSeq) HasGene() bool {
	if o != nil && !IsNil(o.Gene) {
		return true
	}

	return false
}

// SetGene gets a reference to the given Coding and assigns it to the Gene field.
func (o *RnaSeq) SetGene(v Coding) {
	o.Gene = &v
}

// GetPatient returns the Patient field value
func (o *RnaSeq) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *RnaSeq) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *RnaSeq) SetPatient(v Reference) {
	o.Patient = v
}

func (o RnaSeq) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RnaSeq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LibrarySize) {
		toSerialize["librarySize"] = o.LibrarySize
	}
	if !IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	toSerialize["transcriptsPerMillion"] = o.TranscriptsPerMillion
	if !IsNil(o.CohortRanking) {
		toSerialize["cohortRanking"] = o.CohortRanking
	}
	toSerialize["rawCounts"] = o.RawCounts
	toSerialize["variant"] = o.Variant
	if !IsNil(o.TranscriptId) {
		toSerialize["transcriptId"] = o.TranscriptId
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.TissueCorrectedExpression) {
		toSerialize["tissueCorrectedExpression"] = o.TissueCorrectedExpression
	}
	if !IsNil(o.Gene) {
		toSerialize["gene"] = o.Gene
	}
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RnaSeq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transcriptsPerMillion",
		"rawCounts",
		"variant",
		"id",
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

	varRnaSeq := _RnaSeq{}

	err = json.Unmarshal(data, &varRnaSeq)

	if err != nil {
		return err
	}

	*o = RnaSeq(varRnaSeq)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "librarySize")
		delete(additionalProperties, "localization")
		delete(additionalProperties, "transcriptsPerMillion")
		delete(additionalProperties, "cohortRanking")
		delete(additionalProperties, "rawCounts")
		delete(additionalProperties, "variant")
		delete(additionalProperties, "transcriptId")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "id")
		delete(additionalProperties, "tissueCorrectedExpression")
		delete(additionalProperties, "gene")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRnaSeq struct {
	value *RnaSeq
	isSet bool
}

func (v NullableRnaSeq) Get() *RnaSeq {
	return v.value
}

func (v *NullableRnaSeq) Set(val *RnaSeq) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaSeq) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaSeq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaSeq(val *RnaSeq) *NullableRnaSeq {
	return &NullableRnaSeq{value: val, isSet: true}
}

func (v NullableRnaSeq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaSeq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
