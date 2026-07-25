package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class PerformanceStatus {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_EFFECTIVE_DATE = "effectiveDate";

  private Date effectiveDate;

  public static final String JSON_PROPERTY_VALUE = "value";

  private EcogCoding value;

  public PerformanceStatus() {}

  public PerformanceStatus id(String id) {
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

  public PerformanceStatus patient(Reference patient) {
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

  public PerformanceStatus effectiveDate(Date effectiveDate) {
    this.effectiveDate = effectiveDate;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EFFECTIVE_DATE, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getEffectiveDate() {
    return effectiveDate;
  }

  @JsonProperty(value = JSON_PROPERTY_EFFECTIVE_DATE, required = true)
  public void setEffectiveDate(Date effectiveDate) {
    this.effectiveDate = effectiveDate;
  }

  public PerformanceStatus value(EcogCoding value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public EcogCoding getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(EcogCoding value) {
    this.value = value;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PerformanceStatus performanceStatus = (PerformanceStatus) o;
    return Objects.equals(this.id, performanceStatus.id)
        && Objects.equals(this.patient, performanceStatus.patient)
        && Objects.equals(this.effectiveDate, performanceStatus.effectiveDate)
        && Objects.equals(this.value, performanceStatus.value);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, patient, effectiveDate, value);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PerformanceStatus {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    effectiveDate: ").append(toIndentedString(effectiveDate)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PerformanceStatus instance;

    public Builder() {
      this(new PerformanceStatus());
    }

    protected Builder(PerformanceStatus instance) {
      this.instance = instance;
    }

    public PerformanceStatus.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public PerformanceStatus.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public PerformanceStatus.Builder effectiveDate(Date effectiveDate) {
      this.instance.effectiveDate = effectiveDate;
      return this;
    }

    public PerformanceStatus.Builder value(EcogCoding value) {
      this.instance.value = value;
      return this;
    }

    public PerformanceStatus build() {
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

  public static PerformanceStatus.Builder builder() {
    return new PerformanceStatus.Builder();
  }

  public PerformanceStatus.Builder toBuilder() {
    return new PerformanceStatus.Builder()
        .id(getId())
        .patient(getPatient())
        .effectiveDate(getEffectiveDate())
        .value(getValue());
  }
}
