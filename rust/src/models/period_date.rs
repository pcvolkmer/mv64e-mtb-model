use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PeriodDate {
    #[serde(rename = "start")]
    pub start: String,
    #[serde(rename = "end", skip_serializing_if = "Option::is_none")]
    pub end: Option<String>,
}

impl PeriodDate {
    pub fn new(start: String) -> PeriodDate {
        PeriodDate { start, end: None }
    }
}
