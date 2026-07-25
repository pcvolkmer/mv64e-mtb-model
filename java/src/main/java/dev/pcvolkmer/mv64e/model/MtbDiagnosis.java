package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.ArrayList;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class MtbDiagnosis {
  public static final String JSON_PROPERTY_GRADING = "grading";

  private @Nullable MtbDiagnosisGrading grading;

  public static final String JSON_PROPERTY_GERMLINE_CODES = "germlineCodes";

  private @Nullable Set<Coding> germlineCodes;

  public static final String JSON_PROPERTY_CODE = "code";

  private Coding code;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RECORDED_ON = "recordedOn";

  private Date recordedOn;

  public static final String JSON_PROPERTY_TYPE = "type";

  private MtbDiagnosisType type;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_TOPOGRAPHY = "topography";

  private Coding topography;

  public static final String JSON_PROPERTY_STAGING = "staging";

  private @Nullable MtbDiagnosisStaging staging;

  public static final String JSON_PROPERTY_HISTOLOGY = "histology";

  private @Nullable List<Reference> histology;

  public static final String JSON_PROPERTY_GUIDELINE_TREATMENT_STATUS = "guidelineTreatmentStatus";

  private @Nullable MtbDiagnosisGuidelineTreatmentStatusCoding guidelineTreatmentStatus;

  public static final String JSON_PROPERTY_NOTES = "notes";

  private @Nullable List<String> notes;

  public MtbDiagnosis() {}

  public MtbDiagnosis grading(@Nullable MtbDiagnosisGrading grading) {
    this.grading = grading;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GRADING, required = false)
  public @Nullable MtbDiagnosisGrading getGrading() {
    return grading;
  }

  @JsonProperty(value = JSON_PROPERTY_GRADING, required = false)
  public void setGrading(@Nullable MtbDiagnosisGrading grading) {
    this.grading = grading;
  }

  public MtbDiagnosis germlineCodes(@Nullable Set<Coding> germlineCodes) {
    this.germlineCodes = germlineCodes;
    return this;
  }

  public MtbDiagnosis addGermlineCodesItem(Coding germlineCodesItem) {
    if (this.germlineCodes == null) {
      this.germlineCodes = new LinkedHashSet<>();
    }
    this.germlineCodes.add(germlineCodesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GERMLINE_CODES, required = false)
  public @Nullable Set<Coding> getGermlineCodes() {
    return germlineCodes;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_GERMLINE_CODES, required = false)
  public void setGermlineCodes(@Nullable Set<Coding> germlineCodes) {
    this.germlineCodes = germlineCodes;
  }

  public MtbDiagnosis code(Coding code) {
    this.code = code;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public Coding getCode() {
    return code;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public void setCode(Coding code) {
    this.code = code;
  }

  public MtbDiagnosis id(String id) {
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

  public MtbDiagnosis recordedOn(Date recordedOn) {
    this.recordedOn = recordedOn;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RECORDED_ON, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getRecordedOn() {
    return recordedOn;
  }

  @JsonProperty(value = JSON_PROPERTY_RECORDED_ON, required = true)
  public void setRecordedOn(Date recordedOn) {
    this.recordedOn = recordedOn;
  }

  public MtbDiagnosis type(MtbDiagnosisType type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public MtbDiagnosisType getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public void setType(MtbDiagnosisType type) {
    this.type = type;
  }

  public MtbDiagnosis patient(Reference patient) {
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

  public MtbDiagnosis topography(Coding topography) {
    this.topography = topography;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TOPOGRAPHY, required = true)
  public Coding getTopography() {
    return topography;
  }

  @JsonProperty(value = JSON_PROPERTY_TOPOGRAPHY, required = true)
  public void setTopography(Coding topography) {
    this.topography = topography;
  }

  public MtbDiagnosis staging(@Nullable MtbDiagnosisStaging staging) {
    this.staging = staging;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STAGING, required = false)
  public @Nullable MtbDiagnosisStaging getStaging() {
    return staging;
  }

  @JsonProperty(value = JSON_PROPERTY_STAGING, required = false)
  public void setStaging(@Nullable MtbDiagnosisStaging staging) {
    this.staging = staging;
  }

  public MtbDiagnosis histology(@Nullable List<Reference> histology) {
    this.histology = histology;
    return this;
  }

  public MtbDiagnosis addHistologyItem(Reference histologyItem) {
    if (this.histology == null) {
      this.histology = new ArrayList<>();
    }
    this.histology.add(histologyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY, required = false)
  public @Nullable List<Reference> getHistology() {
    return histology;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY, required = false)
  public void setHistology(@Nullable List<Reference> histology) {
    this.histology = histology;
  }

  public MtbDiagnosis guidelineTreatmentStatus(
      @Nullable MtbDiagnosisGuidelineTreatmentStatusCoding guidelineTreatmentStatus) {
    this.guidelineTreatmentStatus = guidelineTreatmentStatus;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_TREATMENT_STATUS, required = false)
  public @Nullable MtbDiagnosisGuidelineTreatmentStatusCoding getGuidelineTreatmentStatus() {
    return guidelineTreatmentStatus;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_TREATMENT_STATUS, required = false)
  public void setGuidelineTreatmentStatus(
      @Nullable MtbDiagnosisGuidelineTreatmentStatusCoding guidelineTreatmentStatus) {
    this.guidelineTreatmentStatus = guidelineTreatmentStatus;
  }

  public MtbDiagnosis notes(@Nullable List<String> notes) {
    this.notes = notes;
    return this;
  }

  public MtbDiagnosis addNotesItem(String notesItem) {
    if (this.notes == null) {
      this.notes = new ArrayList<>();
    }
    this.notes.add(notesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_NOTES, required = false)
  public @Nullable List<String> getNotes() {
    return notes;
  }

  @JsonProperty(value = JSON_PROPERTY_NOTES, required = false)
  public void setNotes(@Nullable List<String> notes) {
    this.notes = notes;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbDiagnosis mtbDiagnosis = (MtbDiagnosis) o;
    return Objects.equals(this.grading, mtbDiagnosis.grading)
        && Objects.equals(this.germlineCodes, mtbDiagnosis.germlineCodes)
        && Objects.equals(this.code, mtbDiagnosis.code)
        && Objects.equals(this.id, mtbDiagnosis.id)
        && Objects.equals(this.recordedOn, mtbDiagnosis.recordedOn)
        && Objects.equals(this.type, mtbDiagnosis.type)
        && Objects.equals(this.patient, mtbDiagnosis.patient)
        && Objects.equals(this.topography, mtbDiagnosis.topography)
        && Objects.equals(this.staging, mtbDiagnosis.staging)
        && Objects.equals(this.histology, mtbDiagnosis.histology)
        && Objects.equals(this.guidelineTreatmentStatus, mtbDiagnosis.guidelineTreatmentStatus)
        && Objects.equals(this.notes, mtbDiagnosis.notes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        grading,
        germlineCodes,
        code,
        id,
        recordedOn,
        type,
        patient,
        topography,
        staging,
        histology,
        guidelineTreatmentStatus,
        notes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbDiagnosis {\n");
    sb.append("    grading: ").append(toIndentedString(grading)).append("\n");
    sb.append("    germlineCodes: ").append(toIndentedString(germlineCodes)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    recordedOn: ").append(toIndentedString(recordedOn)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    topography: ").append(toIndentedString(topography)).append("\n");
    sb.append("    staging: ").append(toIndentedString(staging)).append("\n");
    sb.append("    histology: ").append(toIndentedString(histology)).append("\n");
    sb.append("    guidelineTreatmentStatus: ")
        .append(toIndentedString(guidelineTreatmentStatus))
        .append("\n");
    sb.append("    notes: ").append(toIndentedString(notes)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbDiagnosis instance;

    public Builder() {
      this(new MtbDiagnosis());
    }

    protected Builder(MtbDiagnosis instance) {
      this.instance = instance;
    }

    public MtbDiagnosis.Builder grading(MtbDiagnosisGrading grading) {
      this.instance.grading = grading;
      return this;
    }

    public MtbDiagnosis.Builder germlineCodes(Set<Coding> germlineCodes) {
      this.instance.germlineCodes = germlineCodes;
      return this;
    }

    public MtbDiagnosis.Builder code(Coding code) {
      this.instance.code = code;
      return this;
    }

    public MtbDiagnosis.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbDiagnosis.Builder recordedOn(Date recordedOn) {
      this.instance.recordedOn = recordedOn;
      return this;
    }

    public MtbDiagnosis.Builder type(MtbDiagnosisType type) {
      this.instance.type = type;
      return this;
    }

    public MtbDiagnosis.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbDiagnosis.Builder topography(Coding topography) {
      this.instance.topography = topography;
      return this;
    }

    public MtbDiagnosis.Builder staging(MtbDiagnosisStaging staging) {
      this.instance.staging = staging;
      return this;
    }

    public MtbDiagnosis.Builder histology(List<Reference> histology) {
      this.instance.histology = histology;
      return this;
    }

    public MtbDiagnosis.Builder guidelineTreatmentStatus(
        MtbDiagnosisGuidelineTreatmentStatusCoding guidelineTreatmentStatus) {
      this.instance.guidelineTreatmentStatus = guidelineTreatmentStatus;
      return this;
    }

    public MtbDiagnosis.Builder notes(List<String> notes) {
      this.instance.notes = notes;
      return this;
    }

    public MtbDiagnosis build() {
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

  public static MtbDiagnosis.Builder builder() {
    return new MtbDiagnosis.Builder();
  }

  public MtbDiagnosis.Builder toBuilder() {
    return new MtbDiagnosis.Builder()
        .grading(getGrading())
        .germlineCodes(getGermlineCodes())
        .code(getCode())
        .id(getId())
        .recordedOn(getRecordedOn())
        .type(getType())
        .patient(getPatient())
        .topography(getTopography())
        .staging(getStaging())
        .histology(getHistology())
        .guidelineTreatmentStatus(getGuidelineTreatmentStatus())
        .notes(getNotes());
  }
}
