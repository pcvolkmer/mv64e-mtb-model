package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;

public class PatientAddress {
  public static final String JSON_PROPERTY_MUNICIPALITY_CODE = "municipalityCode";

  private String municipalityCode;

  public PatientAddress() {}

  public PatientAddress municipalityCode(String municipalityCode) {
    this.municipalityCode = municipalityCode;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MUNICIPALITY_CODE, required = true)
  public String getMunicipalityCode() {
    return municipalityCode;
  }

  @JsonProperty(value = JSON_PROPERTY_MUNICIPALITY_CODE, required = true)
  public void setMunicipalityCode(String municipalityCode) {
    this.municipalityCode = municipalityCode;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PatientAddress patientAddress = (PatientAddress) o;
    return Objects.equals(this.municipalityCode, patientAddress.municipalityCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(municipalityCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PatientAddress {\n");
    sb.append("    municipalityCode: ").append(toIndentedString(municipalityCode)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PatientAddress instance;

    public Builder() {
      this(new PatientAddress());
    }

    protected Builder(PatientAddress instance) {
      this.instance = instance;
    }

    public PatientAddress.Builder municipalityCode(String municipalityCode) {
      this.instance.municipalityCode = municipalityCode;
      return this;
    }

    public PatientAddress build() {
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

  public static PatientAddress.Builder builder() {
    return new PatientAddress.Builder();
  }

  public PatientAddress.Builder toBuilder() {
    return new PatientAddress.Builder().municipalityCode(getMunicipalityCode());
  }
}
