use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct IhcReport {
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "results")]
    pub results: models::IhcReportResults,
}

impl IhcReport {
    pub fn new(
        specimen: models::Reference,
        issued_on: chrono::NaiveDate,
        patient: models::Reference,
        id: String,
        results: models::IhcReportResults,
    ) -> IhcReport {
        IhcReport {
            specimen,
            issued_on,
            patient,
            id,
            results,
        }
    }
}
