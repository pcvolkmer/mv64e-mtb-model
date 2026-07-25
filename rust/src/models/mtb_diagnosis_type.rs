use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosisType {
    #[serde(rename = "history")]
    pub history: Vec<models::MtbDiagnosisTypeHistoryInner>,
}

impl MtbDiagnosisType {
    pub fn new(history: Vec<models::MtbDiagnosisTypeHistoryInner>) -> MtbDiagnosisType {
        MtbDiagnosisType { history }
    }
}
