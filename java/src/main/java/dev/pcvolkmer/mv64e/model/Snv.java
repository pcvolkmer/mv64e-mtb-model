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

public class Snv {
  public static final String JSON_PROPERTY_LOCALIZATION = "localization";

  private @Nullable Set<BaseVariantLocalizationCoding> localization;

  public static final String JSON_PROPERTY_PROTEIN_CHANGE = "proteinChange";

  private @Nullable String proteinChange;

  public static final String JSON_PROPERTY_INTERPRETATION = "interpretation";

  private @Nullable ClinVarCoding interpretation;

  public static final String JSON_PROPERTY_TRANSCRIPT_ID = "transcriptId";

  private TranscriptId transcriptId;

  public static final String JSON_PROPERTY_EXTERNAL_IDS = "externalIds";

  private @Nullable List<VariantExternalId> externalIds;

  public static final String JSON_PROPERTY_CHROMOSOME = "chromosome";

  private Chromosome chromosome;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_ALT_ALLELE = "altAllele";

  private String altAllele;

  public static final String JSON_PROPERTY_POSITION = "position";

  private CnvEndRange position;

  public static final String JSON_PROPERTY_GENE = "gene";

  private Coding gene;

  public static final String JSON_PROPERTY_EXON_ID = "exonId";

  private @Nullable String exonId;

  public static final String JSON_PROPERTY_DNA_CHANGE = "dnaChange";

  private String dnaChange;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_REF_ALLELE = "refAllele";

  private String refAllele;

  public static final String JSON_PROPERTY_READ_DEPTH = "readDepth";

  private Integer readDepth;

  public static final String JSON_PROPERTY_ALLELIC_FREQUENCY = "allelicFrequency";

  private BigDecimal allelicFrequency;

  public Snv() {}

  public Snv localization(@Nullable Set<BaseVariantLocalizationCoding> localization) {
    this.localization = localization;
    return this;
  }

  public Snv addLocalizationItem(BaseVariantLocalizationCoding localizationItem) {
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

  public Snv proteinChange(@Nullable String proteinChange) {
    this.proteinChange = proteinChange;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN_CHANGE, required = false)
  public @Nullable String getProteinChange() {
    return proteinChange;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN_CHANGE, required = false)
  public void setProteinChange(@Nullable String proteinChange) {
    this.proteinChange = proteinChange;
  }

  public Snv interpretation(@Nullable ClinVarCoding interpretation) {
    this.interpretation = interpretation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  public @Nullable ClinVarCoding getInterpretation() {
    return interpretation;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  public void setInterpretation(@Nullable ClinVarCoding interpretation) {
    this.interpretation = interpretation;
  }

  public Snv transcriptId(TranscriptId transcriptId) {
    this.transcriptId = transcriptId;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPT_ID, required = true)
  public TranscriptId getTranscriptId() {
    return transcriptId;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSCRIPT_ID, required = true)
  public void setTranscriptId(TranscriptId transcriptId) {
    this.transcriptId = transcriptId;
  }

  public Snv externalIds(@Nullable List<VariantExternalId> externalIds) {
    this.externalIds = externalIds;
    return this;
  }

  public Snv addExternalIdsItem(VariantExternalId externalIdsItem) {
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

  public Snv chromosome(Chromosome chromosome) {
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

  public Snv id(String id) {
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

  public Snv altAllele(String altAllele) {
    this.altAllele = altAllele;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ALT_ALLELE, required = true)
  public String getAltAllele() {
    return altAllele;
  }

  @JsonProperty(value = JSON_PROPERTY_ALT_ALLELE, required = true)
  public void setAltAllele(String altAllele) {
    this.altAllele = altAllele;
  }

  public Snv position(CnvEndRange position) {
    this.position = position;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_POSITION, required = true)
  public CnvEndRange getPosition() {
    return position;
  }

  @JsonProperty(value = JSON_PROPERTY_POSITION, required = true)
  public void setPosition(CnvEndRange position) {
    this.position = position;
  }

  public Snv gene(Coding gene) {
    this.gene = gene;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GENE, required = true)
  public Coding getGene() {
    return gene;
  }

  @JsonProperty(value = JSON_PROPERTY_GENE, required = true)
  public void setGene(Coding gene) {
    this.gene = gene;
  }

  public Snv exonId(@Nullable String exonId) {
    this.exonId = exonId;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EXON_ID, required = false)
  public @Nullable String getExonId() {
    return exonId;
  }

  @JsonProperty(value = JSON_PROPERTY_EXON_ID, required = false)
  public void setExonId(@Nullable String exonId) {
    this.exonId = exonId;
  }

  public Snv dnaChange(String dnaChange) {
    this.dnaChange = dnaChange;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DNA_CHANGE, required = true)
  public String getDnaChange() {
    return dnaChange;
  }

  @JsonProperty(value = JSON_PROPERTY_DNA_CHANGE, required = true)
  public void setDnaChange(String dnaChange) {
    this.dnaChange = dnaChange;
  }

  public Snv patient(Reference patient) {
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

  public Snv refAllele(String refAllele) {
    this.refAllele = refAllele;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REF_ALLELE, required = true)
  public String getRefAllele() {
    return refAllele;
  }

  @JsonProperty(value = JSON_PROPERTY_REF_ALLELE, required = true)
  public void setRefAllele(String refAllele) {
    this.refAllele = refAllele;
  }

  public Snv readDepth(Integer readDepth) {
    this.readDepth = readDepth;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_READ_DEPTH, required = true)
  public Integer getReadDepth() {
    return readDepth;
  }

  @JsonProperty(value = JSON_PROPERTY_READ_DEPTH, required = true)
  public void setReadDepth(Integer readDepth) {
    this.readDepth = readDepth;
  }

  public Snv allelicFrequency(BigDecimal allelicFrequency) {
    this.allelicFrequency = allelicFrequency;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ALLELIC_FREQUENCY, required = true)
  public BigDecimal getAllelicFrequency() {
    return allelicFrequency;
  }

  @JsonProperty(value = JSON_PROPERTY_ALLELIC_FREQUENCY, required = true)
  public void setAllelicFrequency(BigDecimal allelicFrequency) {
    this.allelicFrequency = allelicFrequency;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Snv snv = (Snv) o;
    return Objects.equals(this.localization, snv.localization)
        && Objects.equals(this.proteinChange, snv.proteinChange)
        && Objects.equals(this.interpretation, snv.interpretation)
        && Objects.equals(this.transcriptId, snv.transcriptId)
        && Objects.equals(this.externalIds, snv.externalIds)
        && Objects.equals(this.chromosome, snv.chromosome)
        && Objects.equals(this.id, snv.id)
        && Objects.equals(this.altAllele, snv.altAllele)
        && Objects.equals(this.position, snv.position)
        && Objects.equals(this.gene, snv.gene)
        && Objects.equals(this.exonId, snv.exonId)
        && Objects.equals(this.dnaChange, snv.dnaChange)
        && Objects.equals(this.patient, snv.patient)
        && Objects.equals(this.refAllele, snv.refAllele)
        && Objects.equals(this.readDepth, snv.readDepth)
        && Objects.equals(this.allelicFrequency, snv.allelicFrequency);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        localization,
        proteinChange,
        interpretation,
        transcriptId,
        externalIds,
        chromosome,
        id,
        altAllele,
        position,
        gene,
        exonId,
        dnaChange,
        patient,
        refAllele,
        readDepth,
        allelicFrequency);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Snv {\n");
    sb.append("    localization: ").append(toIndentedString(localization)).append("\n");
    sb.append("    proteinChange: ").append(toIndentedString(proteinChange)).append("\n");
    sb.append("    interpretation: ").append(toIndentedString(interpretation)).append("\n");
    sb.append("    transcriptId: ").append(toIndentedString(transcriptId)).append("\n");
    sb.append("    externalIds: ").append(toIndentedString(externalIds)).append("\n");
    sb.append("    chromosome: ").append(toIndentedString(chromosome)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    altAllele: ").append(toIndentedString(altAllele)).append("\n");
    sb.append("    position: ").append(toIndentedString(position)).append("\n");
    sb.append("    gene: ").append(toIndentedString(gene)).append("\n");
    sb.append("    exonId: ").append(toIndentedString(exonId)).append("\n");
    sb.append("    dnaChange: ").append(toIndentedString(dnaChange)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    refAllele: ").append(toIndentedString(refAllele)).append("\n");
    sb.append("    readDepth: ").append(toIndentedString(readDepth)).append("\n");
    sb.append("    allelicFrequency: ").append(toIndentedString(allelicFrequency)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Snv instance;

    public Builder() {
      this(new Snv());
    }

    protected Builder(Snv instance) {
      this.instance = instance;
    }

    public Snv.Builder localization(Set<BaseVariantLocalizationCoding> localization) {
      this.instance.localization = localization;
      return this;
    }

    public Snv.Builder proteinChange(String proteinChange) {
      this.instance.proteinChange = proteinChange;
      return this;
    }

    public Snv.Builder interpretation(ClinVarCoding interpretation) {
      this.instance.interpretation = interpretation;
      return this;
    }

    public Snv.Builder transcriptId(TranscriptId transcriptId) {
      this.instance.transcriptId = transcriptId;
      return this;
    }

    public Snv.Builder externalIds(List<VariantExternalId> externalIds) {
      this.instance.externalIds = externalIds;
      return this;
    }

    public Snv.Builder chromosome(Chromosome chromosome) {
      this.instance.chromosome = chromosome;
      return this;
    }

    public Snv.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Snv.Builder altAllele(String altAllele) {
      this.instance.altAllele = altAllele;
      return this;
    }

    public Snv.Builder position(CnvEndRange position) {
      this.instance.position = position;
      return this;
    }

    public Snv.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public Snv.Builder exonId(String exonId) {
      this.instance.exonId = exonId;
      return this;
    }

    public Snv.Builder dnaChange(String dnaChange) {
      this.instance.dnaChange = dnaChange;
      return this;
    }

    public Snv.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Snv.Builder refAllele(String refAllele) {
      this.instance.refAllele = refAllele;
      return this;
    }

    public Snv.Builder readDepth(Integer readDepth) {
      this.instance.readDepth = readDepth;
      return this;
    }

    public Snv.Builder allelicFrequency(BigDecimal allelicFrequency) {
      this.instance.allelicFrequency = allelicFrequency;
      return this;
    }

    public Snv build() {
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

  public static Snv.Builder builder() {
    return new Snv.Builder();
  }

  public Snv.Builder toBuilder() {
    return new Snv.Builder()
        .localization(getLocalization())
        .proteinChange(getProteinChange())
        .interpretation(getInterpretation())
        .transcriptId(getTranscriptId())
        .externalIds(getExternalIds())
        .chromosome(getChromosome())
        .id(getId())
        .altAllele(getAltAllele())
        .position(getPosition())
        .gene(getGene())
        .exonId(getExonId())
        .dnaChange(getDnaChange())
        .patient(getPatient())
        .refAllele(getRefAllele())
        .readDepth(getReadDepth())
        .allelicFrequency(getAllelicFrequency());
  }
}
