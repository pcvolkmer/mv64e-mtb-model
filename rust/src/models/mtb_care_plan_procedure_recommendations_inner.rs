use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbCarePlanProcedureRecommendationsInner {
    #[serde(rename = "priority")]
    pub priority: models::RecommendationPriorityCoding,
    #[serde(rename = "supportingVariants", skip_serializing_if = "Option::is_none")]
    pub supporting_variants: Option<Vec<models::GeneAlterationReference>>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(rename = "code")]
    pub code: models::MtbProcedureRecommendationCategoryCoding,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "supportingFindings", skip_serializing_if = "Option::is_none")]
    pub supporting_findings: Option<Vec<models::Reference>>,
    #[serde(rename = "levelOfEvidence", skip_serializing_if = "Option::is_none")]
    pub level_of_evidence: Option<models::LevelOfEvidence>,
    #[serde(rename = "id")]
    pub id: String,
}

impl MtbCarePlanProcedureRecommendationsInner {
    pub fn new(
        priority: models::RecommendationPriorityCoding,
        code: models::MtbProcedureRecommendationCategoryCoding,
        issued_on: String,
        patient: models::Reference,
        id: String,
    ) -> MtbCarePlanProcedureRecommendationsInner {
        MtbCarePlanProcedureRecommendationsInner {
            priority,
            supporting_variants: None,
            reason: None,
            code,
            issued_on,
            patient,
            supporting_findings: None,
            level_of_evidence: None,
            id,
        }
    }
}
