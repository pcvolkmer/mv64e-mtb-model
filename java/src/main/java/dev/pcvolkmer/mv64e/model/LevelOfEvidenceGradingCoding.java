package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class LevelOfEvidenceGradingCoding {
  public enum CodeEnum {
    M1_A(String.valueOf("m1A")),

    M2_A(String.valueOf("m2A")),

    M2_C(String.valueOf("m2C")),

    M2_B(String.valueOf("m2B")),

    M1_B(String.valueOf("m1B")),

    UNDEFINED(String.valueOf("undefined")),

    M4(String.valueOf("m4")),

    M1_C(String.valueOf("m1C")),

    M3(String.valueOf("m3"));

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

  public LevelOfEvidenceGradingCoding() {}

  public LevelOfEvidenceGradingCoding code(CodeEnum code) {
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

  public LevelOfEvidenceGradingCoding display(@Nullable String display) {
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

  public LevelOfEvidenceGradingCoding system(@Nullable String system) {
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

  public LevelOfEvidenceGradingCoding version(@Nullable String version) {
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
    LevelOfEvidenceGradingCoding levelOfEvidenceGradingCoding = (LevelOfEvidenceGradingCoding) o;
    return Objects.equals(this.code, levelOfEvidenceGradingCoding.code)
        && Objects.equals(this.display, levelOfEvidenceGradingCoding.display)
        && Objects.equals(this.system, levelOfEvidenceGradingCoding.system)
        && Objects.equals(this.version, levelOfEvidenceGradingCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LevelOfEvidenceGradingCoding {\n");
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

    private LevelOfEvidenceGradingCoding instance;

    public Builder() {
      this(new LevelOfEvidenceGradingCoding());
    }

    protected Builder(LevelOfEvidenceGradingCoding instance) {
      this.instance = instance;
    }

    public LevelOfEvidenceGradingCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public LevelOfEvidenceGradingCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public LevelOfEvidenceGradingCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public LevelOfEvidenceGradingCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public LevelOfEvidenceGradingCoding build() {
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

  public static LevelOfEvidenceGradingCoding.Builder builder() {
    return new LevelOfEvidenceGradingCoding.Builder();
  }

  public LevelOfEvidenceGradingCoding.Builder toBuilder() {
    return new LevelOfEvidenceGradingCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
