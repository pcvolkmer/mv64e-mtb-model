package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MtbCarePlan {
  public static final String JSON_PROPERTY_BOARD_TYPE = "boardType";

  private @Nullable CarePlanBoardTypeCoding boardType;

  public static final String JSON_PROPERTY_RECOMMENDATIONS_MISSING_REASON =
      "recommendationsMissingReason";

  private @Nullable MtbCarePlanRecommendationsMissingReasonCoding recommendationsMissingReason;

  public static final String JSON_PROPERTY_STUDY_ENROLLMENT_RECOMMENDATIONS =
      "studyEnrollmentRecommendations";

  private @Nullable List<MtbStudyEnrollmentRecommendation> studyEnrollmentRecommendations;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_GENETIC_COUNSELING_RECOMMENDATION =
      "geneticCounselingRecommendation";

  private @Nullable GeneticCounselingRecommendation geneticCounselingRecommendation;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_MEDICATION_RECOMMENDATIONS = "medicationRecommendations";

  private @Nullable List<MtbMedicationRecommendation> medicationRecommendations;

  public static final String JSON_PROPERTY_REBIOPSY_REQUESTS = "rebiopsyRequests";

  private @Nullable List<MtbCarePlanRebiopsyRequestsInner> rebiopsyRequests;

  public static final String JSON_PROPERTY_HISTOLOGY_REEVALUATION_REQUESTS =
      "histologyReevaluationRequests";

  private @Nullable List<MtbCarePlanHistologyReevaluationRequestsInner>
      histologyReevaluationRequests;

  public static final String JSON_PROPERTY_NOTES = "notes";

  private @Nullable List<String> notes;

  public static final String JSON_PROPERTY_PROCEDURE_RECOMMENDATIONS = "procedureRecommendations";

  private @Nullable List<MtbCarePlanProcedureRecommendationsInner> procedureRecommendations;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_NO_SEQUENCING_PERFORMED_REASON =
      "noSequencingPerformedReason";

  private @Nullable CarePlanNoSequencingPerformedReasonCoding noSequencingPerformedReason;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public MtbCarePlan() {}

  public MtbCarePlan boardType(@Nullable CarePlanBoardTypeCoding boardType) {
    this.boardType = boardType;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_BOARD_TYPE, required = false)
  public @Nullable CarePlanBoardTypeCoding getBoardType() {
    return boardType;
  }

  @JsonProperty(value = JSON_PROPERTY_BOARD_TYPE, required = false)
  public void setBoardType(@Nullable CarePlanBoardTypeCoding boardType) {
    this.boardType = boardType;
  }

  public MtbCarePlan recommendationsMissingReason(
      @Nullable MtbCarePlanRecommendationsMissingReasonCoding recommendationsMissingReason) {
    this.recommendationsMissingReason = recommendationsMissingReason;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATIONS_MISSING_REASON, required = false)
  public @Nullable MtbCarePlanRecommendationsMissingReasonCoding getRecommendationsMissingReason() {
    return recommendationsMissingReason;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATIONS_MISSING_REASON, required = false)
  public void setRecommendationsMissingReason(
      @Nullable MtbCarePlanRecommendationsMissingReasonCoding recommendationsMissingReason) {
    this.recommendationsMissingReason = recommendationsMissingReason;
  }

  public MtbCarePlan studyEnrollmentRecommendations(
      @Nullable List<MtbStudyEnrollmentRecommendation> studyEnrollmentRecommendations) {
    this.studyEnrollmentRecommendations = studyEnrollmentRecommendations;
    return this;
  }

  public MtbCarePlan addStudyEnrollmentRecommendationsItem(
      MtbStudyEnrollmentRecommendation studyEnrollmentRecommendationsItem) {
    if (this.studyEnrollmentRecommendations == null) {
      this.studyEnrollmentRecommendations = new ArrayList<>();
    }
    this.studyEnrollmentRecommendations.add(studyEnrollmentRecommendationsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STUDY_ENROLLMENT_RECOMMENDATIONS, required = false)
  public @Nullable List<MtbStudyEnrollmentRecommendation> getStudyEnrollmentRecommendations() {
    return studyEnrollmentRecommendations;
  }

  @JsonProperty(value = JSON_PROPERTY_STUDY_ENROLLMENT_RECOMMENDATIONS, required = false)
  public void setStudyEnrollmentRecommendations(
      @Nullable List<MtbStudyEnrollmentRecommendation> studyEnrollmentRecommendations) {
    this.studyEnrollmentRecommendations = studyEnrollmentRecommendations;
  }

  public MtbCarePlan reason(@Nullable Reference reason) {
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

  public MtbCarePlan geneticCounselingRecommendation(
      @Nullable GeneticCounselingRecommendation geneticCounselingRecommendation) {
    this.geneticCounselingRecommendation = geneticCounselingRecommendation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GENETIC_COUNSELING_RECOMMENDATION, required = false)
  public @Nullable GeneticCounselingRecommendation getGeneticCounselingRecommendation() {
    return geneticCounselingRecommendation;
  }

  @JsonProperty(value = JSON_PROPERTY_GENETIC_COUNSELING_RECOMMENDATION, required = false)
  public void setGeneticCounselingRecommendation(
      @Nullable GeneticCounselingRecommendation geneticCounselingRecommendation) {
    this.geneticCounselingRecommendation = geneticCounselingRecommendation;
  }

  public MtbCarePlan id(String id) {
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

  public MtbCarePlan medicationRecommendations(
      @Nullable List<MtbMedicationRecommendation> medicationRecommendations) {
    this.medicationRecommendations = medicationRecommendations;
    return this;
  }

  public MtbCarePlan addMedicationRecommendationsItem(
      MtbMedicationRecommendation medicationRecommendationsItem) {
    if (this.medicationRecommendations == null) {
      this.medicationRecommendations = new ArrayList<>();
    }
    this.medicationRecommendations.add(medicationRecommendationsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MEDICATION_RECOMMENDATIONS, required = false)
  public @Nullable List<MtbMedicationRecommendation> getMedicationRecommendations() {
    return medicationRecommendations;
  }

  @JsonProperty(value = JSON_PROPERTY_MEDICATION_RECOMMENDATIONS, required = false)
  public void setMedicationRecommendations(
      @Nullable List<MtbMedicationRecommendation> medicationRecommendations) {
    this.medicationRecommendations = medicationRecommendations;
  }

  public MtbCarePlan rebiopsyRequests(
      @Nullable List<MtbCarePlanRebiopsyRequestsInner> rebiopsyRequests) {
    this.rebiopsyRequests = rebiopsyRequests;
    return this;
  }

  public MtbCarePlan addRebiopsyRequestsItem(
      MtbCarePlanRebiopsyRequestsInner rebiopsyRequestsItem) {
    if (this.rebiopsyRequests == null) {
      this.rebiopsyRequests = new ArrayList<>();
    }
    this.rebiopsyRequests.add(rebiopsyRequestsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REBIOPSY_REQUESTS, required = false)
  public @Nullable List<MtbCarePlanRebiopsyRequestsInner> getRebiopsyRequests() {
    return rebiopsyRequests;
  }

  @JsonProperty(value = JSON_PROPERTY_REBIOPSY_REQUESTS, required = false)
  public void setRebiopsyRequests(
      @Nullable List<MtbCarePlanRebiopsyRequestsInner> rebiopsyRequests) {
    this.rebiopsyRequests = rebiopsyRequests;
  }

  public MtbCarePlan histologyReevaluationRequests(
      @Nullable List<MtbCarePlanHistologyReevaluationRequestsInner> histologyReevaluationRequests) {
    this.histologyReevaluationRequests = histologyReevaluationRequests;
    return this;
  }

  public MtbCarePlan addHistologyReevaluationRequestsItem(
      MtbCarePlanHistologyReevaluationRequestsInner histologyReevaluationRequestsItem) {
    if (this.histologyReevaluationRequests == null) {
      this.histologyReevaluationRequests = new ArrayList<>();
    }
    this.histologyReevaluationRequests.add(histologyReevaluationRequestsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY_REEVALUATION_REQUESTS, required = false)
  public @Nullable List<MtbCarePlanHistologyReevaluationRequestsInner>
      getHistologyReevaluationRequests() {
    return histologyReevaluationRequests;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY_REEVALUATION_REQUESTS, required = false)
  public void setHistologyReevaluationRequests(
      @Nullable List<MtbCarePlanHistologyReevaluationRequestsInner> histologyReevaluationRequests) {
    this.histologyReevaluationRequests = histologyReevaluationRequests;
  }

  public MtbCarePlan notes(@Nullable List<String> notes) {
    this.notes = notes;
    return this;
  }

  public MtbCarePlan addNotesItem(String notesItem) {
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

  public MtbCarePlan procedureRecommendations(
      @Nullable List<MtbCarePlanProcedureRecommendationsInner> procedureRecommendations) {
    this.procedureRecommendations = procedureRecommendations;
    return this;
  }

  public MtbCarePlan addProcedureRecommendationsItem(
      MtbCarePlanProcedureRecommendationsInner procedureRecommendationsItem) {
    if (this.procedureRecommendations == null) {
      this.procedureRecommendations = new ArrayList<>();
    }
    this.procedureRecommendations.add(procedureRecommendationsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PROCEDURE_RECOMMENDATIONS, required = false)
  public @Nullable List<MtbCarePlanProcedureRecommendationsInner> getProcedureRecommendations() {
    return procedureRecommendations;
  }

  @JsonProperty(value = JSON_PROPERTY_PROCEDURE_RECOMMENDATIONS, required = false)
  public void setProcedureRecommendations(
      @Nullable List<MtbCarePlanProcedureRecommendationsInner> procedureRecommendations) {
    this.procedureRecommendations = procedureRecommendations;
  }

  public MtbCarePlan issuedOn(Date issuedOn) {
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

  public MtbCarePlan noSequencingPerformedReason(
      @Nullable CarePlanNoSequencingPerformedReasonCoding noSequencingPerformedReason) {
    this.noSequencingPerformedReason = noSequencingPerformedReason;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_NO_SEQUENCING_PERFORMED_REASON, required = false)
  public @Nullable CarePlanNoSequencingPerformedReasonCoding getNoSequencingPerformedReason() {
    return noSequencingPerformedReason;
  }

  @JsonProperty(value = JSON_PROPERTY_NO_SEQUENCING_PERFORMED_REASON, required = false)
  public void setNoSequencingPerformedReason(
      @Nullable CarePlanNoSequencingPerformedReasonCoding noSequencingPerformedReason) {
    this.noSequencingPerformedReason = noSequencingPerformedReason;
  }

  public MtbCarePlan patient(Reference patient) {
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
    MtbCarePlan mtbCarePlan = (MtbCarePlan) o;
    return Objects.equals(this.boardType, mtbCarePlan.boardType)
        && Objects.equals(
            this.recommendationsMissingReason, mtbCarePlan.recommendationsMissingReason)
        && Objects.equals(
            this.studyEnrollmentRecommendations, mtbCarePlan.studyEnrollmentRecommendations)
        && Objects.equals(this.reason, mtbCarePlan.reason)
        && Objects.equals(
            this.geneticCounselingRecommendation, mtbCarePlan.geneticCounselingRecommendation)
        && Objects.equals(this.id, mtbCarePlan.id)
        && Objects.equals(this.medicationRecommendations, mtbCarePlan.medicationRecommendations)
        && Objects.equals(this.rebiopsyRequests, mtbCarePlan.rebiopsyRequests)
        && Objects.equals(
            this.histologyReevaluationRequests, mtbCarePlan.histologyReevaluationRequests)
        && Objects.equals(this.notes, mtbCarePlan.notes)
        && Objects.equals(this.procedureRecommendations, mtbCarePlan.procedureRecommendations)
        && Objects.equals(this.issuedOn, mtbCarePlan.issuedOn)
        && Objects.equals(this.noSequencingPerformedReason, mtbCarePlan.noSequencingPerformedReason)
        && Objects.equals(this.patient, mtbCarePlan.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        boardType,
        recommendationsMissingReason,
        studyEnrollmentRecommendations,
        reason,
        geneticCounselingRecommendation,
        id,
        medicationRecommendations,
        rebiopsyRequests,
        histologyReevaluationRequests,
        notes,
        procedureRecommendations,
        issuedOn,
        noSequencingPerformedReason,
        patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbCarePlan {\n");
    sb.append("    boardType: ").append(toIndentedString(boardType)).append("\n");
    sb.append("    recommendationsMissingReason: ")
        .append(toIndentedString(recommendationsMissingReason))
        .append("\n");
    sb.append("    studyEnrollmentRecommendations: ")
        .append(toIndentedString(studyEnrollmentRecommendations))
        .append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    geneticCounselingRecommendation: ")
        .append(toIndentedString(geneticCounselingRecommendation))
        .append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    medicationRecommendations: ")
        .append(toIndentedString(medicationRecommendations))
        .append("\n");
    sb.append("    rebiopsyRequests: ").append(toIndentedString(rebiopsyRequests)).append("\n");
    sb.append("    histologyReevaluationRequests: ")
        .append(toIndentedString(histologyReevaluationRequests))
        .append("\n");
    sb.append("    notes: ").append(toIndentedString(notes)).append("\n");
    sb.append("    procedureRecommendations: ")
        .append(toIndentedString(procedureRecommendations))
        .append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    noSequencingPerformedReason: ")
        .append(toIndentedString(noSequencingPerformedReason))
        .append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbCarePlan instance;

    public Builder() {
      this(new MtbCarePlan());
    }

    protected Builder(MtbCarePlan instance) {
      this.instance = instance;
    }

    public MtbCarePlan.Builder boardType(CarePlanBoardTypeCoding boardType) {
      this.instance.boardType = boardType;
      return this;
    }

    public MtbCarePlan.Builder recommendationsMissingReason(
        MtbCarePlanRecommendationsMissingReasonCoding recommendationsMissingReason) {
      this.instance.recommendationsMissingReason = recommendationsMissingReason;
      return this;
    }

    public MtbCarePlan.Builder studyEnrollmentRecommendations(
        List<MtbStudyEnrollmentRecommendation> studyEnrollmentRecommendations) {
      this.instance.studyEnrollmentRecommendations = studyEnrollmentRecommendations;
      return this;
    }

    public MtbCarePlan.Builder reason(Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public MtbCarePlan.Builder geneticCounselingRecommendation(
        GeneticCounselingRecommendation geneticCounselingRecommendation) {
      this.instance.geneticCounselingRecommendation = geneticCounselingRecommendation;
      return this;
    }

    public MtbCarePlan.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbCarePlan.Builder medicationRecommendations(
        List<MtbMedicationRecommendation> medicationRecommendations) {
      this.instance.medicationRecommendations = medicationRecommendations;
      return this;
    }

    public MtbCarePlan.Builder rebiopsyRequests(
        List<MtbCarePlanRebiopsyRequestsInner> rebiopsyRequests) {
      this.instance.rebiopsyRequests = rebiopsyRequests;
      return this;
    }

    public MtbCarePlan.Builder histologyReevaluationRequests(
        List<MtbCarePlanHistologyReevaluationRequestsInner> histologyReevaluationRequests) {
      this.instance.histologyReevaluationRequests = histologyReevaluationRequests;
      return this;
    }

    public MtbCarePlan.Builder notes(List<String> notes) {
      this.instance.notes = notes;
      return this;
    }

    public MtbCarePlan.Builder procedureRecommendations(
        List<MtbCarePlanProcedureRecommendationsInner> procedureRecommendations) {
      this.instance.procedureRecommendations = procedureRecommendations;
      return this;
    }

    public MtbCarePlan.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public MtbCarePlan.Builder noSequencingPerformedReason(
        CarePlanNoSequencingPerformedReasonCoding noSequencingPerformedReason) {
      this.instance.noSequencingPerformedReason = noSequencingPerformedReason;
      return this;
    }

    public MtbCarePlan.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbCarePlan build() {
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

  public static MtbCarePlan.Builder builder() {
    return new MtbCarePlan.Builder();
  }

  public MtbCarePlan.Builder toBuilder() {
    return new MtbCarePlan.Builder()
        .boardType(getBoardType())
        .recommendationsMissingReason(getRecommendationsMissingReason())
        .studyEnrollmentRecommendations(getStudyEnrollmentRecommendations())
        .reason(getReason())
        .geneticCounselingRecommendation(getGeneticCounselingRecommendation())
        .id(getId())
        .medicationRecommendations(getMedicationRecommendations())
        .rebiopsyRequests(getRebiopsyRequests())
        .histologyReevaluationRequests(getHistologyReevaluationRequests())
        .notes(getNotes())
        .procedureRecommendations(getProcedureRecommendations())
        .issuedOn(getIssuedOn())
        .noSequencingPerformedReason(getNoSequencingPerformedReason())
        .patient(getPatient());
  }
}
