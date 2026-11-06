package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the Snv type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snv{}

// Snv struct for Snv
type Snv struct {
	Localization         []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	ProteinChange        *string                         `json:"proteinChange,omitempty"`
	Interpretation       *ClinVarCoding                  `json:"interpretation,omitempty"`
	TranscriptId         TranscriptId                    `json:"transcriptId"`
	ExternalIds          []VariantExternalId             `json:"externalIds,omitempty"`
	Chromosome           Chromosome                      `json:"chromosome"`
	Id                   string                          `json:"id"`
	AltAllele            string                          `json:"altAllele"`
	Position             CnvEndRange                     `json:"position"`
	Gene                 Coding                          `json:"gene"`
	ExonId               *string                         `json:"exonId,omitempty"`
	DnaChange            string                          `json:"dnaChange"`
	Patient              Reference                       `json:"patient"`
	RefAllele            string                          `json:"refAllele"`
	ReadDepth            int32                           `json:"readDepth"`
	AllelicFrequency     float32                         `json:"allelicFrequency"`
	AdditionalProperties map[string]interface{}
}

type _Snv Snv

// NewSnv instantiates a new Snv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnv(transcriptId TranscriptId, chromosome Chromosome, id string, altAllele string, position CnvEndRange, gene Coding, dnaChange string, patient Reference, refAllele string, readDepth int32, allelicFrequency float32) *Snv {
	this := Snv{}
	this.TranscriptId = transcriptId
	this.Chromosome = chromosome
	this.Id = id
	this.AltAllele = altAllele
	this.Position = position
	this.Gene = gene
	this.DnaChange = dnaChange
	this.Patient = patient
	this.RefAllele = refAllele
	this.ReadDepth = readDepth
	this.AllelicFrequency = allelicFrequency
	return &this
}

// NewSnvWithDefaults instantiates a new Snv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnvWithDefaults() *Snv {
	this := Snv{}
	return &this
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *Snv) GetLocalization() []BaseVariantLocalizationCoding {
	if o == nil || IsNil(o.Localization) {
		var ret []BaseVariantLocalizationCoding
		return ret
	}
	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snv) GetLocalizationOk() ([]BaseVariantLocalizationCoding, bool) {
	if o == nil || IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *Snv) HasLocalization() bool {
	if o != nil && !IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given []BaseVariantLocalizationCoding and assigns it to the Localization field.
func (o *Snv) SetLocalization(v []BaseVariantLocalizationCoding) {
	o.Localization = v
}

// GetProteinChange returns the ProteinChange field value if set, zero value otherwise.
func (o *Snv) GetProteinChange() string {
	if o == nil || IsNil(o.ProteinChange) {
		var ret string
		return ret
	}
	return *o.ProteinChange
}

// GetProteinChangeOk returns a tuple with the ProteinChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snv) GetProteinChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ProteinChange) {
		return nil, false
	}
	return o.ProteinChange, true
}

// HasProteinChange returns a boolean if a field has been set.
func (o *Snv) HasProteinChange() bool {
	if o != nil && !IsNil(o.ProteinChange) {
		return true
	}

	return false
}

// SetProteinChange gets a reference to the given string and assigns it to the ProteinChange field.
func (o *Snv) SetProteinChange(v string) {
	o.ProteinChange = &v
}

// GetInterpretation returns the Interpretation field value if set, zero value otherwise.
func (o *Snv) GetInterpretation() ClinVarCoding {
	if o == nil || IsNil(o.Interpretation) {
		var ret ClinVarCoding
		return ret
	}
	return *o.Interpretation
}

// GetInterpretationOk returns a tuple with the Interpretation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snv) GetInterpretationOk() (*ClinVarCoding, bool) {
	if o == nil || IsNil(o.Interpretation) {
		return nil, false
	}
	return o.Interpretation, true
}

// HasInterpretation returns a boolean if a field has been set.
func (o *Snv) HasInterpretation() bool {
	if o != nil && !IsNil(o.Interpretation) {
		return true
	}

	return false
}

// SetInterpretation gets a reference to the given ClinVarCoding and assigns it to the Interpretation field.
func (o *Snv) SetInterpretation(v ClinVarCoding) {
	o.Interpretation = &v
}

// GetTranscriptId returns the TranscriptId field value
func (o *Snv) GetTranscriptId() TranscriptId {
	if o == nil {
		var ret TranscriptId
		return ret
	}

	return o.TranscriptId
}

// GetTranscriptIdOk returns a tuple with the TranscriptId field value
// and a boolean to check if the value has been set.
func (o *Snv) GetTranscriptIdOk() (*TranscriptId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranscriptId, true
}

// SetTranscriptId sets field value
func (o *Snv) SetTranscriptId(v TranscriptId) {
	o.TranscriptId = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *Snv) GetExternalIds() []VariantExternalId {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []VariantExternalId
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snv) GetExternalIdsOk() ([]VariantExternalId, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *Snv) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []VariantExternalId and assigns it to the ExternalIds field.
func (o *Snv) SetExternalIds(v []VariantExternalId) {
	o.ExternalIds = v
}

// GetChromosome returns the Chromosome field value
func (o *Snv) GetChromosome() Chromosome {
	if o == nil {
		var ret Chromosome
		return ret
	}

	return o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value
// and a boolean to check if the value has been set.
func (o *Snv) GetChromosomeOk() (*Chromosome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chromosome, true
}

// SetChromosome sets field value
func (o *Snv) SetChromosome(v Chromosome) {
	o.Chromosome = v
}

// GetId returns the Id field value
func (o *Snv) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Snv) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Snv) SetId(v string) {
	o.Id = v
}

// GetAltAllele returns the AltAllele field value
func (o *Snv) GetAltAllele() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AltAllele
}

// GetAltAlleleOk returns a tuple with the AltAllele field value
// and a boolean to check if the value has been set.
func (o *Snv) GetAltAlleleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AltAllele, true
}

// SetAltAllele sets field value
func (o *Snv) SetAltAllele(v string) {
	o.AltAllele = v
}

// GetPosition returns the Position field value
func (o *Snv) GetPosition() CnvEndRange {
	if o == nil {
		var ret CnvEndRange
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *Snv) GetPositionOk() (*CnvEndRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *Snv) SetPosition(v CnvEndRange) {
	o.Position = v
}

// GetGene returns the Gene field value
func (o *Snv) GetGene() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Gene
}

// GetGeneOk returns a tuple with the Gene field value
// and a boolean to check if the value has been set.
func (o *Snv) GetGeneOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gene, true
}

// SetGene sets field value
func (o *Snv) SetGene(v Coding) {
	o.Gene = v
}

// GetExonId returns the ExonId field value if set, zero value otherwise.
func (o *Snv) GetExonId() string {
	if o == nil || IsNil(o.ExonId) {
		var ret string
		return ret
	}
	return *o.ExonId
}

// GetExonIdOk returns a tuple with the ExonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snv) GetExonIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExonId) {
		return nil, false
	}
	return o.ExonId, true
}

// HasExonId returns a boolean if a field has been set.
func (o *Snv) HasExonId() bool {
	if o != nil && !IsNil(o.ExonId) {
		return true
	}

	return false
}

// SetExonId gets a reference to the given string and assigns it to the ExonId field.
func (o *Snv) SetExonId(v string) {
	o.ExonId = &v
}

// GetDnaChange returns the DnaChange field value
func (o *Snv) GetDnaChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DnaChange
}

// GetDnaChangeOk returns a tuple with the DnaChange field value
// and a boolean to check if the value has been set.
func (o *Snv) GetDnaChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnaChange, true
}

// SetDnaChange sets field value
func (o *Snv) SetDnaChange(v string) {
	o.DnaChange = v
}

// GetPatient returns the Patient field value
func (o *Snv) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Snv) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Snv) SetPatient(v Reference) {
	o.Patient = v
}

// GetRefAllele returns the RefAllele field value
func (o *Snv) GetRefAllele() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefAllele
}

// GetRefAlleleOk returns a tuple with the RefAllele field value
// and a boolean to check if the value has been set.
func (o *Snv) GetRefAlleleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefAllele, true
}

// SetRefAllele sets field value
func (o *Snv) SetRefAllele(v string) {
	o.RefAllele = v
}

// GetReadDepth returns the ReadDepth field value
func (o *Snv) GetReadDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReadDepth
}

// GetReadDepthOk returns a tuple with the ReadDepth field value
// and a boolean to check if the value has been set.
func (o *Snv) GetReadDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadDepth, true
}

// SetReadDepth sets field value
func (o *Snv) SetReadDepth(v int32) {
	o.ReadDepth = v
}

// GetAllelicFrequency returns the AllelicFrequency field value
func (o *Snv) GetAllelicFrequency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AllelicFrequency
}

// GetAllelicFrequencyOk returns a tuple with the AllelicFrequency field value
// and a boolean to check if the value has been set.
func (o *Snv) GetAllelicFrequencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllelicFrequency, true
}

// SetAllelicFrequency sets field value
func (o *Snv) SetAllelicFrequency(v float32) {
	o.AllelicFrequency = v
}

func (o Snv) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	if !IsNil(o.ProteinChange) {
		toSerialize["proteinChange"] = o.ProteinChange
	}
	if !IsNil(o.Interpretation) {
		toSerialize["interpretation"] = o.Interpretation
	}
	toSerialize["transcriptId"] = o.TranscriptId
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	toSerialize["chromosome"] = o.Chromosome
	toSerialize["id"] = o.Id
	toSerialize["altAllele"] = o.AltAllele
	toSerialize["position"] = o.Position
	toSerialize["gene"] = o.Gene
	if !IsNil(o.ExonId) {
		toSerialize["exonId"] = o.ExonId
	}
	toSerialize["dnaChange"] = o.DnaChange
	toSerialize["patient"] = o.Patient
	toSerialize["refAllele"] = o.RefAllele
	toSerialize["readDepth"] = o.ReadDepth
	toSerialize["allelicFrequency"] = o.AllelicFrequency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Snv) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transcriptId",
		"chromosome",
		"id",
		"altAllele",
		"position",
		"gene",
		"dnaChange",
		"patient",
		"refAllele",
		"readDepth",
		"allelicFrequency",
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

	varSnv := _Snv{}

	err = json.Unmarshal(data, &varSnv)

	if err != nil {
		return err
	}

	*o = Snv(varSnv)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "localization")
		delete(additionalProperties, "proteinChange")
		delete(additionalProperties, "interpretation")
		delete(additionalProperties, "transcriptId")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "chromosome")
		delete(additionalProperties, "id")
		delete(additionalProperties, "altAllele")
		delete(additionalProperties, "position")
		delete(additionalProperties, "gene")
		delete(additionalProperties, "exonId")
		delete(additionalProperties, "dnaChange")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "refAllele")
		delete(additionalProperties, "readDepth")
		delete(additionalProperties, "allelicFrequency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnv struct {
	value *Snv
	isSet bool
}

func (v NullableSnv) Get() *Snv {
	return v.value
}

func (v *NullableSnv) Set(val *Snv) {
	v.value = val
	v.isSet = true
}

func (v NullableSnv) IsSet() bool {
	return v.isSet
}

func (v *NullableSnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnv(val *Snv) *NullableSnv {
	return &NullableSnv{value: val, isSet: true}
}

func (v NullableSnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
