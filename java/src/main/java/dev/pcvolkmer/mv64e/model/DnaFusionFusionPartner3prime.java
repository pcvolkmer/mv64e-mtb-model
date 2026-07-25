package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class DnaFusionFusionPartner3prime {
  public static final String JSON_PROPERTY_CHROMOSOME = "chromosome";

  private Chromosome chromosome;

  public static final String JSON_PROPERTY_GENE = "gene";

  private Coding gene;

  public static final String JSON_PROPERTY_POSITION = "position";

  private BigDecimal position;

  public DnaFusionFusionPartner3prime() {}

  public DnaFusionFusionPartner3prime chromosome(Chromosome chromosome) {
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

  public DnaFusionFusionPartner3prime gene(Coding gene) {
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

  public DnaFusionFusionPartner3prime position(BigDecimal position) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DnaFusionFusionPartner3prime dnaFusionFusionPartner3prime = (DnaFusionFusionPartner3prime) o;
    return Objects.equals(this.chromosome, dnaFusionFusionPartner3prime.chromosome)
        && Objects.equals(this.gene, dnaFusionFusionPartner3prime.gene)
        && Objects.equals(this.position, dnaFusionFusionPartner3prime.position);
  }

  @Override
  public int hashCode() {
    return Objects.hash(chromosome, gene, position);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DnaFusionFusionPartner3prime {\n");
    sb.append("    chromosome: ").append(toIndentedString(chromosome)).append("\n");
    sb.append("    gene: ").append(toIndentedString(gene)).append("\n");
    sb.append("    position: ").append(toIndentedString(position)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private DnaFusionFusionPartner3prime instance;

    public Builder() {
      this(new DnaFusionFusionPartner3prime());
    }

    protected Builder(DnaFusionFusionPartner3prime instance) {
      this.instance = instance;
    }

    public DnaFusionFusionPartner3prime.Builder chromosome(Chromosome chromosome) {
      this.instance.chromosome = chromosome;
      return this;
    }

    public DnaFusionFusionPartner3prime.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public DnaFusionFusionPartner3prime.Builder position(BigDecimal position) {
      this.instance.position = position;
      return this;
    }

    public DnaFusionFusionPartner3prime build() {
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

  public static DnaFusionFusionPartner3prime.Builder builder() {
    return new DnaFusionFusionPartner3prime.Builder();
  }

  public DnaFusionFusionPartner3prime.Builder toBuilder() {
    return new DnaFusionFusionPartner3prime.Builder()
        .chromosome(getChromosome())
        .gene(getGene())
        .position(getPosition());
  }
}
