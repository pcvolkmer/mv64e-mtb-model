use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbDiagnosisGuidelineTreatmentStatusCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl MtbDiagnosisGuidelineTreatmentStatusCoding {
    pub fn new(code: Code) -> MtbDiagnosisGuidelineTreatmentStatusCoding {
        MtbDiagnosisGuidelineTreatmentStatusCoding {
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
    #[serde(rename = "exhausted")]
    Exhausted,
    #[serde(rename = "non-exhausted")]
    NonExhausted,
    #[serde(rename = "no-guidelines-available")]
    NoGuidelinesAvailable,
    #[serde(rename = "impossible")]
    Impossible,
}

impl Default for Code {
    fn default() -> Code {
        Self::Unknown
    }
}
