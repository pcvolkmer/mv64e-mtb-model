package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class RnaSeq {
  public static final String JSON_PROPERTY_LIBRARY_SIZE = "librarySize";

  private @Nullable Integer librarySize;

  public static final String JSON_PROPERTY_LOCALIZATION = "localization";

  private @Nullable Set<BaseVariantLocalizationCoding> localization;

  public static final String JSON_PROPERTY_TRANSCRIPTS_PER_MILLION = "transcriptsPerMillion";

  private BigDecimal transcriptsPerMillion;

  public static final String JSON_PROPERTY_COHORT_RANKING = "cohortRanking";

  private @Nullable Integer cohortRanking;

  public static final String JSON_PROPERTY_RAW_COUNTS = "rawCounts";

  private Integer rawCounts;

  public static final String JSON_PROPERTY_VARIANT = "variant";

  private Reference variant;

  public static final String JSON_PROPERTY_TRANSCRIPT_ID = "transcriptId";

  private @Nullable TranscriptId transcriptId;

  public static final String JSON_PROPERTY_EXTERNAL_IDS = "externalIds";

  private @Nullable List<VariantExternalId> externalIds;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_TISSUE_CORRECTED_EXPRESSION =
      "tissueCorrectedExpression";

  private @Nullable Boolean tissueCorrectedExpression;

  public static final String JSON_PROPERTY_GENE = "gene";

  private @Nullable Coding gene;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public RnaSeq() {}

  public RnaSeq librarySize(@Nullable Integer librarySize) {
    this.librarySize = librarySize;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LIBRARY_SIZE, required = false)
  public @Nullable Integer getLibrarySize() {
    return librarySize;
  }

  @JsonProperty(value = JSON_PROPERTY_LIBRARY_SIZE, required = false)
  public void setLibrarySize(@Nullable Integer librarySize) {
    this.librarySize = librarySize;
  }

  public RnaSeq localization(@Nullable Set<BaseVariantLocalizationCoding> localization) {
    this.localization = localization;
    return this;
  }

  public RnaSeq addLocalizationItem(BaseVariantLocalizationCoding localizationItem) {
    if (this.localization == null) {
      this.localization = new LinkedHashSet<>();
    }
    this.localization.add(localizationItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LOCALIZATION, required = false)
  public @Nullable Set<BaseVariantLocalizationCoding> getLocalization() {
    return localization;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_LOCALIZATION, required = false)
  public void setLocalization(@Nullable Set<BaseVariantLocalizationCoding> localization) {
    this.localization = localization;
  }

  public RnaSeq transcriptsPerMillion(BigDecimal transcriptsPerMillion) {
    this.transcriptsPerMillion = transcriptsPerMillion;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPTS_PER_MILLION, required = true)
  public BigDecimal getTranscriptsPerMillion() {
    return transcriptsPerMillion;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPTS_PER_MILLION, required = true)
  public void setTranscriptsPerMillion(BigDecimal transcriptsPerMillion) {
    this.transcriptsPerMillion = transcriptsPerMillion;
  }

  public RnaSeq cohortRanking(@Nullable Integer cohortRanking) {
    this.cohortRanking = cohortRanking;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_COHORT_RANKING, required = false)
  public @Nullable Integer getCohortRanking() {
    return cohortRanking;
  }

  @JsonProperty(value = JSON_PROPERTY_COHORT_RANKING, required = false)
  public void setCohortRanking(@Nullable Integer cohortRanking) {
    this.cohortRanking = cohortRanking;
  }

  public RnaSeq rawCounts(Integer rawCounts) {
    this.rawCounts = rawCounts;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RAW_COUNTS, required = true)
  public Integer getRawCounts() {
    return rawCounts;
  }

  @JsonProperty(value = JSON_PROPERTY_RAW_COUNTS, required = true)
  public void setRawCounts(Integer rawCounts) {
    this.rawCounts = rawCounts;
  }

  public RnaSeq variant(Reference variant) {
    this.variant = variant;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VARIANT, required = true)
  public Reference getVariant() {
    return variant;
  }

  @JsonProperty(value = JSON_PROPERTY_VARIANT, required = true)
  public void setVariant(Reference variant) {
    this.variant = variant;
  }

  public RnaSeq transcriptId(@Nullable TranscriptId transcriptId) {
    this.transcriptId = transcriptId;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPT_ID, required = false)
  public @Nullable TranscriptId getTranscriptId() {
    return transcriptId;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPT_ID, required = false)
  public void setTranscriptId(@Nullable TranscriptId transcriptId) {
    this.transcriptId = transcriptId;
  }

  public RnaSeq externalIds(@Nullable List<VariantExternalId> externalIds) {
    this.externalIds = externalIds;
    return this;
  }

  public RnaSeq addExternalIdsItem(VariantExternalId externalIdsItem) {
    if (this.externalIds == null) {
      this.externalIds = new ArrayList<>();
    }
    this.externalIds.add(externalIdsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EXTERNAL_IDS, required = false)
  public @Nullable List<VariantExternalId> getExternalIds() {
    return externalIds;
  }

  @JsonProperty(value = JSON_PROPERTY_EXTERNAL_IDS, required = false)
  public void setExternalIds(@Nullable List<VariantExternalId> externalIds) {
    this.externalIds = externalIds;
  }

  public RnaSeq id(String id) {
    this.id = id;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public String getId() {
    return id;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public void setId(String id) {
    this.id = id;
  }

  public RnaSeq tissueCorrectedExpression(@Nullable Boolean tissueCorrectedExpression) {
    this.tissueCorrectedExpression = tissueCorrectedExpression;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TISSUE_CORRECTED_EXPRESSION, required = false)
  public @Nullable Boolean getTissueCorrectedExpression() {
    return tissueCorrectedExpression;
  }

  @JsonProperty(value = JSON_PROPERTY_TISSUE_CORRECTED_EXPRESSION, required = false)
  public void setTissueCorrectedExpression(@Nullable Boolean tissueCorrectedExpression) {
    this.tissueCorrectedExpression = tissueCorrectedExpression;
  }

  public RnaSeq gene(@Nullable Coding gene) {
    this.gene = gene;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GENE, required = false)
  public @Nullable Coding getGene() {
    return gene;
  }

  @JsonProperty(value = JSON_PROPERTY_GENE, required = false)
  public void setGene(@Nullable Coding gene) {
    this.gene = gene;
  }

  public RnaSeq patient(Reference patient) {
    this.patient = patient;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  public Reference getPatient() {
    return patient;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  public void setPatient(Reference patient) {
    this.patient = patient;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RnaSeq rnaSeq = (RnaSeq) o;
    return Objects.equals(this.librarySize, rnaSeq.librarySize)
        && Objects.equals(this.localization, rnaSeq.localization)
        && Objects.equals(this.transcriptsPerMillion, rnaSeq.transcriptsPerMillion)
        && Objects.equals(this.cohortRanking, rnaSeq.cohortRanking)
        && Objects.equals(this.rawCounts, rnaSeq.rawCounts)
        && Objects.equals(this.variant, rnaSeq.variant)
        && Objects.equals(this.transcriptId, rnaSeq.transcriptId)
        && Objects.equals(this.externalIds, rnaSeq.externalIds)
        && Objects.equals(this.id, rnaSeq.id)
        && Objects.equals(this.tissueCorrectedExpression, rnaSeq.tissueCorrectedExpression)
        && Objects.equals(this.gene, rnaSeq.gene)
        && Objects.equals(this.patient, rnaSeq.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        librarySize,
        localization,
        transcriptsPerMillion,
        cohortRanking,
        rawCounts,
        variant,
        transcriptId,
        externalIds,
        id,
        tissueCorrectedExpression,
        gene,
        patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RnaSeq {\n");
    sb.append("    librarySize: ").append(toIndentedString(librarySize)).append("\n");
    sb.append("    localization: ").append(toIndentedString(localization)).append("\n");
    sb.append("    transcriptsPerMillion: ")
        .append(toIndentedString(transcriptsPerMillion))
        .append("\n");
    sb.append("    cohortRanking: ").append(toIndentedString(cohortRanking)).append("\n");
    sb.append("    rawCounts: ").append(toIndentedString(rawCounts)).append("\n");
    sb.append("    variant: ").append(toIndentedString(variant)).append("\n");
    sb.append("    transcriptId: ").append(toIndentedString(transcriptId)).append("\n");
    sb.append("    externalIds: ").append(toIndentedString(externalIds)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    tissueCorrectedExpression: ")
        .append(toIndentedString(tissueCorrectedExpression))
        .append("\n");
    sb.append("    gene: ").append(toIndentedString(gene)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private RnaSeq instance;

    public Builder() {
      this(new RnaSeq());
    }

    protected Builder(RnaSeq instance) {
      this.instance = instance;
    }

    public RnaSeq.Builder librarySize(Integer librarySize) {
      this.instance.librarySize = librarySize;
      return this;
    }

    public RnaSeq.Builder localization(Set<BaseVariantLocalizationCoding> localization) {
      this.instance.localization = localization;
      return this;
    }

    public RnaSeq.Builder transcriptsPerMillion(BigDecimal transcriptsPerMillion) {
      this.instance.transcriptsPerMillion = transcriptsPerMillion;
      return this;
    }

    public RnaSeq.Builder cohortRanking(Integer cohortRanking) {
      this.instance.cohortRanking = cohortRanking;
      return this;
    }

    public RnaSeq.Builder rawCounts(Integer rawCounts) {
      this.instance.rawCounts = rawCounts;
      return this;
    }

    public RnaSeq.Builder variant(Reference variant) {
      this.instance.variant = variant;
      return this;
    }

    public RnaSeq.Builder transcriptId(TranscriptId transcriptId) {
      this.instance.transcriptId = transcriptId;
      return this;
    }

    public RnaSeq.Builder externalIds(List<VariantExternalId> externalIds) {
      this.instance.externalIds = externalIds;
      return this;
    }

    public RnaSeq.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public RnaSeq.Builder tissueCorrectedExpression(Boolean tissueCorrectedExpression) {
      this.instance.tissueCorrectedExpression = tissueCorrectedExpression;
      return this;
    }

    public RnaSeq.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public RnaSeq.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public RnaSeq build() {
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

  public static RnaSeq.Builder builder() {
    return new RnaSeq.Builder();
  }

  public RnaSeq.Builder toBuilder() {
    return new RnaSeq.Builder()
        .librarySize(getLibrarySize())
        .localization(getLocalization())
        .transcriptsPerMillion(getTranscriptsPerMillion())
        .cohortRanking(getCohortRanking())
        .rawCounts(getRawCounts())
        .variant(getVariant())
        .transcriptId(getTranscriptId())
        .externalIds(getExternalIds())
        .id(getId())
        .tissueCorrectedExpression(getTissueCorrectedExpression())
        .gene(getGene())
        .patient(getPatient());
  }
}
