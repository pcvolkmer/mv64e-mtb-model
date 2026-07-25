use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct HistologyReportResultsTumorMorphology {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "value")]
    pub value: models::Coding,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl HistologyReportResultsTumorMorphology {
    pub fn new(
        id: String,
        specimen: models::Reference,
        value: models::Coding,
        patient: models::Reference,
    ) -> HistologyReportResultsTumorMorphology {
        HistologyReportResultsTumorMorphology {
            id,
            note: None,
            specimen,
            value,
            patient,
        }
    }
}
