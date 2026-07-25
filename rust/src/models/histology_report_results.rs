use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct HistologyReportResults {
    #[serde(rename = "tumorMorphology")]
    pub tumor_morphology: models::HistologyReportResultsTumorMorphology,
    #[serde(rename = "tumorCellContent", skip_serializing_if = "Option::is_none")]
    pub tumor_cell_content: Option<models::TumorCellContent>,
}

impl HistologyReportResults {
    pub fn new(
        tumor_morphology: models::HistologyReportResultsTumorMorphology,
    ) -> HistologyReportResults {
        HistologyReportResults {
            tumor_morphology,
            tumor_cell_content: None,
        }
    }
}
