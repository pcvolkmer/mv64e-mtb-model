package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MolecularDiagnosticReportTypeCoding {
  public enum CodeEnum {
    FUSION_PANEL(String.valueOf("fusion-panel")),

    GENOME_SHORT_READ(String.valueOf("genome-short-read")),

    ARRAY(String.valueOf("array")),

    FISH(String.valueOf("FISH")),

    EXOME(String.valueOf("exome")),

    GENE_PANEL(String.valueOf("gene-panel")),

    PCR(String.valueOf("PCR")),

    KARYOTYPING(String.valueOf("karyotyping")),

    PANEL(String.valueOf("panel")),

    SINGLE(String.valueOf("single")),

    GENOME_LONG_READ(String.valueOf("genome-long-read")),

    OTHER(String.valueOf("other"));

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

  public MolecularDiagnosticReportTypeCoding() {}

  public MolecularDiagnosticReportTypeCoding code(CodeEnum code) {
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

  public MolecularDiagnosticReportTypeCoding display(@Nullable String display) {
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

  public MolecularDiagnosticReportTypeCoding system(@Nullable String system) {
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

  public MolecularDiagnosticReportTypeCoding version(@Nullable String version) {
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
    MolecularDiagnosticReportTypeCoding molecularDiagnosticReportTypeCoding =
        (MolecularDiagnosticReportTypeCoding) o;
    return Objects.equals(this.code, molecularDiagnosticReportTypeCoding.code)
        && Objects.equals(this.display, molecularDiagnosticReportTypeCoding.display)
        && Objects.equals(this.system, molecularDiagnosticReportTypeCoding.system)
        && Objects.equals(this.version, molecularDiagnosticReportTypeCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MolecularDiagnosticReportTypeCoding {\n");
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

    private MolecularDiagnosticReportTypeCoding instance;

    public Builder() {
      this(new MolecularDiagnosticReportTypeCoding());
    }

    protected Builder(MolecularDiagnosticReportTypeCoding instance) {
      this.instance = instance;
    }

    public MolecularDiagnosticReportTypeCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public MolecularDiagnosticReportTypeCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public MolecularDiagnosticReportTypeCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public MolecularDiagnosticReportTypeCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public MolecularDiagnosticReportTypeCoding build() {
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

  public static MolecularDiagnosticReportTypeCoding.Builder builder() {
    return new MolecularDiagnosticReportTypeCoding.Builder();
  }

  public MolecularDiagnosticReportTypeCoding.Builder toBuilder() {
    return new MolecularDiagnosticReportTypeCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
