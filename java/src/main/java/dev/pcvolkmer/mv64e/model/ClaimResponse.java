package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class ClaimResponse {
  public static final String JSON_PROPERTY_STATUS_REASON = "statusReason";

  private @Nullable Set<ClaimResponseStatusReasonCoding> statusReason;

  public static final String JSON_PROPERTY_CLAIM = "claim";

  private Reference claim;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_STATUS = "status";

  private @Nullable ClaimResponseStatusCoding status;

  public static final String JSON_PROPERTY_ISSUED_ON = "issuedOn";

  private Date issuedOn;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public ClaimResponse() {}

  public ClaimResponse statusReason(@Nullable Set<ClaimResponseStatusReasonCoding> statusReason) {
    this.statusReason = statusReason;
    return this;
  }

  public ClaimResponse addStatusReasonItem(ClaimResponseStatusReasonCoding statusReasonItem) {
    if (this.statusReason == null) {
      this.statusReason = new LinkedHashSet<>();
    }
    this.statusReason.add(statusReasonItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS_REASON, required = false)
  public @Nullable Set<ClaimResponseStatusReasonCoding> getStatusReason() {
    return statusReason;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_STATUS_REASON, required = false)
  public void setStatusReason(@Nullable Set<ClaimResponseStatusReasonCoding> statusReason) {
    this.statusReason = statusReason;
  }

  public ClaimResponse claim(Reference claim) {
    this.claim = claim;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIM, required = true)
  public Reference getClaim() {
    return claim;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIM, required = true)
  public void setClaim(Reference claim) {
    this.claim = claim;
  }

  public ClaimResponse id(String id) {
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

  public ClaimResponse status(@Nullable ClaimResponseStatusCoding status) {
    this.status = status;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS, required = false)
  public @Nullable ClaimResponseStatusCoding getStatus() {
    return status;
  }

  @JsonProperty(value = JSON_PROPERTY_STATUS, required = false)
  public void setStatus(@Nullable ClaimResponseStatusCoding status) {
    this.status = status;
  }

  public ClaimResponse issuedOn(Date issuedOn) {
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

  public ClaimResponse patient(Reference patient) {
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
    ClaimResponse claimResponse = (ClaimResponse) o;
    return Objects.equals(this.statusReason, claimResponse.statusReason)
        && Objects.equals(this.claim, claimResponse.claim)
        && Objects.equals(this.id, claimResponse.id)
        && Objects.equals(this.status, claimResponse.status)
        && Objects.equals(this.issuedOn, claimResponse.issuedOn)
        && Objects.equals(this.patient, claimResponse.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(statusReason, claim, id, status, issuedOn, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClaimResponse {\n");
    sb.append("    statusReason: ").append(toIndentedString(statusReason)).append("\n");
    sb.append("    claim: ").append(toIndentedString(claim)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    issuedOn: ").append(toIndentedString(issuedOn)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private ClaimResponse instance;

    public Builder() {
      this(new ClaimResponse());
    }

    protected Builder(ClaimResponse instance) {
      this.instance = instance;
    }

    public ClaimResponse.Builder statusReason(Set<ClaimResponseStatusReasonCoding> statusReason) {
      this.instance.statusReason = statusReason;
      return this;
    }

    public ClaimResponse.Builder claim(Reference claim) {
      this.instance.claim = claim;
      return this;
    }

    public ClaimResponse.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public ClaimResponse.Builder status(ClaimResponseStatusCoding status) {
      this.instance.status = status;
      return this;
    }

    public ClaimResponse.Builder issuedOn(Date issuedOn) {
      this.instance.issuedOn = issuedOn;
      return this;
    }

    public ClaimResponse.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public ClaimResponse build() {
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

  public static ClaimResponse.Builder builder() {
    return new ClaimResponse.Builder();
  }

  public ClaimResponse.Builder toBuilder() {
    return new ClaimResponse.Builder()
        .statusReason(getStatusReason())
        .claim(getClaim())
        .id(getId())
        .status(getStatus())
        .issuedOn(getIssuedOn())
        .patient(getPatient());
  }
}
