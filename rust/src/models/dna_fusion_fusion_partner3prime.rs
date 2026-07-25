use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct DnaFusionFusionPartner3prime {
    #[serde(rename = "chromosome")]
    pub chromosome: models::Chromosome,
    #[serde(rename = "gene")]
    pub gene: models::Coding,
    #[serde(rename = "position")]
    pub position: f64,
}

impl DnaFusionFusionPartner3prime {
    pub fn new(
        chromosome: models::Chromosome,
        gene: models::Coding,
        position: f64,
    ) -> DnaFusionFusionPartner3prime {
        DnaFusionFusionPartner3prime {
            chromosome,
            gene,
            position,
        }
    }
}
