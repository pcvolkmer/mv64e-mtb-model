use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct HrdScoreComponents {
    #[serde(rename = "lst")]
    pub lst: f64,
    #[serde(rename = "loh")]
    pub loh: f64,
    #[serde(rename = "tai")]
    pub tai: f64,
}

impl HrdScoreComponents {
    pub fn new(lst: f64, loh: f64, tai: f64) -> HrdScoreComponents {
        HrdScoreComponents { lst, loh, tai }
    }
}
