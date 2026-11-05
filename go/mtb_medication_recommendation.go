package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MtbMedicationRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtbMedicationRecommendation{}

// MtbMedicationRecommendation struct for MtbMedicationRecommendation
type MtbMedicationRecommendation struct {
	Priority             MtbRecommendationPriorityCoding             `json:"priority"`
	Medication           []AtcUnregisteredMedicationCoding           `json:"medication"`
	SupportingVariants   []GeneAlterationReference                   `json:"supportingVariants,omitempty"`
	UseType              *MtbMedicationRecommendationUseTypeCoding   `json:"useType,omitempty"`
	Reason               *Reference                                  `json:"reason,omitempty"`
	Category             []MtbMedicationRecommendationCategoryCoding `json:"category,omitempty"`
	IssuedOn             string                                      `json:"issuedOn"`
	Patient              Reference                                   `json:"patient"`
	SupportingFindings   []Reference                                 `json:"supportingFindings,omitempty"`
	LevelOfEvidence      *LevelOfEvidence                            `json:"levelOfEvidence,omitempty"`
	Id                   string                                      `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _MtbMedicationRecommendation MtbMedicationRecommendation

// NewMtbMedicationRecommendation instantiates a new MtbMedicationRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtbMedicationRecommendation(priority MtbRecommendationPriorityCoding, medication []AtcUnregisteredMedicationCoding, issuedOn string, patient Reference, id string) *MtbMedicationRecommendation {
	this := MtbMedicationRecommendation{}
	this.Priority = priority
	this.Medication = medication
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	return &this
}

// NewMtbMedicationRecommendationWithDefaults instantiates a new MtbMedicationRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtbMedicationRecommendationWithDefaults() *MtbMedicationRecommendation {
	this := MtbMedicationRecommendation{}
	return &this
}

// GetPriority returns the Priority field value
func (o *MtbMedicationRecommendation) GetPriority() MtbRecommendationPriorityCoding {
	if o == nil {
		var ret MtbRecommendationPriorityCoding
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetPriorityOk() (*MtbRecommendationPriorityCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *MtbMedicationRecommendation) SetPriority(v MtbRecommendationPriorityCoding) {
	o.Priority = v
}

// GetMedication returns the Medication field value
func (o *MtbMedicationRecommendation) GetMedication() []AtcUnregisteredMedicationCoding {
	if o == nil {
		var ret []AtcUnregisteredMedicationCoding
		return ret
	}

	return o.Medication
}

// GetMedicationOk returns a tuple with the Medication field value
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetMedicationOk() ([]AtcUnregisteredMedicationCoding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medication, true
}

// SetMedication sets field value
func (o *MtbMedicationRecommendation) SetMedication(v []AtcUnregisteredMedicationCoding) {
	o.Medication = v
}

// GetSupportingVariants returns the SupportingVariants field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetSupportingVariants() []GeneAlterationReference {
	if o == nil || IsNil(o.SupportingVariants) {
		var ret []GeneAlterationReference
		return ret
	}
	return o.SupportingVariants
}

// GetSupportingVariantsOk returns a tuple with the SupportingVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetSupportingVariantsOk() ([]GeneAlterationReference, bool) {
	if o == nil || IsNil(o.SupportingVariants) {
		return nil, false
	}
	return o.SupportingVariants, true
}

// HasSupportingVariants returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasSupportingVariants() bool {
	if o != nil && !IsNil(o.SupportingVariants) {
		return true
	}

	return false
}

// SetSupportingVariants gets a reference to the given []GeneAlterationReference and assigns it to the SupportingVariants field.
func (o *MtbMedicationRecommendation) SetSupportingVariants(v []GeneAlterationReference) {
	o.SupportingVariants = v
}

// GetUseType returns the UseType field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetUseType() MtbMedicationRecommendationUseTypeCoding {
	if o == nil || IsNil(o.UseType) {
		var ret MtbMedicationRecommendationUseTypeCoding
		return ret
	}
	return *o.UseType
}

// GetUseTypeOk returns a tuple with the UseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetUseTypeOk() (*MtbMedicationRecommendationUseTypeCoding, bool) {
	if o == nil || IsNil(o.UseType) {
		return nil, false
	}
	return o.UseType, true
}

// HasUseType returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasUseType() bool {
	if o != nil && !IsNil(o.UseType) {
		return true
	}

	return false
}

// SetUseType gets a reference to the given MtbMedicationRecommendationUseTypeCoding and assigns it to the UseType field.
func (o *MtbMedicationRecommendation) SetUseType(v MtbMedicationRecommendationUseTypeCoding) {
	o.UseType = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *MtbMedicationRecommendation) SetReason(v Reference) {
	o.Reason = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetCategory() []MtbMedicationRecommendationCategoryCoding {
	if o == nil || IsNil(o.Category) {
		var ret []MtbMedicationRecommendationCategoryCoding
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetCategoryOk() ([]MtbMedicationRecommendationCategoryCoding, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []MtbMedicationRecommendationCategoryCoding and assigns it to the Category field.
func (o *MtbMedicationRecommendation) SetCategory(v []MtbMedicationRecommendationCategoryCoding) {
	o.Category = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MtbMedicationRecommendation) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MtbMedicationRecommendation) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *MtbMedicationRecommendation) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MtbMedicationRecommendation) SetPatient(v Reference) {
	o.Patient = v
}

// GetSupportingFindings returns the SupportingFindings field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetSupportingFindings() []Reference {
	if o == nil || IsNil(o.SupportingFindings) {
		var ret []Reference
		return ret
	}
	return o.SupportingFindings
}

// GetSupportingFindingsOk returns a tuple with the SupportingFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetSupportingFindingsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.SupportingFindings) {
		return nil, false
	}
	return o.SupportingFindings, true
}

// HasSupportingFindings returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasSupportingFindings() bool {
	if o != nil && !IsNil(o.SupportingFindings) {
		return true
	}

	return false
}

// SetSupportingFindings gets a reference to the given []Reference and assigns it to the SupportingFindings field.
func (o *MtbMedicationRecommendation) SetSupportingFindings(v []Reference) {
	o.SupportingFindings = v
}

// GetLevelOfEvidence returns the LevelOfEvidence field value if set, zero value otherwise.
func (o *MtbMedicationRecommendation) GetLevelOfEvidence() LevelOfEvidence {
	if o == nil || IsNil(o.LevelOfEvidence) {
		var ret LevelOfEvidence
		return ret
	}
	return *o.LevelOfEvidence
}

// GetLevelOfEvidenceOk returns a tuple with the LevelOfEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetLevelOfEvidenceOk() (*LevelOfEvidence, bool) {
	if o == nil || IsNil(o.LevelOfEvidence) {
		return nil, false
	}
	return o.LevelOfEvidence, true
}

// HasLevelOfEvidence returns a boolean if a field has been set.
func (o *MtbMedicationRecommendation) HasLevelOfEvidence() bool {
	if o != nil && !IsNil(o.LevelOfEvidence) {
		return true
	}

	return false
}

// SetLevelOfEvidence gets a reference to the given LevelOfEvidence and assigns it to the LevelOfEvidence field.
func (o *MtbMedicationRecommendation) SetLevelOfEvidence(v LevelOfEvidence) {
	o.LevelOfEvidence = &v
}

// GetId returns the Id field value
func (o *MtbMedicationRecommendation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MtbMedicationRecommendation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MtbMedicationRecommendation) SetId(v string) {
	o.Id = v
}

func (o MtbMedicationRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtbMedicationRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["priority"] = o.Priority
	toSerialize["medication"] = o.Medication
	if !IsNil(o.SupportingVariants) {
		toSerialize["supportingVariants"] = o.SupportingVariants
	}
	if !IsNil(o.UseType) {
		toSerialize["useType"] = o.UseType
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["patient"] = o.Patient
	if !IsNil(o.SupportingFindings) {
		toSerialize["supportingFindings"] = o.SupportingFindings
	}
	if !IsNil(o.LevelOfEvidence) {
		toSerialize["levelOfEvidence"] = o.LevelOfEvidence
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MtbMedicationRecommendation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"priority",
		"medication",
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

	varMtbMedicationRecommendation := _MtbMedicationRecommendation{}

	err = json.Unmarshal(data, &varMtbMedicationRecommendation)

	if err != nil {
		return err
	}

	*o = MtbMedicationRecommendation(varMtbMedicationRecommendation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "priority")
		delete(additionalProperties, "medication")
		delete(additionalProperties, "supportingVariants")
		delete(additionalProperties, "useType")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "category")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "supportingFindings")
		delete(additionalProperties, "levelOfEvidence")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMtbMedicationRecommendation struct {
	value *MtbMedicationRecommendation
	isSet bool
}

func (v NullableMtbMedicationRecommendation) Get() *MtbMedicationRecommendation {
	return v.value
}

func (v *NullableMtbMedicationRecommendation) Set(val *MtbMedicationRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableMtbMedicationRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableMtbMedicationRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtbMedicationRecommendation(val *MtbMedicationRecommendation) *NullableMtbMedicationRecommendation {
	return &NullableMtbMedicationRecommendation{value: val, isSet: true}
}

func (v NullableMtbMedicationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtbMedicationRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
