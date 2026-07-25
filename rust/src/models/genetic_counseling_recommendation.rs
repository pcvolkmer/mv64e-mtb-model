use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct GeneticCounselingRecommendation {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
    #[serde(rename = "reason")]
    pub reason: models::GeneticCounselingRecommendationReasonCoding,
}

impl GeneticCounselingRecommendation {
    pub fn new(
        id: String,
        patient: models::Reference,
        issued_on: String,
        reason: models::GeneticCounselingRecommendationReasonCoding,
    ) -> GeneticCounselingRecommendation {
        GeneticCounselingRecommendation {
            id,
            patient,
            issued_on,
            reason,
        }
    }
}
