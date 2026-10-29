use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct StudyInclusion {
    #[serde(rename = "statusReason", skip_serializing_if = "Option::is_none")]
    pub status_reason: Option<models::MtbTherapyStatusReasonCoding>,
    #[serde(rename = "medication", skip_serializing_if = "Option::is_none")]
    pub medication: Option<Vec<models::AtcUnregisteredMedicationCoding>>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "recordedOn")]
    pub recorded_on: chrono::NaiveDate,
    #[serde(rename = "status")]
    pub status: models::TherapyStatusCoding,
    #[serde(rename = "basedOn", skip_serializing_if = "Option::is_none")]
    pub based_on: Option<models::Reference>,
    #[serde(rename = "study")]
    pub study: Vec<models::StudyReference>,
    #[serde(rename = "period", skip_serializing_if = "Option::is_none")]
    pub period: Option<models::PeriodDate>,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl StudyInclusion {
    pub fn new(
        id: String,
        recorded_on: chrono::NaiveDate,
        status: models::TherapyStatusCoding,
        study: Vec<models::StudyReference>,
        patient: models::Reference,
    ) -> StudyInclusion {
        StudyInclusion {
            status_reason: None,
            medication: None,
            reason: None,
            id,
            recorded_on,
            status,
            based_on: None,
            study,
            period: None,
            patient,
        }
    }
}
