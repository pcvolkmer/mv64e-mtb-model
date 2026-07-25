package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class Msi {
  public static final String JSON_PROPERTY_METHOD = "method";

  private MsiMethodCoding method;

  public static final String JSON_PROPERTY_INTERPRETATION = "interpretation";

  private MsiInterpretationCoding interpretation;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public Msi() {}

  public Msi method(MsiMethodCoding method) {
    this.method = method;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  public MsiMethodCoding getMethod() {
    return method;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  public void setMethod(MsiMethodCoding method) {
    this.method = method;
  }

  public Msi interpretation(MsiInterpretationCoding interpretation) {
    this.interpretation = interpretation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = true)
  public MsiInterpretationCoding getInterpretation() {
    return interpretation;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = true)
  public void setInterpretation(MsiInterpretationCoding interpretation) {
    this.interpretation = interpretation;
  }

  public Msi id(String id) {
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

  public Msi specimen(Reference specimen) {
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

  public Msi value(BigDecimal value) {
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

  public Msi patient(Reference patient) {
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
    Msi msi = (Msi) o;
    return Objects.equals(this.method, msi.method)
        && Objects.equals(this.interpretation, msi.interpretation)
        && Objects.equals(this.id, msi.id)
        && Objects.equals(this.specimen, msi.specimen)
        && Objects.equals(this.value, msi.value)
        && Objects.equals(this.patient, msi.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(method, interpretation, id, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Msi {\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    interpretation: ").append(toIndentedString(interpretation)).append("\n");
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

    private Msi instance;

    public Builder() {
      this(new Msi());
    }

    protected Builder(Msi instance) {
      this.instance = instance;
    }

    public Msi.Builder method(MsiMethodCoding method) {
      this.instance.method = method;
      return this;
    }

    public Msi.Builder interpretation(MsiInterpretationCoding interpretation) {
      this.instance.interpretation = interpretation;
      return this;
    }

    public Msi.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Msi.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public Msi.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public Msi.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Msi build() {
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

  public static Msi.Builder builder() {
    return new Msi.Builder();
  }

  public Msi.Builder toBuilder() {
    return new Msi.Builder()
        .method(getMethod())
        .interpretation(getInterpretation())
        .id(getId())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
