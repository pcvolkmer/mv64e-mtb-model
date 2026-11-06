package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the GeneticCounselingRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneticCounselingRecommendation{}

// GeneticCounselingRecommendation struct for GeneticCounselingRecommendation
type GeneticCounselingRecommendation struct {
	Id                   string                                      `json:"id"`
	Patient              Reference                                   `json:"patient"`
	IssuedOn             string                                      `json:"issuedOn"`
	Reason               GeneticCounselingRecommendationReasonCoding `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _GeneticCounselingRecommendation GeneticCounselingRecommendation

// NewGeneticCounselingRecommendation instantiates a new GeneticCounselingRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneticCounselingRecommendation(id string, patient Reference, issuedOn string, reason GeneticCounselingRecommendationReasonCoding) *GeneticCounselingRecommendation {
	this := GeneticCounselingRecommendation{}
	this.Id = id
	this.Patient = patient
	this.IssuedOn = issuedOn
	this.Reason = reason
	return &this
}

// NewGeneticCounselingRecommendationWithDefaults instantiates a new GeneticCounselingRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneticCounselingRecommendationWithDefaults() *GeneticCounselingRecommendation {
	this := GeneticCounselingRecommendation{}
	return &this
}

// GetId returns the Id field value
func (o *GeneticCounselingRecommendation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GeneticCounselingRecommendation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GeneticCounselingRecommendation) SetId(v string) {
	o.Id = v
}

// GetPatient returns the Patient field value
func (o *GeneticCounselingRecommendation) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *GeneticCounselingRecommendation) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *GeneticCounselingRecommendation) SetPatient(v Reference) {
	o.Patient = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *GeneticCounselingRecommendation) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *GeneticCounselingRecommendation) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *GeneticCounselingRecommendation) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetReason returns the Reason field value
func (o *GeneticCounselingRecommendation) GetReason() GeneticCounselingRecommendationReasonCoding {
	if o == nil {
		var ret GeneticCounselingRecommendationReasonCoding
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *GeneticCounselingRecommendation) GetReasonOk() (*GeneticCounselingRecommendationReasonCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *GeneticCounselingRecommendation) SetReason(v GeneticCounselingRecommendationReasonCoding) {
	o.Reason = v
}

func (o GeneticCounselingRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneticCounselingRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["patient"] = o.Patient
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GeneticCounselingRecommendation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"patient",
		"issuedOn",
		"reason",
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

	varGeneticCounselingRecommendation := _GeneticCounselingRecommendation{}

	err = json.Unmarshal(data, &varGeneticCounselingRecommendation)

	if err != nil {
		return err
	}

	*o = GeneticCounselingRecommendation(varGeneticCounselingRecommendation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeneticCounselingRecommendation struct {
	value *GeneticCounselingRecommendation
	isSet bool
}

func (v NullableGeneticCounselingRecommendation) Get() *GeneticCounselingRecommendation {
	return v.value
}

func (v *NullableGeneticCounselingRecommendation) Set(val *GeneticCounselingRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneticCounselingRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneticCounselingRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneticCounselingRecommendation(val *GeneticCounselingRecommendation) *NullableGeneticCounselingRecommendation {
	return &NullableGeneticCounselingRecommendation{value: val, isSet: true}
}

func (v NullableGeneticCounselingRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneticCounselingRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
