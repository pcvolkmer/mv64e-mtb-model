package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the DnaFusion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnaFusion{}

// DnaFusion struct for DnaFusion
type DnaFusion struct {
	FusionPartner3prime  DnaFusionFusionPartner3prime    `json:"fusionPartner3prime"`
	Localization         []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	FusionPartner5prime  DnaFusionFusionPartner5prime    `json:"fusionPartner5prime"`
	ExternalIds          []VariantExternalId             `json:"externalIds,omitempty"`
	Id                   string                          `json:"id"`
	ReportedNumReads     int32                           `json:"reportedNumReads"`
	Patient              Reference                       `json:"patient"`
	AdditionalProperties map[string]interface{}
}

type _DnaFusion DnaFusion

// NewDnaFusion instantiates a new DnaFusion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaFusion(fusionPartner3prime DnaFusionFusionPartner3prime, fusionPartner5prime DnaFusionFusionPartner5prime, id string, reportedNumReads int32, patient Reference) *DnaFusion {
	this := DnaFusion{}
	this.FusionPartner3prime = fusionPartner3prime
	this.FusionPartner5prime = fusionPartner5prime
	this.Id = id
	this.ReportedNumReads = reportedNumReads
	this.Patient = patient
	return &this
}

// NewDnaFusionWithDefaults instantiates a new DnaFusion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaFusionWithDefaults() *DnaFusion {
	this := DnaFusion{}
	return &this
}

// GetFusionPartner3prime returns the FusionPartner3prime field value
func (o *DnaFusion) GetFusionPartner3prime() DnaFusionFusionPartner3prime {
	if o == nil {
		var ret DnaFusionFusionPartner3prime
		return ret
	}

	return o.FusionPartner3prime
}

// GetFusionPartner3primeOk returns a tuple with the FusionPartner3prime field value
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetFusionPartner3primeOk() (*DnaFusionFusionPartner3prime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FusionPartner3prime, true
}

// SetFusionPartner3prime sets field value
func (o *DnaFusion) SetFusionPartner3prime(v DnaFusionFusionPartner3prime) {
	o.FusionPartner3prime = v
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *DnaFusion) GetLocalization() []BaseVariantLocalizationCoding {
	if o == nil || IsNil(o.Localization) {
		var ret []BaseVariantLocalizationCoding
		return ret
	}
	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetLocalizationOk() ([]BaseVariantLocalizationCoding, bool) {
	if o == nil || IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *DnaFusion) HasLocalization() bool {
	if o != nil && !IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given []BaseVariantLocalizationCoding and assigns it to the Localization field.
func (o *DnaFusion) SetLocalization(v []BaseVariantLocalizationCoding) {
	o.Localization = v
}

// GetFusionPartner5prime returns the FusionPartner5prime field value
func (o *DnaFusion) GetFusionPartner5prime() DnaFusionFusionPartner5prime {
	if o == nil {
		var ret DnaFusionFusionPartner5prime
		return ret
	}

	return o.FusionPartner5prime
}

// GetFusionPartner5primeOk returns a tuple with the FusionPartner5prime field value
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetFusionPartner5primeOk() (*DnaFusionFusionPartner5prime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FusionPartner5prime, true
}

// SetFusionPartner5prime sets field value
func (o *DnaFusion) SetFusionPartner5prime(v DnaFusionFusionPartner5prime) {
	o.FusionPartner5prime = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *DnaFusion) GetExternalIds() []VariantExternalId {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []VariantExternalId
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetExternalIdsOk() ([]VariantExternalId, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *DnaFusion) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []VariantExternalId and assigns it to the ExternalIds field.
func (o *DnaFusion) SetExternalIds(v []VariantExternalId) {
	o.ExternalIds = v
}

// GetId returns the Id field value
func (o *DnaFusion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DnaFusion) SetId(v string) {
	o.Id = v
}

// GetReportedNumReads returns the ReportedNumReads field value
func (o *DnaFusion) GetReportedNumReads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReportedNumReads
}

// GetReportedNumReadsOk returns a tuple with the ReportedNumReads field value
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetReportedNumReadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportedNumReads, true
}

// SetReportedNumReads sets field value
func (o *DnaFusion) SetReportedNumReads(v int32) {
	o.ReportedNumReads = v
}

// GetPatient returns the Patient field value
func (o *DnaFusion) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *DnaFusion) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *DnaFusion) SetPatient(v Reference) {
	o.Patient = v
}

func (o DnaFusion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnaFusion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fusionPartner3prime"] = o.FusionPartner3prime
	if !IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	toSerialize["fusionPartner5prime"] = o.FusionPartner5prime
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

func (o *DnaFusion) UnmarshalJSON(data []byte) (err error) {
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

	varDnaFusion := _DnaFusion{}

	err = json.Unmarshal(data, &varDnaFusion)

	if err != nil {
		return err
	}

	*o = DnaFusion(varDnaFusion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fusionPartner3prime")
		delete(additionalProperties, "localization")
		delete(additionalProperties, "fusionPartner5prime")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "id")
		delete(additionalProperties, "reportedNumReads")
		delete(additionalProperties, "patient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDnaFusion struct {
	value *DnaFusion
	isSet bool
}

func (v NullableDnaFusion) Get() *DnaFusion {
	return v.value
}

func (v *NullableDnaFusion) Set(val *DnaFusion) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaFusion) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaFusion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaFusion(val *DnaFusion) *NullableDnaFusion {
	return &NullableDnaFusion{value: val, isSet: true}
}

func (v NullableDnaFusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaFusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
