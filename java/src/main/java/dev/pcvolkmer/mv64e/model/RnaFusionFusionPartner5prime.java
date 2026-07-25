package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class RnaFusionFusionPartner5prime {
  public static final String JSON_PROPERTY_TRANSCRIPT_ID = "transcriptId";

  private TranscriptId transcriptId;

  public static final String JSON_PROPERTY_POSITION = "position";

  private BigDecimal position;

  public static final String JSON_PROPERTY_EXON_ID = "exonId";

  private String exonId;

  public static final String JSON_PROPERTY_STRAND = "strand";

  private RnaFusionStrand strand;

  public static final String JSON_PROPERTY_GENE = "gene";

  private Coding gene;

  public RnaFusionFusionPartner5prime() {}

  public RnaFusionFusionPartner5prime transcriptId(TranscriptId transcriptId) {
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

  public RnaFusionFusionPartner5prime position(BigDecimal position) {
    this.position = position;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_POSITION, required = true)
  public BigDecimal getPosition() {
    return position;
  }

  @JsonProperty(value = JSON_PROPERTY_POSITION, required = true)
  public void setPosition(BigDecimal position) {
    this.position = position;
  }

  public RnaFusionFusionPartner5prime exonId(String exonId) {
    this.exonId = exonId;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EXON_ID, required = true)
  public String getExonId() {
    return exonId;
  }

  @JsonProperty(value = JSON_PROPERTY_EXON_ID, required = true)
  public void setExonId(String exonId) {
    this.exonId = exonId;
  }

  public RnaFusionFusionPartner5prime strand(RnaFusionStrand strand) {
    this.strand = strand;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STRAND, required = true)
  public RnaFusionStrand getStrand() {
    return strand;
  }

  @JsonProperty(value = JSON_PROPERTY_STRAND, required = true)
  public void setStrand(RnaFusionStrand strand) {
    this.strand = strand;
  }

  public RnaFusionFusionPartner5prime gene(Coding gene) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RnaFusionFusionPartner5prime rnaFusionFusionPartner5prime = (RnaFusionFusionPartner5prime) o;
    return Objects.equals(this.transcriptId, rnaFusionFusionPartner5prime.transcriptId)
        && Objects.equals(this.position, rnaFusionFusionPartner5prime.position)
        && Objects.equals(this.exonId, rnaFusionFusionPartner5prime.exonId)
        && Objects.equals(this.strand, rnaFusionFusionPartner5prime.strand)
        && Objects.equals(this.gene, rnaFusionFusionPartner5prime.gene);
  }

  @Override
  public int hashCode() {
    return Objects.hash(transcriptId, position, exonId, strand, gene);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RnaFusionFusionPartner5prime {\n");
    sb.append("    transcriptId: ").append(toIndentedString(transcriptId)).append("\n");
    sb.append("    position: ").append(toIndentedString(position)).append("\n");
    sb.append("    exonId: ").append(toIndentedString(exonId)).append("\n");
    sb.append("    strand: ").append(toIndentedString(strand)).append("\n");
    sb.append("    gene: ").append(toIndentedString(gene)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private RnaFusionFusionPartner5prime instance;

    public Builder() {
      this(new RnaFusionFusionPartner5prime());
    }

    protected Builder(RnaFusionFusionPartner5prime instance) {
      this.instance = instance;
    }

    public RnaFusionFusionPartner5prime.Builder transcriptId(TranscriptId transcriptId) {
      this.instance.transcriptId = transcriptId;
      return this;
    }

    public RnaFusionFusionPartner5prime.Builder position(BigDecimal position) {
      this.instance.position = position;
      return this;
    }

    public RnaFusionFusionPartner5prime.Builder exonId(String exonId) {
      this.instance.exonId = exonId;
      return this;
    }

    public RnaFusionFusionPartner5prime.Builder strand(RnaFusionStrand strand) {
      this.instance.strand = strand;
      return this;
    }

    public RnaFusionFusionPartner5prime.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public RnaFusionFusionPartner5prime build() {
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

  public static RnaFusionFusionPartner5prime.Builder builder() {
    return new RnaFusionFusionPartner5prime.Builder();
  }

  public RnaFusionFusionPartner5prime.Builder toBuilder() {
    return new RnaFusionFusionPartner5prime.Builder()
        .transcriptId(getTranscriptId())
        .position(getPosition())
        .exonId(getExonId())
        .strand(getStrand())
        .gene(getGene());
  }
}
