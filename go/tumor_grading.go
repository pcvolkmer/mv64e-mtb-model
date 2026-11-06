package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the TumorGrading type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorGrading{}

// TumorGrading struct for TumorGrading
type TumorGrading struct {
	Date                 string   `json:"date"`
	Codes                []Coding `json:"codes"`
	AdditionalProperties map[string]interface{}
}

type _TumorGrading TumorGrading

// NewTumorGrading instantiates a new TumorGrading object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorGrading(date string, codes []Coding) *TumorGrading {
	this := TumorGrading{}
	this.Date = date
	this.Codes = codes
	return &this
}

// NewTumorGradingWithDefaults instantiates a new TumorGrading object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorGradingWithDefaults() *TumorGrading {
	this := TumorGrading{}
	return &this
}

// GetDate returns the Date field value
func (o *TumorGrading) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *TumorGrading) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *TumorGrading) SetDate(v string) {
	o.Date = v
}

// GetCodes returns the Codes field value
func (o *TumorGrading) GetCodes() []Coding {
	if o == nil {
		var ret []Coding
		return ret
	}

	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value
// and a boolean to check if the value has been set.
func (o *TumorGrading) GetCodesOk() ([]Coding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codes, true
}

// SetCodes sets field value
func (o *TumorGrading) SetCodes(v []Coding) {
	o.Codes = v
}

func (o TumorGrading) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorGrading) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["codes"] = o.Codes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TumorGrading) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"codes",
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

	varTumorGrading := _TumorGrading{}

	err = json.Unmarshal(data, &varTumorGrading)

	if err != nil {
		return err
	}

	*o = TumorGrading(varTumorGrading)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "codes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTumorGrading struct {
	value *TumorGrading
	isSet bool
}

func (v NullableTumorGrading) Get() *TumorGrading {
	return v.value
}

func (v *NullableTumorGrading) Set(val *TumorGrading) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorGrading) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorGrading) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorGrading(val *TumorGrading) *NullableTumorGrading {
	return &NullableTumorGrading{value: val, isSet: true}
}

func (v NullableTumorGrading) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorGrading) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
