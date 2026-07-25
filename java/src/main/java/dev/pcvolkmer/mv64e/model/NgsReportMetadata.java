package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.Objects;

public class NgsReportMetadata {
  public static final String JSON_PROPERTY_KIT_MANUFACTURER = "kitManufacturer";

  private String kitManufacturer;

  public static final String JSON_PROPERTY_PIPELINE = "pipeline";

  private URI pipeline;

  public static final String JSON_PROPERTY_KIT_TYPE = "kitType";

  private String kitType;

  public static final String JSON_PROPERTY_SEQUENCER = "sequencer";

  private String sequencer;

  public static final String JSON_PROPERTY_REFERENCE_GENOME = "referenceGenome";

  private String referenceGenome;

  public NgsReportMetadata() {}

  public NgsReportMetadata kitManufacturer(String kitManufacturer) {
    this.kitManufacturer = kitManufacturer;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_KIT_MANUFACTURER, required = true)
  public String getKitManufacturer() {
    return kitManufacturer;
  }

  @JsonProperty(value = JSON_PROPERTY_KIT_MANUFACTURER, required = true)
  public void setKitManufacturer(String kitManufacturer) {
    this.kitManufacturer = kitManufacturer;
  }

  public NgsReportMetadata pipeline(URI pipeline) {
    this.pipeline = pipeline;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PIPELINE, required = true)
  public URI getPipeline() {
    return pipeline;
  }

  @JsonProperty(value = JSON_PROPERTY_PIPELINE, required = true)
  public void setPipeline(URI pipeline) {
    this.pipeline = pipeline;
  }

  public NgsReportMetadata kitType(String kitType) {
    this.kitType = kitType;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_KIT_TYPE, required = true)
  public String getKitType() {
    return kitType;
  }

  @JsonProperty(value = JSON_PROPERTY_KIT_TYPE, required = true)
  public void setKitType(String kitType) {
    this.kitType = kitType;
  }

  public NgsReportMetadata sequencer(String sequencer) {
    this.sequencer = sequencer;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SEQUENCER, required = true)
  public String getSequencer() {
    return sequencer;
  }

  @JsonProperty(value = JSON_PROPERTY_SEQUENCER, required = true)
  public void setSequencer(String sequencer) {
    this.sequencer = sequencer;
  }

  public NgsReportMetadata referenceGenome(String referenceGenome) {
    this.referenceGenome = referenceGenome;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REFERENCE_GENOME, required = true)
  public String getReferenceGenome() {
    return referenceGenome;
  }

  @JsonProperty(value = JSON_PROPERTY_REFERENCE_GENOME, required = true)
  public void setReferenceGenome(String referenceGenome) {
    this.referenceGenome = referenceGenome;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NgsReportMetadata ngsReportMetadata = (NgsReportMetadata) o;
    return Objects.equals(this.kitManufacturer, ngsReportMetadata.kitManufacturer)
        && Objects.equals(this.pipeline, ngsReportMetadata.pipeline)
        && Objects.equals(this.kitType, ngsReportMetadata.kitType)
        && Objects.equals(this.sequencer, ngsReportMetadata.sequencer)
        && Objects.equals(this.referenceGenome, ngsReportMetadata.referenceGenome);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kitManufacturer, pipeline, kitType, sequencer, referenceGenome);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NgsReportMetadata {\n");
    sb.append("    kitManufacturer: ").append(toIndentedString(kitManufacturer)).append("\n");
    sb.append("    pipeline: ").append(toIndentedString(pipeline)).append("\n");
    sb.append("    kitType: ").append(toIndentedString(kitType)).append("\n");
    sb.append("    sequencer: ").append(toIndentedString(sequencer)).append("\n");
    sb.append("    referenceGenome: ").append(toIndentedString(referenceGenome)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private NgsReportMetadata instance;

    public Builder() {
      this(new NgsReportMetadata());
    }

    protected Builder(NgsReportMetadata instance) {
      this.instance = instance;
    }

    public NgsReportMetadata.Builder kitManufacturer(String kitManufacturer) {
      this.instance.kitManufacturer = kitManufacturer;
      return this;
    }

    public NgsReportMetadata.Builder pipeline(URI pipeline) {
      this.instance.pipeline = pipeline;
      return this;
    }

    public NgsReportMetadata.Builder kitType(String kitType) {
      this.instance.kitType = kitType;
      return this;
    }

    public NgsReportMetadata.Builder sequencer(String sequencer) {
      this.instance.sequencer = sequencer;
      return this;
    }

    public NgsReportMetadata.Builder referenceGenome(String referenceGenome) {
      this.instance.referenceGenome = referenceGenome;
      return this;
    }

    public NgsReportMetadata build() {
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

  public static NgsReportMetadata.Builder builder() {
    return new NgsReportMetadata.Builder();
  }

  public NgsReportMetadata.Builder toBuilder() {
    return new NgsReportMetadata.Builder()
        .kitManufacturer(getKitManufacturer())
        .pipeline(getPipeline())
        .kitType(getKitType())
        .sequencer(getSequencer())
        .referenceGenome(getReferenceGenome());
  }
}
