package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class ProteinExpressionResultCoding {
  public enum CodeEnum {
    EXP(String.valueOf("exp")),

    UNKNOWN(String.valueOf("unknown")),

    _1_(String.valueOf("1+")),

    _3_(String.valueOf("3+")),

    NOT_EXP(String.valueOf("not-exp")),

    _2_(String.valueOf("2+"));

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

  public ProteinExpressionResultCoding() {}

  public ProteinExpressionResultCoding code(CodeEnum code) {
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

  public ProteinExpressionResultCoding display(@Nullable String display) {
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

  public ProteinExpressionResultCoding system(@Nullable String system) {
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

  public ProteinExpressionResultCoding version(@Nullable String version) {
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
    ProteinExpressionResultCoding proteinExpressionResultCoding = (ProteinExpressionResultCoding) o;
    return Objects.equals(this.code, proteinExpressionResultCoding.code)
        && Objects.equals(this.display, proteinExpressionResultCoding.display)
        && Objects.equals(this.system, proteinExpressionResultCoding.system)
        && Objects.equals(this.version, proteinExpressionResultCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProteinExpressionResultCoding {\n");
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

    private ProteinExpressionResultCoding instance;

    public Builder() {
      this(new ProteinExpressionResultCoding());
    }

    protected Builder(ProteinExpressionResultCoding instance) {
      this.instance = instance;
    }

    public ProteinExpressionResultCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public ProteinExpressionResultCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public ProteinExpressionResultCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public ProteinExpressionResultCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public ProteinExpressionResultCoding build() {
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

  public static ProteinExpressionResultCoding.Builder builder() {
    return new ProteinExpressionResultCoding.Builder();
  }

  public ProteinExpressionResultCoding.Builder toBuilder() {
    return new ProteinExpressionResultCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
