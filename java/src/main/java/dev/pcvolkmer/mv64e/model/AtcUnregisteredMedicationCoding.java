package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class AtcUnregisteredMedicationCoding {
  public static final String JSON_PROPERTY_CODE = "code";

  private String code;

  public static final String JSON_PROPERTY_DISPLAY = "display";

  private @Nullable String display;

  public enum SystemEnum {
    UNDEFINED(String.valueOf("undefined")),

    HTTP_FHIR_DE_CODE_SYSTEM_BFARM_ATC(String.valueOf("http://fhir.de/CodeSystem/bfarm/atc"));

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

  public static final String JSON_PROPERTY_VERSION = "version";

  private @Nullable String version;

  public AtcUnregisteredMedicationCoding() {}

  public AtcUnregisteredMedicationCoding code(String code) {
    this.code = code;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public String getCode() {
    return code;
  }

  @JsonProperty(value = JSON_PROPERTY_CODE, required = true)
  public void setCode(String code) {
    this.code = code;
  }

  public AtcUnregisteredMedicationCoding display(@Nullable String display) {
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

  public AtcUnregisteredMedicationCoding system(SystemEnum system) {
    this.system = system;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = true)
  public SystemEnum getSystem() {
    return system;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEM, required = true)
  public void setSystem(SystemEnum system) {
    this.system = system;
  }

  public AtcUnregisteredMedicationCoding version(@Nullable String version) {
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
    AtcUnregisteredMedicationCoding atcUnregisteredMedicationCoding =
        (AtcUnregisteredMedicationCoding) o;
    return Objects.equals(this.code, atcUnregisteredMedicationCoding.code)
        && Objects.equals(this.display, atcUnregisteredMedicationCoding.display)
        && Objects.equals(this.system, atcUnregisteredMedicationCoding.system)
        && Objects.equals(this.version, atcUnregisteredMedicationCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AtcUnregisteredMedicationCoding {\n");
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

    private AtcUnregisteredMedicationCoding instance;

    public Builder() {
      this(new AtcUnregisteredMedicationCoding());
    }

    protected Builder(AtcUnregisteredMedicationCoding instance) {
      this.instance = instance;
    }

    public AtcUnregisteredMedicationCoding.Builder code(String code) {
      this.instance.code = code;
      return this;
    }

    public AtcUnregisteredMedicationCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public AtcUnregisteredMedicationCoding.Builder system(SystemEnum system) {
      this.instance.system = system;
      return this;
    }

    public AtcUnregisteredMedicationCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public AtcUnregisteredMedicationCoding build() {
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

  public static AtcUnregisteredMedicationCoding.Builder builder() {
    return new AtcUnregisteredMedicationCoding.Builder();
  }

  public AtcUnregisteredMedicationCoding.Builder toBuilder() {
    return new AtcUnregisteredMedicationCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
