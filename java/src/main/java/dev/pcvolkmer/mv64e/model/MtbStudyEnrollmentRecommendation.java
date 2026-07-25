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

public class MtbStudyEnrollmentRecommendation {
  public static final String JSON_PROPERTY_PRIORITY = "priority";

  private RecommendationPriorityCoding priority;

  public static final String JSON_PROPERTY_MEDICATION = "medication";

  private @Nullable Set<AtcUnregisteredMedicationCoding> medication;

  public static final String JSON_PROPERTY_SUPPORTING_VARIANTS = "supportingVariants";

  private @Nullable List<GeneAlterationReference> supportingVariants;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_SUPPORTING_FINDINGS = "supportingFindings";

  private @Nullable List<Reference> supportingFindings;

  public static final String JSON_PROPERTY_LEVEL_OF_EVIDENCE = "levelOfEvidence";

  private @Nullable LevelOfEvidence levelOfEvidence;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_STUDY = "study";

  private List<StudyReference> study;

  public MtbStudyEnrollmentRecommendation() {}

  public MtbStudyEnrollmentRecommendation priority(RecommendationPriorityCoding priority) {
    this.priority = priority;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PRIORITY, required = true)
  public RecommendationPriorityCoding getPriority() {
    return priority;
  }

  @JsonProperty(value = JSON_PROPERTY_PRIORITY, required = true)
  public void setPriority(RecommendationPriorityCoding priority) {
    this.priority = priority;
  }

  public MtbStudyEnrollmentRecommendation medication(
      @Nullable Set<AtcUnregisteredMedicationCoding> medication) {
    this.medication = medication;
    return this;
  }

  public MtbStudyEnrollmentRecommendation addMedicationItem(
      AtcUnregisteredMedicationCoding medicationItem) {
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

  public MtbStudyEnrollmentRecommendation supportingVariants(
      @Nullable List<GeneAlterationReference> supportingVariants) {
    this.supportingVariants = supportingVariants;
    return this;
  }

  public MtbStudyEnrollmentRecommendation addSupportingVariantsItem(
      GeneAlterationReference supportingVariantsItem) {
    if (this.supportingVariants == null) {
      this.supportingVariants = new ArrayList<>();
    }
    this.supportingVariants.add(supportingVariantsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SUPPORTING_VARIANTS, required = false)
  public @Nullable List<GeneAlterationReference> getSupportingVariants() {
    return supportingVariants;
  }

  @JsonProperty(value = JSON_PROPERTY_SUPPORTING_VARIANTS, required = false)
  public void setSupportingVariants(@Nullable List<GeneAlterationReference> supportingVariants) {
    this.supportingVariants = supportingVariants;
  }

  public MtbStudyEnrollmentRecommendation reason(@Nullable Reference reason) {
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

  public MtbStudyEnrollmentRecommendation patient(Reference patient) {
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

  public MtbStudyEnrollmentRecommendation supportingFindings(
      @Nullable List<Reference> supportingFindings) {
    this.supportingFindings = supportingFindings;
    return this;
  }

  public MtbStudyEnrollmentRecommendation addSupportingFindingsItem(
      Reference supportingFindingsItem) {
    if (this.supportingFindings == null) {
      this.supportingFindings = new ArrayList<>();
    }
    this.supportingFindings.add(supportingFindingsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SUPPORTING_FINDINGS, required = false)
  public @Nullable List<Reference> getSupportingFindings() {
    return supportingFindings;
  }

  @JsonProperty(value = JSON_PROPERTY_SUPPORTING_FINDINGS, required = false)
  public void setSupportingFindings(@Nullable List<Reference> supportingFindings) {
    this.supportingFindings = supportingFindings;
  }

  public MtbStudyEnrollmentRecommendation levelOfEvidence(
      @Nullable LevelOfEvidence levelOfEvidence) {
    this.levelOfEvidence = levelOfEvidence;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LEVEL_OF_EVIDENCE, required = false)
  public @Nullable LevelOfEvidence getLevelOfEvidence() {
    return levelOfEvidence;
  }

  @JsonProperty(value = JSON_PROPERTY_LEVEL_OF_EVIDENCE, required = false)
  public void setLevelOfEvidence(@Nullable LevelOfEvidence levelOfEvidence) {
    this.levelOfEvidence = levelOfEvidence;
  }

  public MtbStudyEnrollmentRecommendation id(String id) {
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

  public MtbStudyEnrollmentRecommendation issuedOn(Date issuedOn) {
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

  public MtbStudyEnrollmentRecommendation study(List<StudyReference> study) {
    this.study = study;
    return this;
  }

  public MtbStudyEnrollmentRecommendation addStudyItem(StudyReference studyItem) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbStudyEnrollmentRecommendation mtbStudyEnrollmentRecommendation =
        (MtbStudyEnrollmentRecommendation) o;
    return Objects.equals(this.priority, mtbStudyEnrollmentRecommendation.priority)
        && Objects.equals(this.medication, mtbStudyEnrollmentRecommendation.medication)
        && Objects.equals(
            this.supportingVariants, mtbStudyEnrollmentRecommendation.supportingVariants)
        && Objects.equals(this.reason, mtbStudyEnrollmentRecommendation.reason)
        && Objects.equals(this.patient, mtbStudyEnrollmentRecommendation.patient)
        && Objects.equals(
            this.supportingFindings, mtbStudyEnrollmentRecommendation.supportingFindings)
        && Objects.equals(this.levelOfEvidence, mtbStudyEnrollmentRecommendation.levelOfEvidence)
        && Objects.equals(this.id, mtbStudyEnrollmentRecommendation.id)
        && Objects.equals(this.issuedOn, mtbStudyEnrollmentRecommendation.issuedOn)
        && Objects.equals(this.study, mtbStudyEnrollmentRecommendation.study);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        priority,
        medication,
        supportingVariants,
        reason,
        patient,
        supportingFindings,
        levelOfEvidence,
        id,
        issuedOn,
        study);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbStudyEnrollmentRecommendation {\n");
    sb.append("    priority: ").append(toIndentedString(priority)).append("\n");
    sb.append("    medication: ").append(toIndentedString(medication)).append("\n");
    sb.append("    supportingVariants: ").append(toIndentedString(supportingVariants)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    supportingFindings: ").append(toIndentedString(supportingFindings)).append("\n");
    sb.append("    levelOfEvidence: ").append(toIndentedString(levelOfEvidence)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    study: ").append(toIndentedString(study)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbStudyEnrollmentRecommendation instance;

    public Builder() {
      this(new MtbStudyEnrollmentRecommendation());
    }

    protected Builder(MtbStudyEnrollmentRecommendation instance) {
      this.instance = instance;
    }

    public MtbStudyEnrollmentRecommendation.Builder priority(
        RecommendationPriorityCoding priority) {
      this.instance.priority = priority;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder medication(
        Set<AtcUnregisteredMedicationCoding> medication) {
      this.instance.medication = medication;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder supportingVariants(
        List<GeneAlterationReference> supportingVariants) {
      this.instance.supportingVariants = supportingVariants;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder reason(Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder supportingFindings(
        List<Reference> supportingFindings) {
      this.instance.supportingFindings = supportingFindings;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder levelOfEvidence(
        LevelOfEvidence levelOfEvidence) {
      this.instance.levelOfEvidence = levelOfEvidence;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public MtbStudyEnrollmentRecommendation.Builder study(List<StudyReference> study) {
      this.instance.study = study;
      return this;
    }

    public MtbStudyEnrollmentRecommendation build() {
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

  public static MtbStudyEnrollmentRecommendation.Builder builder() {
    return new MtbStudyEnrollmentRecommendation.Builder();
  }

  public MtbStudyEnrollmentRecommendation.Builder toBuilder() {
    return new MtbStudyEnrollmentRecommendation.Builder()
        .priority(getPriority())
        .medication(getMedication())
        .supportingVariants(getSupportingVariants())
        .reason(getReason())
        .patient(getPatient())
        .supportingFindings(getSupportingFindings())
        .levelOfEvidence(getLevelOfEvidence())
        .id(getId())
        .issuedOn(getIssuedOn())
        .study(getStudy());
  }
}
