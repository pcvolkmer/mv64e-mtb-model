use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbEpisodeOfCare {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "period")]
    pub period: models::PeriodDate,
    #[serde(rename = "diagnoses", skip_serializing_if = "Option::is_none")]
    pub diagnoses: Option<Vec<models::Reference>>,
}

impl MtbEpisodeOfCare {
    pub fn new(
        id: String,
        patient: models::Reference,
        period: models::PeriodDate,
    ) -> MtbEpisodeOfCare {
        MtbEpisodeOfCare {
            id,
            patient,
            period,
            diagnoses: None,
        }
    }
}
