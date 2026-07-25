use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct NgsReportMetadata {
    #[serde(rename = "kitManufacturer")]
    pub kit_manufacturer: String,
    #[serde(rename = "pipeline")]
    pub pipeline: String,
    #[serde(rename = "kitType")]
    pub kit_type: String,
    #[serde(rename = "sequencer")]
    pub sequencer: String,
    #[serde(rename = "referenceGenome")]
    pub reference_genome: String,
}

impl NgsReportMetadata {
    pub fn new(
        kit_manufacturer: String,
        pipeline: String,
        kit_type: String,
        sequencer: String,
        reference_genome: String,
    ) -> NgsReportMetadata {
        NgsReportMetadata {
            kit_manufacturer,
            pipeline,
            kit_type,
            sequencer,
            reference_genome,
        }
    }
}
