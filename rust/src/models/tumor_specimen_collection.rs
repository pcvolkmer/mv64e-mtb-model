use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorSpecimenCollection {
    #[serde(rename = "date", skip_serializing_if = "Option::is_none")]
    pub date: Option<String>,
    #[serde(rename = "method")]
    pub method: models::TumorSpecimenCollectionMethodCoding,
    #[serde(rename = "localization")]
    pub localization: models::TumorSpecimenCollectionLocalizationCoding,
}

impl TumorSpecimenCollection {
    pub fn new(
        method: models::TumorSpecimenCollectionMethodCoding,
        localization: models::TumorSpecimenCollectionLocalizationCoding,
    ) -> TumorSpecimenCollection {
        TumorSpecimenCollection {
            date: None,
            method,
            localization,
        }
    }
}
