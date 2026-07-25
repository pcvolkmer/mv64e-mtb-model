package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class HealthInsuranceTypeCoding {
  public enum CodeEnum {
    GKV(String.valueOf("GKV")),

    BG(String.valueOf("BG")),

    PKV(String.valueOf("PKV")),

    SKT(String.valueOf("SKT")),

    SEL(String.valueOf("SEL")),

    SOZ(String.valueOf("SOZ")),

    GPV(String.valueOf("GPV")),

    UNK(String.valueOf("UNK")),

    BEI(String.valueOf("BEI")),

    PPV(String.valueOf("PPV"));

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

  public HealthInsuranceTypeCoding() {}

  public HealthInsuranceTypeCoding code(CodeEnum code) {
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

  public HealthInsuranceTypeCoding display(@Nullable String display) {
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

  public HealthInsuranceTypeCoding system(@Nullable String system) {
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

  public HealthInsuranceTypeCoding version(@Nullable String version) {
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
    HealthInsuranceTypeCoding healthInsuranceTypeCoding = (HealthInsuranceTypeCoding) o;
    return Objects.equals(this.code, healthInsuranceTypeCoding.code)
        && Objects.equals(this.display, healthInsuranceTypeCoding.display)
        && Objects.equals(this.system, healthInsuranceTypeCoding.system)
        && Objects.equals(this.version, healthInsuranceTypeCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthInsuranceTypeCoding {\n");
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

    private HealthInsuranceTypeCoding instance;

    public Builder() {
      this(new HealthInsuranceTypeCoding());
    }

    protected Builder(HealthInsuranceTypeCoding instance) {
      this.instance = instance;
    }

    public HealthInsuranceTypeCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public HealthInsuranceTypeCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public HealthInsuranceTypeCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public HealthInsuranceTypeCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public HealthInsuranceTypeCoding build() {
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

  public static HealthInsuranceTypeCoding.Builder builder() {
    return new HealthInsuranceTypeCoding.Builder();
  }

  public HealthInsuranceTypeCoding.Builder toBuilder() {
    return new HealthInsuranceTypeCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
