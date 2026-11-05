package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TumorSpecimenCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TumorSpecimenCollection{}

// TumorSpecimenCollection struct for TumorSpecimenCollection
type TumorSpecimenCollection struct {
	Date         *string                                   `json:"date,omitempty"`
	Method       TumorSpecimenCollectionMethodCoding       `json:"method"`
	Localization TumorSpecimenCollectionLocalizationCoding `json:"localization"`
}

type _TumorSpecimenCollection TumorSpecimenCollection

// NewTumorSpecimenCollection instantiates a new TumorSpecimenCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTumorSpecimenCollection(method TumorSpecimenCollectionMethodCoding, localization TumorSpecimenCollectionLocalizationCoding) *TumorSpecimenCollection {
	this := TumorSpecimenCollection{}
	this.Method = method
	this.Localization = localization
	return &this
}

// NewTumorSpecimenCollectionWithDefaults instantiates a new TumorSpecimenCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTumorSpecimenCollectionWithDefaults() *TumorSpecimenCollection {
	this := TumorSpecimenCollection{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TumorSpecimenCollection) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TumorSpecimenCollection) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TumorSpecimenCollection) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *TumorSpecimenCollection) SetDate(v string) {
	o.Date = &v
}

// GetMethod returns the Method field value
func (o *TumorSpecimenCollection) GetMethod() TumorSpecimenCollectionMethodCoding {
	if o == nil {
		var ret TumorSpecimenCollectionMethodCoding
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimenCollection) GetMethodOk() (*TumorSpecimenCollectionMethodCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TumorSpecimenCollection) SetMethod(v TumorSpecimenCollectionMethodCoding) {
	o.Method = v
}

// GetLocalization returns the Localization field value
func (o *TumorSpecimenCollection) GetLocalization() TumorSpecimenCollectionLocalizationCoding {
	if o == nil {
		var ret TumorSpecimenCollectionLocalizationCoding
		return ret
	}

	return o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value
// and a boolean to check if the value has been set.
func (o *TumorSpecimenCollection) GetLocalizationOk() (*TumorSpecimenCollectionLocalizationCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Localization, true
}

// SetLocalization sets field value
func (o *TumorSpecimenCollection) SetLocalization(v TumorSpecimenCollectionLocalizationCoding) {
	o.Localization = v
}

func (o TumorSpecimenCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TumorSpecimenCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	toSerialize["method"] = o.Method
	toSerialize["localization"] = o.Localization
	return toSerialize, nil
}

func (o *TumorSpecimenCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"localization",
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

	varTumorSpecimenCollection := _TumorSpecimenCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTumorSpecimenCollection)

	if err != nil {
		return err
	}

	*o = TumorSpecimenCollection(varTumorSpecimenCollection)

	return err
}

type NullableTumorSpecimenCollection struct {
	value *TumorSpecimenCollection
	isSet bool
}

func (v NullableTumorSpecimenCollection) Get() *TumorSpecimenCollection {
	return v.value
}

func (v *NullableTumorSpecimenCollection) Set(val *TumorSpecimenCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTumorSpecimenCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTumorSpecimenCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTumorSpecimenCollection(val *TumorSpecimenCollection) *NullableTumorSpecimenCollection {
	return &NullableTumorSpecimenCollection{value: val, isSet: true}
}

func (v NullableTumorSpecimenCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTumorSpecimenCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
