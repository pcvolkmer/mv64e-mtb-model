package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the LevelOfEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LevelOfEvidence{}

// LevelOfEvidence struct for LevelOfEvidence
type LevelOfEvidence struct {
	Grading              LevelOfEvidenceGradingCoding    `json:"grading"`
	Addendums            []LevelOfEvidenceAddendumCoding `json:"addendums,omitempty"`
	Publications         []PublicationReference          `json:"publications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LevelOfEvidence LevelOfEvidence

// NewLevelOfEvidence instantiates a new LevelOfEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevelOfEvidence(grading LevelOfEvidenceGradingCoding) *LevelOfEvidence {
	this := LevelOfEvidence{}
	this.Grading = grading
	return &this
}

// NewLevelOfEvidenceWithDefaults instantiates a new LevelOfEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelOfEvidenceWithDefaults() *LevelOfEvidence {
	this := LevelOfEvidence{}
	return &this
}

// GetGrading returns the Grading field value
func (o *LevelOfEvidence) GetGrading() LevelOfEvidenceGradingCoding {
	if o == nil {
		var ret LevelOfEvidenceGradingCoding
		return ret
	}

	return o.Grading
}

// GetGradingOk returns a tuple with the Grading field value
// and a boolean to check if the value has been set.
func (o *LevelOfEvidence) GetGradingOk() (*LevelOfEvidenceGradingCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grading, true
}

// SetGrading sets field value
func (o *LevelOfEvidence) SetGrading(v LevelOfEvidenceGradingCoding) {
	o.Grading = v
}

// GetAddendums returns the Addendums field value if set, zero value otherwise.
func (o *LevelOfEvidence) GetAddendums() []LevelOfEvidenceAddendumCoding {
	if o == nil || IsNil(o.Addendums) {
		var ret []LevelOfEvidenceAddendumCoding
		return ret
	}
	return o.Addendums
}

// GetAddendumsOk returns a tuple with the Addendums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LevelOfEvidence) GetAddendumsOk() ([]LevelOfEvidenceAddendumCoding, bool) {
	if o == nil || IsNil(o.Addendums) {
		return nil, false
	}
	return o.Addendums, true
}

// HasAddendums returns a boolean if a field has been set.
func (o *LevelOfEvidence) HasAddendums() bool {
	if o != nil && !IsNil(o.Addendums) {
		return true
	}

	return false
}

// SetAddendums gets a reference to the given []LevelOfEvidenceAddendumCoding and assigns it to the Addendums field.
func (o *LevelOfEvidence) SetAddendums(v []LevelOfEvidenceAddendumCoding) {
	o.Addendums = v
}

// GetPublications returns the Publications field value if set, zero value otherwise.
func (o *LevelOfEvidence) GetPublications() []PublicationReference {
	if o == nil || IsNil(o.Publications) {
		var ret []PublicationReference
		return ret
	}
	return o.Publications
}

// GetPublicationsOk returns a tuple with the Publications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LevelOfEvidence) GetPublicationsOk() ([]PublicationReference, bool) {
	if o == nil || IsNil(o.Publications) {
		return nil, false
	}
	return o.Publications, true
}

// HasPublications returns a boolean if a field has been set.
func (o *LevelOfEvidence) HasPublications() bool {
	if o != nil && !IsNil(o.Publications) {
		return true
	}

	return false
}

// SetPublications gets a reference to the given []PublicationReference and assigns it to the Publications field.
func (o *LevelOfEvidence) SetPublications(v []PublicationReference) {
	o.Publications = v
}

func (o LevelOfEvidence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LevelOfEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grading"] = o.Grading
	if !IsNil(o.Addendums) {
		toSerialize["addendums"] = o.Addendums
	}
	if !IsNil(o.Publications) {
		toSerialize["publications"] = o.Publications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LevelOfEvidence) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grading",
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

	varLevelOfEvidence := _LevelOfEvidence{}

	err = json.Unmarshal(data, &varLevelOfEvidence)

	if err != nil {
		return err
	}

	*o = LevelOfEvidence(varLevelOfEvidence)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grading")
		delete(additionalProperties, "addendums")
		delete(additionalProperties, "publications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLevelOfEvidence struct {
	value *LevelOfEvidence
	isSet bool
}

func (v NullableLevelOfEvidence) Get() *LevelOfEvidence {
	return v.value
}

func (v *NullableLevelOfEvidence) Set(val *LevelOfEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableLevelOfEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableLevelOfEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevelOfEvidence(val *LevelOfEvidence) *NullableLevelOfEvidence {
	return &NullableLevelOfEvidence{value: val, isSet: true}
}

func (v NullableLevelOfEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevelOfEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
