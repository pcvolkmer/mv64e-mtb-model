use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MvhMetadataModelProjectConsentProvisionsInner {
    #[serde(rename = "date")]
    pub date: String,
    #[serde(rename = "purpose")]
    pub purpose: models::ModelProjectConsentPurpose,
    #[serde(rename = "type")]
    pub r#type: models::ConsentProvisionType,
}

impl MvhMetadataModelProjectConsentProvisionsInner {
    pub fn new(
        date: String,
        purpose: models::ModelProjectConsentPurpose,
        r#type: models::ConsentProvisionType,
    ) -> MvhMetadataModelProjectConsentProvisionsInner {
        MvhMetadataModelProjectConsentProvisionsInner {
            date,
            purpose,
            r#type,
        }
    }
}
