package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the HistologyReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistologyReport{}

// HistologyReport struct for HistologyReport
type HistologyReport struct {
	Specimen             Reference              `json:"specimen"`
	IssuedOn             string                 `json:"issuedOn"`
	Patient              Reference              `json:"patient"`
	Id                   string                 `json:"id"`
	Results              HistologyReportResults `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _HistologyReport HistologyReport

// NewHistologyReport instantiates a new HistologyReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistologyReport(specimen Reference, issuedOn string, patient Reference, id string, results HistologyReportResults) *HistologyReport {
	this := HistologyReport{}
	this.Specimen = specimen
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	this.Results = results
	return &this
}

// NewHistologyReportWithDefaults instantiates a new HistologyReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistologyReportWithDefaults() *HistologyReport {
	this := HistologyReport{}
	return &this
}

// GetSpecimen returns the Specimen field value
func (o *HistologyReport) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *HistologyReport) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *HistologyReport) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *HistologyReport) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *HistologyReport) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *HistologyReport) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *HistologyReport) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *HistologyReport) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *HistologyReport) SetPatient(v Reference) {
	o.Patient = v
}

// GetId returns the Id field value
func (o *HistologyReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistologyReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistologyReport) SetId(v string) {
	o.Id = v
}

// GetResults returns the Results field value
func (o *HistologyReport) GetResults() HistologyReportResults {
	if o == nil {
		var ret HistologyReportResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *HistologyReport) GetResultsOk() (*HistologyReportResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *HistologyReport) SetResults(v HistologyReportResults) {
	o.Results = v
}

func (o HistologyReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistologyReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *HistologyReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varHistologyReport := _HistologyReport{}

	err = json.Unmarshal(data, &varHistologyReport)

	if err != nil {
		return err
	}

	*o = HistologyReport(varHistologyReport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "specimen")
		delete(additionalProperties, "issuedOn")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "id")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistologyReport struct {
	value *HistologyReport
	isSet bool
}

func (v NullableHistologyReport) Get() *HistologyReport {
	return v.value
}

func (v *NullableHistologyReport) Set(val *HistologyReport) {
	v.value = val
	v.isSet = true
}

func (v NullableHistologyReport) IsSet() bool {
	return v.isSet
}

func (v *NullableHistologyReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistologyReport(val *HistologyReport) *NullableHistologyReport {
	return &NullableHistologyReport{value: val, isSet: true}
}

func (v NullableHistologyReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistologyReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
