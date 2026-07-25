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

public class MtbSystemicTherapy {
  public static final String JSON_PROPERTY_STATUS_REASON = "statusReason";

  private @Nullable MtbTherapyStatusReasonCoding statusReason;

  public static final String JSON_PROPERTY_MEDICATION = "medication";

  private @Nullable Set<AtcUnregisteredMedicationCoding> medication;

  public static final String JSON_PROPERTY_INTENT = "intent";

  private @Nullable MtbTherapyIntentCoding intent;

  public static final String JSON_PROPERTY_THERAPY_LINE = "therapyLine";

  private @Nullable Integer therapyLine;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_RECOMMENDATION_FULFILLMENT_STATUS =
      "recommendationFulfillmentStatus";

  private @Nullable MtbSystemicTherapyRecommendationFulfillmentStatusCoding
      recommendationFulfillmentStatus;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RECORDED_ON = "recordedOn";

  private Date recordedOn;

  public static final String JSON_PROPERTY_STATUS = "status";

  private TherapyStatusCoding status;

  public static final String JSON_PROPERTY_NOTES = "notes";

  private @Nullable List<String> notes;

  public static final String JSON_PROPERTY_DOSAGE = "dosage";

  private @Nullable MtbSystemicTherapyDosageDensityCoding dosage;

  public static final String JSON_PROPERTY_PERIOD = "period";

  private @Nullable PeriodDate period;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_CATEGORY = "category";

  private @Nullable MtbSystemicTherapyCategoryCoding category;

  public static final String JSON_PROPERTY_BASED_ON = "basedOn";

  private @Nullable Reference basedOn;

  public MtbSystemicTherapy() {}

  public MtbSystemicTherapy statusReason(@Nullable MtbTherapyStatusReasonCoding statusReason) {
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

  public MtbSystemicTherapy medication(@Nullable Set<AtcUnregisteredMedicationCoding> medication) {
    this.medication = medication;
    return this;
  }

  public MtbSystemicTherapy addMedicationItem(AtcUnregisteredMedicationCoding medicationItem) {
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

  public MtbSystemicTherapy intent(@Nullable MtbTherapyIntentCoding intent) {
    this.intent = intent;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_INTENT, required = false)
  public @Nullable MtbTherapyIntentCoding getIntent() {
    return intent;
  }

  @JsonProperty(value = JSON_PROPERTY_INTENT, required = false)
  public void setIntent(@Nullable MtbTherapyIntentCoding intent) {
    this.intent = intent;
  }

  public MtbSystemicTherapy therapyLine(@Nullable Integer therapyLine) {
    this.therapyLine = therapyLine;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_THERAPY_LINE, required = false)
  public @Nullable Integer getTherapyLine() {
    return therapyLine;
  }

  @JsonProperty(value = JSON_PROPERTY_THERAPY_LINE, required = false)
  public void setTherapyLine(@Nullable Integer therapyLine) {
    this.therapyLine = therapyLine;
  }

  public MtbSystemicTherapy reason(@Nullable Reference reason) {
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

  public MtbSystemicTherapy recommendationFulfillmentStatus(
      @Nullable MtbSystemicTherapyRecommendationFulfillmentStatusCoding
          recommendationFulfillmentStatus) {
    this.recommendationFulfillmentStatus = recommendationFulfillmentStatus;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATION_FULFILLMENT_STATUS, required = false)
  public @Nullable MtbSystemicTherapyRecommendationFulfillmentStatusCoding
      getRecommendationFulfillmentStatus() {
    return recommendationFulfillmentStatus;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATION_FULFILLMENT_STATUS, required = false)
  public void setRecommendationFulfillmentStatus(
      @Nullable MtbSystemicTherapyRecommendationFulfillmentStatusCoding
          recommendationFulfillmentStatus) {
    this.recommendationFulfillmentStatus = recommendationFulfillmentStatus;
  }

  public MtbSystemicTherapy id(String id) {
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

  public MtbSystemicTherapy recordedOn(Date recordedOn) {
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

  public MtbSystemicTherapy status(TherapyStatusCoding status) {
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

  public MtbSystemicTherapy notes(@Nullable List<String> notes) {
    this.notes = notes;
    return this;
  }

  public MtbSystemicTherapy addNotesItem(String notesItem) {
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

  public MtbSystemicTherapy dosage(@Nullable MtbSystemicTherapyDosageDensityCoding dosage) {
    this.dosage = dosage;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DOSAGE, required = false)
  public @Nullable MtbSystemicTherapyDosageDensityCoding getDosage() {
    return dosage;
  }

  @JsonProperty(value = JSON_PROPERTY_DOSAGE, required = false)
  public void setDosage(@Nullable MtbSystemicTherapyDosageDensityCoding dosage) {
    this.dosage = dosage;
  }

  public MtbSystemicTherapy period(@Nullable PeriodDate period) {
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

  public MtbSystemicTherapy patient(Reference patient) {
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

  public MtbSystemicTherapy category(@Nullable MtbSystemicTherapyCategoryCoding category) {
    this.category = category;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CATEGORY, required = false)
  public @Nullable MtbSystemicTherapyCategoryCoding getCategory() {
    return category;
  }

  @JsonProperty(value = JSON_PROPERTY_CATEGORY, required = false)
  public void setCategory(@Nullable MtbSystemicTherapyCategoryCoding category) {
    this.category = category;
  }

  public MtbSystemicTherapy basedOn(@Nullable Reference basedOn) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbSystemicTherapy mtbSystemicTherapy = (MtbSystemicTherapy) o;
    return Objects.equals(this.statusReason, mtbSystemicTherapy.statusReason)
        && Objects.equals(this.medication, mtbSystemicTherapy.medication)
        && Objects.equals(this.intent, mtbSystemicTherapy.intent)
        && Objects.equals(this.therapyLine, mtbSystemicTherapy.therapyLine)
        && Objects.equals(this.reason, mtbSystemicTherapy.reason)
        && Objects.equals(
            this.recommendationFulfillmentStatus,
            mtbSystemicTherapy.recommendationFulfillmentStatus)
        && Objects.equals(this.id, mtbSystemicTherapy.id)
        && Objects.equals(this.recordedOn, mtbSystemicTherapy.recordedOn)
        && Objects.equals(this.status, mtbSystemicTherapy.status)
        && Objects.equals(this.notes, mtbSystemicTherapy.notes)
        && Objects.equals(this.dosage, mtbSystemicTherapy.dosage)
        && Objects.equals(this.period, mtbSystemicTherapy.period)
        && Objects.equals(this.patient, mtbSystemicTherapy.patient)
        && Objects.equals(this.category, mtbSystemicTherapy.category)
        && Objects.equals(this.basedOn, mtbSystemicTherapy.basedOn);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        statusReason,
        medication,
        intent,
        therapyLine,
        reason,
        recommendationFulfillmentStatus,
        id,
        recordedOn,
        status,
        notes,
        dosage,
        period,
        patient,
        category,
        basedOn);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbSystemicTherapy {\n");
    sb.append("    statusReason: ").append(toIndentedString(statusReason)).append("\n");
    sb.append("    medication: ").append(toIndentedString(medication)).append("\n");
    sb.append("    intent: ").append(toIndentedString(intent)).append("\n");
    sb.append("    therapyLine: ").append(toIndentedString(therapyLine)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    recommendationFulfillmentStatus: ")
        .append(toIndentedString(recommendationFulfillmentStatus))
        .append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    recordedOn: ").append(toIndentedString(recordedOn)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    notes: ").append(toIndentedString(notes)).append("\n");
    sb.append("    dosage: ").append(toIndentedString(dosage)).append("\n");
    sb.append("    period: ").append(toIndentedString(period)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    basedOn: ").append(toIndentedString(basedOn)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbSystemicTherapy instance;

    public Builder() {
      this(new MtbSystemicTherapy());
    }

    protected Builder(MtbSystemicTherapy instance) {
      this.instance = instance;
    }

    public MtbSystemicTherapy.Builder statusReason(MtbTherapyStatusReasonCoding statusReason) {
      this.instance.statusReason = statusReason;
      return this;
    }

    public MtbSystemicTherapy.Builder medication(Set<AtcUnregisteredMedicationCoding> medication) {
      this.instance.medication = medication;
      return this;
    }

    public MtbSystemicTherapy.Builder intent(MtbTherapyIntentCoding intent) {
      this.instance.intent = intent;
      return this;
    }

    public MtbSystemicTherapy.Builder therapyLine(Integer therapyLine) {
      this.instance.therapyLine = therapyLine;
      return this;
    }

    public MtbSystemicTherapy.Builder reason(Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public MtbSystemicTherapy.Builder recommendationFulfillmentStatus(
        MtbSystemicTherapyRecommendationFulfillmentStatusCoding recommendationFulfillmentStatus) {
      this.instance.recommendationFulfillmentStatus = recommendationFulfillmentStatus;
      return this;
    }

    public MtbSystemicTherapy.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbSystemicTherapy.Builder recordedOn(Date recordedOn) {
      this.instance.recordedOn = recordedOn;
      return this;
    }

    public MtbSystemicTherapy.Builder status(TherapyStatusCoding status) {
      this.instance.status = status;
      return this;
    }

    public MtbSystemicTherapy.Builder notes(List<String> notes) {
      this.instance.notes = notes;
      return this;
    }

    public MtbSystemicTherapy.Builder dosage(MtbSystemicTherapyDosageDensityCoding dosage) {
      this.instance.dosage = dosage;
      return this;
    }

    public MtbSystemicTherapy.Builder period(PeriodDate period) {
      this.instance.period = period;
      return this;
    }

    public MtbSystemicTherapy.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbSystemicTherapy.Builder category(MtbSystemicTherapyCategoryCoding category) {
      this.instance.category = category;
      return this;
    }

    public MtbSystemicTherapy.Builder basedOn(Reference basedOn) {
      this.instance.basedOn = basedOn;
      return this;
    }

    public MtbSystemicTherapy build() {
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

  public static MtbSystemicTherapy.Builder builder() {
    return new MtbSystemicTherapy.Builder();
  }

  public MtbSystemicTherapy.Builder toBuilder() {
    return new MtbSystemicTherapy.Builder()
        .statusReason(getStatusReason())
        .medication(getMedication())
        .intent(getIntent())
        .therapyLine(getTherapyLine())
        .reason(getReason())
        .recommendationFulfillmentStatus(getRecommendationFulfillmentStatus())
        .id(getId())
        .recordedOn(getRecordedOn())
        .status(getStatus())
        .notes(getNotes())
        .dosage(getDosage())
        .period(getPeriod())
        .patient(getPatient())
        .category(getCategory())
        .basedOn(getBasedOn());
  }
}
