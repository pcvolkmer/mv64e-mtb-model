package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class TumorSpecimen {
  public static final String JSON_PROPERTY_DIAGNOSIS = "diagnosis";

  private Reference diagnosis;

  public static final String JSON_PROPERTY_COLLECTION = "collection";

  private @Nullable TumorSpecimenCollection collection;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_TYPE = "type";

  private TumorSpecimenTypeCoding type;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public TumorSpecimen() {}

  public TumorSpecimen diagnosis(Reference diagnosis) {
    this.diagnosis = diagnosis;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSIS, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public Reference getDiagnosis() {
    return diagnosis;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSIS, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDiagnosis(Reference diagnosis) {
    this.diagnosis = diagnosis;
  }

  public TumorSpecimen collection(@Nullable TumorSpecimenCollection collection) {
    this.collection = collection;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_COLLECTION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public @Nullable TumorSpecimenCollection getCollection() {
    return collection;
  }

  @JsonProperty(value = JSON_PROPERTY_COLLECTION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setCollection(@Nullable TumorSpecimenCollection collection) {
    this.collection = collection;
  }

  public TumorSpecimen id(String id) {
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

  public TumorSpecimen type(TumorSpecimenTypeCoding type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TumorSpecimenTypeCoding getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setType(TumorSpecimenTypeCoding type) {
    this.type = type;
  }

  public TumorSpecimen patient(Reference patient) {
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
    TumorSpecimen tumorSpecimen = (TumorSpecimen) o;
    return Objects.equals(this.diagnosis, tumorSpecimen.diagnosis)
        && Objects.equals(this.collection, tumorSpecimen.collection)
        && Objects.equals(this.id, tumorSpecimen.id)
        && Objects.equals(this.type, tumorSpecimen.type)
        && Objects.equals(this.patient, tumorSpecimen.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(diagnosis, collection, id, type, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorSpecimen {\n");
    sb.append("    diagnosis: ").append(toIndentedString(diagnosis)).append("\n");
    sb.append("    collection: ").append(toIndentedString(collection)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TumorSpecimen instance;

    public Builder() {
      this(new TumorSpecimen());
    }

    protected Builder(TumorSpecimen instance) {
      this.instance = instance;
    }

    public TumorSpecimen.Builder diagnosis(Reference diagnosis) {
      this.instance.diagnosis = diagnosis;
      return this;
    }

    public TumorSpecimen.Builder collection(TumorSpecimenCollection collection) {
      this.instance.collection = collection;
      return this;
    }

    public TumorSpecimen.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public TumorSpecimen.Builder type(TumorSpecimenTypeCoding type) {
      this.instance.type = type;
      return this;
    }

    public TumorSpecimen.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public TumorSpecimen build() {
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

  public static TumorSpecimen.Builder builder() {
    return new TumorSpecimen.Builder();
  }

  public TumorSpecimen.Builder toBuilder() {
    return new TumorSpecimen.Builder()
        .diagnosis(getDiagnosis())
        .collection(getCollection())
        .id(getId())
        .type(getType())
        .patient(getPatient());
  }
}
