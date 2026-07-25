package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MtbSystemicTherapyDosageDensityCoding {
  public enum CodeEnum {
    OVER_50_(String.valueOf("over-50%")),

    UNDER_50_(String.valueOf("under-50%"));

    private String value;

    CodeEnum(String value) {
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
    public static CodeEnum fromValue(String value) {
      for (CodeEnum b : CodeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_CODE = "code";

  private CodeEnum code;

  public static final String JSON_PROPERTY_DISPLAY = "display";

  private @Nullable String display;

  public static final String JSON_PROPERTY_SYSTEM = "system";

  private @Nullable String system;

  public static final String JSON_PROPERTY_VERSION = "version";

  private @Nullable String version;

  public MtbSystemicTherapyDosageDensityCoding() {}

  public MtbSystemicTherapyDosageDensityCoding code(CodeEnum code) {
    this.code = code;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public CodeEnum getCode() {
    return code;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public void setCode(CodeEnum code) {
    this.code = code;
  }

  public MtbSystemicTherapyDosageDensityCoding display(@Nullable String display) {
    this.display = display;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DISPLAY, required = false)
  public @Nullable String getDisplay() {
    return display;
  }

  @JsonProperty(value = JSON_PROPERTY_DISPLAY, required = false)
  public void setDisplay(@Nullable String display) {
    this.display = display;
  }

  public MtbSystemicTherapyDosageDensityCoding system(@Nullable String system) {
    this.system = system;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = false)
  public @Nullable String getSystem() {
    return system;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = false)
  public void setSystem(@Nullable String system) {
    this.system = system;
  }

  public MtbSystemicTherapyDosageDensityCoding version(@Nullable String version) {
    this.version = version;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VERSION, required = false)
  public @Nullable String getVersion() {
    return version;
  }

  @JsonProperty(value = JSON_PROPERTY_VERSION, required = false)
  public void setVersion(@Nullable String version) {
    this.version = version;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbSystemicTherapyDosageDensityCoding mtbSystemicTherapyDosageDensityCoding =
        (MtbSystemicTherapyDosageDensityCoding) o;
    return Objects.equals(this.code, mtbSystemicTherapyDosageDensityCoding.code)
        && Objects.equals(this.display, mtbSystemicTherapyDosageDensityCoding.display)
        && Objects.equals(this.system, mtbSystemicTherapyDosageDensityCoding.system)
        && Objects.equals(this.version, mtbSystemicTherapyDosageDensityCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbSystemicTherapyDosageDensityCoding {\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    display: ").append(toIndentedString(display)).append("\n");
    sb.append("    system: ").append(toIndentedString(system)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbSystemicTherapyDosageDensityCoding instance;

    public Builder() {
      this(new MtbSystemicTherapyDosageDensityCoding());
    }

    protected Builder(MtbSystemicTherapyDosageDensityCoding instance) {
      this.instance = instance;
    }

    public MtbSystemicTherapyDosageDensityCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public MtbSystemicTherapyDosageDensityCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public MtbSystemicTherapyDosageDensityCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public MtbSystemicTherapyDosageDensityCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public MtbSystemicTherapyDosageDensityCoding build() {
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

  public static MtbSystemicTherapyDosageDensityCoding.Builder builder() {
    return new MtbSystemicTherapyDosageDensityCoding.Builder();
  }

  public MtbSystemicTherapyDosageDensityCoding.Builder toBuilder() {
    return new MtbSystemicTherapyDosageDensityCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
