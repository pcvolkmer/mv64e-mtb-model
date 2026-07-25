use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct CarePlanNoSequencingPerformedReasonCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl CarePlanNoSequencingPerformedReasonCoding {
    pub fn new(code: Code) -> CarePlanNoSequencingPerformedReasonCoding {
        CarePlanNoSequencingPerformedReasonCoding {
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
    #[serde(rename = "not-rare-disease")]
    NotRareDisease,
    #[serde(rename = "targeted-diagnostics-recommended")]
    TargetedDiagnosticsRecommended,
    #[serde(rename = "psychosomatic")]
    Psychosomatic,
    #[serde(rename = "non-genetic-cause")]
    NonGeneticCause,
    #[serde(rename = "other")]
    Other,
}

impl Default for Code {
    fn default() -> Code {
        Self::NotRareDisease
    }
}
