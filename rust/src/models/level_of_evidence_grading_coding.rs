use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct LevelOfEvidenceGradingCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl LevelOfEvidenceGradingCoding {
    pub fn new(code: Code) -> LevelOfEvidenceGradingCoding {
        LevelOfEvidenceGradingCoding {
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
    #[serde(rename = "m1A")]
    M1A,
    #[serde(rename = "m2A")]
    M2A,
    #[serde(rename = "m2C")]
    M2C,
    #[serde(rename = "m2B")]
    M2B,
    #[serde(rename = "m1B")]
    M1B,
    #[serde(rename = "undefined")]
    Undefined,
    #[serde(rename = "m4")]
    M4,
    #[serde(rename = "m1C")]
    M1C,
    #[serde(rename = "m3")]
    M3,
}

impl Default for Code {
    fn default() -> Code {
        Self::M1A
    }
}
