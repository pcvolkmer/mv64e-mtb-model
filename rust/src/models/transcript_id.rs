use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct TranscriptId {
    #[serde(rename = "value")]
    pub value: String,
    #[serde(rename = "system")]
    pub system: System,
}

impl TranscriptId {
    pub fn new(value: String, system: System) -> TranscriptId {
        TranscriptId { value, system }
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub enum System {
    #[serde(rename = "https://www.ncbi.nlm.nih.gov/refseq")]
    HttpsColonSlashSlashWwwNcbiNlmNihGovSlashRefseq,
    #[serde(rename = "https://www.ensembl.org")]
    HttpsColonSlashSlashWwwEnsemblOrg,
}

impl Default for System {
    fn default() -> System {
        Self::HttpsColonSlashSlashWwwNcbiNlmNihGovSlashRefseq
    }
}
