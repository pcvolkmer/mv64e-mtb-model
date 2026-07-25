package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class GeneticCounselingRecommendation {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_REASON = "reason";

  private GeneticCounselingRecommendationReasonCoding reason;

  public GeneticCounselingRecommendation() {}

  public GeneticCounselingRecommendation id(String id) {
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

  public GeneticCounselingRecommendation patient(Reference patient) {
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

  public GeneticCounselingRecommendation issuedOn(Date issuedOn) {
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

  public GeneticCounselingRecommendation reason(
      GeneticCounselingRecommendationReasonCoding reason) {
    this.reason = reason;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON, required = true)
  public GeneticCounselingRecommendationReasonCoding getReason() {
    return reason;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON, required = true)
  public void setReason(GeneticCounselingRecommendationReasonCoding reason) {
    this.reason = reason;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GeneticCounselingRecommendation geneticCounselingRecommendation =
        (GeneticCounselingRecommendation) o;
    return Objects.equals(this.id, geneticCounselingRecommendation.id)
        && Objects.equals(this.patient, geneticCounselingRecommendation.patient)
        && Objects.equals(this.issuedOn, geneticCounselingRecommendation.issuedOn)
        && Objects.equals(this.reason, geneticCounselingRecommendation.reason);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, patient, issuedOn, reason);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GeneticCounselingRecommendation {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private GeneticCounselingRecommendation instance;

    public Builder() {
      this(new GeneticCounselingRecommendation());
    }

    protected Builder(GeneticCounselingRecommendation instance) {
      this.instance = instance;
    }

    public GeneticCounselingRecommendation.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public GeneticCounselingRecommendation.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public GeneticCounselingRecommendation.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public GeneticCounselingRecommendation.Builder reason(
        GeneticCounselingRecommendationReasonCoding reason) {
      this.instance.reason = reason;
      return this;
    }

    public GeneticCounselingRecommendation build() {
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

  public static GeneticCounselingRecommendation.Builder builder() {
    return new GeneticCounselingRecommendation.Builder();
  }

  public GeneticCounselingRecommendation.Builder toBuilder() {
    return new GeneticCounselingRecommendation.Builder()
        .id(getId())
        .patient(getPatient())
        .issuedOn(getIssuedOn())
        .reason(getReason());
  }
}
