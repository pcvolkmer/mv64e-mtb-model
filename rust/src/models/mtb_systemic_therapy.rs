use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbSystemicTherapy {
    #[serde(rename = "statusReason", skip_serializing_if = "Option::is_none")]
    pub status_reason: Option<models::MtbTherapyStatusReasonCoding>,
    #[serde(rename = "medication", skip_serializing_if = "Option::is_none")]
    pub medication: Option<Vec<models::AtcUnregisteredMedicationCoding>>,
    #[serde(rename = "intent", skip_serializing_if = "Option::is_none")]
    pub intent: Option<models::MtbTherapyIntentCoding>,
    #[serde(rename = "therapyLine", skip_serializing_if = "Option::is_none")]
    pub therapy_line: Option<i32>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(
        rename = "recommendationFulfillmentStatus",
        skip_serializing_if = "Option::is_none"
    )]
    pub recommendation_fulfillment_status:
        Option<models::MtbSystemicTherapyRecommendationFulfillmentStatusCoding>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "recordedOn")]
    pub recorded_on: chrono::NaiveDate,
    #[serde(rename = "status")]
    pub status: models::TherapyStatusCoding,
    #[serde(rename = "notes", skip_serializing_if = "Option::is_none")]
    pub notes: Option<Vec<String>>,
    #[serde(rename = "dosage", skip_serializing_if = "Option::is_none")]
    pub dosage: Option<models::MtbSystemicTherapyDosageDensityCoding>,
    #[serde(rename = "period", skip_serializing_if = "Option::is_none")]
    pub period: Option<models::PeriodDate>,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "category", skip_serializing_if = "Option::is_none")]
    pub category: Option<models::MtbSystemicTherapyCategoryCoding>,
    #[serde(rename = "basedOn", skip_serializing_if = "Option::is_none")]
    pub based_on: Option<models::Reference>,
}

impl MtbSystemicTherapy {
    pub fn new(
        id: String,
        recorded_on: chrono::NaiveDate,
        status: models::TherapyStatusCoding,
        patient: models::Reference,
    ) -> MtbSystemicTherapy {
        MtbSystemicTherapy {
            status_reason: None,
            medication: None,
            intent: None,
            therapy_line: None,
            reason: None,
            recommendation_fulfillment_status: None,
            id,
            recorded_on,
            status,
            notes: None,
            dosage: None,
            period: None,
            patient,
            category: None,
            based_on: None,
        }
    }
}
