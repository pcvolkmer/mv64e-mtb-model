package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class HrdScore {
  public static final String JSON_PROPERTY_INTERPRETATION = "interpretation";

  private @Nullable HrdScoreInterpretationCoding interpretation;

  public static final String JSON_PROPERTY_COMPONENTS = "components";

  private HrdScoreComponents components;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public HrdScore() {}

  public HrdScore interpretation(@Nullable HrdScoreInterpretationCoding interpretation) {
    this.interpretation = interpretation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  public @Nullable HrdScoreInterpretationCoding getInterpretation() {
    return interpretation;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  public void setInterpretation(@Nullable HrdScoreInterpretationCoding interpretation) {
    this.interpretation = interpretation;
  }

  public HrdScore components(HrdScoreComponents components) {
    this.components = components;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_COMPONENTS, required = true)
  public HrdScoreComponents getComponents() {
    return components;
  }

  @JsonProperty(value = JSON_PROPERTY_COMPONENTS, required = true)
  public void setComponents(HrdScoreComponents components) {
    this.components = components;
  }

  public HrdScore id(String id) {
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

  public HrdScore specimen(Reference specimen) {
    this.specimen = specimen;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  public Reference getSpecimen() {
    return specimen;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  public void setSpecimen(Reference specimen) {
    this.specimen = specimen;
  }

  public HrdScore value(BigDecimal value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public BigDecimal getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(BigDecimal value) {
    this.value = value;
  }

  public HrdScore patient(Reference patient) {
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
    HrdScore hrdScore = (HrdScore) o;
    return Objects.equals(this.interpretation, hrdScore.interpretation)
        && Objects.equals(this.components, hrdScore.components)
        && Objects.equals(this.id, hrdScore.id)
        && Objects.equals(this.specimen, hrdScore.specimen)
        && Objects.equals(this.value, hrdScore.value)
        && Objects.equals(this.patient, hrdScore.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(interpretation, components, id, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HrdScore {\n");
    sb.append("    interpretation: ").append(toIndentedString(interpretation)).append("\n");
    sb.append("    components: ").append(toIndentedString(components)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    specimen: ").append(toIndentedString(specimen)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private HrdScore instance;

    public Builder() {
      this(new HrdScore());
    }

    protected Builder(HrdScore instance) {
      this.instance = instance;
    }

    public HrdScore.Builder interpretation(HrdScoreInterpretationCoding interpretation) {
      this.instance.interpretation = interpretation;
      return this;
    }

    public HrdScore.Builder components(HrdScoreComponents components) {
      this.instance.components = components;
      return this;
    }

    public HrdScore.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public HrdScore.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public HrdScore.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public HrdScore.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public HrdScore build() {
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

  public static HrdScore.Builder builder() {
    return new HrdScore.Builder();
  }

  public HrdScore.Builder toBuilder() {
    return new HrdScore.Builder()
        .interpretation(getInterpretation())
        .components(getComponents())
        .id(getId())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
