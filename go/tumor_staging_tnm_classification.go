package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorStagingTnmClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorStagingTnmClassification{}

// TumorStagingTnmClassification struct for TumorStagingTnmClassification
type TumorStagingTnmClassification struct {
	Tumor                Coding `json:"tumor"`
	Nodes                Coding `json:"nodes"`
	Metastasis           Coding `json:"metastasis"`
	AdditionalProperties map[string]interface{}
}

type _TumorStagingTnmClassification TumorStagingTnmClassification

// NewTumorStagingTnmClassification instantiates a new TumorStagingTnmClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorStagingTnmClassification(tumor Coding, nodes Coding, metastasis Coding) *TumorStagingTnmClassification {
	this := TumorStagingTnmClassification{}
	this.Tumor = tumor
	this.Nodes = nodes
	this.Metastasis = metastasis
	return &this
}

// NewTumorStagingTnmClassificationWithDefaults instantiates a new TumorStagingTnmClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorStagingTnmClassificationWithDefaults() *TumorStagingTnmClassification {
	this := TumorStagingTnmClassification{}
	return &this
}

// GetTumor returns the Tumor field value
func (o *TumorStagingTnmClassification) GetTumor() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Tumor
}

// GetTumorOk returns a tuple with the Tumor field value
// and a boolean to check if the value has been set.
func (o *TumorStagingTnmClassification) GetTumorOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tumor, true
}

// SetTumor sets field value
func (o *TumorStagingTnmClassification) SetTumor(v Coding) {
	o.Tumor = v
}

// GetNodes returns the Nodes field value
func (o *TumorStagingTnmClassification) GetNodes() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *TumorStagingTnmClassification) GetNodesOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *TumorStagingTnmClassification) SetNodes(v Coding) {
	o.Nodes = v
}

// GetMetastasis returns the Metastasis field value
func (o *TumorStagingTnmClassification) GetMetastasis() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Metastasis
}

// GetMetastasisOk returns a tuple with the Metastasis field value
// and a boolean to check if the value has been set.
func (o *TumorStagingTnmClassification) GetMetastasisOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metastasis, true
}

// SetMetastasis sets field value
func (o *TumorStagingTnmClassification) SetMetastasis(v Coding) {
	o.Metastasis = v
}

func (o TumorStagingTnmClassification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorStagingTnmClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tumor"] = o.Tumor
	toSerialize["nodes"] = o.Nodes
	toSerialize["metastasis"] = o.Metastasis

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TumorStagingTnmClassification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tumor",
		"nodes",
		"metastasis",
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

	varTumorStagingTnmClassification := _TumorStagingTnmClassification{}

	err = json.Unmarshal(data, &varTumorStagingTnmClassification)

	if err != nil {
		return err
	}

	*o = TumorStagingTnmClassification(varTumorStagingTnmClassification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tumor")
		delete(additionalProperties, "nodes")
		delete(additionalProperties, "metastasis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTumorStagingTnmClassification struct {
	value *TumorStagingTnmClassification
	isSet bool
}

func (v NullableTumorStagingTnmClassification) Get() *TumorStagingTnmClassification {
	return v.value
}

func (v *NullableTumorStagingTnmClassification) Set(val *TumorStagingTnmClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorStagingTnmClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorStagingTnmClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorStagingTnmClassification(val *TumorStagingTnmClassification) *NullableTumorStagingTnmClassification {
	return &NullableTumorStagingTnmClassification{value: val, isSet: true}
}

func (v NullableTumorStagingTnmClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorStagingTnmClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
