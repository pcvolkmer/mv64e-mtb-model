package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ClaimResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimResponse{}

// ClaimResponse struct for ClaimResponse
type ClaimResponse struct {
	StatusReason []ClaimResponseStatusReasonCoding `json:"statusReason,omitempty"`
	Claim        Reference                         `json:"claim"`
	Id           string                            `json:"id"`
	Status       *ClaimResponseStatusCoding        `json:"status,omitempty"`
	IssuedOn     string                            `json:"issuedOn"`
	Patient      Reference                         `json:"patient"`
}

type _ClaimResponse ClaimResponse

// NewClaimResponse instantiates a new ClaimResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimResponse(claim Reference, id string, issuedOn string, patient Reference) *ClaimResponse {
	this := ClaimResponse{}
	this.Claim = claim
	this.Id = id
	this.IssuedOn = issuedOn
	this.Patient = patient
	return &this
}

// NewClaimResponseWithDefaults instantiates a new ClaimResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimResponseWithDefaults() *ClaimResponse {
	this := ClaimResponse{}
	return &this
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *ClaimResponse) GetStatusReason() []ClaimResponseStatusReasonCoding {
	if o == nil || IsNil(o.StatusReason) {
		var ret []ClaimResponseStatusReasonCoding
		return ret
	}
	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetStatusReasonOk() ([]ClaimResponseStatusReasonCoding, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *ClaimResponse) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given []ClaimResponseStatusReasonCoding and assigns it to the StatusReason field.
func (o *ClaimResponse) SetStatusReason(v []ClaimResponseStatusReasonCoding) {
	o.StatusReason = v
}

// GetClaim returns the Claim field value
func (o *ClaimResponse) GetClaim() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Claim
}

// GetClaimOk returns a tuple with the Claim field value
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetClaimOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Claim, true
}

// SetClaim sets field value
func (o *ClaimResponse) SetClaim(v Reference) {
	o.Claim = v
}

// GetId returns the Id field value
func (o *ClaimResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClaimResponse) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClaimResponse) GetStatus() ClaimResponseStatusCoding {
	if o == nil || IsNil(o.Status) {
		var ret ClaimResponseStatusCoding
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetStatusOk() (*ClaimResponseStatusCoding, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClaimResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClaimResponseStatusCoding and assigns it to the Status field.
func (o *ClaimResponse) SetStatus(v ClaimResponseStatusCoding) {
	o.Status = &v
}

// GetIssuedOn returns the IssuedOn field value
func (o *ClaimResponse) GetIssuedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedOn, true
}

// SetIssuedOn sets field value
func (o *ClaimResponse) SetIssuedOn(v string) {
	o.IssuedOn = v
}

// GetPatient returns the Patient field value
func (o *ClaimResponse) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *ClaimResponse) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *ClaimResponse) SetPatient(v Reference) {
	o.Patient = v
}

func (o ClaimResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	toSerialize["claim"] = o.Claim
	toSerialize["id"] = o.Id
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["issuedOn"] = o.IssuedOn
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *ClaimResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"claim",
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

	varClaimResponse := _ClaimResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClaimResponse)

	if err != nil {
		return err
	}

	*o = ClaimResponse(varClaimResponse)

	return err
}

type NullableClaimResponse struct {
	value *ClaimResponse
	isSet bool
}

func (v NullableClaimResponse) Get() *ClaimResponse {
	return v.value
}

func (v *NullableClaimResponse) Set(val *ClaimResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimResponse(val *ClaimResponse) *NullableClaimResponse {
	return &NullableClaimResponse{value: val, isSet: true}
}

func (v NullableClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
