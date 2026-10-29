use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PatientRecordStudyInclusionsInner {
    #[serde(rename = "history")]
    pub history: Vec<models::StudyInclusion>,
}

impl PatientRecordStudyInclusionsInner {
    pub fn new(history: Vec<models::StudyInclusion>) -> PatientRecordStudyInclusionsInner {
        PatientRecordStudyInclusionsInner { history }
    }
}
