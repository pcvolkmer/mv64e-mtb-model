package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MtbCarePlanProcedureRecommendationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbCarePlanProcedureRecommendationsInner{}

// MtbCarePlanProcedureRecommendationsInner struct for MtbCarePlanProcedureRecommendationsInner
type MtbCarePlanProcedureRecommendationsInner struct {
	Priority           MtbRecommendationPriorityCoding          `json:"priority"`
	SupportingVariants []GeneAlterationReference                `json:"supportingVariants,omitempty"`
	Reason             *Reference                               `json:"reason,omitempty"`
	Code               MtbProcedureRecommendationCategoryCoding `json:"code"`
	IssuedOn           string                                   `json:"issuedOn"`
	Patient            Reference                                `json:"patient"`
	SupportingFindings []Reference                              `json:"supportingFindings,omitempty"`
	LevelOfEvidence    *LevelOfEvidence                         `json:"levelOfEvidence,omitempty"`
	Id                 string                                   `json:"id"`
}

type _MtbCarePlanProcedureRecommendationsInner MtbCarePlanProcedureRecommendationsInner

// NewMtbCarePlanProcedureRecommendationsInner instantiates a new MtbCarePlanProcedureRecommendationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbCarePlanProcedureRecommendationsInner(priority MtbRecommendationPriorityCoding, code MtbProcedureRecommendationCategoryCoding, issuedOn string, patient Reference, id string) *MtbCarePlanProcedureRecommendationsInner {
	this := MtbCarePlanProcedureRecommendationsInner{}
	this.Priority = priority
	this.Code = code
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	return &this
}

// NewMtbCarePlanProcedureRecommendationsInnerWithDefaults instantiates a new MtbCarePlanProcedureRecommendationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbCarePlanProcedureRecommendationsInnerWithDefaults() *MtbCarePlanProcedureRecommendationsInner {
	this := MtbCarePlanProcedureRecommendationsInner{}
	return &this
}

// GetPriority returns the Priority field value
func (o *MtbCarePlanProcedureRecommendationsInner) GetPriority() MtbRecommendationPriorityCoding {
	if o == nil {
		var ret MtbRecommendationPriorityCoding
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetPriorityOk() (*MtbRecommendationPriorityCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *MtbCarePlanProcedureRecommendationsInner) SetPriority(v MtbRecommendationPriorityCoding) {
	o.Priority = v
}

// GetSupportingVariants returns the SupportingVariants field value if set, zero value otherwise.
func (o *MtbCarePlanProcedureRecommendationsInner) GetSupportingVariants() []GeneAlterationReference {
	if o == nil || IsNil(o.SupportingVariants) {
		var ret []GeneAlterationReference
		return ret
	}
	return o.SupportingVariants
}

// GetSupportingVariantsOk returns a tuple with the SupportingVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetSupportingVariantsOk() ([]GeneAlterationReference, bool) {
	if o == nil || IsNil(o.SupportingVariants) {
		return nil, false
	}
	return o.SupportingVariants, true
}

// HasSupportingVariants returns a boolean if a field has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) HasSupportingVariants() bool {
	if o != nil && !IsNil(o.SupportingVariants) {
		return true
	}

	return false
}

// SetSupportingVariants gets a reference to the given []GeneAlterationReference and assigns it to the SupportingVariants field.
func (o *MtbCarePlanProcedureRecommendationsInner) SetSupportingVariants(v []GeneAlterationReference) {
	o.SupportingVariants = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MtbCarePlanProcedureRecommendationsInner) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *MtbCarePlanProcedureRecommendationsInner) SetReason(v Reference) {
	o.Reason = &v
}

// GetCode returns the Code field value
func (o *MtbCarePlanProcedureRecommendationsInner) GetCode() MtbProcedureRecommendationCategoryCoding {
	if o == nil {
		var ret MtbProcedureRecommendationCategoryCoding
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetCodeOk() (*MtbProcedureRecommendationCategoryCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *MtbCarePlanProcedureRecommendationsInner) SetCode(v MtbProcedureRecommendationCategoryCoding) {
	o.Code = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbCarePlanProcedureRecommendationsInner) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbCarePlanProcedureRecommendationsInner) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *MtbCarePlanProcedureRecommendationsInner) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbCarePlanProcedureRecommendationsInner) SetPatient(v Reference) {
	o.Patient = v
}

// GetSupportingFindings returns the SupportingFindings field value if set, zero value otherwise.
func (o *MtbCarePlanProcedureRecommendationsInner) GetSupportingFindings() []Reference {
	if o == nil || IsNil(o.SupportingFindings) {
		var ret []Reference
		return ret
	}
	return o.SupportingFindings
}

// GetSupportingFindingsOk returns a tuple with the SupportingFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetSupportingFindingsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.SupportingFindings) {
		return nil, false
	}
	return o.SupportingFindings, true
}

// HasSupportingFindings returns a boolean if a field has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) HasSupportingFindings() bool {
	if o != nil && !IsNil(o.SupportingFindings) {
		return true
	}

	return false
}

// SetSupportingFindings gets a reference to the given []Reference and assigns it to the SupportingFindings field.
func (o *MtbCarePlanProcedureRecommendationsInner) SetSupportingFindings(v []Reference) {
	o.SupportingFindings = v
}

// GetLevelOfEvidence returns the LevelOfEvidence field value if set, zero value otherwise.
func (o *MtbCarePlanProcedureRecommendationsInner) GetLevelOfEvidence() LevelOfEvidence {
	if o == nil || IsNil(o.LevelOfEvidence) {
		var ret LevelOfEvidence
		return ret
	}
	return *o.LevelOfEvidence
}

// GetLevelOfEvidenceOk returns a tuple with the LevelOfEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetLevelOfEvidenceOk() (*LevelOfEvidence, bool) {
	if o == nil || IsNil(o.LevelOfEvidence) {
		return nil, false
	}
	return o.LevelOfEvidence, true
}

// HasLevelOfEvidence returns a boolean if a field has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) HasLevelOfEvidence() bool {
	if o != nil && !IsNil(o.LevelOfEvidence) {
		return true
	}

	return false
}

// SetLevelOfEvidence gets a reference to the given LevelOfEvidence and assigns it to the LevelOfEvidence field.
func (o *MtbCarePlanProcedureRecommendationsInner) SetLevelOfEvidence(v LevelOfEvidence) {
	o.LevelOfEvidence = &v
}

// GetId returns the Id field value
func (o *MtbCarePlanProcedureRecommendationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbCarePlanProcedureRecommendationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbCarePlanProcedureRecommendationsInner) SetId(v string) {
	o.Id = v
}

func (o MtbCarePlanProcedureRecommendationsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbCarePlanProcedureRecommendationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.SupportingVariants) {
		toSerialize["supportingVariants"] = o.SupportingVariants
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["code"] = o.Code
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["patient"] = o.Patient
	if !IsNil(o.SupportingFindings) {
		toSerialize["supportingFindings"] = o.SupportingFindings
	}
	if !IsNil(o.LevelOfEvidence) {
		toSerialize["levelOfEvidence"] = o.LevelOfEvidence
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *MtbCarePlanProcedureRecommendationsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"priority",
		"code",
		"issuedOn",
		"patient",
		"id",
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

	varMtbCarePlanProcedureRecommendationsInner := _MtbCarePlanProcedureRecommendationsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMtbCarePlanProcedureRecommendationsInner)

	if err != nil {
		return err
	}

	*o = MtbCarePlanProcedureRecommendationsInner(varMtbCarePlanProcedureRecommendationsInner)

	return err
}

type NullableMtbCarePlanProcedureRecommendationsInner struct {
	value *MtbCarePlanProcedureRecommendationsInner
	isSet bool
}

func (v NullableMtbCarePlanProcedureRecommendationsInner) Get() *MtbCarePlanProcedureRecommendationsInner {
	return v.value
}

func (v *NullableMtbCarePlanProcedureRecommendationsInner) Set(val *MtbCarePlanProcedureRecommendationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbCarePlanProcedureRecommendationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbCarePlanProcedureRecommendationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbCarePlanProcedureRecommendationsInner(val *MtbCarePlanProcedureRecommendationsInner) *NullableMtbCarePlanProcedureRecommendationsInner {
	return &NullableMtbCarePlanProcedureRecommendationsInner{value: val, isSet: true}
}

func (v NullableMtbCarePlanProcedureRecommendationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbCarePlanProcedureRecommendationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
