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

public class MtbMedicationRecommendation {
  public static final String JSON_PROPERTY_PRIORITY = "priority";

  private RecommendationPriorityCoding priority;

  public static final String JSON_PROPERTY_MEDICATION = "medication";

  private Set<AtcUnregisteredMedicationCoding> medication;

  public static final String JSON_PROPERTY_SUPPORTING_VARIANTS = "supportingVariants";

  private @Nullable List<GeneAlterationReference> supportingVariants;

  public static final String JSON_PROPERTY_USE_TYPE = "useType";

  private @Nullable MtbMedicationRecommendationUseTypeCoding useType;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_CATEGORY = "category";

  private @Nullable Set<MtbMedicationRecommendationCategoryCoding> category;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_SUPPORTING_FINDINGS = "supportingFindings";

  private @Nullable List<Reference> supportingFindings;

  public static final String JSON_PROPERTY_LEVEL_OF_EVIDENCE = "levelOfEvidence";

  private @Nullable LevelOfEvidence levelOfEvidence;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public MtbMedicationRecommendation() {}

  public MtbMedicationRecommendation priority(RecommendationPriorityCoding priority) {
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

  public MtbMedicationRecommendation medication(Set<AtcUnregisteredMedicationCoding> medication) {
    this.medication = medication;
    return this;
  }

  public MtbMedicationRecommendation addMedicationItem(
      AtcUnregisteredMedicationCoding medicationItem) {
    if (this.medication == null) {
      this.medication = new LinkedHashSet<>();
    }
    this.medication.add(medicationItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MEDICATION, required = true)
  public Set<AtcUnregisteredMedicationCoding> getMedication() {
    return medication;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_MEDICATION, required = true)
  public void setMedication(Set<AtcUnregisteredMedicationCoding> medication) {
    this.medication = medication;
  }

  public MtbMedicationRecommendation supportingVariants(
      @Nullable List<GeneAlterationReference> supportingVariants) {
    this.supportingVariants = supportingVariants;
    return this;
  }

  public MtbMedicationRecommendation addSupportingVariantsItem(
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

  public MtbMedicationRecommendation useType(
      @Nullable MtbMedicationRecommendationUseTypeCoding useType) {
    this.useType = useType;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_USE_TYPE, required = false)
  public @Nullable MtbMedicationRecommendationUseTypeCoding getUseType() {
    return useType;
  }

  @JsonProperty(value = JSON_PROPERTY_USE_TYPE, required = false)
  public void setUseType(@Nullable MtbMedicationRecommendationUseTypeCoding useType) {
    this.useType = useType;
  }

  public MtbMedicationRecommendation reason(@Nullable Reference reason) {
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

  public MtbMedicationRecommendation category(
      @Nullable Set<MtbMedicationRecommendationCategoryCoding> category) {
    this.category = category;
    return this;
  }

  public MtbMedicationRecommendation addCategoryItem(
      MtbMedicationRecommendationCategoryCoding categoryItem) {
    if (this.category == null) {
      this.category = new LinkedHashSet<>();
    }
    this.category.add(categoryItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CATEGORY, required = false)
  public @Nullable Set<MtbMedicationRecommendationCategoryCoding> getCategory() {
    return category;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_CATEGORY, required = false)
  public void setCategory(@Nullable Set<MtbMedicationRecommendationCategoryCoding> category) {
    this.category = category;
  }

  public MtbMedicationRecommendation issuedOn(Date issuedOn) {
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

  public MtbMedicationRecommendation patient(Reference patient) {
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

  public MtbMedicationRecommendation supportingFindings(
      @Nullable List<Reference> supportingFindings) {
    this.supportingFindings = supportingFindings;
    return this;
  }

  public MtbMedicationRecommendation addSupportingFindingsItem(Reference supportingFindingsItem) {
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

  public MtbMedicationRecommendation levelOfEvidence(@Nullable LevelOfEvidence levelOfEvidence) {
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

  public MtbMedicationRecommendation id(String id) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbMedicationRecommendation mtbMedicationRecommendation = (MtbMedicationRecommendation) o;
    return Objects.equals(this.priority, mtbMedicationRecommendation.priority)
        && Objects.equals(this.medication, mtbMedicationRecommendation.medication)
        && Objects.equals(this.supportingVariants, mtbMedicationRecommendation.supportingVariants)
        && Objects.equals(this.useType, mtbMedicationRecommendation.useType)
        && Objects.equals(this.reason, mtbMedicationRecommendation.reason)
        && Objects.equals(this.category, mtbMedicationRecommendation.category)
        && Objects.equals(this.issuedOn, mtbMedicationRecommendation.issuedOn)
        && Objects.equals(this.patient, mtbMedicationRecommendation.patient)
        && Objects.equals(this.supportingFindings, mtbMedicationRecommendation.supportingFindings)
        && Objects.equals(this.levelOfEvidence, mtbMedicationRecommendation.levelOfEvidence)
        && Objects.equals(this.id, mtbMedicationRecommendation.id);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        priority,
        medication,
        supportingVariants,
        useType,
        reason,
        category,
        issuedOn,
        patient,
        supportingFindings,
        levelOfEvidence,
        id);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbMedicationRecommendation {\n");
    sb.append("    priority: ").append(toIndentedString(priority)).append("\n");
    sb.append("    medication: ").append(toIndentedString(medication)).append("\n");
    sb.append("    supportingVariants: ").append(toIndentedString(supportingVariants)).append("\n");
    sb.append("    useType: ").append(toIndentedString(useType)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    supportingFindings: ").append(toIndentedString(supportingFindings)).append("\n");
    sb.append("    levelOfEvidence: ").append(toIndentedString(levelOfEvidence)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbMedicationRecommendation instance;

    public Builder() {
      this(new MtbMedicationRecommendation());
    }

    protected Builder(MtbMedicationRecommendation instance) {
      this.instance = instance;
    }

    public MtbMedicationRecommendation.Builder priority(RecommendationPriorityCoding priority) {
      this.instance.priority = priority;
      return this;
    }

    public MtbMedicationRecommendation.Builder medication(
        Set<AtcUnregisteredMedicationCoding> medication) {
      this.instance.medication = medication;
      return this;
    }

    public MtbMedicationRecommendation.Builder supportingVariants(
        List<GeneAlterationReference> supportingVariants) {
      this.instance.supportingVariants = supportingVariants;
      return this;
    }

    public MtbMedicationRecommendation.Builder useType(
        MtbMedicationRecommendationUseTypeCoding useType) {
      this.instance.useType = useType;
      return this;
    }

    public MtbMedicationRecommendation.Builder reason(Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public MtbMedicationRecommendation.Builder category(
        Set<MtbMedicationRecommendationCategoryCoding> category) {
      this.instance.category = category;
      return this;
    }

    public MtbMedicationRecommendation.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public MtbMedicationRecommendation.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbMedicationRecommendation.Builder supportingFindings(
        List<Reference> supportingFindings) {
      this.instance.supportingFindings = supportingFindings;
      return this;
    }

    public MtbMedicationRecommendation.Builder levelOfEvidence(LevelOfEvidence levelOfEvidence) {
      this.instance.levelOfEvidence = levelOfEvidence;
      return this;
    }

    public MtbMedicationRecommendation.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbMedicationRecommendation build() {
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

  public static MtbMedicationRecommendation.Builder builder() {
    return new MtbMedicationRecommendation.Builder();
  }

  public MtbMedicationRecommendation.Builder toBuilder() {
    return new MtbMedicationRecommendation.Builder()
        .priority(getPriority())
        .medication(getMedication())
        .supportingVariants(getSupportingVariants())
        .useType(getUseType())
        .reason(getReason())
        .category(getCategory())
        .issuedOn(getIssuedOn())
        .patient(getPatient())
        .supportingFindings(getSupportingFindings())
        .levelOfEvidence(getLevelOfEvidence())
        .id(getId());
  }
}
