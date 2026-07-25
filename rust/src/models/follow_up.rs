use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct FollowUp {
    #[serde(rename = "date")]
    pub date: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "lastContactDate", skip_serializing_if = "Option::is_none")]
    pub last_contact_date: Option<String>,
    #[serde(rename = "patientStatus", skip_serializing_if = "Option::is_none")]
    pub patient_status: Option<models::FollowUpPatientStatusCoding>,
}

impl FollowUp {
    pub fn new(date: String, patient: models::Reference) -> FollowUp {
        FollowUp {
            date,
            patient,
            last_contact_date: None,
            patient_status: None,
        }
    }
}
