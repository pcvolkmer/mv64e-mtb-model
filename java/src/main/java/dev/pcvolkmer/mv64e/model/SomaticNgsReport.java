package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;

public class SomaticNgsReport {
  public static final String JSON_PROPERTY_METADATA = "metadata";

  private List<NgsReportMetadata> metadata;

  public static final String JSON_PROPERTY_TYPE = "type";

  private NgsReportTypeCoding type;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RESULTS = "results";

  private SomaticNgsReportResults results;

  public SomaticNgsReport() {}

  public SomaticNgsReport metadata(List<NgsReportMetadata> metadata) {
    this.metadata = metadata;
    return this;
  }

  public SomaticNgsReport addMetadataItem(NgsReportMetadata metadataItem) {
    if (this.metadata == null) {
      this.metadata = new ArrayList<>();
    }
    this.metadata.add(metadataItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METADATA, required = true)
  public List<NgsReportMetadata> getMetadata() {
    return metadata;
  }

  @JsonProperty(value = JSON_PROPERTY_METADATA, required = true)
  public void setMetadata(List<NgsReportMetadata> metadata) {
    this.metadata = metadata;
  }

  public SomaticNgsReport type(NgsReportTypeCoding type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public NgsReportTypeCoding getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public void setType(NgsReportTypeCoding type) {
    this.type = type;
  }

  public SomaticNgsReport specimen(Reference specimen) {
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

  public SomaticNgsReport issuedOn(Date issuedOn) {
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

  public SomaticNgsReport patient(Reference patient) {
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

  public SomaticNgsReport id(String id) {
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

  public SomaticNgsReport results(SomaticNgsReportResults results) {
    this.results = results;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RESULTS, required = true)
  public SomaticNgsReportResults getResults() {
    return results;
  }

  @JsonProperty(value = JSON_PROPERTY_RESULTS, required = true)
  public void setResults(SomaticNgsReportResults results) {
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
    SomaticNgsReport somaticNgsReport = (SomaticNgsReport) o;
    return Objects.equals(this.metadata, somaticNgsReport.metadata)
        && Objects.equals(this.type, somaticNgsReport.type)
        && Objects.equals(this.specimen, somaticNgsReport.specimen)
        && Objects.equals(this.issuedOn, somaticNgsReport.issuedOn)
        && Objects.equals(this.patient, somaticNgsReport.patient)
        && Objects.equals(this.id, somaticNgsReport.id)
        && Objects.equals(this.results, somaticNgsReport.results);
  }

  @Override
  public int hashCode() {
    return Objects.hash(metadata, type, specimen, issuedOn, patient, id, results);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SomaticNgsReport {\n");
    sb.append("    metadata: ").append(toIndentedString(metadata)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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

    private SomaticNgsReport instance;

    public Builder() {
      this(new SomaticNgsReport());
    }

    protected Builder(SomaticNgsReport instance) {
      this.instance = instance;
    }

    public SomaticNgsReport.Builder metadata(List<NgsReportMetadata> metadata) {
      this.instance.metadata = metadata;
      return this;
    }

    public SomaticNgsReport.Builder type(NgsReportTypeCoding type) {
      this.instance.type = type;
      return this;
    }

    public SomaticNgsReport.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public SomaticNgsReport.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public SomaticNgsReport.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public SomaticNgsReport.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public SomaticNgsReport.Builder results(SomaticNgsReportResults results) {
      this.instance.results = results;
      return this;
    }

    public SomaticNgsReport build() {
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

  public static SomaticNgsReport.Builder builder() {
    return new SomaticNgsReport.Builder();
  }

  public SomaticNgsReport.Builder toBuilder() {
    return new SomaticNgsReport.Builder()
        .metadata(getMetadata())
        .type(getType())
        .specimen(getSpecimen())
        .issuedOn(getIssuedOn())
        .patient(getPatient())
        .id(getId())
        .results(getResults());
  }
}
