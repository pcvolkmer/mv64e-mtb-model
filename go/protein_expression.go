package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the ProteinExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProteinExpression{}

// ProteinExpression struct for ProteinExpression
type ProteinExpression struct {
	CpsScore             *int32                          `json:"cpsScore,omitempty"`
	IcScore              *ProteinExpressionIcScoreCoding `json:"icScore,omitempty"`
	Value                ProteinExpressionResultCoding   `json:"value"`
	Patient              Reference                       `json:"patient"`
	Protein              Coding                          `json:"protein"`
	TcScore              *ProteinExpressionTcScoreCoding `json:"tcScore,omitempty"`
	TpsScore             *int32                          `json:"tpsScore,omitempty"`
	Id                   string                          `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ProteinExpression ProteinExpression

// NewProteinExpression instantiates a new ProteinExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProteinExpression(value ProteinExpressionResultCoding, patient Reference, protein Coding, id string) *ProteinExpression {
	this := ProteinExpression{}
	this.Value = value
	this.Patient = patient
	this.Protein = protein
	this.Id = id
	return &this
}

// NewProteinExpressionWithDefaults instantiates a new ProteinExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProteinExpressionWithDefaults() *ProteinExpression {
	this := ProteinExpression{}
	return &this
}

// GetCpsScore returns the CpsScore field value if set, zero value otherwise.
func (o *ProteinExpression) GetCpsScore() int32 {
	if o == nil || IsNil(o.CpsScore) {
		var ret int32
		return ret
	}
	return *o.CpsScore
}

// GetCpsScoreOk returns a tuple with the CpsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetCpsScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CpsScore) {
		return nil, false
	}
	return o.CpsScore, true
}

// HasCpsScore returns a boolean if a field has been set.
func (o *ProteinExpression) HasCpsScore() bool {
	if o != nil && !IsNil(o.CpsScore) {
		return true
	}

	return false
}

// SetCpsScore gets a reference to the given int32 and assigns it to the CpsScore field.
func (o *ProteinExpression) SetCpsScore(v int32) {
	o.CpsScore = &v
}

// GetIcScore returns the IcScore field value if set, zero value otherwise.
func (o *ProteinExpression) GetIcScore() ProteinExpressionIcScoreCoding {
	if o == nil || IsNil(o.IcScore) {
		var ret ProteinExpressionIcScoreCoding
		return ret
	}
	return *o.IcScore
}

// GetIcScoreOk returns a tuple with the IcScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetIcScoreOk() (*ProteinExpressionIcScoreCoding, bool) {
	if o == nil || IsNil(o.IcScore) {
		return nil, false
	}
	return o.IcScore, true
}

// HasIcScore returns a boolean if a field has been set.
func (o *ProteinExpression) HasIcScore() bool {
	if o != nil && !IsNil(o.IcScore) {
		return true
	}

	return false
}

// SetIcScore gets a reference to the given ProteinExpressionIcScoreCoding and assigns it to the IcScore field.
func (o *ProteinExpression) SetIcScore(v ProteinExpressionIcScoreCoding) {
	o.IcScore = &v
}

// GetValue returns the Value field value
func (o *ProteinExpression) GetValue() ProteinExpressionResultCoding {
	if o == nil {
		var ret ProteinExpressionResultCoding
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetValueOk() (*ProteinExpressionResultCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ProteinExpression) SetValue(v ProteinExpressionResultCoding) {
	o.Value = v
}

// GetPatient returns the Patient field value
func (o *ProteinExpression) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *ProteinExpression) SetPatient(v Reference) {
	o.Patient = v
}

// GetProtein returns the Protein field value
func (o *ProteinExpression) GetProtein() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Protein
}

// GetProteinOk returns a tuple with the Protein field value
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetProteinOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protein, true
}

// SetProtein sets field value
func (o *ProteinExpression) SetProtein(v Coding) {
	o.Protein = v
}

// GetTcScore returns the TcScore field value if set, zero value otherwise.
func (o *ProteinExpression) GetTcScore() ProteinExpressionTcScoreCoding {
	if o == nil || IsNil(o.TcScore) {
		var ret ProteinExpressionTcScoreCoding
		return ret
	}
	return *o.TcScore
}

// GetTcScoreOk returns a tuple with the TcScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetTcScoreOk() (*ProteinExpressionTcScoreCoding, bool) {
	if o == nil || IsNil(o.TcScore) {
		return nil, false
	}
	return o.TcScore, true
}

// HasTcScore returns a boolean if a field has been set.
func (o *ProteinExpression) HasTcScore() bool {
	if o != nil && !IsNil(o.TcScore) {
		return true
	}

	return false
}

// SetTcScore gets a reference to the given ProteinExpressionTcScoreCoding and assigns it to the TcScore field.
func (o *ProteinExpression) SetTcScore(v ProteinExpressionTcScoreCoding) {
	o.TcScore = &v
}

// GetTpsScore returns the TpsScore field value if set, zero value otherwise.
func (o *ProteinExpression) GetTpsScore() int32 {
	if o == nil || IsNil(o.TpsScore) {
		var ret int32
		return ret
	}
	return *o.TpsScore
}

// GetTpsScoreOk returns a tuple with the TpsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetTpsScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.TpsScore) {
		return nil, false
	}
	return o.TpsScore, true
}

// HasTpsScore returns a boolean if a field has been set.
func (o *ProteinExpression) HasTpsScore() bool {
	if o != nil && !IsNil(o.TpsScore) {
		return true
	}

	return false
}

// SetTpsScore gets a reference to the given int32 and assigns it to the TpsScore field.
func (o *ProteinExpression) SetTpsScore(v int32) {
	o.TpsScore = &v
}

// GetId returns the Id field value
func (o *ProteinExpression) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProteinExpression) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProteinExpression) SetId(v string) {
	o.Id = v
}

func (o ProteinExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProteinExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpsScore) {
		toSerialize["cpsScore"] = o.CpsScore
	}
	if !IsNil(o.IcScore) {
		toSerialize["icScore"] = o.IcScore
	}
	toSerialize["value"] = o.Value
	toSerialize["patient"] = o.Patient
	toSerialize["protein"] = o.Protein
	if !IsNil(o.TcScore) {
		toSerialize["tcScore"] = o.TcScore
	}
	if !IsNil(o.TpsScore) {
		toSerialize["tpsScore"] = o.TpsScore
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProteinExpression) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"patient",
		"protein",
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

	varProteinExpression := _ProteinExpression{}

	err = json.Unmarshal(data, &varProteinExpression)

	if err != nil {
		return err
	}

	*o = ProteinExpression(varProteinExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cpsScore")
		delete(additionalProperties, "icScore")
		delete(additionalProperties, "value")
		delete(additionalProperties, "patient")
		delete(additionalProperties, "protein")
		delete(additionalProperties, "tcScore")
		delete(additionalProperties, "tpsScore")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProteinExpression struct {
	value *ProteinExpression
	isSet bool
}

func (v NullableProteinExpression) Get() *ProteinExpression {
	return v.value
}

func (v *NullableProteinExpression) Set(val *ProteinExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableProteinExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableProteinExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProteinExpression(val *ProteinExpression) *NullableProteinExpression {
	return &NullableProteinExpression{value: val, isSet: true}
}

func (v NullableProteinExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProteinExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
