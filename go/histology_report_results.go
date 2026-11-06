package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HistologyReportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistologyReportResults{}

// HistologyReportResults struct for HistologyReportResults
type HistologyReportResults struct {
	TumorMorphology  HistologyReportResultsTumorMorphology `json:"tumorMorphology"`
	TumorCellContent *TumorCellContent                     `json:"tumorCellContent,omitempty"`
}

type _HistologyReportResults HistologyReportResults

// NewHistologyReportResults instantiates a new HistologyReportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistologyReportResults(tumorMorphology HistologyReportResultsTumorMorphology) *HistologyReportResults {
	this := HistologyReportResults{}
	this.TumorMorphology = tumorMorphology
	return &this
}

// NewHistologyReportResultsWithDefaults instantiates a new HistologyReportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistologyReportResultsWithDefaults() *HistologyReportResults {
	this := HistologyReportResults{}
	return &this
}

// GetTumorMorphology returns the TumorMorphology field value
func (o *HistologyReportResults) GetTumorMorphology() HistologyReportResultsTumorMorphology {
	if o == nil {
		var ret HistologyReportResultsTumorMorphology
		return ret
	}

	return o.TumorMorphology
}

// GetTumorMorphologyOk returns a tuple with the TumorMorphology field value
// and a boolean to check if the value has been set.
func (o *HistologyReportResults) GetTumorMorphologyOk() (*HistologyReportResultsTumorMorphology, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TumorMorphology, true
}

// SetTumorMorphology sets field value
func (o *HistologyReportResults) SetTumorMorphology(v HistologyReportResultsTumorMorphology) {
	o.TumorMorphology = v
}

// GetTumorCellContent returns the TumorCellContent field value if set, zero value otherwise.
func (o *HistologyReportResults) GetTumorCellContent() TumorCellContent {
	if o == nil || IsNil(o.TumorCellContent) {
		var ret TumorCellContent
		return ret
	}
	return *o.TumorCellContent
}

// GetTumorCellContentOk returns a tuple with the TumorCellContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistologyReportResults) GetTumorCellContentOk() (*TumorCellContent, bool) {
	if o == nil || IsNil(o.TumorCellContent) {
		return nil, false
	}
	return o.TumorCellContent, true
}

// HasTumorCellContent returns a boolean if a field has been set.
func (o *HistologyReportResults) HasTumorCellContent() bool {
	if o != nil && !IsNil(o.TumorCellContent) {
		return true
	}

	return false
}

// SetTumorCellContent gets a reference to the given TumorCellContent and assigns it to the TumorCellContent field.
func (o *HistologyReportResults) SetTumorCellContent(v TumorCellContent) {
	o.TumorCellContent = &v
}

func (o HistologyReportResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistologyReportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tumorMorphology"] = o.TumorMorphology
	if !IsNil(o.TumorCellContent) {
		toSerialize["tumorCellContent"] = o.TumorCellContent
	}
	return toSerialize, nil
}

func (o *HistologyReportResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tumorMorphology",
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

	varHistologyReportResults := _HistologyReportResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHistologyReportResults)

	if err != nil {
		return err
	}

	*o = HistologyReportResults(varHistologyReportResults)

	return err
}

type NullableHistologyReportResults struct {
	value *HistologyReportResults
	isSet bool
}

func (v NullableHistologyReportResults) Get() *HistologyReportResults {
	return v.value
}

func (v *NullableHistologyReportResults) Set(val *HistologyReportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableHistologyReportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableHistologyReportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistologyReportResults(val *HistologyReportResults) *NullableHistologyReportResults {
	return &NullableHistologyReportResults{value: val, isSet: true}
}

func (v NullableHistologyReportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistologyReportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
