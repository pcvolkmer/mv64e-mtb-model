use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosisStaging {
    #[serde(rename = "history")]
    pub history: Vec<models::TumorStaging>,
}

impl MtbDiagnosisStaging {
    pub fn new(history: Vec<models::TumorStaging>) -> MtbDiagnosisStaging {
        MtbDiagnosisStaging { history }
    }
}
