use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorStagingTnmClassification {
    #[serde(rename = "tumor")]
    pub tumor: models::Coding,
    #[serde(rename = "nodes")]
    pub nodes: models::Coding,
    #[serde(rename = "metastasis")]
    pub metastasis: models::Coding,
}

impl TumorStagingTnmClassification {
    pub fn new(
        tumor: models::Coding,
        nodes: models::Coding,
        metastasis: models::Coding,
    ) -> TumorStagingTnmClassification {
        TumorStagingTnmClassification {
            tumor,
            nodes,
            metastasis,
        }
    }
}
