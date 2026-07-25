package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class FollowUp {
  public static final String JSON_PROPERTY_DATE = "date";

  private Date date;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_LAST_CONTACT_DATE = "lastContactDate";

  private @Nullable Date lastContactDate;

  public static final String JSON_PROPERTY_PATIENT_STATUS = "patientStatus";

  private @Nullable FollowUpPatientStatusCoding patientStatus;

  public FollowUp() {}

  public FollowUp date(Date date) {
    this.date = date;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getDate() {
    return date;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  public void setDate(Date date) {
    this.date = date;
  }

  public FollowUp patient(Reference patient) {
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

  public FollowUp lastContactDate(@Nullable Date lastContactDate) {
    this.lastContactDate = lastContactDate;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LAST_CONTACT_DATE, required = false)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public @Nullable Date getLastContactDate() {
    return lastContactDate;
  }

  @JsonProperty(value = JSON_PROPERTY_LAST_CONTACT_DATE, required = false)
  public void setLastContactDate(@Nullable Date lastContactDate) {
    this.lastContactDate = lastContactDate;
  }

  public FollowUp patientStatus(@Nullable FollowUpPatientStatusCoding patientStatus) {
    this.patientStatus = patientStatus;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT_STATUS, required = false)
  public @Nullable FollowUpPatientStatusCoding getPatientStatus() {
    return patientStatus;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT_STATUS, required = false)
  public void setPatientStatus(@Nullable FollowUpPatientStatusCoding patientStatus) {
    this.patientStatus = patientStatus;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FollowUp followUp = (FollowUp) o;
    return Objects.equals(this.date, followUp.date)
        && Objects.equals(this.patient, followUp.patient)
        && Objects.equals(this.lastContactDate, followUp.lastContactDate)
        && Objects.equals(this.patientStatus, followUp.patientStatus);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, patient, lastContactDate, patientStatus);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FollowUp {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    lastContactDate: ").append(toIndentedString(lastContactDate)).append("\n");
    sb.append("    patientStatus: ").append(toIndentedString(patientStatus)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private FollowUp instance;

    public Builder() {
      this(new FollowUp());
    }

    protected Builder(FollowUp instance) {
      this.instance = instance;
    }

    public FollowUp.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public FollowUp.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public FollowUp.Builder lastContactDate(Date lastContactDate) {
      this.instance.lastContactDate = lastContactDate;
      return this;
    }

    public FollowUp.Builder patientStatus(FollowUpPatientStatusCoding patientStatus) {
      this.instance.patientStatus = patientStatus;
      return this;
    }

    public FollowUp build() {
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

  public static FollowUp.Builder builder() {
    return new FollowUp.Builder();
  }

  public FollowUp.Builder toBuilder() {
    return new FollowUp.Builder()
        .date(getDate())
        .patient(getPatient())
        .lastContactDate(getLastContactDate())
        .patientStatus(getPatientStatus());
  }
}
