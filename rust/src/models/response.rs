use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Response {
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<models::ResponseMethodCoding>,
    #[serde(rename = "therapy")]
    pub therapy: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "effectiveDate")]
    pub effective_date: chrono::NaiveDate,
    #[serde(rename = "value")]
    pub value: models::RecistCoding,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl Response {
    pub fn new(
        therapy: models::Reference,
        id: String,
        effective_date: chrono::NaiveDate,
        value: models::RecistCoding,
        patient: models::Reference,
    ) -> Response {
        Response {
            method: None,
            therapy,
            id,
            effective_date,
            value,
            patient,
        }
    }
}
