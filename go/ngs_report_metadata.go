package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the NgsReportMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgsReportMetadata{}

// NgsReportMetadata struct for NgsReportMetadata
type NgsReportMetadata struct {
	KitManufacturer      string `json:"kitManufacturer"`
	Pipeline             string `json:"pipeline"`
	KitType              string `json:"kitType"`
	Sequencer            string `json:"sequencer"`
	ReferenceGenome      string `json:"referenceGenome"`
	AdditionalProperties map[string]interface{}
}

type _NgsReportMetadata NgsReportMetadata

// NewNgsReportMetadata instantiates a new NgsReportMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgsReportMetadata(kitManufacturer string, pipeline string, kitType string, sequencer string, referenceGenome string) *NgsReportMetadata {
	this := NgsReportMetadata{}
	this.KitManufacturer = kitManufacturer
	this.Pipeline = pipeline
	this.KitType = kitType
	this.Sequencer = sequencer
	this.ReferenceGenome = referenceGenome
	return &this
}

// NewNgsReportMetadataWithDefaults instantiates a new NgsReportMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgsReportMetadataWithDefaults() *NgsReportMetadata {
	this := NgsReportMetadata{}
	return &this
}

// GetKitManufacturer returns the KitManufacturer field value
func (o *NgsReportMetadata) GetKitManufacturer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KitManufacturer
}

// GetKitManufacturerOk returns a tuple with the KitManufacturer field value
// and a boolean to check if the value has been set.
func (o *NgsReportMetadata) GetKitManufacturerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KitManufacturer, true
}

// SetKitManufacturer sets field value
func (o *NgsReportMetadata) SetKitManufacturer(v string) {
	o.KitManufacturer = v
}

// GetPipeline returns the Pipeline field value
func (o *NgsReportMetadata) GetPipeline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *NgsReportMetadata) GetPipelineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *NgsReportMetadata) SetPipeline(v string) {
	o.Pipeline = v
}

// GetKitType returns the KitType field value
func (o *NgsReportMetadata) GetKitType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KitType
}

// GetKitTypeOk returns a tuple with the KitType field value
// and a boolean to check if the value has been set.
func (o *NgsReportMetadata) GetKitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KitType, true
}

// SetKitType sets field value
func (o *NgsReportMetadata) SetKitType(v string) {
	o.KitType = v
}

// GetSequencer returns the Sequencer field value
func (o *NgsReportMetadata) GetSequencer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequencer
}

// GetSequencerOk returns a tuple with the Sequencer field value
// and a boolean to check if the value has been set.
func (o *NgsReportMetadata) GetSequencerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequencer, true
}

// SetSequencer sets field value
func (o *NgsReportMetadata) SetSequencer(v string) {
	o.Sequencer = v
}

// GetReferenceGenome returns the ReferenceGenome field value
func (o *NgsReportMetadata) GetReferenceGenome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceGenome
}

// GetReferenceGenomeOk returns a tuple with the ReferenceGenome field value
// and a boolean to check if the value has been set.
func (o *NgsReportMetadata) GetReferenceGenomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceGenome, true
}

// SetReferenceGenome sets field value
func (o *NgsReportMetadata) SetReferenceGenome(v string) {
	o.ReferenceGenome = v
}

func (o NgsReportMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgsReportMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kitManufacturer"] = o.KitManufacturer
	toSerialize["pipeline"] = o.Pipeline
	toSerialize["kitType"] = o.KitType
	toSerialize["sequencer"] = o.Sequencer
	toSerialize["referenceGenome"] = o.ReferenceGenome

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NgsReportMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kitManufacturer",
		"pipeline",
		"kitType",
		"sequencer",
		"referenceGenome",
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

	varNgsReportMetadata := _NgsReportMetadata{}

	err = json.Unmarshal(data, &varNgsReportMetadata)

	if err != nil {
		return err
	}

	*o = NgsReportMetadata(varNgsReportMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kitManufacturer")
		delete(additionalProperties, "pipeline")
		delete(additionalProperties, "kitType")
		delete(additionalProperties, "sequencer")
		delete(additionalProperties, "referenceGenome")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNgsReportMetadata struct {
	value *NgsReportMetadata
	isSet bool
}

func (v NullableNgsReportMetadata) Get() *NgsReportMetadata {
	return v.value
}

func (v *NullableNgsReportMetadata) Set(val *NgsReportMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNgsReportMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNgsReportMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgsReportMetadata(val *NgsReportMetadata) *NullableNgsReportMetadata {
	return &NullableNgsReportMetadata{value: val, isSet: true}
}

func (v NullableNgsReportMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgsReportMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
