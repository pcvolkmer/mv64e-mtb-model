package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MtbEpisodeOfCare {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_PERIOD = "period";

  private PeriodDate period;

  public static final String JSON_PROPERTY_DIAGNOSES = "diagnoses";

  private @Nullable List<Reference> diagnoses;

  public MtbEpisodeOfCare() {}

  public MtbEpisodeOfCare id(String id) {
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

  public MtbEpisodeOfCare patient(Reference patient) {
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

  public MtbEpisodeOfCare period(PeriodDate period) {
    this.period = period;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PERIOD, required = true)
  public PeriodDate getPeriod() {
    return period;
  }

  @JsonProperty(value = JSON_PROPERTY_PERIOD, required = true)
  public void setPeriod(PeriodDate period) {
    this.period = period;
  }

  public MtbEpisodeOfCare diagnoses(@Nullable List<Reference> diagnoses) {
    this.diagnoses = diagnoses;
    return this;
  }

  public MtbEpisodeOfCare addDiagnosesItem(Reference diagnosesItem) {
    if (this.diagnoses == null) {
      this.diagnoses = new ArrayList<>();
    }
    this.diagnoses.add(diagnosesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSES, required = false)
  public @Nullable List<Reference> getDiagnoses() {
    return diagnoses;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSES, required = false)
  public void setDiagnoses(@Nullable List<Reference> diagnoses) {
    this.diagnoses = diagnoses;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbEpisodeOfCare mtbEpisodeOfCare = (MtbEpisodeOfCare) o;
    return Objects.equals(this.id, mtbEpisodeOfCare.id)
        && Objects.equals(this.patient, mtbEpisodeOfCare.patient)
        && Objects.equals(this.period, mtbEpisodeOfCare.period)
        && Objects.equals(this.diagnoses, mtbEpisodeOfCare.diagnoses);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, patient, period, diagnoses);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbEpisodeOfCare {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    period: ").append(toIndentedString(period)).append("\n");
    sb.append("    diagnoses: ").append(toIndentedString(diagnoses)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbEpisodeOfCare instance;

    public Builder() {
      this(new MtbEpisodeOfCare());
    }

    protected Builder(MtbEpisodeOfCare instance) {
      this.instance = instance;
    }

    public MtbEpisodeOfCare.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public MtbEpisodeOfCare.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public MtbEpisodeOfCare.Builder period(PeriodDate period) {
      this.instance.period = period;
      return this;
    }

    public MtbEpisodeOfCare.Builder diagnoses(List<Reference> diagnoses) {
      this.instance.diagnoses = diagnoses;
      return this;
    }

    public MtbEpisodeOfCare build() {
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

  public static MtbEpisodeOfCare.Builder builder() {
    return new MtbEpisodeOfCare.Builder();
  }

  public MtbEpisodeOfCare.Builder toBuilder() {
    return new MtbEpisodeOfCare.Builder()
        .id(getId())
        .patient(getPatient())
        .period(getPeriod())
        .diagnoses(getDiagnoses());
  }
}
