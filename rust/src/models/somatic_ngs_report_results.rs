use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct SomaticNgsReportResults {
    #[serde(rename = "rnaSeqs", skip_serializing_if = "Option::is_none")]
    pub rna_seqs: Option<Vec<models::RnaSeq>>,
    #[serde(rename = "tmb", skip_serializing_if = "Option::is_none")]
    pub tmb: Option<models::Tmb>,
    #[serde(rename = "simpleVariants", skip_serializing_if = "Option::is_none")]
    pub simple_variants: Option<Vec<models::Snv>>,
    #[serde(rename = "hrdScore", skip_serializing_if = "Option::is_none")]
    pub hrd_score: Option<models::HrdScore>,
    #[serde(rename = "copyNumberVariants", skip_serializing_if = "Option::is_none")]
    pub copy_number_variants: Option<Vec<models::Cnv>>,
    #[serde(rename = "tumorCellContent", skip_serializing_if = "Option::is_none")]
    pub tumor_cell_content: Option<models::TumorCellContent>,
    #[serde(rename = "dnaFusions", skip_serializing_if = "Option::is_none")]
    pub dna_fusions: Option<Vec<models::DnaFusion>>,
    #[serde(rename = "rnaFusions", skip_serializing_if = "Option::is_none")]
    pub rna_fusions: Option<Vec<models::RnaFusion>>,
    #[serde(rename = "brcaness", skip_serializing_if = "Option::is_none")]
    pub brcaness: Option<models::Brcaness>,
}

impl SomaticNgsReportResults {
    pub fn new() -> SomaticNgsReportResults {
        SomaticNgsReportResults {
            rna_seqs: None,
            tmb: None,
            simple_variants: None,
            hrd_score: None,
            copy_number_variants: None,
            tumor_cell_content: None,
            dna_fusions: None,
            rna_fusions: None,
            brcaness: None,
        }
    }
}
