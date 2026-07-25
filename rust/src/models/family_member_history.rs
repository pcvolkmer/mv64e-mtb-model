use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct FamilyMemberHistory {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "relationship")]
    pub relationship: models::FamilyMemberHistoryRelationshipTypeCoding,
}

impl FamilyMemberHistory {
    pub fn new(
        id: String,
        patient: models::Reference,
        relationship: models::FamilyMemberHistoryRelationshipTypeCoding,
    ) -> FamilyMemberHistory {
        FamilyMemberHistory {
            id,
            patient,
            relationship,
        }
    }
}
