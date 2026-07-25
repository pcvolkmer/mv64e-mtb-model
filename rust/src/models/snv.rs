use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Snv {
    #[serde(rename = "localization", skip_serializing_if = "Option::is_none")]
    pub localization: Option<Vec<models::BaseVariantLocalizationCoding>>,
    #[serde(rename = "proteinChange", skip_serializing_if = "Option::is_none")]
    pub protein_change: Option<String>,
    #[serde(rename = "interpretation", skip_serializing_if = "Option::is_none")]
    pub interpretation: Option<models::ClinVarCoding>,
    #[serde(rename = "transcriptId")]
    pub transcript_id: models::TranscriptId,
    #[serde(rename = "externalIds", skip_serializing_if = "Option::is_none")]
    pub external_ids: Option<Vec<models::VariantExternalId>>,
    #[serde(rename = "chromosome")]
    pub chromosome: models::Chromosome,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "altAllele")]
    pub alt_allele: String,
    #[serde(rename = "position")]
    pub position: models::CnvEndRange,
    #[serde(rename = "gene")]
    pub gene: models::Coding,
    #[serde(rename = "exonId", skip_serializing_if = "Option::is_none")]
    pub exon_id: Option<String>,
    #[serde(rename = "dnaChange")]
    pub dna_change: String,
    #[serde(rename = "patient")]
    pub patient: models::Reference,
    #[serde(rename = "refAllele")]
    pub ref_allele: String,
    #[serde(rename = "readDepth")]
    pub read_depth: i32,
    #[serde(rename = "allelicFrequency")]
    pub allelic_frequency: f64,
}

impl Snv {
    pub fn new(
        transcript_id: models::TranscriptId,
        chromosome: models::Chromosome,
        id: String,
        alt_allele: String,
        position: models::CnvEndRange,
        gene: models::Coding,
        dna_change: String,
        patient: models::Reference,
        ref_allele: String,
        read_depth: i32,
        allelic_frequency: f64,
    ) -> Snv {
        Snv {
            localization: None,
            protein_change: None,
            interpretation: None,
            transcript_id,
            external_ids: None,
            chromosome,
            id,
            alt_allele,
            position,
            gene,
            exon_id: None,
            dna_change,
            patient,
            ref_allele,
            read_depth,
            allelic_frequency,
        }
    }
}
