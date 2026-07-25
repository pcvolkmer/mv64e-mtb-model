use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosisGrading {
    #[serde(rename = "history")]
    pub history: Vec<models::TumorGrading>,
}

impl MtbDiagnosisGrading {
    pub fn new(history: Vec<models::TumorGrading>) -> MtbDiagnosisGrading {
        MtbDiagnosisGrading { history }
    }
}
