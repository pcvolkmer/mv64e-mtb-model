use crate::models;
use serde::{Deserialize, Serialize};

///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum ConsentProvisionType {
    #[serde(rename = "deny")]
    Deny,
    #[serde(rename = "permit")]
    Permit,
}

impl std::fmt::Display for ConsentProvisionType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Deny => write!(f, "deny"),
            Self::Permit => write!(f, "permit"),
        }
    }
}

impl Default for ConsentProvisionType {
    fn default() -> ConsentProvisionType {
        Self::Deny
    }
}
