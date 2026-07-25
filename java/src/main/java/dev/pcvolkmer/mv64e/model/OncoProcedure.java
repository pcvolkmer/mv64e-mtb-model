package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class OncoProcedure {
  public static final String JSON_PROPERTY_STATUS_REASON = "statusReason";

  private @Nullable MtbTherapyStatusReasonCoding statusReason;

  public static final String JSON_PROPERTY_INTENT = "intent";

  private @Nullable MtbTherapyIntentCoding intent;

  public static final String JSON_PROPERTY_THERAPY_LINE = "therapyLine";

  private @Nullable Integer therapyLine;

  public static final String JSON_PROPERTY_REASON = "reason";

  private @Nullable Reference reason;

  public static final String JSON_PROPERTY_CODE = "code";

  private OncoProcedureTypeCoding code;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_RECORDED_ON = "recordedOn";

  private Date recordedOn;

  public static final String JSON_PROPERTY_STATUS = "status";

  private TherapyStatusCoding status;

  public static final String JSON_PROPERTY_BASED_ON = "basedOn";

  private @Nullable Reference basedOn;

  public static final String JSON_PROPERTY_NOTES = "notes";

  private @Nullable List<String> notes;

  public static final String JSON_PROPERTY_PERIOD = "period";

  private @Nullable PeriodDate period;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public OncoProcedure() {}

  public OncoProcedure statusReason(@Nullable MtbTherapyStatusReasonCoding statusReason) {
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

  public OncoProcedure intent(@Nullable MtbTherapyIntentCoding intent) {
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

  public OncoProcedure therapyLine(@Nullable Integer therapyLine) {
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

  public OncoProcedure reason(@Nullable Reference reason) {
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

  public OncoProcedure code(OncoProcedureTypeCoding code) {
    this.code = code;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public OncoProcedureTypeCoding getCode() {
    return code;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public void setCode(OncoProcedureTypeCoding code) {
    this.code = code;
  }

  public OncoProcedure id(String id) {
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

  public OncoProcedure recordedOn(Date recordedOn) {
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

  public OncoProcedure status(TherapyStatusCoding status) {
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

  public OncoProcedure basedOn(@Nullable Reference basedOn) {
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

  public OncoProcedure notes(@Nullable List<String> notes) {
    this.notes = notes;
    return this;
  }

  public OncoProcedure addNotesItem(String notesItem) {
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

  public OncoProcedure period(@Nullable PeriodDate period) {
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

  public OncoProcedure patient(Reference patient) {
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
    OncoProcedure oncoProcedure = (OncoProcedure) o;
    return Objects.equals(this.statusReason, oncoProcedure.statusReason)
        && Objects.equals(this.intent, oncoProcedure.intent)
        && Objects.equals(this.therapyLine, oncoProcedure.therapyLine)
        && Objects.equals(this.reason, oncoProcedure.reason)
        && Objects.equals(this.code, oncoProcedure.code)
        && Objects.equals(this.id, oncoProcedure.id)
        && Objects.equals(this.recordedOn, oncoProcedure.recordedOn)
        && Objects.equals(this.status, oncoProcedure.status)
        && Objects.equals(this.basedOn, oncoProcedure.basedOn)
        && Objects.equals(this.notes, oncoProcedure.notes)
        && Objects.equals(this.period, oncoProcedure.period)
        && Objects.equals(this.patient, oncoProcedure.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        statusReason,
        intent,
        therapyLine,
        reason,
        code,
        id,
        recordedOn,
        status,
        basedOn,
        notes,
        period,
        patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OncoProcedure {\n");
    sb.append("    statusReason: ").append(toIndentedString(statusReason)).append("\n");
    sb.append("    intent: ").append(toIndentedString(intent)).append("\n");
    sb.append("    therapyLine: ").append(toIndentedString(therapyLine)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    recordedOn: ").append(toIndentedString(recordedOn)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    basedOn: ").append(toIndentedString(basedOn)).append("\n");
    sb.append("    notes: ").append(toIndentedString(notes)).append("\n");
    sb.append("    period: ").append(toIndentedString(period)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private OncoProcedure instance;

    public Builder() {
      this(new OncoProcedure());
    }

    protected Builder(OncoProcedure instance) {
      this.instance = instance;
    }

    public OncoProcedure.Builder statusReason(MtbTherapyStatusReasonCoding statusReason) {
      this.instance.statusReason = statusReason;
      return this;
    }

    public OncoProcedure.Builder intent(MtbTherapyIntentCoding intent) {
      this.instance.intent = intent;
      return this;
    }

    public OncoProcedure.Builder therapyLine(Integer therapyLine) {
      this.instance.therapyLine = therapyLine;
      return this;
    }

    public OncoProcedure.Builder reason(Reference reason) {
      this.instance.reason = reason;
      return this;
    }

    public OncoProcedure.Builder code(OncoProcedureTypeCoding code) {
      this.instance.code = code;
      return this;
    }

    public OncoProcedure.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public OncoProcedure.Builder recordedOn(Date recordedOn) {
      this.instance.recordedOn = recordedOn;
      return this;
    }

    public OncoProcedure.Builder status(TherapyStatusCoding status) {
      this.instance.status = status;
      return this;
    }

    public OncoProcedure.Builder basedOn(Reference basedOn) {
      this.instance.basedOn = basedOn;
      return this;
    }

    public OncoProcedure.Builder notes(List<String> notes) {
      this.instance.notes = notes;
      return this;
    }

    public OncoProcedure.Builder period(PeriodDate period) {
      this.instance.period = period;
      return this;
    }

    public OncoProcedure.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public OncoProcedure build() {
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

  public static OncoProcedure.Builder builder() {
    return new OncoProcedure.Builder();
  }

  public OncoProcedure.Builder toBuilder() {
    return new OncoProcedure.Builder()
        .statusReason(getStatusReason())
        .intent(getIntent())
        .therapyLine(getTherapyLine())
        .reason(getReason())
        .code(getCode())
        .id(getId())
        .recordedOn(getRecordedOn())
        .status(getStatus())
        .basedOn(getBasedOn())
        .notes(getNotes())
        .period(getPeriod())
        .patient(getPatient());
  }
}
