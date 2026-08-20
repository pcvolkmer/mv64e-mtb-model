use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbRecommendationPriorityCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl MtbRecommendationPriorityCoding {
    pub fn new(code: Code) -> MtbRecommendationPriorityCoding {
        MtbRecommendationPriorityCoding {
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
    #[serde(rename = "6")]
    Variant6,
    #[serde(rename = "8")]
    Variant8,
    #[serde(rename = "4")]
    Variant4,
    #[serde(rename = "12")]
    Variant12,
    #[serde(rename = "11")]
    Variant11,
    #[serde(rename = "9")]
    Variant9,
    #[serde(rename = "7")]
    Variant7,
    #[serde(rename = "3")]
    Variant3,
    #[serde(rename = "10")]
    Variant10,
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
