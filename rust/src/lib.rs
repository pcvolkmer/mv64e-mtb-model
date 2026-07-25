//! This crate provides structs to serialize and deserialize DNPM:DIP MTB DTOs.
//! The base struct is `PatientRecord`.

#![allow(clippy::needless_doctest_main)]

use crate::models::{
    GenderCoding, HealthInsuranceTypeCoding, Patient, PatientHealthInsurance, PatientRecord,
    gender_coding, health_insurance_type_coding,
};
use std::error::Error;
use std::fmt::{Debug, Display, Formatter};
use std::str::FromStr;

pub mod models;

mod year_month_format {
    use regex::Regex;
    use serde::{Serialize, Serializer};
    use std::str::FromStr;

    pub(crate) fn serialize_year_month_format<S>(
        date: &String,
        serializer: S,
    ) -> Result<S::Ok, S::Error>
    where
        S: Serializer,
    {
        let new_format = Regex::from_str(r"^\d{4}-\d{2}$").expect("new format regex");
        let old_format = Regex::from_str(r"^\d{4}-\d{2}-\d{2}$").expect("old format regex");

        if new_format.is_match(date) {
            date.serialize(serializer)
        } else if old_format.is_match(date) {
            date.as_str()[0..7].to_string().serialize(serializer)
        } else {
            Err(serde::ser::Error::custom("Must be a valid date format"))
        }
    }

    pub(crate) fn serialize_option_year_month_format<S>(
        date: &Option<String>,
        serializer: S,
    ) -> Result<S::Ok, S::Error>
    where
        S: Serializer,
    {
        if let Some(s) = date {
            let new_format = Regex::from_str(r"^\d{4}-\d{2}$").unwrap();
            let old_format = Regex::from_str(r"^\d{4}-\d{2}-\d{2}$").unwrap();

            return if new_format.is_match(s) {
                Option::<String>::serialize(&Some(s.into()), serializer)
            } else if old_format.is_match(s) {
                Option::<String>::serialize(&Some(s.as_str()[0..7].into()), serializer)
            } else {
                Err(serde::ser::Error::custom("Must be a valid date format"))
            };
        }
        Option::<String>::serialize(&None, serializer)
    }
}

#[derive(Debug)]
pub struct SerdeError(String);

impl Display for SerdeError {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "MtbFile Serde Error: {}", self.0)
    }
}

impl Error for SerdeError {}

impl FromStr for PatientRecord {
    type Err = SerdeError;

    /// Deserialize an instance of `MtbFile` from a string of JSON text.
    ///
    /// # Example
    ///
    /// ```
    /// use mv64e_mtb_model::models::patient_record::PatientRecord;
    /// use std::str::FromStr;
    ///
    /// fn main() {
    ///     const JSON: &str = include_str!("../tests/mv64e-mtb-fake-patient.json");
    ///
    ///     let mtb_file = PatientRecord::from_str(JSON).unwrap();
    ///     println!("{:#?}", mtb_file);
    /// }
    /// ```
    ///
    /// # Errors
    ///
    /// If the conversion fails, an `SerdeError` will be returned.
    fn from_str(value: &str) -> Result<PatientRecord, SerdeError> {
        serde_json::from_str(value).map_err(|err| SerdeError(err.to_string()))
    }
}

impl PatientRecord {
    /// Creates "dummy" PatientRecord with consent status `REJECTED`.
    /// The created PatientRecord does not contain all information, just enough to contain the
    /// information, that the patient with given ID has rejected the consent.
    pub fn new_with_consent_rejected(patient_id: &str) -> PatientRecord {
        PatientRecord {
            care_plans: None,
            claim_responses: None,
            claims: None,
            diagnoses: None,
            performance_status: None,
            episodes_of_care: None,
            histology_reports: None,
            guideline_therapies: None,
            ngs_reports: None,
            ihc_reports: None,
            patient: Some(Patient {
                address: None,
                age: None,
                birth_date: "1900-01".to_string(),
                date_of_death: None,
                gender: GenderCoding {
                    code: gender_coding::Code::Female,
                    display: None,
                    system: None,
                    version: None,
                },
                id: patient_id.to_string(),
                health_insurance: PatientHealthInsurance {
                    r#type: HealthInsuranceTypeCoding {
                        code: health_insurance_type_coding::Code::Unk,
                        display: None,
                        system: None,
                        version: None,
                    },
                    reference: None,
                },
                vital_status: None,
                managing_site: None,
            }),
            responses: None,
            specimens: None,
            guideline_procedures: None,
            family_member_histories: None,
            follow_ups: None,
            prior_diagnostic_reports: None,
            systemic_therapies: None,
            metadata: None,
            msi_findings: None,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use regex::Regex;

    const MTB_JSON: &str = include_str!("../tests/mv64e-mtb-fake-patient.json");

    #[test]
    fn should_deserialize_json_string() {
        let mtbfile = PatientRecord::from_str(MTB_JSON);
        assert!(mtbfile.is_ok())
    }

    #[test]
    fn should_keep_patient_birthdate_format_in_year_month() {
        let m = Regex::from_str(r#""birthDate":\s?"\d{4}-\d{2}(-\d{2})?""#)
            .unwrap()
            .replace(MTB_JSON, r#""birthDate": "2025-03""#);

        let mtbfile = PatientRecord::from_str(&m).unwrap();
        let actual = serde_json::to_string(&mtbfile).unwrap();
        assert!(actual.contains(r#""birthDate":"2025-03""#));
    }

    #[test]
    fn should_convert_patient_birthdate_format_in_year_month() {
        let m = Regex::from_str(r#""birthDate":\s?"\d{4}-\d{2}(-\d{2})?""#)
            .unwrap()
            .replace(MTB_JSON, r#""birthDate": "2025-03-19""#);

        let mtbfile = PatientRecord::from_str(&m).unwrap();
        let actual = serde_json::to_string(&mtbfile).unwrap();
        assert!(actual.contains(r#""birthDate":"2025-03""#));
    }
}
