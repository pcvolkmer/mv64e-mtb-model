package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;

public class TranscriptId {
  public static final String JSON_PROPERTY_VALUE = "value";

  private String value;

  public enum SystemEnum {
    HTTPS_WWW_NCBI_NLM_NIH_GOV_REFSEQ(String.valueOf("https://www.ncbi.nlm.nih.gov/refseq")),

    HTTPS_WWW_ENSEMBL_ORG(String.valueOf("https://www.ensembl.org"));

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

  public TranscriptId() {}

  public TranscriptId value(String value) {
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

  public TranscriptId system(SystemEnum system) {
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
    TranscriptId transcriptId = (TranscriptId) o;
    return Objects.equals(this.value, transcriptId.value)
        && Objects.equals(this.system, transcriptId.system);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, system);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TranscriptId {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    system: ").append(toIndentedString(system)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TranscriptId instance;

    public Builder() {
      this(new TranscriptId());
    }

    protected Builder(TranscriptId instance) {
      this.instance = instance;
    }

    public TranscriptId.Builder value(String value) {
      this.instance.value = value;
      return this;
    }

    public TranscriptId.Builder system(SystemEnum system) {
      this.instance.system = system;
      return this;
    }

    public TranscriptId build() {
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

  public static TranscriptId.Builder builder() {
    return new TranscriptId.Builder();
  }

  public TranscriptId.Builder toBuilder() {
    return new TranscriptId.Builder().value(getValue()).system(getSystem());
  }
}
