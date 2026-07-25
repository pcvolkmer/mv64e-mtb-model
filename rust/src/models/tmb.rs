use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Tmb {
    #[serde(rename = "interpretation", skip_serializing_if = "Option::is_none")]
    pub interpretation: Option<models::TmbInterpretationCoding>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: models::TmbResult,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl Tmb {
    pub fn new(
        id: String,
        specimen: models::Reference,
        value: models::TmbResult,
        patient: models::Reference,
    ) -> Tmb {
        Tmb {
            interpretation: None,
            id,
            specimen,
            value,
            patient,
        }
    }
}
