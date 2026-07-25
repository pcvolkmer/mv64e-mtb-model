package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class Tmb {
  public static final String JSON_PROPERTY_INTERPRETATION = "interpretation";

  private @Nullable TmbInterpretationCoding interpretation;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_SPECIMEN = "specimen";

  private Reference specimen;

  public static final String JSON_PROPERTY_VALUE = "value";

  private TmbResult value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public Tmb() {}

  public Tmb interpretation(@Nullable TmbInterpretationCoding interpretation) {
    this.interpretation = interpretation;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public @Nullable TmbInterpretationCoding getInterpretation() {
    return interpretation;
  }

  @JsonProperty(value = JSON_PROPERTY_INTERPRETATION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setInterpretation(@Nullable TmbInterpretationCoding interpretation) {
    this.interpretation = interpretation;
  }

  public Tmb id(String id) {
    this.id = id;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public String getId() {
    return id;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setId(String id) {
    this.id = id;
  }

  public Tmb specimen(Reference specimen) {
    this.specimen = specimen;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Reference getSpecimen() {
    return specimen;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMEN, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSpecimen(Reference specimen) {
    this.specimen = specimen;
  }

  public Tmb value(TmbResult value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TmbResult getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setValue(TmbResult value) {
    this.value = value;
  }

  public Tmb patient(Reference patient) {
    this.patient = patient;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Reference getPatient() {
    return patient;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
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
    Tmb tmb = (Tmb) o;
    return Objects.equals(this.interpretation, tmb.interpretation)
        && Objects.equals(this.id, tmb.id)
        && Objects.equals(this.specimen, tmb.specimen)
        && Objects.equals(this.value, tmb.value)
        && Objects.equals(this.patient, tmb.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(interpretation, id, specimen, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Tmb {\n");
    sb.append("    interpretation: ").append(toIndentedString(interpretation)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    specimen: ").append(toIndentedString(specimen)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Tmb instance;

    public Builder() {
      this(new Tmb());
    }

    protected Builder(Tmb instance) {
      this.instance = instance;
    }

    public Tmb.Builder interpretation(TmbInterpretationCoding interpretation) {
      this.instance.interpretation = interpretation;
      return this;
    }

    public Tmb.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Tmb.Builder specimen(Reference specimen) {
      this.instance.specimen = specimen;
      return this;
    }

    public Tmb.Builder value(TmbResult value) {
      this.instance.value = value;
      return this;
    }

    public Tmb.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Tmb build() {
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

  public static Tmb.Builder builder() {
    return new Tmb.Builder();
  }

  public Tmb.Builder toBuilder() {
    return new Tmb.Builder()
        .interpretation(getInterpretation())
        .id(getId())
        .specimen(getSpecimen())
        .value(getValue())
        .patient(getPatient());
  }
}
