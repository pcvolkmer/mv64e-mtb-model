package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// Chromosome the model 'Chromosome'
type Chromosome string

// List of Chromosome
const (
	CHR5   Chromosome = "chr5"
	CHR6   Chromosome = "chr6"
	CHR8   Chromosome = "chr8"
	CHR1   Chromosome = "chr1"
	CHR_X  Chromosome = "chrX"
	CHR12  Chromosome = "chr12"
	CHR11  Chromosome = "chr11"
	CHR17  Chromosome = "chr17"
	CHR21  Chromosome = "chr21"
	CHR19  Chromosome = "chr19"
	CHR14  Chromosome = "chr14"
	CHR9   Chromosome = "chr9"
	CHR4   Chromosome = "chr4"
	CHR_MT Chromosome = "chrMT"
	CHR2   Chromosome = "chr2"
	CHR_Y  Chromosome = "chrY"
	CHR3   Chromosome = "chr3"
	CHR7   Chromosome = "chr7"
	CHR13  Chromosome = "chr13"
	CHR10  Chromosome = "chr10"
	CHR15  Chromosome = "chr15"
	CHR22  Chromosome = "chr22"
	CHR18  Chromosome = "chr18"
	CHR16  Chromosome = "chr16"
	CHR20  Chromosome = "chr20"
)

// All allowed values of Chromosome enum
var AllowedChromosomeEnumValues = []Chromosome{
	"chr5",
	"chr6",
	"chr8",
	"chr1",
	"chrX",
	"chr12",
	"chr11",
	"chr17",
	"chr21",
	"chr19",
	"chr14",
	"chr9",
	"chr4",
	"chrMT",
	"chr2",
	"chrY",
	"chr3",
	"chr7",
	"chr13",
	"chr10",
	"chr15",
	"chr22",
	"chr18",
	"chr16",
	"chr20",
}

func (v *Chromosome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Chromosome(value)
	for _, existing := range AllowedChromosomeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Chromosome", value)
}

// NewChromosomeFromValue returns a pointer to a valid Chromosome
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChromosomeFromValue(v string) (*Chromosome, error) {
	ev := Chromosome(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Chromosome: valid values are %v", v, AllowedChromosomeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Chromosome) IsValid() bool {
	for _, existing := range AllowedChromosomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Chromosome value
func (v Chromosome) Ptr() *Chromosome {
	return &v
}

type NullableChromosome struct {
	value *Chromosome
	isSet bool
}

func (v NullableChromosome) Get() *Chromosome {
	return v.value
}

func (v *NullableChromosome) Set(val *Chromosome) {
	v.value = val
	v.isSet = true
}

func (v NullableChromosome) IsSet() bool {
	return v.isSet
}

func (v *NullableChromosome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChromosome(val *Chromosome) *NullableChromosome {
	return &NullableChromosome{value: val, isSet: true}
}

func (v NullableChromosome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChromosome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
