package mv64e_model

import (
	"encoding/json"
)

// checks if the SomaticNgsReportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SomaticNgsReportResults{}

// SomaticNgsReportResults struct for SomaticNgsReportResults
type SomaticNgsReportResults struct {
	RnaSeqs            []RnaSeq          `json:"rnaSeqs,omitempty"`
	Tmb                *Tmb              `json:"tmb,omitempty"`
	SimpleVariants     []Snv             `json:"simpleVariants,omitempty"`
	HrdScore           *HrdScore         `json:"hrdScore,omitempty"`
	CopyNumberVariants []Cnv             `json:"copyNumberVariants,omitempty"`
	TumorCellContent   *TumorCellContent `json:"tumorCellContent,omitempty"`
	DnaFusions         []DnaFusion       `json:"dnaFusions,omitempty"`
	RnaFusions         []RnaFusion       `json:"rnaFusions,omitempty"`
	Brcaness           *Brcaness         `json:"brcaness,omitempty"`
}

// NewSomaticNgsReportResults instantiates a new SomaticNgsReportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSomaticNgsReportResults() *SomaticNgsReportResults {
	this := SomaticNgsReportResults{}
	return &this
}

// NewSomaticNgsReportResultsWithDefaults instantiates a new SomaticNgsReportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSomaticNgsReportResultsWithDefaults() *SomaticNgsReportResults {
	this := SomaticNgsReportResults{}
	return &this
}

// GetRnaSeqs returns the RnaSeqs field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetRnaSeqs() []RnaSeq {
	if o == nil || IsNil(o.RnaSeqs) {
		var ret []RnaSeq
		return ret
	}
	return o.RnaSeqs
}

// GetRnaSeqsOk returns a tuple with the RnaSeqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetRnaSeqsOk() ([]RnaSeq, bool) {
	if o == nil || IsNil(o.RnaSeqs) {
		return nil, false
	}
	return o.RnaSeqs, true
}

// HasRnaSeqs returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasRnaSeqs() bool {
	if o != nil && !IsNil(o.RnaSeqs) {
		return true
	}

	return false
}

// SetRnaSeqs gets a reference to the given []RnaSeq and assigns it to the RnaSeqs field.
func (o *SomaticNgsReportResults) SetRnaSeqs(v []RnaSeq) {
	o.RnaSeqs = v
}

// GetTmb returns the Tmb field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetTmb() Tmb {
	if o == nil || IsNil(o.Tmb) {
		var ret Tmb
		return ret
	}
	return *o.Tmb
}

// GetTmbOk returns a tuple with the Tmb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetTmbOk() (*Tmb, bool) {
	if o == nil || IsNil(o.Tmb) {
		return nil, false
	}
	return o.Tmb, true
}

// HasTmb returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasTmb() bool {
	if o != nil && !IsNil(o.Tmb) {
		return true
	}

	return false
}

// SetTmb gets a reference to the given Tmb and assigns it to the Tmb field.
func (o *SomaticNgsReportResults) SetTmb(v Tmb) {
	o.Tmb = &v
}

// GetSimpleVariants returns the SimpleVariants field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetSimpleVariants() []Snv {
	if o == nil || IsNil(o.SimpleVariants) {
		var ret []Snv
		return ret
	}
	return o.SimpleVariants
}

// GetSimpleVariantsOk returns a tuple with the SimpleVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetSimpleVariantsOk() ([]Snv, bool) {
	if o == nil || IsNil(o.SimpleVariants) {
		return nil, false
	}
	return o.SimpleVariants, true
}

// HasSimpleVariants returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasSimpleVariants() bool {
	if o != nil && !IsNil(o.SimpleVariants) {
		return true
	}

	return false
}

// SetSimpleVariants gets a reference to the given []Snv and assigns it to the SimpleVariants field.
func (o *SomaticNgsReportResults) SetSimpleVariants(v []Snv) {
	o.SimpleVariants = v
}

// GetHrdScore returns the HrdScore field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetHrdScore() HrdScore {
	if o == nil || IsNil(o.HrdScore) {
		var ret HrdScore
		return ret
	}
	return *o.HrdScore
}

// GetHrdScoreOk returns a tuple with the HrdScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetHrdScoreOk() (*HrdScore, bool) {
	if o == nil || IsNil(o.HrdScore) {
		return nil, false
	}
	return o.HrdScore, true
}

// HasHrdScore returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasHrdScore() bool {
	if o != nil && !IsNil(o.HrdScore) {
		return true
	}

	return false
}

// SetHrdScore gets a reference to the given HrdScore and assigns it to the HrdScore field.
func (o *SomaticNgsReportResults) SetHrdScore(v HrdScore) {
	o.HrdScore = &v
}

// GetCopyNumberVariants returns the CopyNumberVariants field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetCopyNumberVariants() []Cnv {
	if o == nil || IsNil(o.CopyNumberVariants) {
		var ret []Cnv
		return ret
	}
	return o.CopyNumberVariants
}

// GetCopyNumberVariantsOk returns a tuple with the CopyNumberVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetCopyNumberVariantsOk() ([]Cnv, bool) {
	if o == nil || IsNil(o.CopyNumberVariants) {
		return nil, false
	}
	return o.CopyNumberVariants, true
}

// HasCopyNumberVariants returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasCopyNumberVariants() bool {
	if o != nil && !IsNil(o.CopyNumberVariants) {
		return true
	}

	return false
}

// SetCopyNumberVariants gets a reference to the given []Cnv and assigns it to the CopyNumberVariants field.
func (o *SomaticNgsReportResults) SetCopyNumberVariants(v []Cnv) {
	o.CopyNumberVariants = v
}

// GetTumorCellContent returns the TumorCellContent field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetTumorCellContent() TumorCellContent {
	if o == nil || IsNil(o.TumorCellContent) {
		var ret TumorCellContent
		return ret
	}
	return *o.TumorCellContent
}

// GetTumorCellContentOk returns a tuple with the TumorCellContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetTumorCellContentOk() (*TumorCellContent, bool) {
	if o == nil || IsNil(o.TumorCellContent) {
		return nil, false
	}
	return o.TumorCellContent, true
}

// HasTumorCellContent returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasTumorCellContent() bool {
	if o != nil && !IsNil(o.TumorCellContent) {
		return true
	}

	return false
}

// SetTumorCellContent gets a reference to the given TumorCellContent and assigns it to the TumorCellContent field.
func (o *SomaticNgsReportResults) SetTumorCellContent(v TumorCellContent) {
	o.TumorCellContent = &v
}

// GetDnaFusions returns the DnaFusions field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetDnaFusions() []DnaFusion {
	if o == nil || IsNil(o.DnaFusions) {
		var ret []DnaFusion
		return ret
	}
	return o.DnaFusions
}

// GetDnaFusionsOk returns a tuple with the DnaFusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetDnaFusionsOk() ([]DnaFusion, bool) {
	if o == nil || IsNil(o.DnaFusions) {
		return nil, false
	}
	return o.DnaFusions, true
}

// HasDnaFusions returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasDnaFusions() bool {
	if o != nil && !IsNil(o.DnaFusions) {
		return true
	}

	return false
}

// SetDnaFusions gets a reference to the given []DnaFusion and assigns it to the DnaFusions field.
func (o *SomaticNgsReportResults) SetDnaFusions(v []DnaFusion) {
	o.DnaFusions = v
}

// GetRnaFusions returns the RnaFusions field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetRnaFusions() []RnaFusion {
	if o == nil || IsNil(o.RnaFusions) {
		var ret []RnaFusion
		return ret
	}
	return o.RnaFusions
}

// GetRnaFusionsOk returns a tuple with the RnaFusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetRnaFusionsOk() ([]RnaFusion, bool) {
	if o == nil || IsNil(o.RnaFusions) {
		return nil, false
	}
	return o.RnaFusions, true
}

// HasRnaFusions returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasRnaFusions() bool {
	if o != nil && !IsNil(o.RnaFusions) {
		return true
	}

	return false
}

// SetRnaFusions gets a reference to the given []RnaFusion and assigns it to the RnaFusions field.
func (o *SomaticNgsReportResults) SetRnaFusions(v []RnaFusion) {
	o.RnaFusions = v
}

// GetBrcaness returns the Brcaness field value if set, zero value otherwise.
func (o *SomaticNgsReportResults) GetBrcaness() Brcaness {
	if o == nil || IsNil(o.Brcaness) {
		var ret Brcaness
		return ret
	}
	return *o.Brcaness
}

// GetBrcanessOk returns a tuple with the Brcaness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SomaticNgsReportResults) GetBrcanessOk() (*Brcaness, bool) {
	if o == nil || IsNil(o.Brcaness) {
		return nil, false
	}
	return o.Brcaness, true
}

// HasBrcaness returns a boolean if a field has been set.
func (o *SomaticNgsReportResults) HasBrcaness() bool {
	if o != nil && !IsNil(o.Brcaness) {
		return true
	}

	return false
}

// SetBrcaness gets a reference to the given Brcaness and assigns it to the Brcaness field.
func (o *SomaticNgsReportResults) SetBrcaness(v Brcaness) {
	o.Brcaness = &v
}

func (o SomaticNgsReportResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SomaticNgsReportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RnaSeqs) {
		toSerialize["rnaSeqs"] = o.RnaSeqs
	}
	if !IsNil(o.Tmb) {
		toSerialize["tmb"] = o.Tmb
	}
	if !IsNil(o.SimpleVariants) {
		toSerialize["simpleVariants"] = o.SimpleVariants
	}
	if !IsNil(o.HrdScore) {
		toSerialize["hrdScore"] = o.HrdScore
	}
	if !IsNil(o.CopyNumberVariants) {
		toSerialize["copyNumberVariants"] = o.CopyNumberVariants
	}
	if !IsNil(o.TumorCellContent) {
		toSerialize["tumorCellContent"] = o.TumorCellContent
	}
	if !IsNil(o.DnaFusions) {
		toSerialize["dnaFusions"] = o.DnaFusions
	}
	if !IsNil(o.RnaFusions) {
		toSerialize["rnaFusions"] = o.RnaFusions
	}
	if !IsNil(o.Brcaness) {
		toSerialize["brcaness"] = o.Brcaness
	}
	return toSerialize, nil
}

type NullableSomaticNgsReportResults struct {
	value *SomaticNgsReportResults
	isSet bool
}

func (v NullableSomaticNgsReportResults) Get() *SomaticNgsReportResults {
	return v.value
}

func (v *NullableSomaticNgsReportResults) Set(val *SomaticNgsReportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableSomaticNgsReportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableSomaticNgsReportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSomaticNgsReportResults(val *SomaticNgsReportResults) *NullableSomaticNgsReportResults {
	return &NullableSomaticNgsReportResults{value: val, isSet: true}
}

func (v NullableSomaticNgsReportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSomaticNgsReportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
