use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct RecistCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl RecistCoding {
    pub fn new(code: Code) -> RecistCoding {
        RecistCoding {
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
    #[serde(rename = "SD")]
    Sd,
    #[serde(rename = "PR")]
    Pr,
    #[serde(rename = "CR")]
    Cr,
    #[serde(rename = "NA")]
    Na,
    #[serde(rename = "MR")]
    Mr,
    #[serde(rename = "PD")]
    Pd,
}

impl Default for Code {
    fn default() -> Code {
        Self::Sd
    }
}
