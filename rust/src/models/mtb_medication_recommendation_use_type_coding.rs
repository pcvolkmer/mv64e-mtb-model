use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbMedicationRecommendationUseTypeCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl MtbMedicationRecommendationUseTypeCoding {
    pub fn new(code: Code) -> MtbMedicationRecommendationUseTypeCoding {
        MtbMedicationRecommendationUseTypeCoding {
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
    #[serde(rename = "unknown")]
    Unknown,
    #[serde(rename = "sec-preventive")]
    SecPreventive,
    #[serde(rename = "off-label")]
    OffLabel,
    #[serde(rename = "in-label")]
    InLabel,
    #[serde(rename = "compassionate")]
    Compassionate,
}

impl Default for Code {
    fn default() -> Code {
        Self::Unknown
    }
}
