package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the MvhMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MvhMetadata{}

// MvhMetadata This DTO corresponds to a \"Prüfbericht\" for the Model Project. The structure is slightly different from the schema defined for the BfArM API: the conversion/mapping onto the later takes place in the central data node. Also, this DTO contains additional fields not represented in the BfArM \"Prüfbericht\" itself, but which are required for compilation of appendices 1/2 of the quarterly report.
type MvhMetadata struct {
	ModelProjectConsent          MvhMetadataModelProjectConsent `json:"modelProjectConsent"`
	ResearchConsents             []map[string]interface{}       `json:"researchConsents,omitempty"`
	EpisodeOfCare                *Reference                     `json:"episodeOfCare,omitempty"`
	ReasonResearchConsentMissing *BroadConsentReasonMissing     `json:"reasonResearchConsentMissing,omitempty"`
	TransferTAN                  string                         `json:"transferTAN"`
	Type                         MvhSubmissionType              `json:"type"`
	AdditionalProperties         map[string]interface{}
}

type _MvhMetadata MvhMetadata

// NewMvhMetadata instantiates a new MvhMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMvhMetadata(modelProjectConsent MvhMetadataModelProjectConsent, transferTAN string, type_ MvhSubmissionType) *MvhMetadata {
	this := MvhMetadata{}
	this.ModelProjectConsent = modelProjectConsent
	this.TransferTAN = transferTAN
	this.Type = type_
	return &this
}

// NewMvhMetadataWithDefaults instantiates a new MvhMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMvhMetadataWithDefaults() *MvhMetadata {
	this := MvhMetadata{}
	return &this
}

// GetModelProjectConsent returns the ModelProjectConsent field value
func (o *MvhMetadata) GetModelProjectConsent() MvhMetadataModelProjectConsent {
	if o == nil {
		var ret MvhMetadataModelProjectConsent
		return ret
	}

	return o.ModelProjectConsent
}

// GetModelProjectConsentOk returns a tuple with the ModelProjectConsent field value
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetModelProjectConsentOk() (*MvhMetadataModelProjectConsent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelProjectConsent, true
}

// SetModelProjectConsent sets field value
func (o *MvhMetadata) SetModelProjectConsent(v MvhMetadataModelProjectConsent) {
	o.ModelProjectConsent = v
}

// GetResearchConsents returns the ResearchConsents field value if set, zero value otherwise.
func (o *MvhMetadata) GetResearchConsents() []map[string]interface{} {
	if o == nil || IsNil(o.ResearchConsents) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ResearchConsents
}

// GetResearchConsentsOk returns a tuple with the ResearchConsents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetResearchConsentsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResearchConsents) {
		return nil, false
	}
	return o.ResearchConsents, true
}

// HasResearchConsents returns a boolean if a field has been set.
func (o *MvhMetadata) HasResearchConsents() bool {
	if o != nil && !IsNil(o.ResearchConsents) {
		return true
	}

	return false
}

// SetResearchConsents gets a reference to the given []map[string]interface{} and assigns it to the ResearchConsents field.
func (o *MvhMetadata) SetResearchConsents(v []map[string]interface{}) {
	o.ResearchConsents = v
}

// GetEpisodeOfCare returns the EpisodeOfCare field value if set, zero value otherwise.
func (o *MvhMetadata) GetEpisodeOfCare() Reference {
	if o == nil || IsNil(o.EpisodeOfCare) {
		var ret Reference
		return ret
	}
	return *o.EpisodeOfCare
}

// GetEpisodeOfCareOk returns a tuple with the EpisodeOfCare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetEpisodeOfCareOk() (*Reference, bool) {
	if o == nil || IsNil(o.EpisodeOfCare) {
		return nil, false
	}
	return o.EpisodeOfCare, true
}

// HasEpisodeOfCare returns a boolean if a field has been set.
func (o *MvhMetadata) HasEpisodeOfCare() bool {
	if o != nil && !IsNil(o.EpisodeOfCare) {
		return true
	}

	return false
}

// SetEpisodeOfCare gets a reference to the given Reference and assigns it to the EpisodeOfCare field.
func (o *MvhMetadata) SetEpisodeOfCare(v Reference) {
	o.EpisodeOfCare = &v
}

// GetReasonResearchConsentMissing returns the ReasonResearchConsentMissing field value if set, zero value otherwise.
func (o *MvhMetadata) GetReasonResearchConsentMissing() BroadConsentReasonMissing {
	if o == nil || IsNil(o.ReasonResearchConsentMissing) {
		var ret BroadConsentReasonMissing
		return ret
	}
	return *o.ReasonResearchConsentMissing
}

// GetReasonResearchConsentMissingOk returns a tuple with the ReasonResearchConsentMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetReasonResearchConsentMissingOk() (*BroadConsentReasonMissing, bool) {
	if o == nil || IsNil(o.ReasonResearchConsentMissing) {
		return nil, false
	}
	return o.ReasonResearchConsentMissing, true
}

// HasReasonResearchConsentMissing returns a boolean if a field has been set.
func (o *MvhMetadata) HasReasonResearchConsentMissing() bool {
	if o != nil && !IsNil(o.ReasonResearchConsentMissing) {
		return true
	}

	return false
}

// SetReasonResearchConsentMissing gets a reference to the given BroadConsentReasonMissing and assigns it to the ReasonResearchConsentMissing field.
func (o *MvhMetadata) SetReasonResearchConsentMissing(v BroadConsentReasonMissing) {
	o.ReasonResearchConsentMissing = &v
}

// GetTransferTAN returns the TransferTAN field value
func (o *MvhMetadata) GetTransferTAN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferTAN
}

// GetTransferTANOk returns a tuple with the TransferTAN field value
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetTransferTANOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferTAN, true
}

// SetTransferTAN sets field value
func (o *MvhMetadata) SetTransferTAN(v string) {
	o.TransferTAN = v
}

// GetType returns the Type field value
func (o *MvhMetadata) GetType() MvhSubmissionType {
	if o == nil {
		var ret MvhSubmissionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MvhMetadata) GetTypeOk() (*MvhSubmissionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MvhMetadata) SetType(v MvhSubmissionType) {
	o.Type = v
}

func (o MvhMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MvhMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modelProjectConsent"] = o.ModelProjectConsent
	if !IsNil(o.ResearchConsents) {
		toSerialize["researchConsents"] = o.ResearchConsents
	}
	if !IsNil(o.EpisodeOfCare) {
		toSerialize["episodeOfCare"] = o.EpisodeOfCare
	}
	if !IsNil(o.ReasonResearchConsentMissing) {
		toSerialize["reasonResearchConsentMissing"] = o.ReasonResearchConsentMissing
	}
	toSerialize["transferTAN"] = o.TransferTAN
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MvhMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modelProjectConsent",
		"transferTAN",
		"type",
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

	varMvhMetadata := _MvhMetadata{}

	err = json.Unmarshal(data, &varMvhMetadata)

	if err != nil {
		return err
	}

	*o = MvhMetadata(varMvhMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "modelProjectConsent")
		delete(additionalProperties, "researchConsents")
		delete(additionalProperties, "episodeOfCare")
		delete(additionalProperties, "reasonResearchConsentMissing")
		delete(additionalProperties, "transferTAN")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMvhMetadata struct {
	value *MvhMetadata
	isSet bool
}

func (v NullableMvhMetadata) Get() *MvhMetadata {
	return v.value
}

func (v *NullableMvhMetadata) Set(val *MvhMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMvhMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMvhMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMvhMetadata(val *MvhMetadata) *NullableMvhMetadata {
	return &NullableMvhMetadata{value: val, isSet: true}
}

func (v NullableMvhMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMvhMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
