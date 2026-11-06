package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the RnaFusionFusionPartner3prime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RnaFusionFusionPartner3prime{}

// RnaFusionFusionPartner3prime struct for RnaFusionFusionPartner3prime
type RnaFusionFusionPartner3prime struct {
	TranscriptId         TranscriptId    `json:"transcriptId"`
	Position             float32         `json:"position"`
	ExonId               string          `json:"exonId"`
	Strand               RnaFusionStrand `json:"strand"`
	Gene                 Coding          `json:"gene"`
	AdditionalProperties map[string]interface{}
}

type _RnaFusionFusionPartner3prime RnaFusionFusionPartner3prime

// NewRnaFusionFusionPartner3prime instantiates a new RnaFusionFusionPartner3prime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaFusionFusionPartner3prime(transcriptId TranscriptId, position float32, exonId string, strand RnaFusionStrand, gene Coding) *RnaFusionFusionPartner3prime {
	this := RnaFusionFusionPartner3prime{}
	this.TranscriptId = transcriptId
	this.Position = position
	this.ExonId = exonId
	this.Strand = strand
	this.Gene = gene
	return &this
}

// NewRnaFusionFusionPartner3primeWithDefaults instantiates a new RnaFusionFusionPartner3prime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaFusionFusionPartner3primeWithDefaults() *RnaFusionFusionPartner3prime {
	this := RnaFusionFusionPartner3prime{}
	return &this
}

// GetTranscriptId returns the TranscriptId field value
func (o *RnaFusionFusionPartner3prime) GetTranscriptId() TranscriptId {
	if o == nil {
		var ret TranscriptId
		return ret
	}

	return o.TranscriptId
}

// GetTranscriptIdOk returns a tuple with the TranscriptId field value
// and a boolean to check if the value has been set.
func (o *RnaFusionFusionPartner3prime) GetTranscriptIdOk() (*TranscriptId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranscriptId, true
}

// SetTranscriptId sets field value
func (o *RnaFusionFusionPartner3prime) SetTranscriptId(v TranscriptId) {
	o.TranscriptId = v
}

// GetPosition returns the Position field value
func (o *RnaFusionFusionPartner3prime) GetPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *RnaFusionFusionPartner3prime) GetPositionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *RnaFusionFusionPartner3prime) SetPosition(v float32) {
	o.Position = v
}

// GetExonId returns the ExonId field value
func (o *RnaFusionFusionPartner3prime) GetExonId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExonId
}

// GetExonIdOk returns a tuple with the ExonId field value
// and a boolean to check if the value has been set.
func (o *RnaFusionFusionPartner3prime) GetExonIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExonId, true
}

// SetExonId sets field value
func (o *RnaFusionFusionPartner3prime) SetExonId(v string) {
	o.ExonId = v
}

// GetStrand returns the Strand field value
func (o *RnaFusionFusionPartner3prime) GetStrand() RnaFusionStrand {
	if o == nil {
		var ret RnaFusionStrand
		return ret
	}

	return o.Strand
}

// GetStrandOk returns a tuple with the Strand field value
// and a boolean to check if the value has been set.
func (o *RnaFusionFusionPartner3prime) GetStrandOk() (*RnaFusionStrand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strand, true
}

// SetStrand sets field value
func (o *RnaFusionFusionPartner3prime) SetStrand(v RnaFusionStrand) {
	o.Strand = v
}

// GetGene returns the Gene field value
func (o *RnaFusionFusionPartner3prime) GetGene() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Gene
}

// GetGeneOk returns a tuple with the Gene field value
// and a boolean to check if the value has been set.
func (o *RnaFusionFusionPartner3prime) GetGeneOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gene, true
}

// SetGene sets field value
func (o *RnaFusionFusionPartner3prime) SetGene(v Coding) {
	o.Gene = v
}

func (o RnaFusionFusionPartner3prime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RnaFusionFusionPartner3prime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transcriptId"] = o.TranscriptId
	toSerialize["position"] = o.Position
	toSerialize["exonId"] = o.ExonId
	toSerialize["strand"] = o.Strand
	toSerialize["gene"] = o.Gene

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RnaFusionFusionPartner3prime) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transcriptId",
		"position",
		"exonId",
		"strand",
		"gene",
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

	varRnaFusionFusionPartner3prime := _RnaFusionFusionPartner3prime{}

	err = json.Unmarshal(data, &varRnaFusionFusionPartner3prime)

	if err != nil {
		return err
	}

	*o = RnaFusionFusionPartner3prime(varRnaFusionFusionPartner3prime)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transcriptId")
		delete(additionalProperties, "position")
		delete(additionalProperties, "exonId")
		delete(additionalProperties, "strand")
		delete(additionalProperties, "gene")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRnaFusionFusionPartner3prime struct {
	value *RnaFusionFusionPartner3prime
	isSet bool
}

func (v NullableRnaFusionFusionPartner3prime) Get() *RnaFusionFusionPartner3prime {
	return v.value
}

func (v *NullableRnaFusionFusionPartner3prime) Set(val *RnaFusionFusionPartner3prime) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaFusionFusionPartner3prime) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaFusionFusionPartner3prime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaFusionFusionPartner3prime(val *RnaFusionFusionPartner3prime) *NullableRnaFusionFusionPartner3prime {
	return &NullableRnaFusionFusionPartner3prime{value: val, isSet: true}
}

func (v NullableRnaFusionFusionPartner3prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaFusionFusionPartner3prime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
