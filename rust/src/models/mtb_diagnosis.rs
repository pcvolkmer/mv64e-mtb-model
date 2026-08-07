use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosis {
    #[serde(rename = "grading", skip_serializing_if = "Option::is_none")]
    pub grading: Option<models::MtbDiagnosisGrading>,
    #[serde(rename = "germlineCodes", skip_serializing_if = "Option::is_none")]
    pub germline_codes: Option<Vec<models::Coding>>,
    #[serde(rename = "code")]
    pub code: models::Coding,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "recordedOn")]
    pub recorded_on: chrono::NaiveDate,
    #[serde(rename = "type")]
    pub r#type: models::MtbDiagnosisType,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "topography")]
    pub topography: models::Coding,
    #[serde(rename = "staging", skip_serializing_if = "Option::is_none")]
    pub staging: Option<models::MtbDiagnosisStaging>,
    #[serde(rename = "histology", skip_serializing_if = "Option::is_none")]
    pub histology: Option<Vec<models::Reference>>,
    #[serde(
        rename = "guidelineTreatmentStatus",
        skip_serializing_if = "Option::is_none"
    )]
    pub guideline_treatment_status: Option<models::MtbDiagnosisGuidelineTreatmentStatusCoding>,
    #[serde(rename = "notes", skip_serializing_if = "Option::is_none")]
    pub notes: Option<Vec<String>>,
}

impl MtbDiagnosis {
    pub fn new(
        code: models::Coding,
        id: String,
        recorded_on: chrono::NaiveDate,
        r#type: models::MtbDiagnosisType,
        patient: models::Reference,
        topography: models::Coding,
    ) -> MtbDiagnosis {
        MtbDiagnosis {
            grading: None,
            germline_codes: None,
            code,
            id,
            recorded_on,
            r#type,
            patient,
            topography,
            staging: None,
            histology: None,
            guideline_treatment_status: None,
            notes: None,
        }
    }
}
