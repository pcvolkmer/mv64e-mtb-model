package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MtbTherapyStatusReasonCoding {
  public enum CodeEnum {
    PATIENT_WISH(String.valueOf("patient-wish")),

    PROGRESSION(String.valueOf("progression")),

    PAYMENT_REFUSED(String.valueOf("payment-refused")),

    DETERIORATION(String.valueOf("deterioration")),

    PATIENT_DEATH(String.valueOf("patient-death")),

    REGULAR_COMPLETION(String.valueOf("regular-completion")),

    REGULAR_COMPLETION_WITH_DOSAGE_REDUCTION(
        String.valueOf("regular-completion-with-dosage-reduction")),

    BEST_SUPPORTIVE_CARE(String.valueOf("best-supportive-care")),

    LOST_TO_FU(String.valueOf("lost-to-fu")),

    CHRONIC_REMISSION(String.valueOf("chronic-remission")),

    TOXICITY(String.valueOf("toxicity")),

    PAYMENT_PENDING(String.valueOf("payment-pending")),

    PATIENT_REFUSAL(String.valueOf("patient-refusal")),

    NO_INDICATION(String.valueOf("no-indication")),

    REGULAR_COMPLETION_WITH_SUBSTANCE_CHANGE(
        String.valueOf("regular-completion-with-substance-change")),

    PAYMENT_ENDED(String.valueOf("payment-ended")),

    OTHER_THERAPY_CHOSEN(String.valueOf("other-therapy-chosen")),

    MEDICAL_REASONS(String.valueOf("medical-reasons")),

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

  public MtbTherapyStatusReasonCoding() {}

  public MtbTherapyStatusReasonCoding code(CodeEnum code) {
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

  public MtbTherapyStatusReasonCoding display(@Nullable String display) {
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

  public MtbTherapyStatusReasonCoding system(@Nullable String system) {
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

  public MtbTherapyStatusReasonCoding version(@Nullable String version) {
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
    MtbTherapyStatusReasonCoding mtbTherapyStatusReasonCoding = (MtbTherapyStatusReasonCoding) o;
    return Objects.equals(this.code, mtbTherapyStatusReasonCoding.code)
        && Objects.equals(this.display, mtbTherapyStatusReasonCoding.display)
        && Objects.equals(this.system, mtbTherapyStatusReasonCoding.system)
        && Objects.equals(this.version, mtbTherapyStatusReasonCoding.version);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, display, system, version);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbTherapyStatusReasonCoding {\n");
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

    private MtbTherapyStatusReasonCoding instance;

    public Builder() {
      this(new MtbTherapyStatusReasonCoding());
    }

    protected Builder(MtbTherapyStatusReasonCoding instance) {
      this.instance = instance;
    }

    public MtbTherapyStatusReasonCoding.Builder code(CodeEnum code) {
      this.instance.code = code;
      return this;
    }

    public MtbTherapyStatusReasonCoding.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public MtbTherapyStatusReasonCoding.Builder system(String system) {
      this.instance.system = system;
      return this;
    }

    public MtbTherapyStatusReasonCoding.Builder version(String version) {
      this.instance.version = version;
      return this;
    }

    public MtbTherapyStatusReasonCoding build() {
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

  public static MtbTherapyStatusReasonCoding.Builder builder() {
    return new MtbTherapyStatusReasonCoding.Builder();
  }

  public MtbTherapyStatusReasonCoding.Builder toBuilder() {
    return new MtbTherapyStatusReasonCoding.Builder()
        .code(getCode())
        .display(getDisplay())
        .system(getSystem())
        .version(getVersion());
  }
}
