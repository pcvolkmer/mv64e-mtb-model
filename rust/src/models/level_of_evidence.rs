use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct LevelOfEvidence {
    #[serde(rename = "grading")]
    pub grading: models::LevelOfEvidenceGradingCoding,
    #[serde(rename = "addendums", skip_serializing_if = "Option::is_none")]
    pub addendums: Option<Vec<models::LevelOfEvidenceAddendumCoding>>,
    #[serde(rename = "publications", skip_serializing_if = "Option::is_none")]
    pub publications: Option<Vec<models::PublicationReference>>,
}

impl LevelOfEvidence {
    pub fn new(grading: models::LevelOfEvidenceGradingCoding) -> LevelOfEvidence {
        LevelOfEvidence {
            grading,
            addendums: None,
            publications: None,
        }
    }
}
