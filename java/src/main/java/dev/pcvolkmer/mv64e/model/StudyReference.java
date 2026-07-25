package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class StudyReference {
  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public enum SystemEnum {
    NCT(String.valueOf("NCT")),

    EUDRA_CT(String.valueOf("Eudra-CT")),

    OTHER(String.valueOf("other")),

    EUDAMED(String.valueOf("EUDAMED")),

    DRKS(String.valueOf("DRKS"));

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

  public static final String JSON_PROPERTY_DISPLAY = "display";

  private @Nullable String display;

  public static final String JSON_PROPERTY_TYPE = "type";

  private @Nullable String type;

  public StudyReference() {}

  public StudyReference id(String id) {
    this.id = id;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public String getId() {
    return id;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public void setId(String id) {
    this.id = id;
  }

  public StudyReference system(SystemEnum system) {
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

  public StudyReference display(@Nullable String display) {
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

  public StudyReference type(@Nullable String type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = false)
  public @Nullable String getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = false)
  public void setType(@Nullable String type) {
    this.type = type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    StudyReference studyReference = (StudyReference) o;
    return Objects.equals(this.id, studyReference.id)
        && Objects.equals(this.system, studyReference.system)
        && Objects.equals(this.display, studyReference.display)
        && Objects.equals(this.type, studyReference.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, system, display, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class StudyReference {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    system: ").append(toIndentedString(system)).append("\n");
    sb.append("    display: ").append(toIndentedString(display)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private StudyReference instance;

    public Builder() {
      this(new StudyReference());
    }

    protected Builder(StudyReference instance) {
      this.instance = instance;
    }

    public StudyReference.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public StudyReference.Builder system(SystemEnum system) {
      this.instance.system = system;
      return this;
    }

    public StudyReference.Builder display(String display) {
      this.instance.display = display;
      return this;
    }

    public StudyReference.Builder type(String type) {
      this.instance.type = type;
      return this;
    }

    public StudyReference build() {
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

  public static StudyReference.Builder builder() {
    return new StudyReference.Builder();
  }

  public StudyReference.Builder toBuilder() {
    return new StudyReference.Builder()
        .id(getId())
        .system(getSystem())
        .display(getDisplay())
        .type(getType());
  }
}
