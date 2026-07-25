use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct CnvEndRange {
    #[serde(rename = "start")]
    pub start: f64,
    #[serde(rename = "end", skip_serializing_if = "Option::is_none")]
    pub end: Option<f64>,
}

impl CnvEndRange {
    pub fn new(start: f64) -> CnvEndRange {
        CnvEndRange { start, end: None }
    }
}
