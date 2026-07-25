use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Msi {
    #[serde(rename = "method")]
    pub method: models::MsiMethodCoding,
    #[serde(rename = "interpretation")]
    pub interpretation: models::MsiInterpretationCoding,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl Msi {
    pub fn new(
        method: models::MsiMethodCoding,
        interpretation: models::MsiInterpretationCoding,
        id: String,
        specimen: models::Reference,
        value: f64,
        patient: models::Reference,
    ) -> Msi {
        Msi {
            method,
            interpretation,
            id,
            specimen,
            value,
            patient,
        }
    }
}
