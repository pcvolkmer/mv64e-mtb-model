use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct EcogCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl EcogCoding {
    pub fn new(code: Code) -> EcogCoding {
        EcogCoding {
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
    #[serde(rename = "1")]
    Variant1,
    #[serde(rename = "4")]
    Variant4,
    #[serde(rename = "0")]
    Variant0,
    #[serde(rename = "3")]
    Variant3,
    #[serde(rename = "2")]
    Variant2,
    #[serde(rename = "5")]
    Variant5,
}

impl Default for Code {
    fn default() -> Code {
        Self::Variant1
    }
}
