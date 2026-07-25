package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class TumorStaging {
  public static final String JSON_PROPERTY_DATE = "date";

  private Date date;

  public static final String JSON_PROPERTY_METHOD = "method";

  private TumorStagingMethodCoding method;

  public static final String JSON_PROPERTY_TNM_CLASSIFICATION = "tnmClassification";

  private @Nullable TumorStagingTnmClassification tnmClassification;

  public static final String JSON_PROPERTY_OTHER_CLASSIFICATIONS = "otherClassifications";

  private @Nullable List<Coding> otherClassifications;

  public TumorStaging() {}

  public TumorStaging date(Date date) {
    this.date = date;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getDate() {
    return date;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDate(Date date) {
    this.date = date;
  }

  public TumorStaging method(TumorStagingMethodCoding method) {
    this.method = method;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public TumorStagingMethodCoding getMethod() {
    return method;
  }

  @JsonProperty(value = JSON_PROPERTY_METHOD, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMethod(TumorStagingMethodCoding method) {
    this.method = method;
  }

  public TumorStaging tnmClassification(@Nullable TumorStagingTnmClassification tnmClassification) {
    this.tnmClassification = tnmClassification;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TNM_CLASSIFICATION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public @Nullable TumorStagingTnmClassification getTnmClassification() {
    return tnmClassification;
  }

  @JsonProperty(value = JSON_PROPERTY_TNM_CLASSIFICATION, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setTnmClassification(@Nullable TumorStagingTnmClassification tnmClassification) {
    this.tnmClassification = tnmClassification;
  }

  public TumorStaging otherClassifications(@Nullable List<Coding> otherClassifications) {
    this.otherClassifications = otherClassifications;
    return this;
  }

  public TumorStaging addOtherClassificationsItem(Coding otherClassificationsItem) {
    if (this.otherClassifications == null) {
      this.otherClassifications = new ArrayList<>();
    }
    this.otherClassifications.add(otherClassificationsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_OTHER_CLASSIFICATIONS, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public @Nullable List<Coding> getOtherClassifications() {
    return otherClassifications;
  }

  @JsonProperty(value = JSON_PROPERTY_OTHER_CLASSIFICATIONS, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setOtherClassifications(@Nullable List<Coding> otherClassifications) {
    this.otherClassifications = otherClassifications;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TumorStaging tumorStaging = (TumorStaging) o;
    return Objects.equals(this.date, tumorStaging.date)
        && Objects.equals(this.method, tumorStaging.method)
        && Objects.equals(this.tnmClassification, tumorStaging.tnmClassification)
        && Objects.equals(this.otherClassifications, tumorStaging.otherClassifications);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, method, tnmClassification, otherClassifications);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorStaging {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    tnmClassification: ").append(toIndentedString(tnmClassification)).append("\n");
    sb.append("    otherClassifications: ")
        .append(toIndentedString(otherClassifications))
        .append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TumorStaging instance;

    public Builder() {
      this(new TumorStaging());
    }

    protected Builder(TumorStaging instance) {
      this.instance = instance;
    }

    public TumorStaging.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public TumorStaging.Builder method(TumorStagingMethodCoding method) {
      this.instance.method = method;
      return this;
    }

    public TumorStaging.Builder tnmClassification(TumorStagingTnmClassification tnmClassification) {
      this.instance.tnmClassification = tnmClassification;
      return this;
    }

    public TumorStaging.Builder otherClassifications(List<Coding> otherClassifications) {
      this.instance.otherClassifications = otherClassifications;
      return this;
    }

    public TumorStaging build() {
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

  public static TumorStaging.Builder builder() {
    return new TumorStaging.Builder();
  }

  public TumorStaging.Builder toBuilder() {
    return new TumorStaging.Builder()
        .date(getDate())
        .method(getMethod())
        .tnmClassification(getTnmClassification())
        .otherClassifications(getOtherClassifications());
  }
}
