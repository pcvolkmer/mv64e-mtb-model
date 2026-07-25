package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;

public class TumorStagingTnmClassification {
  public static final String JSON_PROPERTY_TUMOR = "tumor";

  private Coding tumor;

  public static final String JSON_PROPERTY_NODES = "nodes";

  private Coding nodes;

  public static final String JSON_PROPERTY_METASTASIS = "metastasis";

  private Coding metastasis;

  public TumorStagingTnmClassification() {}

  public TumorStagingTnmClassification tumor(Coding tumor) {
    this.tumor = tumor;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Coding getTumor() {
    return tumor;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setTumor(Coding tumor) {
    this.tumor = tumor;
  }

  public TumorStagingTnmClassification nodes(Coding nodes) {
    this.nodes = nodes;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_NODES, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Coding getNodes() {
    return nodes;
  }

  @JsonProperty(value = JSON_PROPERTY_NODES, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setNodes(Coding nodes) {
    this.nodes = nodes;
  }

  public TumorStagingTnmClassification metastasis(Coding metastasis) {
    this.metastasis = metastasis;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METASTASIS, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Coding getMetastasis() {
    return metastasis;
  }

  @JsonProperty(value = JSON_PROPERTY_METASTASIS, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMetastasis(Coding metastasis) {
    this.metastasis = metastasis;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TumorStagingTnmClassification tumorStagingTnmClassification = (TumorStagingTnmClassification) o;
    return Objects.equals(this.tumor, tumorStagingTnmClassification.tumor)
        && Objects.equals(this.nodes, tumorStagingTnmClassification.nodes)
        && Objects.equals(this.metastasis, tumorStagingTnmClassification.metastasis);
  }

  @Override
  public int hashCode() {
    return Objects.hash(tumor, nodes, metastasis);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorStagingTnmClassification {\n");
    sb.append("    tumor: ").append(toIndentedString(tumor)).append("\n");
    sb.append("    nodes: ").append(toIndentedString(nodes)).append("\n");
    sb.append("    metastasis: ").append(toIndentedString(metastasis)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TumorStagingTnmClassification instance;

    public Builder() {
      this(new TumorStagingTnmClassification());
    }

    protected Builder(TumorStagingTnmClassification instance) {
      this.instance = instance;
    }

    public TumorStagingTnmClassification.Builder tumor(Coding tumor) {
      this.instance.tumor = tumor;
      return this;
    }

    public TumorStagingTnmClassification.Builder nodes(Coding nodes) {
      this.instance.nodes = nodes;
      return this;
    }

    public TumorStagingTnmClassification.Builder metastasis(Coding metastasis) {
      this.instance.metastasis = metastasis;
      return this;
    }

    public TumorStagingTnmClassification build() {
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

  public static TumorStagingTnmClassification.Builder builder() {
    return new TumorStagingTnmClassification.Builder();
  }

  public TumorStagingTnmClassification.Builder toBuilder() {
    return new TumorStagingTnmClassification.Builder()
        .tumor(getTumor())
        .nodes(getNodes())
        .metastasis(getMetastasis());
  }
}
