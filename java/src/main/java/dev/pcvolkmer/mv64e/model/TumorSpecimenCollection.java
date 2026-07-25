package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class TumorSpecimenCollection {
  public static final String JSON_PROPERTY_DATE = "date";

  private @Nullable Date date;

  public static final String JSON_PROPERTY_METHOD = "method";

  private TumorSpecimenCollectionMethodCoding method;

  public static final String JSON_PROPERTY_LOCALIZATION = "localization";

  private TumorSpecimenCollectionLocalizationCoding localization;

  public TumorSpecimenCollection() {}

  public TumorSpecimenCollection date(@Nullable Date date) {
    this.date = date;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public @Nullable Date getDate() {
    return date;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDate(@Nullable Date date) {
    this.date = date;
  }

  public TumorSpecimenCollection method(TumorSpecimenCollectionMethodCoding method) {
    this.method = method;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TumorSpecimenCollectionMethodCoding getMethod() {
    return method;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMethod(TumorSpecimenCollectionMethodCoding method) {
    this.method = method;
  }

  public TumorSpecimenCollection localization(
      TumorSpecimenCollectionLocalizationCoding localization) {
    this.localization = localization;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LOCALIZATION, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TumorSpecimenCollectionLocalizationCoding getLocalization() {
    return localization;
  }

  @JsonProperty(value = JSON_PROPERTY_LOCALIZATION, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLocalization(TumorSpecimenCollectionLocalizationCoding localization) {
    this.localization = localization;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TumorSpecimenCollection tumorSpecimenCollection = (TumorSpecimenCollection) o;
    return Objects.equals(this.date, tumorSpecimenCollection.date)
        && Objects.equals(this.method, tumorSpecimenCollection.method)
        && Objects.equals(this.localization, tumorSpecimenCollection.localization);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, method, localization);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorSpecimenCollection {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    localization: ").append(toIndentedString(localization)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TumorSpecimenCollection instance;

    public Builder() {
      this(new TumorSpecimenCollection());
    }

    protected Builder(TumorSpecimenCollection instance) {
      this.instance = instance;
    }

    public TumorSpecimenCollection.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public TumorSpecimenCollection.Builder method(TumorSpecimenCollectionMethodCoding method) {
      this.instance.method = method;
      return this;
    }

    public TumorSpecimenCollection.Builder localization(
        TumorSpecimenCollectionLocalizationCoding localization) {
      this.instance.localization = localization;
      return this;
    }

    public TumorSpecimenCollection build() {
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

  public static TumorSpecimenCollection.Builder builder() {
    return new TumorSpecimenCollection.Builder();
  }

  public TumorSpecimenCollection.Builder toBuilder() {
    return new TumorSpecimenCollection.Builder()
        .date(getDate())
        .method(getMethod())
        .localization(getLocalization());
  }
}
