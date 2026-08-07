use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosisTypeHistoryInner {
    #[serde(rename = "value")]
    pub value: models::MtbDiagnosisTypeCoding,
    #[serde(rename = "date")]
    pub date: chrono::NaiveDate,
}

impl MtbDiagnosisTypeHistoryInner {
    pub fn new(
        value: models::MtbDiagnosisTypeCoding,
        date: chrono::NaiveDate,
    ) -> MtbDiagnosisTypeHistoryInner {
        MtbDiagnosisTypeHistoryInner { value, date }
    }
}
