package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class TmbResult {
  public static final String JSON_PROPERTY_VALUE = "value";

  private BigDecimal value;

  public static final String JSON_PROPERTY_UNIT = "unit";

  private @Nullable String unit;

  public TmbResult() {}

  public TmbResult value(BigDecimal value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public BigDecimal getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setValue(BigDecimal value) {
    this.value = value;
  }

  public TmbResult unit(@Nullable String unit) {
    this.unit = unit;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_UNIT, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public @Nullable String getUnit() {
    return unit;
  }

  @JsonProperty(value = JSON_PROPERTY_UNIT, required = false)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setUnit(@Nullable String unit) {
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
    TmbResult tmbResult = (TmbResult) o;
    return Objects.equals(this.value, tmbResult.value) && Objects.equals(this.unit, tmbResult.unit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, unit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TmbResult {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    unit: ").append(toIndentedString(unit)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TmbResult instance;

    public Builder() {
      this(new TmbResult());
    }

    protected Builder(TmbResult instance) {
      this.instance = instance;
    }

    public TmbResult.Builder value(BigDecimal value) {
      this.instance.value = value;
      return this;
    }

    public TmbResult.Builder unit(String unit) {
      this.instance.unit = unit;
      return this;
    }

    public TmbResult build() {
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

  public static TmbResult.Builder builder() {
    return new TmbResult.Builder();
  }

  public TmbResult.Builder toBuilder() {
    return new TmbResult.Builder().value(getValue()).unit(getUnit());
  }
}
