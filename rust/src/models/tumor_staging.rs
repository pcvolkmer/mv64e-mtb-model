use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorStaging {
    #[serde(rename = "date")]
    pub date: chrono::NaiveDate,
    #[serde(rename = "method")]
    pub method: models::TumorStagingMethodCoding,
    #[serde(rename = "tnmClassification", skip_serializing_if = "Option::is_none")]
    pub tnm_classification: Option<models::TumorStagingTnmClassification>,
    #[serde(
        rename = "otherClassifications",
        skip_serializing_if = "Option::is_none"
    )]
    pub other_classifications: Option<Vec<models::Coding>>,
}

impl TumorStaging {
    pub fn new(date: chrono::NaiveDate, method: models::TumorStagingMethodCoding) -> TumorStaging {
        TumorStaging {
            date,
            method,
            tnm_classification: None,
            other_classifications: None,
        }
    }
}
