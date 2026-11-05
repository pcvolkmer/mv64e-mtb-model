package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MolecularDiagnosticReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularDiagnosticReport{}

// MolecularDiagnosticReport struct for MolecularDiagnosticReport
type MolecularDiagnosticReport struct {
	Type      MolecularDiagnosticReportTypeCoding `json:"type"`
	Specimen  Reference                           `json:"specimen"`
	IssuedOn  string                              `json:"issuedOn"`
	Performer *Reference                          `json:"performer,omitempty"`
	Patient   Reference                           `json:"patient"`
	Id        string                              `json:"id"`
	Results   []string                            `json:"results,omitempty"`
}

type _MolecularDiagnosticReport MolecularDiagnosticReport

// NewMolecularDiagnosticReport instantiates a new MolecularDiagnosticReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularDiagnosticReport(type_ MolecularDiagnosticReportTypeCoding, specimen Reference, issuedOn string, patient Reference, id string) *MolecularDiagnosticReport {
	this := MolecularDiagnosticReport{}
	this.Type = type_
	this.Specimen = specimen
	this.IssuedOn = issuedOn
	this.Patient = patient
	this.Id = id
	return &this
}

// NewMolecularDiagnosticReportWithDefaults instantiates a new MolecularDiagnosticReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularDiagnosticReportWithDefaults() *MolecularDiagnosticReport {
	this := MolecularDiagnosticReport{}
	return &this
}

// GetType returns the Type field value
func (o *MolecularDiagnosticReport) GetType() MolecularDiagnosticReportTypeCoding {
	if o == nil {
		var ret MolecularDiagnosticReportTypeCoding
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetTypeOk() (*MolecularDiagnosticReportTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MolecularDiagnosticReport) SetType(v MolecularDiagnosticReportTypeCoding) {
	o.Type = v
}

// GetSpecimen returns the Specimen field value
func (o *MolecularDiagnosticReport) GetSpecimen() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetSpecimenOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specimen, true
}

// SetSpecimen sets field value
func (o *MolecularDiagnosticReport) SetSpecimen(v Reference) {
	o.Specimen = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *MolecularDiagnosticReport) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *MolecularDiagnosticReport) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *MolecularDiagnosticReport) GetPerformer() Reference {
	if o == nil || IsNil(o.Performer) {
		var ret Reference
		return ret
	}
	return *o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetPerformerOk() (*Reference, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *MolecularDiagnosticReport) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given Reference and assigns it to the Performer field.
func (o *MolecularDiagnosticReport) SetPerformer(v Reference) {
	o.Performer = &v
}

// GetPatient returns the Patient field value
func (o *MolecularDiagnosticReport) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *MolecularDiagnosticReport) SetPatient(v Reference) {
	o.Patient = v
}

// GetId returns the Id field value
func (o *MolecularDiagnosticReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MolecularDiagnosticReport) SetId(v string) {
	o.Id = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MolecularDiagnosticReport) GetResults() []string {
	if o == nil || IsNil(o.Results) {
		var ret []string
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularDiagnosticReport) GetResultsOk() ([]string, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MolecularDiagnosticReport) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []string and assigns it to the Results field.
func (o *MolecularDiagnosticReport) SetResults(v []string) {
	o.Results = v
}

func (o MolecularDiagnosticReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularDiagnosticReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["specimen"] = o.Specimen
	toSerialize["issuedOn"] = o.IssuedOn
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
	}
	toSerialize["patient"] = o.Patient
	toSerialize["id"] = o.Id
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

func (o *MolecularDiagnosticReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"specimen",
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

	varMolecularDiagnosticReport := _MolecularDiagnosticReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMolecularDiagnosticReport)

	if err != nil {
		return err
	}

	*o = MolecularDiagnosticReport(varMolecularDiagnosticReport)

	return err
}

type NullableMolecularDiagnosticReport struct {
	value *MolecularDiagnosticReport
	isSet bool
}

func (v NullableMolecularDiagnosticReport) Get() *MolecularDiagnosticReport {
	return v.value
}

func (v *NullableMolecularDiagnosticReport) Set(val *MolecularDiagnosticReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularDiagnosticReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularDiagnosticReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularDiagnosticReport(val *MolecularDiagnosticReport) *NullableMolecularDiagnosticReport {
	return &NullableMolecularDiagnosticReport{value: val, isSet: true}
}

func (v NullableMolecularDiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularDiagnosticReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
