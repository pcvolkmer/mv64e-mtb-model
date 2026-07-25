package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;

public class VariantExternalId {
  public static final String JSON_PROPERTY_VALUE = "value";

  private String value;

  public enum SystemEnum {
    HTTPS_WWW_NCBI_NLM_NIH_GOV_ENTREZ(String.valueOf("https://www.ncbi.nlm.nih.gov/entrez")),

    HTTPS_WWW_ENSEMBL_ORG(String.valueOf("https://www.ensembl.org")),

    HTTPS_WWW_NCBI_NLM_NIH_GOV_SNP(String.valueOf("https://www.ncbi.nlm.nih.gov/snp")),

    HTTPS_CANCER_SANGER_AC_UK_COSMIC(String.valueOf("https://cancer.sanger.ac.uk/cosmic"));

    private String value;

    SystemEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static SystemEnum fromValue(String value) {
      for (SystemEnum b : SystemEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_SYSTEM = "system";

  private SystemEnum system;

  public VariantExternalId() {}

  public VariantExternalId value(String value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public String getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setValue(String value) {
    this.value = value;
  }

  public VariantExternalId system(SystemEnum system) {
    this.system = system;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public SystemEnum getSystem() {
    return system;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSystem(SystemEnum system) {
    this.system = system;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    VariantExternalId variantExternalId = (VariantExternalId) o;
    return Objects.equals(this.value, variantExternalId.value)
        && Objects.equals(this.system, variantExternalId.system);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, system);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class VariantExternalId {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    system: ").append(toIndentedString(system)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private VariantExternalId instance;

    public Builder() {
      this(new VariantExternalId());
    }

    protected Builder(VariantExternalId instance) {
      this.instance = instance;
    }

    public VariantExternalId.Builder value(String value) {
      this.instance.value = value;
      return this;
    }

    public VariantExternalId.Builder system(SystemEnum system) {
      this.instance.system = system;
      return this;
    }

    public VariantExternalId build() {
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

  public static VariantExternalId.Builder builder() {
    return new VariantExternalId.Builder();
  }

  public VariantExternalId.Builder toBuilder() {
    return new VariantExternalId.Builder().value(getValue()).system(getSystem());
  }
}
