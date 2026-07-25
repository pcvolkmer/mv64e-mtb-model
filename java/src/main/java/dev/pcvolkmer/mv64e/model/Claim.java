package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class Claim {
  public static final String JSON_PROPERTY_STAGE = "stage";

  private @Nullable ClaimStageCoding stage;

  public static final String JSON_PROPERTY_REQUESTED_MEDICATION = "requestedMedication";

  private @Nullable Set<AtcUnregisteredMedicationCoding> requestedMedication;

  public static final String JSON_PROPERTY_RECOMMENDATION = "recommendation";

  private Reference recommendation;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public Claim() {}

  public Claim stage(@Nullable ClaimStageCoding stage) {
    this.stage = stage;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STAGE, required = false)
  public @Nullable ClaimStageCoding getStage() {
    return stage;
  }

  @JsonProperty(value = JSON_PROPERTY_STAGE, required = false)
  public void setStage(@Nullable ClaimStageCoding stage) {
    this.stage = stage;
  }

  public Claim requestedMedication(
      @Nullable Set<AtcUnregisteredMedicationCoding> requestedMedication) {
    this.requestedMedication = requestedMedication;
    return this;
  }

  public Claim addRequestedMedicationItem(AtcUnregisteredMedicationCoding requestedMedicationItem) {
    if (this.requestedMedication == null) {
      this.requestedMedication = new LinkedHashSet<>();
    }
    this.requestedMedication.add(requestedMedicationItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REQUESTED_MEDICATION, required = false)
  public @Nullable Set<AtcUnregisteredMedicationCoding> getRequestedMedication() {
    return requestedMedication;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_REQUESTED_MEDICATION, required = false)
  public void setRequestedMedication(
      @Nullable Set<AtcUnregisteredMedicationCoding> requestedMedication) {
    this.requestedMedication = requestedMedication;
  }

  public Claim recommendation(Reference recommendation) {
    this.recommendation = recommendation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATION, required = true)
  public Reference getRecommendation() {
    return recommendation;
  }

  @JsonProperty(value = JSON_PROPERTY_RECOMMENDATION, required = true)
  public void setRecommendation(Reference recommendation) {
    this.recommendation = recommendation;
  }

  public Claim id(String id) {
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

  public Claim issuedOn(Date issuedOn) {
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

  public Claim patient(Reference patient) {
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
    Claim claim = (Claim) o;
    return Objects.equals(this.stage, claim.stage)
        && Objects.equals(this.requestedMedication, claim.requestedMedication)
        && Objects.equals(this.recommendation, claim.recommendation)
        && Objects.equals(this.id, claim.id)
        && Objects.equals(this.issuedOn, claim.issuedOn)
        && Objects.equals(this.patient, claim.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(stage, requestedMedication, recommendation, id, issuedOn, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Claim {\n");
    sb.append("    stage: ").append(toIndentedString(stage)).append("\n");
    sb.append("    requestedMedication: ")
        .append(toIndentedString(requestedMedication))
        .append("\n");
    sb.append("    recommendation: ").append(toIndentedString(recommendation)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Claim instance;

    public Builder() {
      this(new Claim());
    }

    protected Builder(Claim instance) {
      this.instance = instance;
    }

    public Claim.Builder stage(ClaimStageCoding stage) {
      this.instance.stage = stage;
      return this;
    }

    public Claim.Builder requestedMedication(
        Set<AtcUnregisteredMedicationCoding> requestedMedication) {
      this.instance.requestedMedication = requestedMedication;
      return this;
    }

    public Claim.Builder recommendation(Reference recommendation) {
      this.instance.recommendation = recommendation;
      return this;
    }

    public Claim.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Claim.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public Claim.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Claim build() {
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

  public static Claim.Builder builder() {
    return new Claim.Builder();
  }

  public Claim.Builder toBuilder() {
    return new Claim.Builder()
        .stage(getStage())
        .requestedMedication(getRequestedMedication())
        .recommendation(getRecommendation())
        .id(getId())
        .issuedOn(getIssuedOn())
        .patient(getPatient());
  }
}
