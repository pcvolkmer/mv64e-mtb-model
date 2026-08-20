use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbMedicationRecommendation {
    #[serde(rename = "priority")]
    pub priority: models::MtbRecommendationPriorityCoding,
    #[serde(rename = "medication")]
    pub medication: Vec<models::AtcUnregisteredMedicationCoding>,
    #[serde(rename = "supportingVariants", skip_serializing_if = "Option::is_none")]
    pub supporting_variants: Option<Vec<models::GeneAlterationReference>>,
    #[serde(rename = "useType", skip_serializing_if = "Option::is_none")]
    pub use_type: Option<models::MtbMedicationRecommendationUseTypeCoding>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(rename = "category", skip_serializing_if = "Option::is_none")]
    pub category: Option<Vec<models::MtbMedicationRecommendationCategoryCoding>>,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "supportingFindings", skip_serializing_if = "Option::is_none")]
    pub supporting_findings: Option<Vec<models::Reference>>,
    #[serde(rename = "levelOfEvidence", skip_serializing_if = "Option::is_none")]
    pub level_of_evidence: Option<models::LevelOfEvidence>,
    #[serde(rename = "id")]
    pub id: String,
}

impl MtbMedicationRecommendation {
    pub fn new(
        priority: models::MtbRecommendationPriorityCoding,
        medication: Vec<models::AtcUnregisteredMedicationCoding>,
        issued_on: chrono::NaiveDate,
        patient: models::Reference,
        id: String,
    ) -> MtbMedicationRecommendation {
        MtbMedicationRecommendation {
            priority,
            medication,
            supporting_variants: None,
            use_type: None,
            reason: None,
            category: None,
            issued_on,
            patient,
            supporting_findings: None,
            level_of_evidence: None,
            id,
        }
    }
}
