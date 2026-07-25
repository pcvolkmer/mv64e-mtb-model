use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct StudyReference {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "system")]
    pub system: System,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "type", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<String>,
}

impl StudyReference {
    pub fn new(id: String, system: System) -> StudyReference {
        StudyReference {
            id,
            system,
            display: None,
            r#type: None,
        }
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum System {
    #[serde(rename = "NCT")]
    Nct,
    #[serde(rename = "Eudra-CT")]
    EudraCt,
    #[serde(rename = "other")]
    Other,
    #[serde(rename = "EUDAMED")]
    Eudamed,
    #[serde(rename = "DRKS")]
    Drks,
}

impl Default for System {
    fn default() -> System {
        Self::Nct
    }
}
