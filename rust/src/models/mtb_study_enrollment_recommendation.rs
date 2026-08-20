use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbStudyEnrollmentRecommendation {
    #[serde(rename = "priority")]
    pub priority: models::MtbRecommendationPriorityCoding,
    #[serde(rename = "medication", skip_serializing_if = "Option::is_none")]
    pub medication: Option<Vec<models::AtcUnregisteredMedicationCoding>>,
    #[serde(rename = "supportingVariants", skip_serializing_if = "Option::is_none")]
    pub supporting_variants: Option<Vec<models::GeneAlterationReference>>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "supportingFindings", skip_serializing_if = "Option::is_none")]
    pub supporting_findings: Option<Vec<models::Reference>>,
    #[serde(rename = "levelOfEvidence", skip_serializing_if = "Option::is_none")]
    pub level_of_evidence: Option<models::LevelOfEvidence>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(rename = "study")]
    pub study: Vec<models::StudyReference>,
}

impl MtbStudyEnrollmentRecommendation {
    pub fn new(
        priority: models::MtbRecommendationPriorityCoding,
        patient: models::Reference,
        id: String,
        issued_on: chrono::NaiveDate,
        study: Vec<models::StudyReference>,
    ) -> MtbStudyEnrollmentRecommendation {
        MtbStudyEnrollmentRecommendation {
            priority,
            medication: None,
            supporting_variants: None,
            reason: None,
            patient,
            supporting_findings: None,
            level_of_evidence: None,
            id,
            issued_on,
            study,
        }
    }
}
