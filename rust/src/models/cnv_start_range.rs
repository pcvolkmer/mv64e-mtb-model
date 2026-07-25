use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct CnvStartRange {
    #[serde(rename = "start")]
    pub start: f64,
    #[serde(rename = "end", skip_serializing_if = "Option::is_none")]
    pub end: Option<f64>,
}

impl CnvStartRange {
    pub fn new(start: f64) -> CnvStartRange {
        CnvStartRange { start, end: None }
    }
}
