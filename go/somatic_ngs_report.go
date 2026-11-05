package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the SomaticNgsReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SomaticNgsReport{}

// SomaticNgsReport struct for SomaticNgsReport
type SomaticNgsReport struct {
	Metadata             []NgsReportMetadata     `json:"metadata"`
	Type                 NgsReportTypeCoding     `json:"type"`
	Specimen             Reference               `json:"specimen"`
	IssuedOn             string                  `json:"issuedOn"`
	Patient              Reference               `json:"patient"`
	Id                   string                  `json:"id"`
	Results              SomaticNgsReportResults `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _SomaticNgsReport SomaticNgsReport

// NewSomaticNgsReport instantiates a new SomaticNgsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSomaticNgsReport(metadata []NgsReportMetadata, type_ NgsReportTypeCoding, specimen Reference, issuedOn string, patient Reference, id string, results SomaticNgsReportResults) *SomaticNgsReport {
	this := SomaticNgsReport{}
	this.Metadata = metadata
	this.Type = type_
	this.Specimen = specimen
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	this.Results = results
	return &this
}

// NewSomaticNgsReportWithDefaults instantiates a new SomaticNgsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSomaticNgsReportWithDefaults() *SomaticNgsReport {
	this := SomaticNgsReport{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *SomaticNgsReport) GetMetadata() []NgsReportMetadata {
	if o == nil {
		var ret []NgsReportMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetMetadataOk() ([]NgsReportMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *SomaticNgsReport) SetMetadata(v []NgsReportMetadata) {
	o.Metadata = v
}

// GetType returns the Type field value
func (o *SomaticNgsReport) GetType() NgsReportTypeCoding {
	if o == nil {
		var ret NgsReportTypeCoding
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetTypeOk() (*NgsReportTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SomaticNgsReport) SetType(v NgsReportTypeCoding) {
	o.Type = v
}

// GetSpecimen returns the Specimen field value
func (o *SomaticNgsReport) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *SomaticNgsReport) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *SomaticNgsReport) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *SomaticNgsReport) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *SomaticNgsReport) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *SomaticNgsReport) SetPatient(v Reference) {
	o.Patient = v
}

// GetId returns the Id field value
func (o *SomaticNgsReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SomaticNgsReport) SetId(v string) {
	o.Id = v
}

// GetResults returns the Results field value
func (o *SomaticNgsReport) GetResults() SomaticNgsReportResults {
	if o == nil {
		var ret SomaticNgsReportResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SomaticNgsReport) GetResultsOk() (*SomaticNgsReportResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *SomaticNgsReport) SetResults(v SomaticNgsReportResults) {
	o.Results = v
}

func (o SomaticNgsReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SomaticNgsReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	toSerialize["type"] = o.Type
	toSerialize["specimen"] = o.Specimen
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["patient"] = o.Patient
	toSerialize["id"] = o.Id
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SomaticNgsReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
		"type",
		"specimen",
		"issuedOn",
		"patient",
		"id",
		"results",
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

	varSomaticNgsReport := _SomaticNgsReport{}

	err = json.Unmarshal(data, &varSomaticNgsReport)

	if err != nil {
		return err
	}

	*o = SomaticNgsReport(varSomaticNgsReport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "type")
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "id")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSomaticNgsReport struct {
	value *SomaticNgsReport
	isSet bool
}

func (v NullableSomaticNgsReport) Get() *SomaticNgsReport {
	return v.value
}

func (v *NullableSomaticNgsReport) Set(val *SomaticNgsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSomaticNgsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSomaticNgsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSomaticNgsReport(val *SomaticNgsReport) *NullableSomaticNgsReport {
	return &NullableSomaticNgsReport{value: val, isSet: true}
}

func (v NullableSomaticNgsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSomaticNgsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
