use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PerformanceStatus {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "effectiveDate")]
    pub effective_date: chrono::NaiveDate,
    #[serde(rename = "value")]
    pub value: models::EcogCoding,
}

impl PerformanceStatus {
    pub fn new(
        id: String,
        patient: models::Reference,
        effective_date: chrono::NaiveDate,
        value: models::EcogCoding,
    ) -> PerformanceStatus {
        PerformanceStatus {
            id,
            patient,
            effective_date,
            value,
        }
    }
}
