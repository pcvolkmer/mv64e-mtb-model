package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorStaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorStaging{}

// TumorStaging struct for TumorStaging
type TumorStaging struct {
	Date                 string                         `json:"date"`
	Method               TumorStagingMethodCoding       `json:"method"`
	TnmClassification    *TumorStagingTnmClassification `json:"tnmClassification,omitempty"`
	OtherClassifications []Coding                       `json:"otherClassifications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TumorStaging TumorStaging

// NewTumorStaging instantiates a new TumorStaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorStaging(date string, method TumorStagingMethodCoding) *TumorStaging {
	this := TumorStaging{}
	this.Date = date
	this.Method = method
	return &this
}

// NewTumorStagingWithDefaults instantiates a new TumorStaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorStagingWithDefaults() *TumorStaging {
	this := TumorStaging{}
	return &this
}

// GetDate returns the Date field value
func (o *TumorStaging) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *TumorStaging) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *TumorStaging) SetDate(v string) {
	o.Date = v
}

// GetMethod returns the Method field value
func (o *TumorStaging) GetMethod() TumorStagingMethodCoding {
	if o == nil {
		var ret TumorStagingMethodCoding
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TumorStaging) GetMethodOk() (*TumorStagingMethodCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TumorStaging) SetMethod(v TumorStagingMethodCoding) {
	o.Method = v
}

// GetTnmClassification returns the TnmClassification field value if set, zero value otherwise.
func (o *TumorStaging) GetTnmClassification() TumorStagingTnmClassification {
	if o == nil || IsNil(o.TnmClassification) {
		var ret TumorStagingTnmClassification
		return ret
	}
	return *o.TnmClassification
}

// GetTnmClassificationOk returns a tuple with the TnmClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorStaging) GetTnmClassificationOk() (*TumorStagingTnmClassification, bool) {
	if o == nil || IsNil(o.TnmClassification) {
		return nil, false
	}
	return o.TnmClassification, true
}

// HasTnmClassification returns a boolean if a field has been set.
func (o *TumorStaging) HasTnmClassification() bool {
	if o != nil && !IsNil(o.TnmClassification) {
		return true
	}

	return false
}

// SetTnmClassification gets a reference to the given TumorStagingTnmClassification and assigns it to the TnmClassification field.
func (o *TumorStaging) SetTnmClassification(v TumorStagingTnmClassification) {
	o.TnmClassification = &v
}

// GetOtherClassifications returns the OtherClassifications field value if set, zero value otherwise.
func (o *TumorStaging) GetOtherClassifications() []Coding {
	if o == nil || IsNil(o.OtherClassifications) {
		var ret []Coding
		return ret
	}
	return o.OtherClassifications
}

// GetOtherClassificationsOk returns a tuple with the OtherClassifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorStaging) GetOtherClassificationsOk() ([]Coding, bool) {
	if o == nil || IsNil(o.OtherClassifications) {
		return nil, false
	}
	return o.OtherClassifications, true
}

// HasOtherClassifications returns a boolean if a field has been set.
func (o *TumorStaging) HasOtherClassifications() bool {
	if o != nil && !IsNil(o.OtherClassifications) {
		return true
	}

	return false
}

// SetOtherClassifications gets a reference to the given []Coding and assigns it to the OtherClassifications field.
func (o *TumorStaging) SetOtherClassifications(v []Coding) {
	o.OtherClassifications = v
}

func (o TumorStaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorStaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["method"] = o.Method
	if !IsNil(o.TnmClassification) {
		toSerialize["tnmClassification"] = o.TnmClassification
	}
	if !IsNil(o.OtherClassifications) {
		toSerialize["otherClassifications"] = o.OtherClassifications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TumorStaging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"method",
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

	varTumorStaging := _TumorStaging{}

	err = json.Unmarshal(data, &varTumorStaging)

	if err != nil {
		return err
	}

	*o = TumorStaging(varTumorStaging)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "method")
		delete(additionalProperties, "tnmClassification")
		delete(additionalProperties, "otherClassifications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTumorStaging struct {
	value *TumorStaging
	isSet bool
}

func (v NullableTumorStaging) Get() *TumorStaging {
	return v.value
}

func (v *NullableTumorStaging) Set(val *TumorStaging) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorStaging) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorStaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorStaging(val *TumorStaging) *NullableTumorStaging {
	return &NullableTumorStaging{value: val, isSet: true}
}

func (v NullableTumorStaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorStaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
