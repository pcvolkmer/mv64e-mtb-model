package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Response{}

// Response struct for Response
type Response struct {
	Method        *ResponseMethodCoding `json:"method,omitempty"`
	Therapy       Reference             `json:"therapy"`
	Id            string                `json:"id"`
	EffectiveDate string                `json:"effectiveDate"`
	Value         RecistCoding          `json:"value"`
	Patient       Reference             `json:"patient"`
}

type _Response Response

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse(therapy Reference, id string, effectiveDate string, value RecistCoding, patient Reference) *Response {
	this := Response{}
	this.Therapy = therapy
	this.Id = id
	this.EffectiveDate = effectiveDate
	this.Value = value
	this.Patient = patient
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *Response) GetMethod() ResponseMethodCoding {
	if o == nil || IsNil(o.Method) {
		var ret ResponseMethodCoding
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetMethodOk() (*ResponseMethodCoding, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *Response) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given ResponseMethodCoding and assigns it to the Method field.
func (o *Response) SetMethod(v ResponseMethodCoding) {
	o.Method = &v
}

// GetTherapy returns the Therapy field value
func (o *Response) GetTherapy() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Therapy
}

// GetTherapyOk returns a tuple with the Therapy field value
// and a boolean to check if the value has been set.
func (o *Response) GetTherapyOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Therapy, true
}

// SetTherapy sets field value
func (o *Response) SetTherapy(v Reference) {
	o.Therapy = v
}

// GetId returns the Id field value
func (o *Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Response) SetId(v string) {
	o.Id = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *Response) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *Response) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *Response) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetValue returns the Value field value
func (o *Response) GetValue() RecistCoding {
	if o == nil {
		var ret RecistCoding
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Response) GetValueOk() (*RecistCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Response) SetValue(v RecistCoding) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *Response) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *Response) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *Response) SetPatient(v Reference) {
	o.Patient = v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	toSerialize["therapy"] = o.Therapy
	toSerialize["id"] = o.Id
	toSerialize["effectiveDate"] = o.EffectiveDate
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"therapy",
		"id",
		"effectiveDate",
		"value",
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

	varResponse := _Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponse)

	if err != nil {
		return err
	}

	*o = Response(varResponse)

	return err
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
