use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PatientAddress {
    #[serde(rename = "municipalityCode")]
    pub municipality_code: String,
}

impl PatientAddress {
    pub fn new(municipality_code: String) -> PatientAddress {
        PatientAddress { municipality_code }
    }
}
