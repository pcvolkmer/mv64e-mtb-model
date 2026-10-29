package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.ArrayList;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class StudyInclusion {
  public static final String JSON_PROPERTY_STATUS_REASON = "statusReason";

  private @Nullable MtbTherapyStatusReasonCoding statusReason;

  public static final String JSON_PROPERTY_MEDICATION = "medication";

  private @Nullable Set<AtcUnregisteredMedicationCoding> medication;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RECORDED_ON = "recordedOn";

  private Date recordedOn;

  public static final String JSON_PROPERTY_STATUS = "status";

  private TherapyStatusCoding status;

  public static final String JSON_PROPERTY_BASED_ON = "basedOn";

  private @Nullable Reference basedOn;

  public static final String JSON_PROPERTY_STUDY = "study";

  private List<StudyReference> study;

  public static final String JSON_PROPERTY_PERIOD = "period";

  private @Nullable PeriodDate period;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public StudyInclusion() {}

  public StudyInclusion statusReason(@Nullable MtbTherapyStatusReasonCoding statusReason) {

    this.statusReason = statusReason;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS_REASON, required = false)
  public @Nullable MtbTherapyStatusReasonCoding getStatusReason() {
    return statusReason;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS_REASON, required = false)
  public void setStatusReason(@Nullable MtbTherapyStatusReasonCoding statusReason) {
    this.statusReason = statusReason;
  }

  public StudyInclusion medication(@Nullable Set<AtcUnregisteredMedicationCoding> medication) {

    this.medication = medication;
    return this;
  }

  public StudyInclusion addMedicationItem(AtcUnregisteredMedicationCoding medicationItem) {
    if (this.medication == null) {
      this.medication = new LinkedHashSet<>();
    }
    this.medication.add(medicationItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MEDICATION, required = false)
  public @Nullable Set<AtcUnregisteredMedicationCoding> getMedication() {
    return medication;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_MEDICATION, required = false)
  public void setMedication(@Nullable Set<AtcUnregisteredMedicationCoding> medication) {
    this.medication = medication;
  }

  public StudyInclusion reason(@Nullable Reference reason) {

    this.reason = reason;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON, required = false)
  public @Nullable Reference getReason() {
    return reason;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON, required = false)
  public void setReason(@Nullable Reference reason) {
    this.reason = reason;
  }

  public StudyInclusion id(String id) {

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

  public StudyInclusion recordedOn(Date recordedOn) {

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

  public StudyInclusion status(TherapyStatusCoding status) {

    this.status = status;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS, required = true)
  public TherapyStatusCoding getStatus() {
    return status;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS, required = true)
  public void setStatus(TherapyStatusCoding status) {
    this.status = status;
  }

  public StudyInclusion basedOn(@Nullable Reference basedOn) {

    this.basedOn = basedOn;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_BASED_ON, required = false)
  public @Nullable Reference getBasedOn() {
    return basedOn;
  }

  @JsonProperty(value = JSON_PROPERTY_BASED_ON, required = false)
  public void setBasedOn(@Nullable Reference basedOn) {
    this.basedOn = basedOn;
  }

  public StudyInclusion study(List<StudyReference> study) {

    this.study = study;
    return this;
  }

  public StudyInclusion addStudyItem(StudyReference studyItem) {
    if (this.study == null) {
      this.study = new ArrayList<>();
    }
    this.study.add(studyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STUDY, required = true)
  public List<StudyReference> getStudy() {
    return study;
  }

  @JsonProperty(value = JSON_PROPERTY_STUDY, required = true)
  public void setStudy(List<StudyReference> study) {
    this.study = study;
  }

  public StudyInclusion period(@Nullable PeriodDate period) {

    this.period = period;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PERIOD, required = false)
  public @Nullable PeriodDate getPeriod() {
    return period;
  }

  @JsonProperty(value = JSON_PROPERTY_PERIOD, required = false)
  public void setPeriod(@Nullable PeriodDate period) {
    this.period = period;
  }

  public StudyInclusion patient(Reference patient) {

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
    StudyInclusion studyInclusion = (StudyInclusion) o;
    return Objects.equals(this.statusReason, studyInclusion.statusReason)
        && Objects.equals(this.medication, studyInclusion.medication)
        && Objects.equals(this.reason, studyInclusion.reason)
        && Objects.equals(this.id, studyInclusion.id)
        && Objects.equals(this.recordedOn, studyInclusion.recordedOn)
        && Objects.equals(this.status, studyInclusion.status)
        && Objects.equals(this.basedOn, studyInclusion.basedOn)
        && Objects.equals(this.study, studyInclusion.study)
        && Objects.equals(this.period, studyInclusion.period)
        && Objects.equals(this.patient, studyInclusion.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        statusReason, medication, reason, id, recordedOn, status, basedOn, study, period, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class StudyInclusion {\n");
    sb.append("    statusReason: ").append(toIndentedString(statusReason)).append("\n");
    sb.append("    medication: ").append(toIndentedString(medication)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    recordedOn: ").append(toIndentedString(recordedOn)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    basedOn: ").append(toIndentedString(basedOn)).append("\n");
    sb.append("    study: ").append(toIndentedString(study)).append("\n");
    sb.append("    period: ").append(toIndentedString(period)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private StudyInclusion instance;

    public Builder() {
      this(new StudyInclusion());
    }

    protected Builder(StudyInclusion instance) {
      this.instance = instance;
    }

    public StudyInclusion.Builder statusReason(
        @Nullable MtbTherapyStatusReasonCoding statusReason) {
      this.instance.statusReason = statusReason;
      return this;
    }

    public StudyInclusion.Builder medication(
        @Nullable Set<AtcUnregisteredMedicationCoding> medication) {
      this.instance.medication = medication;
      return this;
    }

    public StudyInclusion.Builder reason(@Nullable Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public StudyInclusion.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public StudyInclusion.Builder recordedOn(Date recordedOn) {
      this.instance.recordedOn = recordedOn;
      return this;
    }

    public StudyInclusion.Builder status(TherapyStatusCoding status) {
      this.instance.status = status;
      return this;
    }

    public StudyInclusion.Builder basedOn(@Nullable Reference basedOn) {
      this.instance.basedOn = basedOn;
      return this;
    }

    public StudyInclusion.Builder study(List<StudyReference> study) {
      this.instance.study = study;
      return this;
    }

    public StudyInclusion.Builder period(@Nullable PeriodDate period) {
      this.instance.period = period;
      return this;
    }

    public StudyInclusion.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public StudyInclusion build() {
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

  public static StudyInclusion.Builder builder() {
    return new StudyInclusion.Builder();
  }

  public StudyInclusion.Builder toBuilder() {
    return new StudyInclusion.Builder()
        .statusReason(getStatusReason())
        .medication(getMedication())
        .reason(getReason())
        .id(getId())
        .recordedOn(getRecordedOn())
        .status(getStatus())
        .basedOn(getBasedOn())
        .study(getStudy())
        .period(getPeriod())
        .patient(getPatient());
  }
}
