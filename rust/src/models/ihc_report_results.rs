use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct IhcReportResults {
    #[serde(rename = "proteinExpression")]
    pub protein_expression: Vec<models::ProteinExpression>,
    #[serde(rename = "msiMmr")]
    pub msi_mmr: Vec<models::ProteinExpression>,
}

impl IhcReportResults {
    pub fn new(
        protein_expression: Vec<models::ProteinExpression>,
        msi_mmr: Vec<models::ProteinExpression>,
    ) -> IhcReportResults {
        IhcReportResults {
            protein_expression,
            msi_mmr,
        }
    }
}
