use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct BrcanessConfidenceRange {
    #[serde(rename = "min")]
    pub min: f64,
    #[serde(rename = "max")]
    pub max: f64,
}

impl BrcanessConfidenceRange {
    pub fn new(min: f64, max: f64) -> BrcanessConfidenceRange {
        BrcanessConfidenceRange { min, max }
    }
}
