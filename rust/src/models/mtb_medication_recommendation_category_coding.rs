use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbMedicationRecommendationCategoryCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl MtbMedicationRecommendationCategoryCoding {
    pub fn new(code: Code) -> MtbMedicationRecommendationCategoryCoding {
        MtbMedicationRecommendationCategoryCoding {
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
    #[serde(rename = "SO")]
    So,
    #[serde(rename = "SZ")]
    Sz,
    #[serde(rename = "IM")]
    Im,
    #[serde(rename = "CH")]
    Ch,
    #[serde(rename = "HO")]
    Ho,
    #[serde(rename = "ZS")]
    Zs,
}

impl Default for Code {
    fn default() -> Code {
        Self::So
    }
}
