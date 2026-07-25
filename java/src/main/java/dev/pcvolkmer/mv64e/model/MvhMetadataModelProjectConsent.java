package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MvhMetadataModelProjectConsent {
  public static final String JSON_PROPERTY_VERSION = "version";

  private String version;

  public static final String JSON_PROPERTY_DATE = "date";

  private @Nullable Date date;

  public static final String JSON_PROPERTY_PROVISIONS = "provisions";

  private List<MvhMetadataModelProjectConsentProvisionsInner> provisions;

  public MvhMetadataModelProjectConsent() {}

  public MvhMetadataModelProjectConsent version(String version) {
    this.version = version;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VERSION, required = true)
  public String getVersion() {
    return version;
  }

  @JsonProperty(value = JSON_PROPERTY_VERSION, required = true)
  public void setVersion(String version) {
    this.version = version;
  }

  public MvhMetadataModelProjectConsent date(@Nullable Date date) {
    this.date = date;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = false)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public @Nullable Date getDate() {
    return date;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = false)
  public void setDate(@Nullable Date date) {
    this.date = date;
  }

  public MvhMetadataModelProjectConsent provisions(
      List<MvhMetadataModelProjectConsentProvisionsInner> provisions) {
    this.provisions = provisions;
    return this;
  }

  public MvhMetadataModelProjectConsent addProvisionsItem(
      MvhMetadataModelProjectConsentProvisionsInner provisionsItem) {
    if (this.provisions == null) {
      this.provisions = new ArrayList<>();
    }
    this.provisions.add(provisionsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PROVISIONS, required = true)
  public List<MvhMetadataModelProjectConsentProvisionsInner> getProvisions() {
    return provisions;
  }

  @JsonProperty(value = JSON_PROPERTY_PROVISIONS, required = true)
  public void setProvisions(List<MvhMetadataModelProjectConsentProvisionsInner> provisions) {
    this.provisions = provisions;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MvhMetadataModelProjectConsent mvhMetadataModelProjectConsent =
        (MvhMetadataModelProjectConsent) o;
    return Objects.equals(this.version, mvhMetadataModelProjectConsent.version)
        && Objects.equals(this.date, mvhMetadataModelProjectConsent.date)
        && Objects.equals(this.provisions, mvhMetadataModelProjectConsent.provisions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(version, date, provisions);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MvhMetadataModelProjectConsent {\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    provisions: ").append(toIndentedString(provisions)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MvhMetadataModelProjectConsent instance;

    public Builder() {
      this(new MvhMetadataModelProjectConsent());
    }

    protected Builder(MvhMetadataModelProjectConsent instance) {
      this.instance = instance;
    }

    public MvhMetadataModelProjectConsent.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public MvhMetadataModelProjectConsent.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public MvhMetadataModelProjectConsent.Builder provisions(
        List<MvhMetadataModelProjectConsentProvisionsInner> provisions) {
      this.instance.provisions = provisions;
      return this;
    }

    public MvhMetadataModelProjectConsent build() {
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

  public static MvhMetadataModelProjectConsent.Builder builder() {
    return new MvhMetadataModelProjectConsent.Builder();
  }

  public MvhMetadataModelProjectConsent.Builder toBuilder() {
    return new MvhMetadataModelProjectConsent.Builder()
        .version(getVersion())
        .date(getDate())
        .provisions(getProvisions());
  }
}
