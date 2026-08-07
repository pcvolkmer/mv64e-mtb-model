use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PeriodDate {
    #[serde(rename = "start")]
    pub start: chrono::NaiveDate,
    #[serde(rename = "end", skip_serializing_if = "Option::is_none")]
    pub end: Option<chrono::NaiveDate>,
}

impl PeriodDate {
    pub fn new(start: chrono::NaiveDate) -> PeriodDate {
        PeriodDate { start, end: None }
    }
}
