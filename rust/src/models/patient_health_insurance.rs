use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PatientHealthInsurance {
    #[serde(rename = "type")]
    pub r#type: models::HealthInsuranceTypeCoding,
    #[serde(rename = "reference", skip_serializing_if = "Option::is_none")]
    pub reference: Option<models::Reference>,
}

impl PatientHealthInsurance {
    pub fn new(r#type: models::HealthInsuranceTypeCoding) -> PatientHealthInsurance {
        PatientHealthInsurance {
            r#type,
            reference: None,
        }
    }
}
