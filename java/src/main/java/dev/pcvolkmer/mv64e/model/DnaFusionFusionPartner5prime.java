package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class DnaFusionFusionPartner5prime {
  public static final String JSON_PROPERTY_CHROMOSOME = "chromosome";

  private Chromosome chromosome;

  public static final String JSON_PROPERTY_GENE = "gene";

  private Coding gene;

  public static final String JSON_PROPERTY_POSITION = "position";

  private BigDecimal position;

  public DnaFusionFusionPartner5prime() {}

  public DnaFusionFusionPartner5prime chromosome(Chromosome chromosome) {
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

  public DnaFusionFusionPartner5prime gene(Coding gene) {
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

  public DnaFusionFusionPartner5prime position(BigDecimal position) {
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
    DnaFusionFusionPartner5prime dnaFusionFusionPartner5prime = (DnaFusionFusionPartner5prime) o;
    return Objects.equals(this.chromosome, dnaFusionFusionPartner5prime.chromosome)
        && Objects.equals(this.gene, dnaFusionFusionPartner5prime.gene)
        && Objects.equals(this.position, dnaFusionFusionPartner5prime.position);
  }

  @Override
  public int hashCode() {
    return Objects.hash(chromosome, gene, position);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DnaFusionFusionPartner5prime {\n");
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

    private DnaFusionFusionPartner5prime instance;

    public Builder() {
      this(new DnaFusionFusionPartner5prime());
    }

    protected Builder(DnaFusionFusionPartner5prime instance) {
      this.instance = instance;
    }

    public DnaFusionFusionPartner5prime.Builder chromosome(Chromosome chromosome) {
      this.instance.chromosome = chromosome;
      return this;
    }

    public DnaFusionFusionPartner5prime.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public DnaFusionFusionPartner5prime.Builder position(BigDecimal position) {
      this.instance.position = position;
      return this;
    }

    public DnaFusionFusionPartner5prime build() {
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

  public static DnaFusionFusionPartner5prime.Builder builder() {
    return new DnaFusionFusionPartner5prime.Builder();
  }

  public DnaFusionFusionPartner5prime.Builder toBuilder() {
    return new DnaFusionFusionPartner5prime.Builder()
        .chromosome(getChromosome())
        .gene(getGene())
        .position(getPosition());
  }
}
