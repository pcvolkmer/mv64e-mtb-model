use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct RnaFusionFusionPartner3prime {
    #[serde(rename = "transcriptId")]
    pub transcript_id: models::TranscriptId,
    #[serde(rename = "position")]
    pub position: f64,
    #[serde(rename = "exonId")]
    pub exon_id: String,
    #[serde(rename = "strand")]
    pub strand: models::RnaFusionStrand,
    #[serde(rename = "gene")]
    pub gene: models::Coding,
}

impl RnaFusionFusionPartner3prime {
    pub fn new(
        transcript_id: models::TranscriptId,
        position: f64,
        exon_id: String,
        strand: models::RnaFusionStrand,
        gene: models::Coding,
    ) -> RnaFusionFusionPartner3prime {
        RnaFusionFusionPartner3prime {
            transcript_id,
            position,
            exon_id,
            strand,
            gene,
        }
    }
}
