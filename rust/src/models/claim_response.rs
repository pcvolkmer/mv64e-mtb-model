use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct ClaimResponse {
    #[serde(rename = "statusReason", skip_serializing_if = "Option::is_none")]
    pub status_reason: Option<Vec<models::ClaimResponseStatusReasonCoding>>,
    #[serde(rename = "claim")]
    pub claim: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<models::ClaimResponseStatusCoding>,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl ClaimResponse {
    pub fn new(
        claim: models::Reference,
        id: String,
        issued_on: String,
        patient: models::Reference,
    ) -> ClaimResponse {
        ClaimResponse {
            status_reason: None,
            claim,
            id,
            status: None,
            issued_on,
            patient,
        }
    }
}
