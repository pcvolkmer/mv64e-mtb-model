use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct HrdScore {
    #[serde(rename = "interpretation", skip_serializing_if = "Option::is_none")]
    pub interpretation: Option<models::HrdScoreInterpretationCoding>,
    #[serde(rename = "components")]
    pub components: models::HrdScoreComponents,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl HrdScore {
    pub fn new(
        components: models::HrdScoreComponents,
        id: String,
        specimen: models::Reference,
        value: f64,
        patient: models::Reference,
    ) -> HrdScore {
        HrdScore {
            interpretation: None,
            components,
            id,
            specimen,
            value,
            patient,
        }
    }
}
