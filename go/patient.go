package mv64e_model

import (
	"encoding/json"
	"fmt"
)

// checks if the Patient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Patient{}

// Patient struct for Patient
type Patient struct {
	DateOfDeath          *string                `json:"dateOfDeath,omitempty"`
	Age                  *Age                   `json:"age,omitempty"`
	HealthInsurance      PatientHealthInsurance `json:"healthInsurance"`
	Address              *PatientAddress        `json:"address,omitempty"`
	VitalStatus          *VitalStatusCoding     `json:"vitalStatus,omitempty"`
	BirthDate            YearMonth              `json:"birthDate"`
	Id                   string                 `json:"id"`
	Gender               GenderCoding           `json:"gender"`
	ManagingSite         *Coding                `json:"managingSite,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Patient Patient

// NewPatient instantiates a new Patient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatient(healthInsurance PatientHealthInsurance, birthDate string, id string, gender GenderCoding) *Patient {
	this := Patient{}
	this.HealthInsurance = healthInsurance
	this.BirthDate = YearMonth{birthDate}
	this.Id = id
	this.Gender = gender
	return &this
}

// NewPatientWithDefaults instantiates a new Patient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientWithDefaults() *Patient {
	this := Patient{}
	return &this
}

// GetDateOfDeath returns the DateOfDeath field value if set, zero value otherwise.
func (o *Patient) GetDateOfDeath() string {
	if o == nil || IsNil(o.DateOfDeath) {
		var ret string
		return ret
	}
	return *o.DateOfDeath
}

// GetDateOfDeathOk returns a tuple with the DateOfDeath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Patient) GetDateOfDeathOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfDeath) {
		return nil, false
	}
	return o.DateOfDeath, true
}

// HasDateOfDeath returns a boolean if a field has been set.
func (o *Patient) HasDateOfDeath() bool {
	if o != nil && !IsNil(o.DateOfDeath) {
		return true
	}

	return false
}

// SetDateOfDeath gets a reference to the given string and assigns it to the DateOfDeath field.
func (o *Patient) SetDateOfDeath(v string) {
	o.DateOfDeath = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *Patient) GetAge() Age {
	if o == nil || IsNil(o.Age) {
		var ret Age
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Patient) GetAgeOk() (*Age, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *Patient) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given Age and assigns it to the Age field.
func (o *Patient) SetAge(v Age) {
	o.Age = &v
}

// GetHealthInsurance returns the HealthInsurance field value
func (o *Patient) GetHealthInsurance() PatientHealthInsurance {
	if o == nil {
		var ret PatientHealthInsurance
		return ret
	}

	return o.HealthInsurance
}

// GetHealthInsuranceOk returns a tuple with the HealthInsurance field value
// and a boolean to check if the value has been set.
func (o *Patient) GetHealthInsuranceOk() (*PatientHealthInsurance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthInsurance, true
}

// SetHealthInsurance sets field value
func (o *Patient) SetHealthInsurance(v PatientHealthInsurance) {
	o.HealthInsurance = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Patient) GetAddress() PatientAddress {
	if o == nil || IsNil(o.Address) {
		var ret PatientAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Patient) GetAddressOk() (*PatientAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Patient) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PatientAddress and assigns it to the Address field.
func (o *Patient) SetAddress(v PatientAddress) {
	o.Address = &v
}

// GetVitalStatus returns the VitalStatus field value if set, zero value otherwise.
func (o *Patient) GetVitalStatus() VitalStatusCoding {
	if o == nil || IsNil(o.VitalStatus) {
		var ret VitalStatusCoding
		return ret
	}
	return *o.VitalStatus
}

// GetVitalStatusOk returns a tuple with the VitalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Patient) GetVitalStatusOk() (*VitalStatusCoding, bool) {
	if o == nil || IsNil(o.VitalStatus) {
		return nil, false
	}
	return o.VitalStatus, true
}

// HasVitalStatus returns a boolean if a field has been set.
func (o *Patient) HasVitalStatus() bool {
	if o != nil && !IsNil(o.VitalStatus) {
		return true
	}

	return false
}

// SetVitalStatus gets a reference to the given VitalStatusCoding and assigns it to the VitalStatus field.
func (o *Patient) SetVitalStatus(v VitalStatusCoding) {
	o.VitalStatus = &v
}

// GetBirthDate returns the BirthDate field value
func (o *Patient) GetBirthDate() YearMonth {
	if o == nil {
		var ret YearMonth
		return ret
	}

	return o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value
// and a boolean to check if the value has been set.
func (o *Patient) GetBirthDateOk() (*YearMonth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BirthDate, true
}

// SetBirthDate sets field value
func (o *Patient) SetBirthDate(v YearMonth) {
	o.BirthDate = v
}

// GetId returns the Id field value
func (o *Patient) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Patient) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Patient) SetId(v string) {
	o.Id = v
}

// GetGender returns the Gender field value
func (o *Patient) GetGender() GenderCoding {
	if o == nil {
		var ret GenderCoding
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *Patient) GetGenderOk() (*GenderCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *Patient) SetGender(v GenderCoding) {
	o.Gender = v
}

// GetManagingSite returns the ManagingSite field value if set, zero value otherwise.
func (o *Patient) GetManagingSite() Coding {
	if o == nil || IsNil(o.ManagingSite) {
		var ret Coding
		return ret
	}
	return *o.ManagingSite
}

// GetManagingSiteOk returns a tuple with the ManagingSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Patient) GetManagingSiteOk() (*Coding, bool) {
	if o == nil || IsNil(o.ManagingSite) {
		return nil, false
	}
	return o.ManagingSite, true
}

// HasManagingSite returns a boolean if a field has been set.
func (o *Patient) HasManagingSite() bool {
	if o != nil && !IsNil(o.ManagingSite) {
		return true
	}

	return false
}

// SetManagingSite gets a reference to the given Coding and assigns it to the ManagingSite field.
func (o *Patient) SetManagingSite(v Coding) {
	o.ManagingSite = &v
}

func (o Patient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Patient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateOfDeath) {
		toSerialize["dateOfDeath"] = o.DateOfDeath
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	toSerialize["healthInsurance"] = o.HealthInsurance
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.VitalStatus) {
		toSerialize["vitalStatus"] = o.VitalStatus
	}
	toSerialize["birthDate"] = o.BirthDate
	toSerialize["id"] = o.Id
	toSerialize["gender"] = o.Gender
	if !IsNil(o.ManagingSite) {
		toSerialize["managingSite"] = o.ManagingSite
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Patient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"healthInsurance",
		"birthDate",
		"id",
		"gender",
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

	varPatient := _Patient{}

	err = json.Unmarshal(data, &varPatient)

	if err != nil {
		return err
	}

	*o = Patient(varPatient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dateOfDeath")
		delete(additionalProperties, "age")
		delete(additionalProperties, "healthInsurance")
		delete(additionalProperties, "address")
		delete(additionalProperties, "vitalStatus")
		delete(additionalProperties, "birthDate")
		delete(additionalProperties, "id")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "managingSite")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatient struct {
	value *Patient
	isSet bool
}

func (v NullablePatient) Get() *Patient {
	return v.value
}

func (v *NullablePatient) Set(val *Patient) {
	v.value = val
	v.isSet = true
}

func (v NullablePatient) IsSet() bool {
	return v.isSet
}

func (v *NullablePatient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatient(val *Patient) *NullablePatient {
	return &NullablePatient{value: val, isSet: true}
}

func (v NullablePatient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
