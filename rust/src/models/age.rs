use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Age {
    #[serde(rename = "value")]
    pub value: f64,
    #[serde(rename = "unit")]
    pub unit: Unit,
}

impl Age {
    pub fn new(value: f64, unit: Unit) -> Age {
        Age { value, unit }
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum Unit {
    #[serde(rename = "Months")]
    Months,
    #[serde(rename = "Years")]
    Years,
}

impl Default for Unit {
    fn default() -> Unit {
        Self::Months
    }
}
