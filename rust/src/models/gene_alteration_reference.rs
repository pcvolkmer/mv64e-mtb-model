use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct GeneAlterationReference {
    #[serde(rename = "variant")]
    pub variant: models::Reference,
    #[serde(rename = "display", skip_serializing_if = "Option::is_none")]
    pub display: Option<String>,
    #[serde(rename = "gene", skip_serializing_if = "Option::is_none")]
    pub gene: Option<models::Coding>,
}

impl GeneAlterationReference {
    pub fn new(variant: models::Reference) -> GeneAlterationReference {
        GeneAlterationReference {
            variant,
            display: None,
            gene: None,
        }
    }
}
