use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MolecularDiagnosticReport {
    #[serde(rename = "type")]
    pub r#type: models::MolecularDiagnosticReportTypeCoding,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(rename = "performer", skip_serializing_if = "Option::is_none")]
    pub performer: Option<models::Reference>,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "results", skip_serializing_if = "Option::is_none")]
    pub results: Option<Vec<String>>,
}

impl MolecularDiagnosticReport {
    pub fn new(
        r#type: models::MolecularDiagnosticReportTypeCoding,
        specimen: models::Reference,
        issued_on: chrono::NaiveDate,
        patient: models::Reference,
        id: String,
    ) -> MolecularDiagnosticReport {
        MolecularDiagnosticReport {
            r#type,
            specimen,
            issued_on,
            performer: None,
            patient,
            id,
            results: None,
        }
    }
}
