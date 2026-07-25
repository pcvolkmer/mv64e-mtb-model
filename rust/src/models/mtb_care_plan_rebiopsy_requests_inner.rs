use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbCarePlanRebiopsyRequestsInner {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "tumorEntity")]
    pub tumor_entity: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
}

impl MtbCarePlanRebiopsyRequestsInner {
    pub fn new(
        id: String,
        patient: models::Reference,
        tumor_entity: models::Reference,
        issued_on: String,
    ) -> MtbCarePlanRebiopsyRequestsInner {
        MtbCarePlanRebiopsyRequestsInner {
            id,
            patient,
            tumor_entity,
            issued_on,
        }
    }
}
