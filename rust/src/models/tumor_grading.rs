use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorGrading {
    #[serde(rename = "date")]
    pub date: chrono::NaiveDate,
    #[serde(rename = "codes")]
    pub codes: Vec<models::Coding>,
}

impl TumorGrading {
    pub fn new(date: chrono::NaiveDate, codes: Vec<models::Coding>) -> TumorGrading {
        TumorGrading { date, codes }
    }
}
