package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;

public class MvhMetadataModelProjectConsentProvisionsInner {
  public static final String JSON_PROPERTY_DATE = "date";

  private Date date;

  public static final String JSON_PROPERTY_PURPOSE = "purpose";

  private ModelProjectConsentPurpose purpose;

  public static final String JSON_PROPERTY_TYPE = "type";

  private ConsentProvisionType type;

  public MvhMetadataModelProjectConsentProvisionsInner() {}

  public MvhMetadataModelProjectConsentProvisionsInner date(Date date) {
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

  public MvhMetadataModelProjectConsentProvisionsInner purpose(ModelProjectConsentPurpose purpose) {
    this.purpose = purpose;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PURPOSE, required = true)
  public ModelProjectConsentPurpose getPurpose() {
    return purpose;
  }

  @JsonProperty(value = JSON_PROPERTY_PURPOSE, required = true)
  public void setPurpose(ModelProjectConsentPurpose purpose) {
    this.purpose = purpose;
  }

  public MvhMetadataModelProjectConsentProvisionsInner type(ConsentProvisionType type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public ConsentProvisionType getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public void setType(ConsentProvisionType type) {
    this.type = type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MvhMetadataModelProjectConsentProvisionsInner mvhMetadataModelProjectConsentProvisionsInner =
        (MvhMetadataModelProjectConsentProvisionsInner) o;
    return Objects.equals(this.date, mvhMetadataModelProjectConsentProvisionsInner.date)
        && Objects.equals(this.purpose, mvhMetadataModelProjectConsentProvisionsInner.purpose)
        && Objects.equals(this.type, mvhMetadataModelProjectConsentProvisionsInner.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, purpose, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MvhMetadataModelProjectConsentProvisionsInner {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    purpose: ").append(toIndentedString(purpose)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MvhMetadataModelProjectConsentProvisionsInner instance;

    public Builder() {
      this(new MvhMetadataModelProjectConsentProvisionsInner());
    }

    protected Builder(MvhMetadataModelProjectConsentProvisionsInner instance) {
      this.instance = instance;
    }

    public MvhMetadataModelProjectConsentProvisionsInner.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public MvhMetadataModelProjectConsentProvisionsInner.Builder purpose(
        ModelProjectConsentPurpose purpose) {
      this.instance.purpose = purpose;
      return this;
    }

    public MvhMetadataModelProjectConsentProvisionsInner.Builder type(ConsentProvisionType type) {
      this.instance.type = type;
      return this;
    }

    public MvhMetadataModelProjectConsentProvisionsInner build() {
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

  public static MvhMetadataModelProjectConsentProvisionsInner.Builder builder() {
    return new MvhMetadataModelProjectConsentProvisionsInner.Builder();
  }

  public MvhMetadataModelProjectConsentProvisionsInner.Builder toBuilder() {
    return new MvhMetadataModelProjectConsentProvisionsInner.Builder()
        .date(getDate())
        .purpose(getPurpose())
        .type(getType());
  }
}
