use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PatientRecordSystemicTherapiesInner {
    #[serde(rename = "history")]
    pub history: Vec<models::MtbSystemicTherapy>,
}

impl PatientRecordSystemicTherapiesInner {
    pub fn new(history: Vec<models::MtbSystemicTherapy>) -> PatientRecordSystemicTherapiesInner {
        PatientRecordSystemicTherapiesInner { history }
    }
}
