package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the RnaFusion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RnaFusion{}

// RnaFusion struct for RnaFusion
type RnaFusion struct {
	FusionPartner3prime  RnaFusionFusionPartner3prime    `json:"fusionPartner3prime"`
	Localization         []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	FusionPartner5prime  RnaFusionFusionPartner5prime    `json:"fusionPartner5prime"`
	Effect               *string                         `json:"effect,omitempty"`
	ExternalIds          []VariantExternalId             `json:"externalIds,omitempty"`
	Id                   string                          `json:"id"`
	ReportedNumReads     int32                           `json:"reportedNumReads"`
	Patient              Reference                       `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _RnaFusion RnaFusion

// NewRnaFusion instantiates a new RnaFusion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaFusion(fusionPartner3prime RnaFusionFusionPartner3prime, fusionPartner5prime RnaFusionFusionPartner5prime, id string, reportedNumReads int32, patient Reference) *RnaFusion {
	this := RnaFusion{}
	this.FusionPartner3prime = fusionPartner3prime
	this.FusionPartner5prime = fusionPartner5prime
	this.Id = id
	this.ReportedNumReads = reportedNumReads
	this.Patient = patient
	return &this
}

// NewRnaFusionWithDefaults instantiates a new RnaFusion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaFusionWithDefaults() *RnaFusion {
	this := RnaFusion{}
	return &this
}

// GetFusionPartner3prime returns the FusionPartner3prime field value
func (o *RnaFusion) GetFusionPartner3prime() RnaFusionFusionPartner3prime {
	if o == nil {
		var ret RnaFusionFusionPartner3prime
		return ret
	}

	return o.FusionPartner3prime
}

// GetFusionPartner3primeOk returns a tuple with the FusionPartner3prime field value
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetFusionPartner3primeOk() (*RnaFusionFusionPartner3prime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FusionPartner3prime, true
}

// SetFusionPartner3prime sets field value
func (o *RnaFusion) SetFusionPartner3prime(v RnaFusionFusionPartner3prime) {
	o.FusionPartner3prime = v
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *RnaFusion) GetLocalization() []BaseVariantLocalizationCoding {
	if o == nil || IsNil(o.Localization) {
		var ret []BaseVariantLocalizationCoding
		return ret
	}
	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetLocalizationOk() ([]BaseVariantLocalizationCoding, bool) {
	if o == nil || IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *RnaFusion) HasLocalization() bool {
	if o != nil && !IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given []BaseVariantLocalizationCoding and assigns it to the Localization field.
func (o *RnaFusion) SetLocalization(v []BaseVariantLocalizationCoding) {
	o.Localization = v
}

// GetFusionPartner5prime returns the FusionPartner5prime field value
func (o *RnaFusion) GetFusionPartner5prime() RnaFusionFusionPartner5prime {
	if o == nil {
		var ret RnaFusionFusionPartner5prime
		return ret
	}

	return o.FusionPartner5prime
}

// GetFusionPartner5primeOk returns a tuple with the FusionPartner5prime field value
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetFusionPartner5primeOk() (*RnaFusionFusionPartner5prime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FusionPartner5prime, true
}

// SetFusionPartner5prime sets field value
func (o *RnaFusion) SetFusionPartner5prime(v RnaFusionFusionPartner5prime) {
	o.FusionPartner5prime = v
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *RnaFusion) GetEffect() string {
	if o == nil || IsNil(o.Effect) {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetEffectOk() (*string, bool) {
	if o == nil || IsNil(o.Effect) {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *RnaFusion) HasEffect() bool {
	if o != nil && !IsNil(o.Effect) {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *RnaFusion) SetEffect(v string) {
	o.Effect = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *RnaFusion) GetExternalIds() []VariantExternalId {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []VariantExternalId
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetExternalIdsOk() ([]VariantExternalId, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *RnaFusion) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []VariantExternalId and assigns it to the ExternalIds field.
func (o *RnaFusion) SetExternalIds(v []VariantExternalId) {
	o.ExternalIds = v
}

// GetId returns the Id field value
func (o *RnaFusion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RnaFusion) SetId(v string) {
	o.Id = v
}

// GetReportedNumReads returns the ReportedNumReads field value
func (o *RnaFusion) GetReportedNumReads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReportedNumReads
}

// GetReportedNumReadsOk returns a tuple with the ReportedNumReads field value
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetReportedNumReadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportedNumReads, true
}

// SetReportedNumReads sets field value
func (o *RnaFusion) SetReportedNumReads(v int32) {
	o.ReportedNumReads = v
}

// GetPatient returns the Patient field value
func (o *RnaFusion) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *RnaFusion) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *RnaFusion) SetPatient(v Reference) {
	o.Patient = v
}

func (o RnaFusion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RnaFusion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fusionPartner3prime"] = o.FusionPartner3prime
	if !IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	toSerialize["fusionPartner5prime"] = o.FusionPartner5prime
	if !IsNil(o.Effect) {
		toSerialize["effect"] = o.Effect
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	toSerialize["id"] = o.Id
	toSerialize["reportedNumReads"] = o.ReportedNumReads
	toSerialize["patient"] = o.Patient

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RnaFusion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fusionPartner3prime",
		"fusionPartner5prime",
		"id",
		"reportedNumReads",
		"patient",
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

	varRnaFusion := _RnaFusion{}

	err = json.Unmarshal(data, &varRnaFusion)

	if err != nil {
		return err
	}

	*o = RnaFusion(varRnaFusion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fusionPartner3prime")
		delete(additionalProperties, "localization")
		delete(additionalProperties, "fusionPartner5prime")
		delete(additionalProperties, "effect")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "id")
		delete(additionalProperties, "reportedNumReads")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRnaFusion struct {
	value *RnaFusion
	isSet bool
}

func (v NullableRnaFusion) Get() *RnaFusion {
	return v.value
}

func (v *NullableRnaFusion) Set(val *RnaFusion) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaFusion) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaFusion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaFusion(val *RnaFusion) *NullableRnaFusion {
	return &NullableRnaFusion{value: val, isSet: true}
}

func (v NullableRnaFusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaFusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
