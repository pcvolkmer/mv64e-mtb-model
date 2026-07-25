use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbCarePlanHistologyReevaluationRequestsInner {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "specimen")]
    pub specimen: models::Reference,
    #[serde(rename = "issuedOn")]
    pub issued_on: String,
}

impl MtbCarePlanHistologyReevaluationRequestsInner {
    pub fn new(
        id: String,
        patient: models::Reference,
        specimen: models::Reference,
        issued_on: String,
    ) -> MtbCarePlanHistologyReevaluationRequestsInner {
        MtbCarePlanHistologyReevaluationRequestsInner {
            id,
            patient,
            specimen,
            issued_on,
        }
    }
}
