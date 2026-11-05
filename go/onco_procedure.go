package mv64e_model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OncoProcedure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OncoProcedure{}

// OncoProcedure struct for OncoProcedure
type OncoProcedure struct {
	StatusReason *MtbTherapyStatusReasonCoding `json:"statusReason,omitempty"`
	Intent       *MtbTherapyIntentCoding       `json:"intent,omitempty"`
	TherapyLine  *int32                        `json:"therapyLine,omitempty"`
	Reason       *Reference                    `json:"reason,omitempty"`
	Code         OncoProcedureTypeCoding       `json:"code"`
	Id           string                        `json:"id"`
	RecordedOn   string                        `json:"recordedOn"`
	Status       TherapyStatusCoding           `json:"status"`
	BasedOn      *Reference                    `json:"basedOn,omitempty"`
	Notes        []string                      `json:"notes,omitempty"`
	Period       *PeriodDate                   `json:"period,omitempty"`
	Patient      Reference                     `json:"patient"`
}

type _OncoProcedure OncoProcedure

// NewOncoProcedure instantiates a new OncoProcedure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOncoProcedure(code OncoProcedureTypeCoding, id string, recordedOn string, status TherapyStatusCoding, patient Reference) *OncoProcedure {
	this := OncoProcedure{}
	this.Code = code
	this.Id = id
	this.RecordedOn = recordedOn
	this.Status = status
	this.Patient = patient
	return &this
}

// NewOncoProcedureWithDefaults instantiates a new OncoProcedure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOncoProcedureWithDefaults() *OncoProcedure {
	this := OncoProcedure{}
	return &this
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *OncoProcedure) GetStatusReason() MtbTherapyStatusReasonCoding {
	if o == nil || IsNil(o.StatusReason) {
		var ret MtbTherapyStatusReasonCoding
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetStatusReasonOk() (*MtbTherapyStatusReasonCoding, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *OncoProcedure) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given MtbTherapyStatusReasonCoding and assigns it to the StatusReason field.
func (o *OncoProcedure) SetStatusReason(v MtbTherapyStatusReasonCoding) {
	o.StatusReason = &v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *OncoProcedure) GetIntent() MtbTherapyIntentCoding {
	if o == nil || IsNil(o.Intent) {
		var ret MtbTherapyIntentCoding
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetIntentOk() (*MtbTherapyIntentCoding, bool) {
	if o == nil || IsNil(o.Intent) {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *OncoProcedure) HasIntent() bool {
	if o != nil && !IsNil(o.Intent) {
		return true
	}

	return false
}

// SetIntent gets a reference to the given MtbTherapyIntentCoding and assigns it to the Intent field.
func (o *OncoProcedure) SetIntent(v MtbTherapyIntentCoding) {
	o.Intent = &v
}

// GetTherapyLine returns the TherapyLine field value if set, zero value otherwise.
func (o *OncoProcedure) GetTherapyLine() int32 {
	if o == nil || IsNil(o.TherapyLine) {
		var ret int32
		return ret
	}
	return *o.TherapyLine
}

// GetTherapyLineOk returns a tuple with the TherapyLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetTherapyLineOk() (*int32, bool) {
	if o == nil || IsNil(o.TherapyLine) {
		return nil, false
	}
	return o.TherapyLine, true
}

// HasTherapyLine returns a boolean if a field has been set.
func (o *OncoProcedure) HasTherapyLine() bool {
	if o != nil && !IsNil(o.TherapyLine) {
		return true
	}

	return false
}

// SetTherapyLine gets a reference to the given int32 and assigns it to the TherapyLine field.
func (o *OncoProcedure) SetTherapyLine(v int32) {
	o.TherapyLine = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OncoProcedure) GetReason() Reference {
	if o == nil || IsNil(o.Reason) {
		var ret Reference
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetReasonOk() (*Reference, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OncoProcedure) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reference and assigns it to the Reason field.
func (o *OncoProcedure) SetReason(v Reference) {
	o.Reason = &v
}

// GetCode returns the Code field value
func (o *OncoProcedure) GetCode() OncoProcedureTypeCoding {
	if o == nil {
		var ret OncoProcedureTypeCoding
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetCodeOk() (*OncoProcedureTypeCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OncoProcedure) SetCode(v OncoProcedureTypeCoding) {
	o.Code = v
}

// GetId returns the Id field value
func (o *OncoProcedure) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OncoProcedure) SetId(v string) {
	o.Id = v
}

// GetRecordedOn returns the RecordedOn field value
func (o *OncoProcedure) GetRecordedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordedOn
}

// GetRecordedOnOk returns a tuple with the RecordedOn field value
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetRecordedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedOn, true
}

// SetRecordedOn sets field value
func (o *OncoProcedure) SetRecordedOn(v string) {
	o.RecordedOn = v
}

// GetStatus returns the Status field value
func (o *OncoProcedure) GetStatus() TherapyStatusCoding {
	if o == nil {
		var ret TherapyStatusCoding
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetStatusOk() (*TherapyStatusCoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OncoProcedure) SetStatus(v TherapyStatusCoding) {
	o.Status = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *OncoProcedure) GetBasedOn() Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret Reference
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetBasedOnOk() (*Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *OncoProcedure) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given Reference and assigns it to the BasedOn field.
func (o *OncoProcedure) SetBasedOn(v Reference) {
	o.BasedOn = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OncoProcedure) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OncoProcedure) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *OncoProcedure) SetNotes(v []string) {
	o.Notes = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *OncoProcedure) GetPeriod() PeriodDate {
	if o == nil || IsNil(o.Period) {
		var ret PeriodDate
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetPeriodOk() (*PeriodDate, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *OncoProcedure) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given PeriodDate and assigns it to the Period field.
func (o *OncoProcedure) SetPeriod(v PeriodDate) {
	o.Period = &v
}

// GetPatient returns the Patient field value
func (o *OncoProcedure) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *OncoProcedure) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *OncoProcedure) SetPatient(v Reference) {
	o.Patient = v
}

func (o OncoProcedure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OncoProcedure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	if !IsNil(o.Intent) {
		toSerialize["intent"] = o.Intent
	}
	if !IsNil(o.TherapyLine) {
		toSerialize["therapyLine"] = o.TherapyLine
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["code"] = o.Code
	toSerialize["id"] = o.Id
	toSerialize["recordedOn"] = o.RecordedOn
	toSerialize["status"] = o.Status
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["patient"] = o.Patient
	return toSerialize, nil
}

func (o *OncoProcedure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"id",
		"recordedOn",
		"status",
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

	varOncoProcedure := _OncoProcedure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOncoProcedure)

	if err != nil {
		return err
	}

	*o = OncoProcedure(varOncoProcedure)

	return err
}

type NullableOncoProcedure struct {
	value *OncoProcedure
	isSet bool
}

func (v NullableOncoProcedure) Get() *OncoProcedure {
	return v.value
}

func (v *NullableOncoProcedure) Set(val *OncoProcedure) {
	v.value = val
	v.isSet = true
}

func (v NullableOncoProcedure) IsSet() bool {
	return v.isSet
}

func (v *NullableOncoProcedure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOncoProcedure(val *OncoProcedure) *NullableOncoProcedure {
	return &NullableOncoProcedure{value: val, isSet: true}
}

func (v NullableOncoProcedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOncoProcedure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
