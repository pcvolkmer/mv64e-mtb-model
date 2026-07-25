package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class GeneAlterationReference {
  public static final String JSON_PROPERTY_VARIANT = "variant";

  private Reference variant;

  public static final String JSON_PROPERTY_DISPLAY = "display";

  private @Nullable String display;

  public static final String JSON_PROPERTY_GENE = "gene";

  private @Nullable Coding gene;

  public GeneAlterationReference() {}

  public GeneAlterationReference variant(Reference variant) {
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

  public GeneAlterationReference display(@Nullable String display) {
    this.display = display;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DISPLAY, required = false)
  public @Nullable String getDisplay() {
    return display;
  }

  @JsonProperty(value = JSON_PROPERTY_DISPLAY, required = false)
  public void setDisplay(@Nullable String display) {
    this.display = display;
  }

  public GeneAlterationReference gene(@Nullable Coding gene) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GeneAlterationReference geneAlterationReference = (GeneAlterationReference) o;
    return Objects.equals(this.variant, geneAlterationReference.variant)
        && Objects.equals(this.display, geneAlterationReference.display)
        && Objects.equals(this.gene, geneAlterationReference.gene);
  }

  @Override
  public int hashCode() {
    return Objects.hash(variant, display, gene);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GeneAlterationReference {\n");
    sb.append("    variant: ").append(toIndentedString(variant)).append("\n");
    sb.append("    display: ").append(toIndentedString(display)).append("\n");
    sb.append("    gene: ").append(toIndentedString(gene)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private GeneAlterationReference instance;

    public Builder() {
      this(new GeneAlterationReference());
    }

    protected Builder(GeneAlterationReference instance) {
      this.instance = instance;
    }

    public GeneAlterationReference.Builder variant(Reference variant) {
      this.instance.variant = variant;
      return this;
    }

    public GeneAlterationReference.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public GeneAlterationReference.Builder gene(Coding gene) {
      this.instance.gene = gene;
      return this;
    }

    public GeneAlterationReference build() {
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

  public static GeneAlterationReference.Builder builder() {
    return new GeneAlterationReference.Builder();
  }

  public GeneAlterationReference.Builder toBuilder() {
    return new GeneAlterationReference.Builder()
        .variant(getVariant())
        .display(getDisplay())
        .gene(getGene());
  }
}
