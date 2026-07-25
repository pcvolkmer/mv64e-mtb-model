use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TumorSpecimen {
    #[serde(rename = "diagnosis")]
    pub diagnosis: models::Reference,
    #[serde(rename = "collection", skip_serializing_if = "Option::is_none")]
    pub collection: Option<models::TumorSpecimenCollection>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "type")]
    pub r#type: models::TumorSpecimenTypeCoding,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl TumorSpecimen {
    pub fn new(
        diagnosis: models::Reference,
        id: String,
        r#type: models::TumorSpecimenTypeCoding,
        patient: models::Reference,
    ) -> TumorSpecimen {
        TumorSpecimen {
            diagnosis,
            collection: None,
            id,
            r#type,
            patient,
        }
    }
}
