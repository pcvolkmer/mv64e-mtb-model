use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorCellContent {
    #[serde(rename = "method")]
    pub method: models::TumorCellContentMethodCoding,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl TumorCellContent {
    pub fn new(
        method: models::TumorCellContentMethodCoding,
        id: String,
        specimen: models::Reference,
        value: f64,
        patient: models::Reference,
    ) -> TumorCellContent {
        TumorCellContent {
            method,
            id,
            specimen,
            value,
            patient,
        }
    }
}
