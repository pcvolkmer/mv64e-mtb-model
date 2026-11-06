package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the DnaFusionFusionPartner3prime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnaFusionFusionPartner3prime{}

// DnaFusionFusionPartner3prime struct for DnaFusionFusionPartner3prime
type DnaFusionFusionPartner3prime struct {
	Chromosome           Chromosome `json:"chromosome"`
	Gene                 Coding     `json:"gene"`
	Position             float32    `json:"position"`
	AdditionalProperties map[string]interface{}
}

type _DnaFusionFusionPartner3prime DnaFusionFusionPartner3prime

// NewDnaFusionFusionPartner3prime instantiates a new DnaFusionFusionPartner3prime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaFusionFusionPartner3prime(chromosome Chromosome, gene Coding, position float32) *DnaFusionFusionPartner3prime {
	this := DnaFusionFusionPartner3prime{}
	this.Chromosome = chromosome
	this.Gene = gene
	this.Position = position
	return &this
}

// NewDnaFusionFusionPartner3primeWithDefaults instantiates a new DnaFusionFusionPartner3prime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaFusionFusionPartner3primeWithDefaults() *DnaFusionFusionPartner3prime {
	this := DnaFusionFusionPartner3prime{}
	return &this
}

// GetChromosome returns the Chromosome field value
func (o *DnaFusionFusionPartner3prime) GetChromosome() Chromosome {
	if o == nil {
		var ret Chromosome
		return ret
	}

	return o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner3prime) GetChromosomeOk() (*Chromosome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chromosome, true
}

// SetChromosome sets field value
func (o *DnaFusionFusionPartner3prime) SetChromosome(v Chromosome) {
	o.Chromosome = v
}

// GetGene returns the Gene field value
func (o *DnaFusionFusionPartner3prime) GetGene() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Gene
}

// GetGeneOk returns a tuple with the Gene field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner3prime) GetGeneOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gene, true
}

// SetGene sets field value
func (o *DnaFusionFusionPartner3prime) SetGene(v Coding) {
	o.Gene = v
}

// GetPosition returns the Position field value
func (o *DnaFusionFusionPartner3prime) GetPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *DnaFusionFusionPartner3prime) GetPositionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *DnaFusionFusionPartner3prime) SetPosition(v float32) {
	o.Position = v
}

func (o DnaFusionFusionPartner3prime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnaFusionFusionPartner3prime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chromosome"] = o.Chromosome
	toSerialize["gene"] = o.Gene
	toSerialize["position"] = o.Position

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnaFusionFusionPartner3prime) UnmarshalJSON(data []byte) (err error) {
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

	varDnaFusionFusionPartner3prime := _DnaFusionFusionPartner3prime{}

	err = json.Unmarshal(data, &varDnaFusionFusionPartner3prime)

	if err != nil {
		return err
	}

	*o = DnaFusionFusionPartner3prime(varDnaFusionFusionPartner3prime)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chromosome")
		delete(additionalProperties, "gene")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDnaFusionFusionPartner3prime struct {
	value *DnaFusionFusionPartner3prime
	isSet bool
}

func (v NullableDnaFusionFusionPartner3prime) Get() *DnaFusionFusionPartner3prime {
	return v.value
}

func (v *NullableDnaFusionFusionPartner3prime) Set(val *DnaFusionFusionPartner3prime) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaFusionFusionPartner3prime) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaFusionFusionPartner3prime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaFusionFusionPartner3prime(val *DnaFusionFusionPartner3prime) *NullableDnaFusionFusionPartner3prime {
	return &NullableDnaFusionFusionPartner3prime{value: val, isSet: true}
}

func (v NullableDnaFusionFusionPartner3prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaFusionFusionPartner3prime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
