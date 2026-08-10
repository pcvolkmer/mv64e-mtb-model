use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct PatientRecord {
    #[serde(rename = "diagnoses")]
    pub diagnoses: Vec<models::MtbDiagnosis>,
    #[serde(rename = "ngsReports", skip_serializing_if = "Option::is_none")]
    pub ngs_reports: Option<Vec<models::SomaticNgsReport>>,
    #[serde(rename = "performanceStatus", skip_serializing_if = "Option::is_none")]
    pub performance_status: Option<Vec<models::PerformanceStatus>>,
    #[serde(
        rename = "guidelineProcedures",
        skip_serializing_if = "Option::is_none"
    )]
    pub guideline_procedures: Option<Vec<models::OncoProcedure>>,
    #[serde(
        rename = "familyMemberHistories",
        skip_serializing_if = "Option::is_none"
    )]
    pub family_member_histories: Option<Vec<models::FamilyMemberHistory>>,
    #[serde(rename = "claimResponses", skip_serializing_if = "Option::is_none")]
    pub claim_responses: Option<Vec<models::ClaimResponse>>,
    #[serde(rename = "msiFindings", skip_serializing_if = "Option::is_none")]
    pub msi_findings: Option<Vec<models::Msi>>,
    #[serde(rename = "carePlans", skip_serializing_if = "Option::is_none")]
    pub care_plans: Option<Vec<models::MtbCarePlan>>,
    #[serde(rename = "specimens", skip_serializing_if = "Option::is_none")]
    pub specimens: Option<Vec<models::TumorSpecimen>>,
    #[serde(rename = "ihcReports", skip_serializing_if = "Option::is_none")]
    pub ihc_reports: Option<Vec<models::IhcReport>>,
    #[serde(rename = "histologyReports", skip_serializing_if = "Option::is_none")]
    pub histology_reports: Option<Vec<models::HistologyReport>>,
    #[serde(rename = "guidelineTherapies", skip_serializing_if = "Option::is_none")]
    pub guideline_therapies: Option<Vec<models::MtbSystemicTherapy>>,
    #[serde(rename = "episodesOfCare")]
    pub episodes_of_care: Vec<models::MtbEpisodeOfCare>,
    #[serde(
        rename = "priorDiagnosticReports",
        skip_serializing_if = "Option::is_none"
    )]
    pub prior_diagnostic_reports: Option<Vec<models::MolecularDiagnosticReport>>,
    #[serde(rename = "systemicTherapies", skip_serializing_if = "Option::is_none")]
    pub systemic_therapies: Option<Vec<models::PatientRecordSystemicTherapiesInner>>,
    #[serde(rename = "followUps", skip_serializing_if = "Option::is_none")]
    pub follow_ups: Option<Vec<models::FollowUp>>,
    #[serde(rename = "claims", skip_serializing_if = "Option::is_none")]
    pub claims: Option<Vec<models::Claim>>,
    #[serde(rename = "metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<models::MvhMetadata>,
    #[serde(rename = "responses", skip_serializing_if = "Option::is_none")]
    pub responses: Option<Vec<models::Response>>,
    #[serde(rename = "patient")]
    pub patient: models::Patient,
}

impl PatientRecord {
    pub fn new(
        diagnoses: Vec<models::MtbDiagnosis>,
        episodes_of_care: Vec<models::MtbEpisodeOfCare>,
        patient: models::Patient,
    ) -> PatientRecord {
        PatientRecord {
            diagnoses,
            ngs_reports: None,
            performance_status: None,
            guideline_procedures: None,
            family_member_histories: None,
            claim_responses: None,
            msi_findings: None,
            care_plans: None,
            specimens: None,
            ihc_reports: None,
            histology_reports: None,
            guideline_therapies: None,
            episodes_of_care,
            prior_diagnostic_reports: None,
            systemic_therapies: None,
            follow_ups: None,
            claims: None,
            metadata: None,
            responses: None,
            patient,
        }
    }
}
