use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MvhMetadataModelProjectConsent {
    #[serde(rename = "version")]
    pub version: String,
    #[serde(rename = "date", skip_serializing_if = "Option::is_none")]
    pub date: Option<String>,
    #[serde(rename = "provisions")]
    pub provisions: Vec<models::MvhMetadataModelProjectConsentProvisionsInner>,
}

impl MvhMetadataModelProjectConsent {
    pub fn new(
        version: String,
        provisions: Vec<models::MvhMetadataModelProjectConsentProvisionsInner>,
    ) -> MvhMetadataModelProjectConsent {
        MvhMetadataModelProjectConsent {
            version,
            date: None,
            provisions,
        }
    }
}
