use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbProcedureRecommendationCategoryCoding {
    #[serde(rename = "code")]
    pub code: Code,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system", skip_serializing_if = "Option::is_none")]
    pub system: Option<String>,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl MtbProcedureRecommendationCategoryCoding {
    pub fn new(code: Code) -> MtbProcedureRecommendationCategoryCoding {
        MtbProcedureRecommendationCategoryCoding {
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
    #[serde(rename = "ST")]
    St,
    #[serde(rename = "WS")]
    Ws,
    #[serde(rename = "WW")]
    Ww,
    #[serde(rename = "AS")]
    As,
    #[serde(rename = "OP")]
    Op,
    #[serde(rename = "SO")]
    So,
}

impl Default for Code {
    fn default() -> Code {
        Self::St
    }
}
