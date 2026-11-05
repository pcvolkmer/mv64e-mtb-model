package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DnaFusionFusionPartner5prime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnaFusionFusionPartner5prime{}

// DnaFusionFusionPartner5prime struct for DnaFusionFusionPartner5prime
type DnaFusionFusionPartner5prime struct {
	Chromosome Chromosome `json:"chromosome"`
	Gene       Coding     `json:"gene"`
	Position   float32    `json:"position"`
}

type _DnaFusionFusionPartner5prime DnaFusionFusionPartner5prime

// NewDnaFusionFusionPartner5prime instantiates a new DnaFusionFusionPartner5prime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaFusionFusionPartner5prime(chromosome Chromosome, gene Coding, position float32) *DnaFusionFusionPartner5prime {
	this := DnaFusionFusionPartner5prime{}
	this.Chromosome = chromosome
	this.Gene = gene
	this.Position = position
	return &this
}

// NewDnaFusionFusionPartner5primeWithDefaults instantiates a new DnaFusionFusionPartner5prime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaFusionFusionPartner5primeWithDefaults() *DnaFusionFusionPartner5prime {
	this := DnaFusionFusionPartner5prime{}
	return &this
}

// GetChromosome returns the Chromosome field value
func (o *DnaFusionFusionPartner5prime) GetChromosome() Chromosome {
	if o == nil {
		var ret Chromosome
		return ret
	}

	return o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner5prime) GetChromosomeOk() (*Chromosome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chromosome, true
}

// SetChromosome sets field value
func (o *DnaFusionFusionPartner5prime) SetChromosome(v Chromosome) {
	o.Chromosome = v
}

// GetGene returns the Gene field value
func (o *DnaFusionFusionPartner5prime) GetGene() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Gene
}

// GetGeneOk returns a tuple with the Gene field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner5prime) GetGeneOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gene, true
}

// SetGene sets field value
func (o *DnaFusionFusionPartner5prime) SetGene(v Coding) {
	o.Gene = v
}

// GetPosition returns the Position field value
func (o *DnaFusionFusionPartner5prime) GetPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner5prime) GetPositionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *DnaFusionFusionPartner5prime) SetPosition(v float32) {
	o.Position = v
}

func (o DnaFusionFusionPartner5prime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnaFusionFusionPartner5prime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chromosome"] = o.Chromosome
	toSerialize["gene"] = o.Gene
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

func (o *DnaFusionFusionPartner5prime) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chromosome",
		"gene",
		"position",
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

	varDnaFusionFusionPartner5prime := _DnaFusionFusionPartner5prime{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDnaFusionFusionPartner5prime)

	if err != nil {
		return err
	}

	*o = DnaFusionFusionPartner5prime(varDnaFusionFusionPartner5prime)

	return err
}

type NullableDnaFusionFusionPartner5prime struct {
	value *DnaFusionFusionPartner5prime
	isSet bool
}

func (v NullableDnaFusionFusionPartner5prime) Get() *DnaFusionFusionPartner5prime {
	return v.value
}

func (v *NullableDnaFusionFusionPartner5prime) Set(val *DnaFusionFusionPartner5prime) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaFusionFusionPartner5prime) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaFusionFusionPartner5prime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaFusionFusionPartner5prime(val *DnaFusionFusionPartner5prime) *NullableDnaFusionFusionPartner5prime {
	return &NullableDnaFusionFusionPartner5prime{value: val, isSet: true}
}

func (v NullableDnaFusionFusionPartner5prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaFusionFusionPartner5prime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
