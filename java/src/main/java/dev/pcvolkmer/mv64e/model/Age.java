package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import java.math.BigDecimal;
import java.util.Objects;

public class Age {
  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public enum UnitEnum {
    MONTHS(String.valueOf("Months")),

    YEARS(String.valueOf("Years"));

    private String value;

    UnitEnum(String value) {
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
    public static UnitEnum fromValue(String value) {
      for (UnitEnum b : UnitEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_UNIT = "unit";

  private UnitEnum unit;

  public Age() {}

  public Age value(BigDecimal value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public BigDecimal getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(BigDecimal value) {
    this.value = value;
  }

  public Age unit(UnitEnum unit) {
    this.unit = unit;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_UNIT, required = true)
  public UnitEnum getUnit() {
    return unit;
  }

  @JsonProperty(value = JSON_PROPERTY_UNIT, required = true)
  public void setUnit(UnitEnum unit) {
    this.unit = unit;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Age age = (Age) o;
    return Objects.equals(this.value, age.value) && Objects.equals(this.unit, age.unit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, unit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Age {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    unit: ").append(toIndentedString(unit)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Age instance;

    public Builder() {
      this(new Age());
    }

    protected Builder(Age instance) {
      this.instance = instance;
    }

    public Age.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public Age.Builder unit(UnitEnum unit) {
      this.instance.unit = unit;
      return this;
    }

    public Age build() {
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

  public static Age.Builder builder() {
    return new Age.Builder();
  }

  public Age.Builder toBuilder() {
    return new Age.Builder().value(getValue()).unit(getUnit());
  }
}
