use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Brcaness {
    #[serde(rename = "confidenceRange")]
    pub confidence_range: models::BrcanessConfidenceRange,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl Brcaness {
    pub fn new(
        confidence_range: models::BrcanessConfidenceRange,
        id: String,
        specimen: models::Reference,
        value: f64,
        patient: models::Reference,
    ) -> Brcaness {
        Brcaness {
            confidence_range,
            id,
            specimen,
            value,
            patient,
        }
    }
}
