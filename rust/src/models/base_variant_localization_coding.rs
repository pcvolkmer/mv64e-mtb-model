use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct BaseVariantLocalizationCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl BaseVariantLocalizationCoding {
    pub fn new(code: Code) -> BaseVariantLocalizationCoding {
        BaseVariantLocalizationCoding {
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
    #[serde(rename = "intronic")]
    Intronic,
    #[serde(rename = "regulatory-region")]
    RegulatoryRegion,
    #[serde(rename = "coding-region")]
    CodingRegion,
    #[serde(rename = "intergenic")]
    Intergenic,
    #[serde(rename = "splicing-region")]
    SplicingRegion,
}

impl Default for Code {
    fn default() -> Code {
        Self::Intronic
    }
}
