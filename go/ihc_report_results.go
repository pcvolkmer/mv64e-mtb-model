package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IhcReportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IhcReportResults{}

// IhcReportResults struct for IhcReportResults
type IhcReportResults struct {
	ProteinExpression []ProteinExpression `json:"proteinExpression"`
	MsiMmr            []ProteinExpression `json:"msiMmr"`
}

type _IhcReportResults IhcReportResults

// NewIhcReportResults instantiates a new IhcReportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIhcReportResults(proteinExpression []ProteinExpression, msiMmr []ProteinExpression) *IhcReportResults {
	this := IhcReportResults{}
	this.ProteinExpression = proteinExpression
	this.MsiMmr = msiMmr
	return &this
}

// NewIhcReportResultsWithDefaults instantiates a new IhcReportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIhcReportResultsWithDefaults() *IhcReportResults {
	this := IhcReportResults{}
	return &this
}

// GetProteinExpression returns the ProteinExpression field value
func (o *IhcReportResults) GetProteinExpression() []ProteinExpression {
	if o == nil {
		var ret []ProteinExpression
		return ret
	}

	return o.ProteinExpression
}

// GetProteinExpressionOk returns a tuple with the ProteinExpression field value
// and a boolean to check if the value has been set.
func (o *IhcReportResults) GetProteinExpressionOk() ([]ProteinExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProteinExpression, true
}

// SetProteinExpression sets field value
func (o *IhcReportResults) SetProteinExpression(v []ProteinExpression) {
	o.ProteinExpression = v
}

// GetMsiMmr returns the MsiMmr field value
func (o *IhcReportResults) GetMsiMmr() []ProteinExpression {
	if o == nil {
		var ret []ProteinExpression
		return ret
	}

	return o.MsiMmr
}

// GetMsiMmrOk returns a tuple with the MsiMmr field value
// and a boolean to check if the value has been set.
func (o *IhcReportResults) GetMsiMmrOk() ([]ProteinExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.MsiMmr, true
}

// SetMsiMmr sets field value
func (o *IhcReportResults) SetMsiMmr(v []ProteinExpression) {
	o.MsiMmr = v
}

func (o IhcReportResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IhcReportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["proteinExpression"] = o.ProteinExpression
	toSerialize["msiMmr"] = o.MsiMmr
	return toSerialize, nil
}

func (o *IhcReportResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"proteinExpression",
		"msiMmr",
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

	varIhcReportResults := _IhcReportResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIhcReportResults)

	if err != nil {
		return err
	}

	*o = IhcReportResults(varIhcReportResults)

	return err
}

type NullableIhcReportResults struct {
	value *IhcReportResults
	isSet bool
}

func (v NullableIhcReportResults) Get() *IhcReportResults {
	return v.value
}

func (v *NullableIhcReportResults) Set(val *IhcReportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableIhcReportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableIhcReportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIhcReportResults(val *IhcReportResults) *NullableIhcReportResults {
	return &NullableIhcReportResults{value: val, isSet: true}
}

func (v NullableIhcReportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIhcReportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
