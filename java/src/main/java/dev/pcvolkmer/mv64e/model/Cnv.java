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

public class Cnv {
  public static final String JSON_PROPERTY_LOCALIZATION = "localization";

  private @Nullable Set<BaseVariantLocalizationCoding> localization;

  public static final String JSON_PROPERTY_REPORTED_FOCALITY = "reportedFocality";

  private @Nullable String reportedFocality;

  public static final String JSON_PROPERTY_REPORTED_AFFECTED_GENES = "reportedAffectedGenes";

  private @Nullable Set<Coding> reportedAffectedGenes;

  public static final String JSON_PROPERTY_TOTAL_COPY_NUMBER = "totalCopyNumber";

  private @Nullable Integer totalCopyNumber;

  public static final String JSON_PROPERTY_CN_A = "cnA";

  private @Nullable BigDecimal cnA;

  public static final String JSON_PROPERTY_EXTERNAL_IDS = "externalIds";

  private @Nullable List<VariantExternalId> externalIds;

  public static final String JSON_PROPERTY_START_RANGE = "startRange";

  private @Nullable CnvStartRange startRange;

  public static final String JSON_PROPERTY_TYPE = "type";

  private CnvTypeCoding type;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_END_RANGE = "endRange";

  private @Nullable CnvEndRange endRange;

  public static final String JSON_PROPERTY_CHROMOSOME = "chromosome";

  private Chromosome chromosome;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_CN_B = "cnB";

  private @Nullable BigDecimal cnB;

  public static final String JSON_PROPERTY_RELATIVE_COPY_NUMBER = "relativeCopyNumber";

  private @Nullable BigDecimal relativeCopyNumber;

  public static final String JSON_PROPERTY_COPY_NUMBER_NEUTRAL_LO_H = "copyNumberNeutralLoH";

  private @Nullable Set<Coding> copyNumberNeutralLoH;

  public Cnv() {}

  public Cnv localization(@Nullable Set<BaseVariantLocalizationCoding> localization) {
    this.localization = localization;
    return this;
  }

  public Cnv addLocalizationItem(BaseVariantLocalizationCoding localizationItem) {
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

  public Cnv reportedFocality(@Nullable String reportedFocality) {
    this.reportedFocality = reportedFocality;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REPORTED_FOCALITY, required = false)
  public @Nullable String getReportedFocality() {
    return reportedFocality;
  }

  @JsonProperty(value = JSON_PROPERTY_REPORTED_FOCALITY, required = false)
  public void setReportedFocality(@Nullable String reportedFocality) {
    this.reportedFocality = reportedFocality;
  }

  public Cnv reportedAffectedGenes(@Nullable Set<Coding> reportedAffectedGenes) {
    this.reportedAffectedGenes = reportedAffectedGenes;
    return this;
  }

  public Cnv addReportedAffectedGenesItem(Coding reportedAffectedGenesItem) {
    if (this.reportedAffectedGenes == null) {
      this.reportedAffectedGenes = new LinkedHashSet<>();
    }
    this.reportedAffectedGenes.add(reportedAffectedGenesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REPORTED_AFFECTED_GENES, required = false)
  public @Nullable Set<Coding> getReportedAffectedGenes() {
    return reportedAffectedGenes;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_REPORTED_AFFECTED_GENES, required = false)
  public void setReportedAffectedGenes(@Nullable Set<Coding> reportedAffectedGenes) {
    this.reportedAffectedGenes = reportedAffectedGenes;
  }

  public Cnv totalCopyNumber(@Nullable Integer totalCopyNumber) {
    this.totalCopyNumber = totalCopyNumber;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TOTAL_COPY_NUMBER, required = false)
  public @Nullable Integer getTotalCopyNumber() {
    return totalCopyNumber;
  }

  @JsonProperty(value = JSON_PROPERTY_TOTAL_COPY_NUMBER, required = false)
  public void setTotalCopyNumber(@Nullable Integer totalCopyNumber) {
    this.totalCopyNumber = totalCopyNumber;
  }

  public Cnv cnA(@Nullable BigDecimal cnA) {
    this.cnA = cnA;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CN_A, required = false)
  public @Nullable BigDecimal getCnA() {
    return cnA;
  }

  @JsonProperty(value = JSON_PROPERTY_CN_A, required = false)
  public void setCnA(@Nullable BigDecimal cnA) {
    this.cnA = cnA;
  }

  public Cnv externalIds(@Nullable List<VariantExternalId> externalIds) {
    this.externalIds = externalIds;
    return this;
  }

  public Cnv addExternalIdsItem(VariantExternalId externalIdsItem) {
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

  public Cnv startRange(@Nullable CnvStartRange startRange) {
    this.startRange = startRange;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_START_RANGE, required = false)
  public @Nullable CnvStartRange getStartRange() {
    return startRange;
  }

  @JsonProperty(value = JSON_PROPERTY_START_RANGE, required = false)
  public void setStartRange(@Nullable CnvStartRange startRange) {
    this.startRange = startRange;
  }

  public Cnv type(CnvTypeCoding type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public CnvTypeCoding getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public void setType(CnvTypeCoding type) {
    this.type = type;
  }

  public Cnv patient(Reference patient) {
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

  public Cnv endRange(@Nullable CnvEndRange endRange) {
    this.endRange = endRange;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_END_RANGE, required = false)
  public @Nullable CnvEndRange getEndRange() {
    return endRange;
  }

  @JsonProperty(value = JSON_PROPERTY_END_RANGE, required = false)
  public void setEndRange(@Nullable CnvEndRange endRange) {
    this.endRange = endRange;
  }

  public Cnv chromosome(Chromosome chromosome) {
    this.chromosome = chromosome;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CHROMOSOME, required = true)
  public Chromosome getChromosome() {
    return chromosome;
  }

  @JsonProperty(value = JSON_PROPERTY_CHROMOSOME, required = true)
  public void setChromosome(Chromosome chromosome) {
    this.chromosome = chromosome;
  }

  public Cnv id(String id) {
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

  public Cnv cnB(@Nullable BigDecimal cnB) {
    this.cnB = cnB;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CN_B, required = false)
  public @Nullable BigDecimal getCnB() {
    return cnB;
  }

  @JsonProperty(value = JSON_PROPERTY_CN_B, required = false)
  public void setCnB(@Nullable BigDecimal cnB) {
    this.cnB = cnB;
  }

  public Cnv relativeCopyNumber(@Nullable BigDecimal relativeCopyNumber) {
    this.relativeCopyNumber = relativeCopyNumber;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RELATIVE_COPY_NUMBER, required = false)
  public @Nullable BigDecimal getRelativeCopyNumber() {
    return relativeCopyNumber;
  }

  @JsonProperty(value = JSON_PROPERTY_RELATIVE_COPY_NUMBER, required = false)
  public void setRelativeCopyNumber(@Nullable BigDecimal relativeCopyNumber) {
    this.relativeCopyNumber = relativeCopyNumber;
  }

  public Cnv copyNumberNeutralLoH(@Nullable Set<Coding> copyNumberNeutralLoH) {
    this.copyNumberNeutralLoH = copyNumberNeutralLoH;
    return this;
  }

  public Cnv addCopyNumberNeutralLoHItem(Coding copyNumberNeutralLoHItem) {
    if (this.copyNumberNeutralLoH == null) {
      this.copyNumberNeutralLoH = new LinkedHashSet<>();
    }
    this.copyNumberNeutralLoH.add(copyNumberNeutralLoHItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_COPY_NUMBER_NEUTRAL_LO_H, required = false)
  public @Nullable Set<Coding> getCopyNumberNeutralLoH() {
    return copyNumberNeutralLoH;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_COPY_NUMBER_NEUTRAL_LO_H, required = false)
  public void setCopyNumberNeutralLoH(@Nullable Set<Coding> copyNumberNeutralLoH) {
    this.copyNumberNeutralLoH = copyNumberNeutralLoH;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Cnv cnv = (Cnv) o;
    return Objects.equals(this.localization, cnv.localization)
        && Objects.equals(this.reportedFocality, cnv.reportedFocality)
        && Objects.equals(this.reportedAffectedGenes, cnv.reportedAffectedGenes)
        && Objects.equals(this.totalCopyNumber, cnv.totalCopyNumber)
        && Objects.equals(this.cnA, cnv.cnA)
        && Objects.equals(this.externalIds, cnv.externalIds)
        && Objects.equals(this.startRange, cnv.startRange)
        && Objects.equals(this.type, cnv.type)
        && Objects.equals(this.patient, cnv.patient)
        && Objects.equals(this.endRange, cnv.endRange)
        && Objects.equals(this.chromosome, cnv.chromosome)
        && Objects.equals(this.id, cnv.id)
        && Objects.equals(this.cnB, cnv.cnB)
        && Objects.equals(this.relativeCopyNumber, cnv.relativeCopyNumber)
        && Objects.equals(this.copyNumberNeutralLoH, cnv.copyNumberNeutralLoH);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        localization,
        reportedFocality,
        reportedAffectedGenes,
        totalCopyNumber,
        cnA,
        externalIds,
        startRange,
        type,
        patient,
        endRange,
        chromosome,
        id,
        cnB,
        relativeCopyNumber,
        copyNumberNeutralLoH);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Cnv {\n");
    sb.append("    localization: ").append(toIndentedString(localization)).append("\n");
    sb.append("    reportedFocality: ").append(toIndentedString(reportedFocality)).append("\n");
    sb.append("    reportedAffectedGenes: ")
        .append(toIndentedString(reportedAffectedGenes))
        .append("\n");
    sb.append("    totalCopyNumber: ").append(toIndentedString(totalCopyNumber)).append("\n");
    sb.append("    cnA: ").append(toIndentedString(cnA)).append("\n");
    sb.append("    externalIds: ").append(toIndentedString(externalIds)).append("\n");
    sb.append("    startRange: ").append(toIndentedString(startRange)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    endRange: ").append(toIndentedString(endRange)).append("\n");
    sb.append("    chromosome: ").append(toIndentedString(chromosome)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    cnB: ").append(toIndentedString(cnB)).append("\n");
    sb.append("    relativeCopyNumber: ").append(toIndentedString(relativeCopyNumber)).append("\n");
    sb.append("    copyNumberNeutralLoH: ")
        .append(toIndentedString(copyNumberNeutralLoH))
        .append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Cnv instance;

    public Builder() {
      this(new Cnv());
    }

    protected Builder(Cnv instance) {
      this.instance = instance;
    }

    public Cnv.Builder localization(Set<BaseVariantLocalizationCoding> localization) {
      this.instance.localization = localization;
      return this;
    }

    public Cnv.Builder reportedFocality(String reportedFocality) {
      this.instance.reportedFocality = reportedFocality;
      return this;
    }

    public Cnv.Builder reportedAffectedGenes(Set<Coding> reportedAffectedGenes) {
      this.instance.reportedAffectedGenes = reportedAffectedGenes;
      return this;
    }

    public Cnv.Builder totalCopyNumber(Integer totalCopyNumber) {
      this.instance.totalCopyNumber = totalCopyNumber;
      return this;
    }

    public Cnv.Builder cnA(BigDecimal cnA) {
      this.instance.cnA = cnA;
      return this;
    }

    public Cnv.Builder externalIds(List<VariantExternalId> externalIds) {
      this.instance.externalIds = externalIds;
      return this;
    }

    public Cnv.Builder startRange(CnvStartRange startRange) {
      this.instance.startRange = startRange;
      return this;
    }

    public Cnv.Builder type(CnvTypeCoding type) {
      this.instance.type = type;
      return this;
    }

    public Cnv.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Cnv.Builder endRange(CnvEndRange endRange) {
      this.instance.endRange = endRange;
      return this;
    }

    public Cnv.Builder chromosome(Chromosome chromosome) {
      this.instance.chromosome = chromosome;
      return this;
    }

    public Cnv.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Cnv.Builder cnB(BigDecimal cnB) {
      this.instance.cnB = cnB;
      return this;
    }

    public Cnv.Builder relativeCopyNumber(BigDecimal relativeCopyNumber) {
      this.instance.relativeCopyNumber = relativeCopyNumber;
      return this;
    }

    public Cnv.Builder copyNumberNeutralLoH(Set<Coding> copyNumberNeutralLoH) {
      this.instance.copyNumberNeutralLoH = copyNumberNeutralLoH;
      return this;
    }

    public Cnv build() {
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

  public static Cnv.Builder builder() {
    return new Cnv.Builder();
  }

  public Cnv.Builder toBuilder() {
    return new Cnv.Builder()
        .localization(getLocalization())
        .reportedFocality(getReportedFocality())
        .reportedAffectedGenes(getReportedAffectedGenes())
        .totalCopyNumber(getTotalCopyNumber())
        .cnA(getCnA())
        .externalIds(getExternalIds())
        .startRange(getStartRange())
        .type(getType())
        .patient(getPatient())
        .endRange(getEndRange())
        .chromosome(getChromosome())
        .id(getId())
        .cnB(getCnB())
        .relativeCopyNumber(getRelativeCopyNumber())
        .copyNumberNeutralLoH(getCopyNumberNeutralLoH());
  }
}
