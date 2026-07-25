package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class TumorCellContent {
  public static final String JSON_PROPERTY_METHOD = "method";

  private TumorCellContentMethodCoding method;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public TumorCellContent() {}

  public TumorCellContent method(TumorCellContentMethodCoding method) {
    this.method = method;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TumorCellContentMethodCoding getMethod() {
    return method;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMethod(TumorCellContentMethodCoding method) {
    this.method = method;
  }

  public TumorCellContent id(String id) {
    this.id = id;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public String getId() {
    return id;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setId(String id) {
    this.id = id;
  }

  public TumorCellContent specimen(Reference specimen) {
    this.specimen = specimen;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Reference getSpecimen() {
    return specimen;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSpecimen(Reference specimen) {
    this.specimen = specimen;
  }

  public TumorCellContent value(BigDecimal value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public BigDecimal getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setValue(BigDecimal value) {
    this.value = value;
  }

  public TumorCellContent patient(Reference patient) {
    this.patient = patient;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Reference getPatient() {
    return patient;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
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
    TumorCellContent tumorCellContent = (TumorCellContent) o;
    return Objects.equals(this.method, tumorCellContent.method)
        && Objects.equals(this.id, tumorCellContent.id)
        && Objects.equals(this.specimen, tumorCellContent.specimen)
        && Objects.equals(this.value, tumorCellContent.value)
        && Objects.equals(this.patient, tumorCellContent.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(method, id, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorCellContent {\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
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

    private TumorCellContent instance;

    public Builder() {
      this(new TumorCellContent());
    }

    protected Builder(TumorCellContent instance) {
      this.instance = instance;
    }

    public TumorCellContent.Builder method(TumorCellContentMethodCoding method) {
      this.instance.method = method;
      return this;
    }

    public TumorCellContent.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public TumorCellContent.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public TumorCellContent.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public TumorCellContent.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public TumorCellContent build() {
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

  public static TumorCellContent.Builder builder() {
    return new TumorCellContent.Builder();
  }

  public TumorCellContent.Builder toBuilder() {
    return new TumorCellContent.Builder()
        .method(getMethod())
        .id(getId())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
