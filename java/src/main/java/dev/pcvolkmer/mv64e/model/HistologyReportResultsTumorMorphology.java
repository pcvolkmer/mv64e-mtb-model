package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class HistologyReportResultsTumorMorphology {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_NOTE = "note";

  private @Nullable String note;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private Coding value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public HistologyReportResultsTumorMorphology() {}

  public HistologyReportResultsTumorMorphology id(String id) {
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

  public HistologyReportResultsTumorMorphology note(@Nullable String note) {
    this.note = note;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_NOTE, required = false)
  public @Nullable String getNote() {
    return note;
  }

  @JsonProperty(value = JSON_PROPERTY_NOTE, required = false)
  public void setNote(@Nullable String note) {
    this.note = note;
  }

  public HistologyReportResultsTumorMorphology specimen(Reference specimen) {
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

  public HistologyReportResultsTumorMorphology value(Coding value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public Coding getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(Coding value) {
    this.value = value;
  }

  public HistologyReportResultsTumorMorphology patient(Reference patient) {
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
    HistologyReportResultsTumorMorphology histologyReportResultsTumorMorphology =
        (HistologyReportResultsTumorMorphology) o;
    return Objects.equals(this.id, histologyReportResultsTumorMorphology.id)
        && Objects.equals(this.note, histologyReportResultsTumorMorphology.note)
        && Objects.equals(this.specimen, histologyReportResultsTumorMorphology.specimen)
        && Objects.equals(this.value, histologyReportResultsTumorMorphology.value)
        && Objects.equals(this.patient, histologyReportResultsTumorMorphology.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, note, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HistologyReportResultsTumorMorphology {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    note: ").append(toIndentedString(note)).append("\n");
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

    private HistologyReportResultsTumorMorphology instance;

    public Builder() {
      this(new HistologyReportResultsTumorMorphology());
    }

    protected Builder(HistologyReportResultsTumorMorphology instance) {
      this.instance = instance;
    }

    public HistologyReportResultsTumorMorphology.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public HistologyReportResultsTumorMorphology.Builder note(String note) {
      this.instance.note = note;
      return this;
    }

    public HistologyReportResultsTumorMorphology.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public HistologyReportResultsTumorMorphology.Builder value(Coding value) {
      this.instance.value = value;
      return this;
    }

    public HistologyReportResultsTumorMorphology.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public HistologyReportResultsTumorMorphology build() {
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

  public static HistologyReportResultsTumorMorphology.Builder builder() {
    return new HistologyReportResultsTumorMorphology.Builder();
  }

  public HistologyReportResultsTumorMorphology.Builder toBuilder() {
    return new HistologyReportResultsTumorMorphology.Builder()
        .id(getId())
        .note(getNote())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
