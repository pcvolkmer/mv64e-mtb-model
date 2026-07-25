use crate::models;
use serde::{Deserialize, Serialize};

///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum ModelProjectConsentPurpose {
    #[serde(rename = "case-identification")]
    CaseIdentification,
    #[serde(rename = "reidentification")]
    Reidentification,
    #[serde(rename = "sequencing")]
    Sequencing,
}

impl std::fmt::Display for ModelProjectConsentPurpose {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::CaseIdentification => write!(f, "case-identification"),
            Self::Reidentification => write!(f, "reidentification"),
            Self::Sequencing => write!(f, "sequencing"),
        }
    }
}

impl Default for ModelProjectConsentPurpose {
    fn default() -> ModelProjectConsentPurpose {
        Self::CaseIdentification
    }
}
