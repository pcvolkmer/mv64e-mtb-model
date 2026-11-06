package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Claim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Claim{}

// Claim struct for Claim
type Claim struct {
	Stage               *ClaimStageCoding                 `json:"stage,omitempty"`
	RequestedMedication []AtcUnregisteredMedicationCoding `json:"requestedMedication,omitempty"`
	Recommendation      Reference                         `json:"recommendation"`
	Id                  string                            `json:"id"`
	IssuedOn            string                            `json:"issuedOn"`
	Patient             Reference                         `json:"patient"`
}

type _Claim Claim

// NewClaim instantiates a new Claim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaim(recommendation Reference, id string, issuedOn string, patient Reference) *Claim {
	this := Claim{}
	this.Recommendation = recommendation
	this.Id = id
	this.IssuedOn = issuedOn
	this.Patient = patient
	return &this
}

// NewClaimWithDefaults instantiates a new Claim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimWithDefaults() *Claim {
	this := Claim{}
	return &this
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *Claim) GetStage() ClaimStageCoding {
	if o == nil || IsNil(o.Stage) {
		var ret ClaimStageCoding
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Claim) GetStageOk() (*ClaimStageCoding, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *Claim) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given ClaimStageCoding and assigns it to the Stage field.
func (o *Claim) SetStage(v ClaimStageCoding) {
	o.Stage = &v
}

// GetRequestedMedication returns the RequestedMedication field value if set, zero value otherwise.
func (o *Claim) GetRequestedMedication() []AtcUnregisteredMedicationCoding {
	if o == nil || IsNil(o.RequestedMedication) {
		var ret []AtcUnregisteredMedicationCoding
		return ret
	}
	return o.RequestedMedication
}

// GetRequestedMedicationOk returns a tuple with the RequestedMedication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Claim) GetRequestedMedicationOk() ([]AtcUnregisteredMedicationCoding, bool) {
	if o == nil || IsNil(o.RequestedMedication) {
		return nil, false
	}
	return o.RequestedMedication, true
}

// HasRequestedMedication returns a boolean if a field has been set.
func (o *Claim) HasRequestedMedication() bool {
	if o != nil && !IsNil(o.RequestedMedication) {
		return true
	}

	return false
}

// SetRequestedMedication gets a reference to the given []AtcUnregisteredMedicationCoding and assigns it to the RequestedMedication field.
func (o *Claim) SetRequestedMedication(v []AtcUnregisteredMedicationCoding) {
	o.RequestedMedication = v
}

// GetRecommendation returns the Recommendation field value
func (o *Claim) GetRecommendation() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value
// and a boolean to check if the value has been set.
func (o *Claim) GetRecommendationOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommendation, true
}

// SetRecommendation sets field value
func (o *Claim) SetRecommendation(v Reference) {
	o.Recommendation = v
}

// GetId returns the Id field value
func (o *Claim) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Claim) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Claim) SetId(v string) {
	o.Id = v
}

// GetIssuedOn returns the IssuedOn field value
func (o *Claim) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *Claim) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *Claim) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *Claim) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Claim) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Claim) SetPatient(v Reference) {
	o.Patient = v
}

func (o Claim) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Claim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.RequestedMedication) {
		toSerialize["requestedMedication"] = o.RequestedMedication
	}
	toSerialize["recommendation"] = o.Recommendation
	toSerialize["id"] = o.Id
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *Claim) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recommendation",
		"id",
		"issuedOn",
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

	varClaim := _Claim{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClaim)

	if err != nil {
		return err
	}

	*o = Claim(varClaim)

	return err
}

type NullableClaim struct {
	value *Claim
	isSet bool
}

func (v NullableClaim) Get() *Claim {
	return v.value
}

func (v *NullableClaim) Set(val *Claim) {
	v.value = val
	v.isSet = true
}

func (v NullableClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaim(val *Claim) *NullableClaim {
	return &NullableClaim{value: val, isSet: true}
}

func (v NullableClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
