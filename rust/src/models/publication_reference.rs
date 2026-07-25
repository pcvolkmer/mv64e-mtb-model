use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PublicationReference {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "system")]
    pub system: System,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "type", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<String>,
}

impl PublicationReference {
    pub fn new(id: String, system: System) -> PublicationReference {
        PublicationReference {
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
    #[serde(rename = "https://www.doi.org")]
    HttpsColonSlashSlashWwwDoiOrg,
    #[serde(rename = "https://pubmed.ncbi.nlm.nih.gov")]
    HttpsColonSlashSlashPubmedNcbiNlmNihGov,
}

impl Default for System {
    fn default() -> System {
        Self::HttpsColonSlashSlashWwwDoiOrg
    }
}
