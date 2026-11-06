package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the IhcReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IhcReport{}

// IhcReport struct for IhcReport
type IhcReport struct {
	Specimen             Reference        `json:"specimen"`
	IssuedOn             string           `json:"issuedOn"`
	Patient              Reference        `json:"patient"`
	Id                   string           `json:"id"`
	Results              IhcReportResults `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _IhcReport IhcReport

// NewIhcReport instantiates a new IhcReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIhcReport(specimen Reference, issuedOn string, patient Reference, id string, results IhcReportResults) *IhcReport {
	this := IhcReport{}
	this.Specimen = specimen
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	this.Results = results
	return &this
}

// NewIhcReportWithDefaults instantiates a new IhcReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIhcReportWithDefaults() *IhcReport {
	this := IhcReport{}
	return &this
}

// GetSpecimen returns the Specimen field value
func (o *IhcReport) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *IhcReport) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *IhcReport) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *IhcReport) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *IhcReport) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *IhcReport) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *IhcReport) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *IhcReport) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *IhcReport) SetPatient(v Reference) {
	o.Patient = v
}

// GetId returns the Id field value
func (o *IhcReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IhcReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IhcReport) SetId(v string) {
	o.Id = v
}

// GetResults returns the Results field value
func (o *IhcReport) GetResults() IhcReportResults {
	if o == nil {
		var ret IhcReportResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *IhcReport) GetResultsOk() (*IhcReportResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *IhcReport) SetResults(v IhcReportResults) {
	o.Results = v
}

func (o IhcReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IhcReport) ToMap() (map[string]interface{}, error) {
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

func (o *IhcReport) UnmarshalJSON(data []byte) (err error) {
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

	varIhcReport := _IhcReport{}

	err = json.Unmarshal(data, &varIhcReport)

	if err != nil {
		return err
	}

	*o = IhcReport(varIhcReport)

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

type NullableIhcReport struct {
	value *IhcReport
	isSet bool
}

func (v NullableIhcReport) Get() *IhcReport {
	return v.value
}

func (v *NullableIhcReport) Set(val *IhcReport) {
	v.value = val
	v.isSet = true
}

func (v NullableIhcReport) IsSet() bool {
	return v.isSet
}

func (v *NullableIhcReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIhcReport(val *IhcReport) *NullableIhcReport {
	return &NullableIhcReport{value: val, isSet: true}
}

func (v NullableIhcReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIhcReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
