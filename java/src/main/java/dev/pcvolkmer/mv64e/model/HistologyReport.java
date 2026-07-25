package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class HistologyReport {
  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RESULTS = "results";

  private HistologyReportResults results;

  public HistologyReport() {}

  public HistologyReport specimen(Reference specimen) {
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

  public HistologyReport issuedOn(Date issuedOn) {
    this.issuedOn = issuedOn;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ISSUED_ON, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getIssuedOn() {
    return issuedOn;
  }

  @JsonProperty(value = JSON_PROPERTY_ISSUED_ON, required = true)
  public void setIssuedOn(Date issuedOn) {
    this.issuedOn = issuedOn;
  }

  public HistologyReport patient(Reference patient) {
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

  public HistologyReport id(String id) {
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

  public HistologyReport results(HistologyReportResults results) {
    this.results = results;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RESULTS, required = true)
  public HistologyReportResults getResults() {
    return results;
  }

  @JsonProperty(value = JSON_PROPERTY_RESULTS, required = true)
  public void setResults(HistologyReportResults results) {
    this.results = results;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HistologyReport histologyReport = (HistologyReport) o;
    return Objects.equals(this.specimen, histologyReport.specimen)
        && Objects.equals(this.issuedOn, histologyReport.issuedOn)
        && Objects.equals(this.patient, histologyReport.patient)
        && Objects.equals(this.id, histologyReport.id)
        && Objects.equals(this.results, histologyReport.results);
  }

  @Override
  public int hashCode() {
    return Objects.hash(specimen, issuedOn, patient, id, results);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HistologyReport {\n");
    sb.append("    specimen: ").append(toIndentedString(specimen)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    results: ").append(toIndentedString(results)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private HistologyReport instance;

    public Builder() {
      this(new HistologyReport());
    }

    protected Builder(HistologyReport instance) {
      this.instance = instance;
    }

    public HistologyReport.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public HistologyReport.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public HistologyReport.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public HistologyReport.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public HistologyReport.Builder results(HistologyReportResults results) {
      this.instance.results = results;
      return this;
    }

    public HistologyReport build() {
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

  public static HistologyReport.Builder builder() {
    return new HistologyReport.Builder();
  }

  public HistologyReport.Builder toBuilder() {
    return new HistologyReport.Builder()
        .specimen(getSpecimen())
        .issuedOn(getIssuedOn())
        .patient(getPatient())
        .id(getId())
        .results(getResults());
  }
}
