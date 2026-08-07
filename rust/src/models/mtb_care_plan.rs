use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct MtbCarePlan {
    #[serde(rename = "boardType", skip_serializing_if = "Option::is_none")]
    pub board_type: Option<models::CarePlanBoardTypeCoding>,
    #[serde(
        rename = "recommendationsMissingReason",
        skip_serializing_if = "Option::is_none"
    )]
    pub recommendations_missing_reason:
        Option<models::MtbCarePlanRecommendationsMissingReasonCoding>,
    #[serde(
        rename = "studyEnrollmentRecommendations",
        skip_serializing_if = "Option::is_none"
    )]
    pub study_enrollment_recommendations: Option<Vec<models::MtbStudyEnrollmentRecommendation>>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<models::Reference>,
    #[serde(
        rename = "geneticCounselingRecommendation",
        skip_serializing_if = "Option::is_none"
    )]
    pub genetic_counseling_recommendation: Option<models::GeneticCounselingRecommendation>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(
        rename = "medicationRecommendations",
        skip_serializing_if = "Option::is_none"
    )]
    pub medication_recommendations: Option<Vec<models::MtbMedicationRecommendation>>,
    #[serde(rename = "rebiopsyRequests", skip_serializing_if = "Option::is_none")]
    pub rebiopsy_requests: Option<Vec<models::MtbCarePlanRebiopsyRequestsInner>>,
    #[serde(
        rename = "histologyReevaluationRequests",
        skip_serializing_if = "Option::is_none"
    )]
    pub histology_reevaluation_requests:
        Option<Vec<models::MtbCarePlanHistologyReevaluationRequestsInner>>,
    #[serde(rename = "notes", skip_serializing_if = "Option::is_none")]
    pub notes: Option<Vec<String>>,
    #[serde(
        rename = "procedureRecommendations",
        skip_serializing_if = "Option::is_none"
    )]
    pub procedure_recommendations: Option<Vec<models::MtbCarePlanProcedureRecommendationsInner>>,
    #[serde(rename = "issuedOn")]
    pub issued_on: chrono::NaiveDate,
    #[serde(
        rename = "noSequencingPerformedReason",
        skip_serializing_if = "Option::is_none"
    )]
    pub no_sequencing_performed_reason: Option<models::CarePlanNoSequencingPerformedReasonCoding>,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
}

impl MtbCarePlan {
    pub fn new(
        id: String,
        issued_on: chrono::NaiveDate,
        patient: models::Reference,
    ) -> MtbCarePlan {
        MtbCarePlan {
            board_type: None,
            recommendations_missing_reason: None,
            study_enrollment_recommendations: None,
            reason: None,
            genetic_counseling_recommendation: None,
            id,
            medication_recommendations: None,
            rebiopsy_requests: None,
            histology_reevaluation_requests: None,
            notes: None,
            procedure_recommendations: None,
            issued_on,
            no_sequencing_performed_reason: None,
            patient,
        }
    }
}
