use crate::models;
use serde::{Deserialize, Serialize};

///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum MvhSubmissionType {
    #[serde(rename = "initial")]
    Initial,
    #[serde(rename = "correction")]
    Correction,
    #[serde(rename = "test")]
    Test,
    #[serde(rename = "addition")]
    Addition,
    #[serde(rename = "followup")]
    Followup,
}

impl std::fmt::Display for MvhSubmissionType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Initial => write!(f, "initial"),
            Self::Correction => write!(f, "correction"),
            Self::Test => write!(f, "test"),
            Self::Addition => write!(f, "addition"),
            Self::Followup => write!(f, "followup"),
        }
    }
}

impl Default for MvhSubmissionType {
    fn default() -> MvhSubmissionType {
        Self::Initial
    }
}
