use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct AtcUnregisteredMedicationCoding {
    #[serde(rename = "code")]
    pub code: String,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "system")]
    pub system: System,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl AtcUnregisteredMedicationCoding {
    pub fn new(code: String, system: System) -> AtcUnregisteredMedicationCoding {
        AtcUnregisteredMedicationCoding {
            code,
            display: None,
            system,
            version: None,
        }
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum System {
    #[serde(rename = "undefined")]
    Undefined,
    #[serde(rename = "http://fhir.de/CodeSystem/bfarm/atc")]
    HttpColonSlashSlashFhirDeSlashCodeSystemSlashBfarmSlashAtc,
}

impl Default for System {
    fn default() -> System {
        Self::Undefined
    }
}
