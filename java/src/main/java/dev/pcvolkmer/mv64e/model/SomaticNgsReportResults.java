package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class SomaticNgsReportResults {
  public static final String JSON_PROPERTY_RNA_SEQS = "rnaSeqs";

  private @Nullable List<RnaSeq> rnaSeqs;

  public static final String JSON_PROPERTY_TMB = "tmb";

  private @Nullable Tmb tmb;

  public static final String JSON_PROPERTY_SIMPLE_VARIANTS = "simpleVariants";

  private @Nullable List<Snv> simpleVariants;

  public static final String JSON_PROPERTY_HRD_SCORE = "hrdScore";

  private @Nullable HrdScore hrdScore;

  public static final String JSON_PROPERTY_COPY_NUMBER_VARIANTS = "copyNumberVariants";

  private @Nullable List<Cnv> copyNumberVariants;

  public static final String JSON_PROPERTY_TUMOR_CELL_CONTENT = "tumorCellContent";

  private @Nullable TumorCellContent tumorCellContent;

  public static final String JSON_PROPERTY_DNA_FUSIONS = "dnaFusions";

  private @Nullable List<DnaFusion> dnaFusions;

  public static final String JSON_PROPERTY_RNA_FUSIONS = "rnaFusions";

  private @Nullable List<RnaFusion> rnaFusions;

  public static final String JSON_PROPERTY_BRCANESS = "brcaness";

  private @Nullable Brcaness brcaness;

  public SomaticNgsReportResults() {}

  public SomaticNgsReportResults rnaSeqs(@Nullable List<RnaSeq> rnaSeqs) {
    this.rnaSeqs = rnaSeqs;
    return this;
  }

  public SomaticNgsReportResults addRnaSeqsItem(RnaSeq rnaSeqsItem) {
    if (this.rnaSeqs == null) {
      this.rnaSeqs = new ArrayList<>();
    }
    this.rnaSeqs.add(rnaSeqsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RNA_SEQS, required = false)
  public @Nullable List<RnaSeq> getRnaSeqs() {
    return rnaSeqs;
  }

  @JsonProperty(value = JSON_PROPERTY_RNA_SEQS, required = false)
  public void setRnaSeqs(@Nullable List<RnaSeq> rnaSeqs) {
    this.rnaSeqs = rnaSeqs;
  }

  public SomaticNgsReportResults tmb(@Nullable Tmb tmb) {
    this.tmb = tmb;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TMB, required = false)
  public @Nullable Tmb getTmb() {
    return tmb;
  }

  @JsonProperty(value = JSON_PROPERTY_TMB, required = false)
  public void setTmb(@Nullable Tmb tmb) {
    this.tmb = tmb;
  }

  public SomaticNgsReportResults simpleVariants(@Nullable List<Snv> simpleVariants) {
    this.simpleVariants = simpleVariants;
    return this;
  }

  public SomaticNgsReportResults addSimpleVariantsItem(Snv simpleVariantsItem) {
    if (this.simpleVariants == null) {
      this.simpleVariants = new ArrayList<>();
    }
    this.simpleVariants.add(simpleVariantsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SIMPLE_VARIANTS, required = false)
  public @Nullable List<Snv> getSimpleVariants() {
    return simpleVariants;
  }

  @JsonProperty(value = JSON_PROPERTY_SIMPLE_VARIANTS, required = false)
  public void setSimpleVariants(@Nullable List<Snv> simpleVariants) {
    this.simpleVariants = simpleVariants;
  }

  public SomaticNgsReportResults hrdScore(@Nullable HrdScore hrdScore) {
    this.hrdScore = hrdScore;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HRD_SCORE, required = false)
  public @Nullable HrdScore getHrdScore() {
    return hrdScore;
  }

  @JsonProperty(value = JSON_PROPERTY_HRD_SCORE, required = false)
  public void setHrdScore(@Nullable HrdScore hrdScore) {
    this.hrdScore = hrdScore;
  }

  public SomaticNgsReportResults copyNumberVariants(@Nullable List<Cnv> copyNumberVariants) {
    this.copyNumberVariants = copyNumberVariants;
    return this;
  }

  public SomaticNgsReportResults addCopyNumberVariantsItem(Cnv copyNumberVariantsItem) {
    if (this.copyNumberVariants == null) {
      this.copyNumberVariants = new ArrayList<>();
    }
    this.copyNumberVariants.add(copyNumberVariantsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_COPY_NUMBER_VARIANTS, required = false)
  public @Nullable List<Cnv> getCopyNumberVariants() {
    return copyNumberVariants;
  }

  @JsonProperty(value = JSON_PROPERTY_COPY_NUMBER_VARIANTS, required = false)
  public void setCopyNumberVariants(@Nullable List<Cnv> copyNumberVariants) {
    this.copyNumberVariants = copyNumberVariants;
  }

  public SomaticNgsReportResults tumorCellContent(@Nullable TumorCellContent tumorCellContent) {
    this.tumorCellContent = tumorCellContent;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_CELL_CONTENT, required = false)
  public @Nullable TumorCellContent getTumorCellContent() {
    return tumorCellContent;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_CELL_CONTENT, required = false)
  public void setTumorCellContent(@Nullable TumorCellContent tumorCellContent) {
    this.tumorCellContent = tumorCellContent;
  }

  public SomaticNgsReportResults dnaFusions(@Nullable List<DnaFusion> dnaFusions) {
    this.dnaFusions = dnaFusions;
    return this;
  }

  public SomaticNgsReportResults addDnaFusionsItem(DnaFusion dnaFusionsItem) {
    if (this.dnaFusions == null) {
      this.dnaFusions = new ArrayList<>();
    }
    this.dnaFusions.add(dnaFusionsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DNA_FUSIONS, required = false)
  public @Nullable List<DnaFusion> getDnaFusions() {
    return dnaFusions;
  }

  @JsonProperty(value = JSON_PROPERTY_DNA_FUSIONS, required = false)
  public void setDnaFusions(@Nullable List<DnaFusion> dnaFusions) {
    this.dnaFusions = dnaFusions;
  }

  public SomaticNgsReportResults rnaFusions(@Nullable List<RnaFusion> rnaFusions) {
    this.rnaFusions = rnaFusions;
    return this;
  }

  public SomaticNgsReportResults addRnaFusionsItem(RnaFusion rnaFusionsItem) {
    if (this.rnaFusions == null) {
      this.rnaFusions = new ArrayList<>();
    }
    this.rnaFusions.add(rnaFusionsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RNA_FUSIONS, required = false)
  public @Nullable List<RnaFusion> getRnaFusions() {
    return rnaFusions;
  }

  @JsonProperty(value = JSON_PROPERTY_RNA_FUSIONS, required = false)
  public void setRnaFusions(@Nullable List<RnaFusion> rnaFusions) {
    this.rnaFusions = rnaFusions;
  }

  public SomaticNgsReportResults brcaness(@Nullable Brcaness brcaness) {
    this.brcaness = brcaness;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_BRCANESS, required = false)
  public @Nullable Brcaness getBrcaness() {
    return brcaness;
  }

  @JsonProperty(value = JSON_PROPERTY_BRCANESS, required = false)
  public void setBrcaness(@Nullable Brcaness brcaness) {
    this.brcaness = brcaness;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SomaticNgsReportResults somaticNgsReportResults = (SomaticNgsReportResults) o;
    return Objects.equals(this.rnaSeqs, somaticNgsReportResults.rnaSeqs)
        && Objects.equals(this.tmb, somaticNgsReportResults.tmb)
        && Objects.equals(this.simpleVariants, somaticNgsReportResults.simpleVariants)
        && Objects.equals(this.hrdScore, somaticNgsReportResults.hrdScore)
        && Objects.equals(this.copyNumberVariants, somaticNgsReportResults.copyNumberVariants)
        && Objects.equals(this.tumorCellContent, somaticNgsReportResults.tumorCellContent)
        && Objects.equals(this.dnaFusions, somaticNgsReportResults.dnaFusions)
        && Objects.equals(this.rnaFusions, somaticNgsReportResults.rnaFusions)
        && Objects.equals(this.brcaness, somaticNgsReportResults.brcaness);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        rnaSeqs,
        tmb,
        simpleVariants,
        hrdScore,
        copyNumberVariants,
        tumorCellContent,
        dnaFusions,
        rnaFusions,
        brcaness);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SomaticNgsReportResults {\n");
    sb.append("    rnaSeqs: ").append(toIndentedString(rnaSeqs)).append("\n");
    sb.append("    tmb: ").append(toIndentedString(tmb)).append("\n");
    sb.append("    simpleVariants: ").append(toIndentedString(simpleVariants)).append("\n");
    sb.append("    hrdScore: ").append(toIndentedString(hrdScore)).append("\n");
    sb.append("    copyNumberVariants: ").append(toIndentedString(copyNumberVariants)).append("\n");
    sb.append("    tumorCellContent: ").append(toIndentedString(tumorCellContent)).append("\n");
    sb.append("    dnaFusions: ").append(toIndentedString(dnaFusions)).append("\n");
    sb.append("    rnaFusions: ").append(toIndentedString(rnaFusions)).append("\n");
    sb.append("    brcaness: ").append(toIndentedString(brcaness)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private SomaticNgsReportResults instance;

    public Builder() {
      this(new SomaticNgsReportResults());
    }

    protected Builder(SomaticNgsReportResults instance) {
      this.instance = instance;
    }

    public SomaticNgsReportResults.Builder rnaSeqs(List<RnaSeq> rnaSeqs) {
      this.instance.rnaSeqs = rnaSeqs;
      return this;
    }

    public SomaticNgsReportResults.Builder tmb(Tmb tmb) {
      this.instance.tmb = tmb;
      return this;
    }

    public SomaticNgsReportResults.Builder simpleVariants(List<Snv> simpleVariants) {
      this.instance.simpleVariants = simpleVariants;
      return this;
    }

    public SomaticNgsReportResults.Builder hrdScore(HrdScore hrdScore) {
      this.instance.hrdScore = hrdScore;
      return this;
    }

    public SomaticNgsReportResults.Builder copyNumberVariants(List<Cnv> copyNumberVariants) {
      this.instance.copyNumberVariants = copyNumberVariants;
      return this;
    }

    public SomaticNgsReportResults.Builder tumorCellContent(TumorCellContent tumorCellContent) {
      this.instance.tumorCellContent = tumorCellContent;
      return this;
    }

    public SomaticNgsReportResults.Builder dnaFusions(List<DnaFusion> dnaFusions) {
      this.instance.dnaFusions = dnaFusions;
      return this;
    }

    public SomaticNgsReportResults.Builder rnaFusions(List<RnaFusion> rnaFusions) {
      this.instance.rnaFusions = rnaFusions;
      return this;
    }

    public SomaticNgsReportResults.Builder brcaness(Brcaness brcaness) {
      this.instance.brcaness = brcaness;
      return this;
    }

    public SomaticNgsReportResults build() {
      try {
        return this.instance;
      } finally {
        // ensure that this.instance is not reused
        this.instance = null;
      }
    }

    @Override
    public String toString() {
      return getClass() + "=(" + instance + ")";
    }
  }

  public static SomaticNgsReportResults.Builder builder() {
    return new SomaticNgsReportResults.Builder();
  }

  public SomaticNgsReportResults.Builder toBuilder() {
    return new SomaticNgsReportResults.Builder()
        .rnaSeqs(getRnaSeqs())
        .tmb(getTmb())
        .simpleVariants(getSimpleVariants())
        .hrdScore(getHrdScore())
        .copyNumberVariants(getCopyNumberVariants())
        .tumorCellContent(getTumorCellContent())
        .dnaFusions(getDnaFusions())
        .rnaFusions(getRnaFusions())
        .brcaness(getBrcaness());
  }
}
