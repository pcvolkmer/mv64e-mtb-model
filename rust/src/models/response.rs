use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Response {
    #[serde(rename = "method")]
    pub method: models::ResponseMethodCoding,
    #[serde(rename = "therapy")]
    pub therapy: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "effectiveDate")]
    pub effective_date: String,
    #[serde(rename = "value")]
    pub value: models::RecistCoding,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl Response {
    pub fn new(
        method: models::ResponseMethodCoding,
        therapy: models::Reference,
        id: String,
        effective_date: String,
        value: models::RecistCoding,
        patient: models::Reference,
    ) -> Response {
        Response {
            method,
            therapy,
            id,
            effective_date,
            value,
            patient,
        }
    }
}
