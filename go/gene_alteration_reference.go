package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the GeneAlterationReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneAlterationReference{}

// GeneAlterationReference struct for GeneAlterationReference
type GeneAlterationReference struct {
	Variant              Reference `json:"variant"`
	Display              *string   `json:"display,omitempty"`
	Gene                 *Coding   `json:"gene,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GeneAlterationReference GeneAlterationReference

// NewGeneAlterationReference instantiates a new GeneAlterationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneAlterationReference(variant Reference) *GeneAlterationReference {
	this := GeneAlterationReference{}
	this.Variant = variant
	return &this
}

// NewGeneAlterationReferenceWithDefaults instantiates a new GeneAlterationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneAlterationReferenceWithDefaults() *GeneAlterationReference {
	this := GeneAlterationReference{}
	return &this
}

// GetVariant returns the Variant field value
func (o *GeneAlterationReference) GetVariant() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *GeneAlterationReference) GetVariantOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *GeneAlterationReference) SetVariant(v Reference) {
	o.Variant = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *GeneAlterationReference) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneAlterationReference) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *GeneAlterationReference) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *GeneAlterationReference) SetDisplay(v string) {
	o.Display = &v
}

// GetGene returns the Gene field value if set, zero value otherwise.
func (o *GeneAlterationReference) GetGene() Coding {
	if o == nil || IsNil(o.Gene) {
		var ret Coding
		return ret
	}
	return *o.Gene
}

// GetGeneOk returns a tuple with the Gene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneAlterationReference) GetGeneOk() (*Coding, bool) {
	if o == nil || IsNil(o.Gene) {
		return nil, false
	}
	return o.Gene, true
}

// HasGene returns a boolean if a field has been set.
func (o *GeneAlterationReference) HasGene() bool {
	if o != nil && !IsNil(o.Gene) {
		return true
	}

	return false
}

// SetGene gets a reference to the given Coding and assigns it to the Gene field.
func (o *GeneAlterationReference) SetGene(v Coding) {
	o.Gene = &v
}

func (o GeneAlterationReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneAlterationReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant"] = o.Variant
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Gene) {
		toSerialize["gene"] = o.Gene
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GeneAlterationReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"variant",
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

	varGeneAlterationReference := _GeneAlterationReference{}

	err = json.Unmarshal(data, &varGeneAlterationReference)

	if err != nil {
		return err
	}

	*o = GeneAlterationReference(varGeneAlterationReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "variant")
		delete(additionalProperties, "display")
		delete(additionalProperties, "gene")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeneAlterationReference struct {
	value *GeneAlterationReference
	isSet bool
}

func (v NullableGeneAlterationReference) Get() *GeneAlterationReference {
	return v.value
}

func (v *NullableGeneAlterationReference) Set(val *GeneAlterationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneAlterationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneAlterationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneAlterationReference(val *GeneAlterationReference) *NullableGeneAlterationReference {
	return &NullableGeneAlterationReference{value: val, isSet: true}
}

func (v NullableGeneAlterationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneAlterationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
