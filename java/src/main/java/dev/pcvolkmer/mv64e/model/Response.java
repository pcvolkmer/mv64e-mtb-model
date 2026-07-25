package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class Response {
  public static final String JSON_PROPERTY_METHOD = "method";

  private ResponseMethodCoding method;

  public static final String JSON_PROPERTY_THERAPY = "therapy";

  private Reference therapy;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_EFFECTIVE_DATE = "effectiveDate";

  private Date effectiveDate;

  public static final String JSON_PROPERTY_VALUE = "value";

  private RecistCoding value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public Response() {}

  public Response method(ResponseMethodCoding method) {
    this.method = method;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  public ResponseMethodCoding getMethod() {
    return method;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  public void setMethod(ResponseMethodCoding method) {
    this.method = method;
  }

  public Response therapy(Reference therapy) {
    this.therapy = therapy;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_THERAPY, required = true)
  public Reference getTherapy() {
    return therapy;
  }

  @JsonProperty(value = JSON_PROPERTY_THERAPY, required = true)
  public void setTherapy(Reference therapy) {
    this.therapy = therapy;
  }

  public Response id(String id) {
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

  public Response effectiveDate(Date effectiveDate) {
    this.effectiveDate = effectiveDate;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EFFECTIVE_DATE, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getEffectiveDate() {
    return effectiveDate;
  }

  @JsonProperty(value = JSON_PROPERTY_EFFECTIVE_DATE, required = true)
  public void setEffectiveDate(Date effectiveDate) {
    this.effectiveDate = effectiveDate;
  }

  public Response value(RecistCoding value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public RecistCoding getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(RecistCoding value) {
    this.value = value;
  }

  public Response patient(Reference patient) {
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
    Response response = (Response) o;
    return Objects.equals(this.method, response.method)
        && Objects.equals(this.therapy, response.therapy)
        && Objects.equals(this.id, response.id)
        && Objects.equals(this.effectiveDate, response.effectiveDate)
        && Objects.equals(this.value, response.value)
        && Objects.equals(this.patient, response.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(method, therapy, id, effectiveDate, value, patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Response {\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    therapy: ").append(toIndentedString(therapy)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    effectiveDate: ").append(toIndentedString(effectiveDate)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Response instance;

    public Builder() {
      this(new Response());
    }

    protected Builder(Response instance) {
      this.instance = instance;
    }

    public Response.Builder method(ResponseMethodCoding method) {
      this.instance.method = method;
      return this;
    }

    public Response.Builder therapy(Reference therapy) {
      this.instance.therapy = therapy;
      return this;
    }

    public Response.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Response.Builder effectiveDate(Date effectiveDate) {
      this.instance.effectiveDate = effectiveDate;
      return this;
    }

    public Response.Builder value(RecistCoding value) {
      this.instance.value = value;
      return this;
    }

    public Response.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public Response build() {
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

  public static Response.Builder builder() {
    return new Response.Builder();
  }

  public Response.Builder toBuilder() {
    return new Response.Builder()
        .method(getMethod())
        .therapy(getTherapy())
        .id(getId())
        .effectiveDate(getEffectiveDate())
        .value(getValue())
        .patient(getPatient());
  }
}
