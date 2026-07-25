package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class Brcaness {
  public static final String JSON_PROPERTY_CONFIDENCE_RANGE = "confidenceRange";

  private BrcanessConfidenceRange confidenceRange;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public Brcaness() {}

  public Brcaness confidenceRange(BrcanessConfidenceRange confidenceRange) {
    this.confidenceRange = confidenceRange;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CONFIDENCE_RANGE, required = true)
  public BrcanessConfidenceRange getConfidenceRange() {
    return confidenceRange;
  }

  @JsonProperty(value = JSON_PROPERTY_CONFIDENCE_RANGE, required = true)
  public void setConfidenceRange(BrcanessConfidenceRange confidenceRange) {
    this.confidenceRange = confidenceRange;
  }

  public Brcaness id(String id) {
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

  public Brcaness specimen(Reference specimen) {
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

  public Brcaness value(BigDecimal value) {
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

  public Brcaness patient(Reference patient) {
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
    Brcaness brcaness = (Brcaness) o;
    return Objects.equals(this.confidenceRange, brcaness.confidenceRange)
        && Objects.equals(this.id, brcaness.id)
        && Objects.equals(this.specimen, brcaness.specimen)
        && Objects.equals(this.value, brcaness.value)
        && Objects.equals(this.patient, brcaness.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(confidenceRange, id, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Brcaness {\n");
    sb.append("    confidenceRange: ").append(toIndentedString(confidenceRange)).append("\n");
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

    private Brcaness instance;

    public Builder() {
      this(new Brcaness());
    }

    protected Builder(Brcaness instance) {
      this.instance = instance;
    }

    public Brcaness.Builder confidenceRange(BrcanessConfidenceRange confidenceRange) {
      this.instance.confidenceRange = confidenceRange;
      return this;
    }

    public Brcaness.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Brcaness.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public Brcaness.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public Brcaness.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Brcaness build() {
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

  public static Brcaness.Builder builder() {
    return new Brcaness.Builder();
  }

  public Brcaness.Builder toBuilder() {
    return new Brcaness.Builder()
        .confidenceRange(getConfidenceRange())
        .id(getId())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
