use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct ClaimResponseStatusReasonCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl ClaimResponseStatusReasonCoding {
    pub fn new(code: Code) -> ClaimResponseStatusReasonCoding {
        ClaimResponseStatusReasonCoding {
            code,
            display: None,
            system: None,
            version: None,
        }
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum Code {
    #[serde(rename = "inclusion-in-study")]
    InclusionInStudy,
    #[serde(rename = "unknown")]
    Unknown,
    #[serde(rename = "insufficient-evidence")]
    InsufficientEvidence,
    #[serde(rename = "formal-reasons")]
    FormalReasons,
    #[serde(rename = "approval-revocation")]
    ApprovalRevocation,
    #[serde(rename = "standard-therapy-not-exhausted")]
    StandardTherapyNotExhausted,
    #[serde(rename = "other-therapy-recommended")]
    OtherTherapyRecommended,
    #[serde(rename = "other")]
    Other,
}

impl Default for Code {
    fn default() -> Code {
        Self::InclusionInStudy
    }
}
