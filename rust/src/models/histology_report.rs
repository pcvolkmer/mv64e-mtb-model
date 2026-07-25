use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct HistologyReport {
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "results")]
    pub results: models::HistologyReportResults,
}

impl HistologyReport {
    pub fn new(
        specimen: models::Reference,
        issued_on: String,
        patient: models::Reference,
        id: String,
        results: models::HistologyReportResults,
    ) -> HistologyReport {
        HistologyReport {
            specimen,
            issued_on,
            patient,
            id,
            results,
        }
    }
}
