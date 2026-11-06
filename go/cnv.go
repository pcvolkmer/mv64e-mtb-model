package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the Cnv type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cnv{}

// Cnv struct for Cnv
type Cnv struct {
	Localization          []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	ReportedFocality      *string                         `json:"reportedFocality,omitempty"`
	ReportedAffectedGenes []Coding                        `json:"reportedAffectedGenes,omitempty"`
	TotalCopyNumber       *int32                          `json:"totalCopyNumber,omitempty"`
	CnA                   *float32                        `json:"cnA,omitempty"`
	ExternalIds           []VariantExternalId             `json:"externalIds,omitempty"`
	StartRange            *CnvStartRange                  `json:"startRange,omitempty"`
	Type                  CnvTypeCoding                   `json:"type"`
	Patient               Reference                       `json:"patient"`
	EndRange              *CnvEndRange                    `json:"endRange,omitempty"`
	Chromosome            Chromosome                      `json:"chromosome"`
	Id                    string                          `json:"id"`
	CnB                   *float32                        `json:"cnB,omitempty"`
	RelativeCopyNumber    *float32                        `json:"relativeCopyNumber,omitempty"`
	CopyNumberNeutralLoH  []Coding                        `json:"copyNumberNeutralLoH,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _Cnv Cnv

// NewCnv instantiates a new Cnv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnv(type_ CnvTypeCoding, patient Reference, chromosome Chromosome, id string) *Cnv {
	this := Cnv{}
	this.Type = type_
	this.Patient = patient
	this.Chromosome = chromosome
	this.Id = id
	return &this
}

// NewCnvWithDefaults instantiates a new Cnv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnvWithDefaults() *Cnv {
	this := Cnv{}
	return &this
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *Cnv) GetLocalization() []BaseVariantLocalizationCoding {
	if o == nil || IsNil(o.Localization) {
		var ret []BaseVariantLocalizationCoding
		return ret
	}
	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetLocalizationOk() ([]BaseVariantLocalizationCoding, bool) {
	if o == nil || IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *Cnv) HasLocalization() bool {
	if o != nil && !IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given []BaseVariantLocalizationCoding and assigns it to the Localization field.
func (o *Cnv) SetLocalization(v []BaseVariantLocalizationCoding) {
	o.Localization = v
}

// GetReportedFocality returns the ReportedFocality field value if set, zero value otherwise.
func (o *Cnv) GetReportedFocality() string {
	if o == nil || IsNil(o.ReportedFocality) {
		var ret string
		return ret
	}
	return *o.ReportedFocality
}

// GetReportedFocalityOk returns a tuple with the ReportedFocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetReportedFocalityOk() (*string, bool) {
	if o == nil || IsNil(o.ReportedFocality) {
		return nil, false
	}
	return o.ReportedFocality, true
}

// HasReportedFocality returns a boolean if a field has been set.
func (o *Cnv) HasReportedFocality() bool {
	if o != nil && !IsNil(o.ReportedFocality) {
		return true
	}

	return false
}

// SetReportedFocality gets a reference to the given string and assigns it to the ReportedFocality field.
func (o *Cnv) SetReportedFocality(v string) {
	o.ReportedFocality = &v
}

// GetReportedAffectedGenes returns the ReportedAffectedGenes field value if set, zero value otherwise.
func (o *Cnv) GetReportedAffectedGenes() []Coding {
	if o == nil || IsNil(o.ReportedAffectedGenes) {
		var ret []Coding
		return ret
	}
	return o.ReportedAffectedGenes
}

// GetReportedAffectedGenesOk returns a tuple with the ReportedAffectedGenes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetReportedAffectedGenesOk() ([]Coding, bool) {
	if o == nil || IsNil(o.ReportedAffectedGenes) {
		return nil, false
	}
	return o.ReportedAffectedGenes, true
}

// HasReportedAffectedGenes returns a boolean if a field has been set.
func (o *Cnv) HasReportedAffectedGenes() bool {
	if o != nil && !IsNil(o.ReportedAffectedGenes) {
		return true
	}

	return false
}

// SetReportedAffectedGenes gets a reference to the given []Coding and assigns it to the ReportedAffectedGenes field.
func (o *Cnv) SetReportedAffectedGenes(v []Coding) {
	o.ReportedAffectedGenes = v
}

// GetTotalCopyNumber returns the TotalCopyNumber field value if set, zero value otherwise.
func (o *Cnv) GetTotalCopyNumber() int32 {
	if o == nil || IsNil(o.TotalCopyNumber) {
		var ret int32
		return ret
	}
	return *o.TotalCopyNumber
}

// GetTotalCopyNumberOk returns a tuple with the TotalCopyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetTotalCopyNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCopyNumber) {
		return nil, false
	}
	return o.TotalCopyNumber, true
}

// HasTotalCopyNumber returns a boolean if a field has been set.
func (o *Cnv) HasTotalCopyNumber() bool {
	if o != nil && !IsNil(o.TotalCopyNumber) {
		return true
	}

	return false
}

// SetTotalCopyNumber gets a reference to the given int32 and assigns it to the TotalCopyNumber field.
func (o *Cnv) SetTotalCopyNumber(v int32) {
	o.TotalCopyNumber = &v
}

// GetCnA returns the CnA field value if set, zero value otherwise.
func (o *Cnv) GetCnA() float32 {
	if o == nil || IsNil(o.CnA) {
		var ret float32
		return ret
	}
	return *o.CnA
}

// GetCnAOk returns a tuple with the CnA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetCnAOk() (*float32, bool) {
	if o == nil || IsNil(o.CnA) {
		return nil, false
	}
	return o.CnA, true
}

// HasCnA returns a boolean if a field has been set.
func (o *Cnv) HasCnA() bool {
	if o != nil && !IsNil(o.CnA) {
		return true
	}

	return false
}

// SetCnA gets a reference to the given float32 and assigns it to the CnA field.
func (o *Cnv) SetCnA(v float32) {
	o.CnA = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *Cnv) GetExternalIds() []VariantExternalId {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []VariantExternalId
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetExternalIdsOk() ([]VariantExternalId, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *Cnv) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []VariantExternalId and assigns it to the ExternalIds field.
func (o *Cnv) SetExternalIds(v []VariantExternalId) {
	o.ExternalIds = v
}

// GetStartRange returns the StartRange field value if set, zero value otherwise.
func (o *Cnv) GetStartRange() CnvStartRange {
	if o == nil || IsNil(o.StartRange) {
		var ret CnvStartRange
		return ret
	}
	return *o.StartRange
}

// GetStartRangeOk returns a tuple with the StartRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetStartRangeOk() (*CnvStartRange, bool) {
	if o == nil || IsNil(o.StartRange) {
		return nil, false
	}
	return o.StartRange, true
}

// HasStartRange returns a boolean if a field has been set.
func (o *Cnv) HasStartRange() bool {
	if o != nil && !IsNil(o.StartRange) {
		return true
	}

	return false
}

// SetStartRange gets a reference to the given CnvStartRange and assigns it to the StartRange field.
func (o *Cnv) SetStartRange(v CnvStartRange) {
	o.StartRange = &v
}

// GetType returns the Type field value
func (o *Cnv) GetType() CnvTypeCoding {
	if o == nil {
		var ret CnvTypeCoding
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Cnv) GetTypeOk() (*CnvTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Cnv) SetType(v CnvTypeCoding) {
	o.Type = v
}

// GetPatient returns the Patient field value
func (o *Cnv) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Cnv) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Cnv) SetPatient(v Reference) {
	o.Patient = v
}

// GetEndRange returns the EndRange field value if set, zero value otherwise.
func (o *Cnv) GetEndRange() CnvEndRange {
	if o == nil || IsNil(o.EndRange) {
		var ret CnvEndRange
		return ret
	}
	return *o.EndRange
}

// GetEndRangeOk returns a tuple with the EndRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetEndRangeOk() (*CnvEndRange, bool) {
	if o == nil || IsNil(o.EndRange) {
		return nil, false
	}
	return o.EndRange, true
}

// HasEndRange returns a boolean if a field has been set.
func (o *Cnv) HasEndRange() bool {
	if o != nil && !IsNil(o.EndRange) {
		return true
	}

	return false
}

// SetEndRange gets a reference to the given CnvEndRange and assigns it to the EndRange field.
func (o *Cnv) SetEndRange(v CnvEndRange) {
	o.EndRange = &v
}

// GetChromosome returns the Chromosome field value
func (o *Cnv) GetChromosome() Chromosome {
	if o == nil {
		var ret Chromosome
		return ret
	}

	return o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value
// and a boolean to check if the value has been set.
func (o *Cnv) GetChromosomeOk() (*Chromosome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chromosome, true
}

// SetChromosome sets field value
func (o *Cnv) SetChromosome(v Chromosome) {
	o.Chromosome = v
}

// GetId returns the Id field value
func (o *Cnv) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cnv) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cnv) SetId(v string) {
	o.Id = v
}

// GetCnB returns the CnB field value if set, zero value otherwise.
func (o *Cnv) GetCnB() float32 {
	if o == nil || IsNil(o.CnB) {
		var ret float32
		return ret
	}
	return *o.CnB
}

// GetCnBOk returns a tuple with the CnB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetCnBOk() (*float32, bool) {
	if o == nil || IsNil(o.CnB) {
		return nil, false
	}
	return o.CnB, true
}

// HasCnB returns a boolean if a field has been set.
func (o *Cnv) HasCnB() bool {
	if o != nil && !IsNil(o.CnB) {
		return true
	}

	return false
}

// SetCnB gets a reference to the given float32 and assigns it to the CnB field.
func (o *Cnv) SetCnB(v float32) {
	o.CnB = &v
}

// GetRelativeCopyNumber returns the RelativeCopyNumber field value if set, zero value otherwise.
func (o *Cnv) GetRelativeCopyNumber() float32 {
	if o == nil || IsNil(o.RelativeCopyNumber) {
		var ret float32
		return ret
	}
	return *o.RelativeCopyNumber
}

// GetRelativeCopyNumberOk returns a tuple with the RelativeCopyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetRelativeCopyNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.RelativeCopyNumber) {
		return nil, false
	}
	return o.RelativeCopyNumber, true
}

// HasRelativeCopyNumber returns a boolean if a field has been set.
func (o *Cnv) HasRelativeCopyNumber() bool {
	if o != nil && !IsNil(o.RelativeCopyNumber) {
		return true
	}

	return false
}

// SetRelativeCopyNumber gets a reference to the given float32 and assigns it to the RelativeCopyNumber field.
func (o *Cnv) SetRelativeCopyNumber(v float32) {
	o.RelativeCopyNumber = &v
}

// GetCopyNumberNeutralLoH returns the CopyNumberNeutralLoH field value if set, zero value otherwise.
func (o *Cnv) GetCopyNumberNeutralLoH() []Coding {
	if o == nil || IsNil(o.CopyNumberNeutralLoH) {
		var ret []Coding
		return ret
	}
	return o.CopyNumberNeutralLoH
}

// GetCopyNumberNeutralLoHOk returns a tuple with the CopyNumberNeutralLoH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cnv) GetCopyNumberNeutralLoHOk() ([]Coding, bool) {
	if o == nil || IsNil(o.CopyNumberNeutralLoH) {
		return nil, false
	}
	return o.CopyNumberNeutralLoH, true
}

// HasCopyNumberNeutralLoH returns a boolean if a field has been set.
func (o *Cnv) HasCopyNumberNeutralLoH() bool {
	if o != nil && !IsNil(o.CopyNumberNeutralLoH) {
		return true
	}

	return false
}

// SetCopyNumberNeutralLoH gets a reference to the given []Coding and assigns it to the CopyNumberNeutralLoH field.
func (o *Cnv) SetCopyNumberNeutralLoH(v []Coding) {
	o.CopyNumberNeutralLoH = v
}

func (o Cnv) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cnv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	if !IsNil(o.ReportedFocality) {
		toSerialize["reportedFocality"] = o.ReportedFocality
	}
	if !IsNil(o.ReportedAffectedGenes) {
		toSerialize["reportedAffectedGenes"] = o.ReportedAffectedGenes
	}
	if !IsNil(o.TotalCopyNumber) {
		toSerialize["totalCopyNumber"] = o.TotalCopyNumber
	}
	if !IsNil(o.CnA) {
		toSerialize["cnA"] = o.CnA
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.StartRange) {
		toSerialize["startRange"] = o.StartRange
	}
	toSerialize["type"] = o.Type
	toSerialize["patient"] = o.Patient
	if !IsNil(o.EndRange) {
		toSerialize["endRange"] = o.EndRange
	}
	toSerialize["chromosome"] = o.Chromosome
	toSerialize["id"] = o.Id
	if !IsNil(o.CnB) {
		toSerialize["cnB"] = o.CnB
	}
	if !IsNil(o.RelativeCopyNumber) {
		toSerialize["relativeCopyNumber"] = o.RelativeCopyNumber
	}
	if !IsNil(o.CopyNumberNeutralLoH) {
		toSerialize["copyNumberNeutralLoH"] = o.CopyNumberNeutralLoH
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cnv) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"patient",
		"chromosome",
		"id",
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

	varCnv := _Cnv{}

	err = json.Unmarshal(data, &varCnv)

	if err != nil {
		return err
	}

	*o = Cnv(varCnv)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "localization")
		delete(additionalProperties, "reportedFocality")
		delete(additionalProperties, "reportedAffectedGenes")
		delete(additionalProperties, "totalCopyNumber")
		delete(additionalProperties, "cnA")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "startRange")
		delete(additionalProperties, "type")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "endRange")
		delete(additionalProperties, "chromosome")
		delete(additionalProperties, "id")
		delete(additionalProperties, "cnB")
		delete(additionalProperties, "relativeCopyNumber")
		delete(additionalProperties, "copyNumberNeutralLoH")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCnv struct {
	value *Cnv
	isSet bool
}

func (v NullableCnv) Get() *Cnv {
	return v.value
}

func (v *NullableCnv) Set(val *Cnv) {
	v.value = val
	v.isSet = true
}

func (v NullableCnv) IsSet() bool {
	return v.isSet
}

func (v *NullableCnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnv(val *Cnv) *NullableCnv {
	return &NullableCnv{value: val, isSet: true}
}

func (v NullableCnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
