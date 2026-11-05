package mv64e_model

import (
	_ "embed"
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

//go:embed tests/mv64e-mtb-fake-patient.json
var fakeMtbData []byte

func TestShouldDeserializeJson(t *testing.T) {
	patientRecord := PatientRecord{}
	err := patientRecord.UnmarshalJSON(fakeMtbData)

	if err != nil {
		t.Errorf("Cannot deserialize MTB file")
	}
}

func TestShouldKeepPatientBirthdateFormatInYearMonth(t *testing.T) {
	re := regexp.MustCompile("\"birthDate\":\"\\d{4}-\\d{2}(-\\d{2})?\"")

	alteredMtbData := re.ReplaceAll(fakeMtbData, []byte("\"birthDate\":\"2025-03\""))

	patientRecord := PatientRecord{}
	_ = patientRecord.UnmarshalJSON(alteredMtbData)

	if patientRecord.Patient.BirthDate.Value != "\"2025-03\"" {
		t.Errorf("Serde issues?")
	}

	actual, err := json.Marshal(patientRecord)

	if err != nil || !strings.Contains(string(actual), `"birthDate":"2025-03"`) {
		t.Errorf("Date does not match. Timezone issues?")
	}
}

func TestShouldConvertPatientBirthdateFormatInYearMonth(t *testing.T) {
	re := regexp.MustCompile("\"birthDate\":\"\\d{4}-\\d{2}(-\\d{2})?\"")

	alteredMtbData := re.ReplaceAll(fakeMtbData, []byte("\"birthDate\":\"2025-03-19\""))

	patientRecord := PatientRecord{}
	_ = patientRecord.UnmarshalJSON(alteredMtbData)

	if patientRecord.Patient.BirthDate.Value != "\"2025-03-19\"" {
		t.Errorf("Serde issues?")
	}

	actual, err := json.Marshal(patientRecord)

	if err != nil || !strings.Contains(string(actual), `"birthDate":"2025-03"`) {
		t.Errorf("Date does not match. Timezone issues?")
	}
}
