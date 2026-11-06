package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the PeriodDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeriodDate{}

// PeriodDate struct for PeriodDate
type PeriodDate struct {
	Start                string  `json:"start"`
	End                  *string `json:"end,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PeriodDate PeriodDate

// NewPeriodDate instantiates a new PeriodDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodDate(start string) *PeriodDate {
	this := PeriodDate{}
	this.Start = start
	return &this
}

// NewPeriodDateWithDefaults instantiates a new PeriodDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodDateWithDefaults() *PeriodDate {
	this := PeriodDate{}
	return &this
}

// GetStart returns the Start field value
func (o *PeriodDate) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *PeriodDate) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *PeriodDate) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *PeriodDate) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodDate) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *PeriodDate) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *PeriodDate) SetEnd(v string) {
	o.End = &v
}

func (o PeriodDate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeriodDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PeriodDate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
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

	varPeriodDate := _PeriodDate{}

	err = json.Unmarshal(data, &varPeriodDate)

	if err != nil {
		return err
	}

	*o = PeriodDate(varPeriodDate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePeriodDate struct {
	value *PeriodDate
	isSet bool
}

func (v NullablePeriodDate) Get() *PeriodDate {
	return v.value
}

func (v *NullablePeriodDate) Set(val *PeriodDate) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodDate) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodDate(val *PeriodDate) *NullablePeriodDate {
	return &NullablePeriodDate{value: val, isSet: true}
}

func (v NullablePeriodDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
