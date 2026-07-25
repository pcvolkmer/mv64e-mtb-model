use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TmbResult {
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "unit", skip_serializing_if = "Option::is_none")]
    pub unit: Option<String>,
}

impl TmbResult {
    pub fn new(value: f64) -> TmbResult {
        TmbResult { value, unit: None }
    }
}
