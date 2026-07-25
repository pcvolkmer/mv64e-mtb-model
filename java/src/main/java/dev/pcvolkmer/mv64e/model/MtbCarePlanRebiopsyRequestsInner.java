package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class MtbCarePlanRebiopsyRequestsInner {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_TUMOR_ENTITY = "tumorEntity";

  private Reference tumorEntity;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public MtbCarePlanRebiopsyRequestsInner() {}

  public MtbCarePlanRebiopsyRequestsInner id(String id) {
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

  public MtbCarePlanRebiopsyRequestsInner patient(Reference patient) {
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

  public MtbCarePlanRebiopsyRequestsInner tumorEntity(Reference tumorEntity) {
    this.tumorEntity = tumorEntity;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_ENTITY, required = true)
  public Reference getTumorEntity() {
    return tumorEntity;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_ENTITY, required = true)
  public void setTumorEntity(Reference tumorEntity) {
    this.tumorEntity = tumorEntity;
  }

  public MtbCarePlanRebiopsyRequestsInner issuedOn(Date issuedOn) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbCarePlanRebiopsyRequestsInner mtbCarePlanRebiopsyRequestsInner =
        (MtbCarePlanRebiopsyRequestsInner) o;
    return Objects.equals(this.id, mtbCarePlanRebiopsyRequestsInner.id)
        && Objects.equals(this.patient, mtbCarePlanRebiopsyRequestsInner.patient)
        && Objects.equals(this.tumorEntity, mtbCarePlanRebiopsyRequestsInner.tumorEntity)
        && Objects.equals(this.issuedOn, mtbCarePlanRebiopsyRequestsInner.issuedOn);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, patient, tumorEntity, issuedOn);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbCarePlanRebiopsyRequestsInner {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    tumorEntity: ").append(toIndentedString(tumorEntity)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbCarePlanRebiopsyRequestsInner instance;

    public Builder() {
      this(new MtbCarePlanRebiopsyRequestsInner());
    }

    protected Builder(MtbCarePlanRebiopsyRequestsInner instance) {
      this.instance = instance;
    }

    public MtbCarePlanRebiopsyRequestsInner.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbCarePlanRebiopsyRequestsInner.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbCarePlanRebiopsyRequestsInner.Builder tumorEntity(Reference tumorEntity) {
      this.instance.tumorEntity = tumorEntity;
      return this;
    }

    public MtbCarePlanRebiopsyRequestsInner.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public MtbCarePlanRebiopsyRequestsInner build() {
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

  public static MtbCarePlanRebiopsyRequestsInner.Builder builder() {
    return new MtbCarePlanRebiopsyRequestsInner.Builder();
  }

  public MtbCarePlanRebiopsyRequestsInner.Builder toBuilder() {
    return new MtbCarePlanRebiopsyRequestsInner.Builder()
        .id(getId())
        .patient(getPatient())
        .tumorEntity(getTumorEntity())
        .issuedOn(getIssuedOn());
  }
}
