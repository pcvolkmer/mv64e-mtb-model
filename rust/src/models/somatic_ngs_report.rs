use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct SomaticNgsReport {
    #[serde(rename = "metadata")]
    pub metadata: Vec<models::NgsReportMetadata>,
    #[serde(rename = "type")]
    pub r#type: models::NgsReportTypeCoding,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "results")]
    pub results: models::SomaticNgsReportResults,
}

impl SomaticNgsReport {
    pub fn new(
        metadata: Vec<models::NgsReportMetadata>,
        r#type: models::NgsReportTypeCoding,
        specimen: models::Reference,
        issued_on: chrono::NaiveDate,
        patient: models::Reference,
        id: String,
        results: models::SomaticNgsReportResults,
    ) -> SomaticNgsReport {
        SomaticNgsReport {
            metadata,
            r#type,
            specimen,
            issued_on,
            patient,
            id,
            results,
        }
    }
}
