package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class CnvStartRange {
  public static final String JSON_PROPERTY_START = "start";

  private BigDecimal start;

  public static final String JSON_PROPERTY_END = "end";

  private @Nullable BigDecimal end;

  public CnvStartRange() {}

  public CnvStartRange start(BigDecimal start) {
    this.start = start;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_START, required = true)
  public BigDecimal getStart() {
    return start;
  }

  @JsonProperty(value = JSON_PROPERTY_START, required = true)
  public void setStart(BigDecimal start) {
    this.start = start;
  }

  public CnvStartRange end(@Nullable BigDecimal end) {
    this.end = end;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_END, required = false)
  public @Nullable BigDecimal getEnd() {
    return end;
  }

  @JsonProperty(value = JSON_PROPERTY_END, required = false)
  public void setEnd(@Nullable BigDecimal end) {
    this.end = end;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CnvStartRange cnvStartRange = (CnvStartRange) o;
    return Objects.equals(this.start, cnvStartRange.start)
        && Objects.equals(this.end, cnvStartRange.end);
  }

  @Override
  public int hashCode() {
    return Objects.hash(start, end);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CnvStartRange {\n");
    sb.append("    start: ").append(toIndentedString(start)).append("\n");
    sb.append("    end: ").append(toIndentedString(end)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private CnvStartRange instance;

    public Builder() {
      this(new CnvStartRange());
    }

    protected Builder(CnvStartRange instance) {
      this.instance = instance;
    }

    public CnvStartRange.Builder start(BigDecimal start) {
      this.instance.start = start;
      return this;
    }

    public CnvStartRange.Builder end(BigDecimal end) {
      this.instance.end = end;
      return this;
    }

    public CnvStartRange build() {
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

  public static CnvStartRange.Builder builder() {
    return new CnvStartRange.Builder();
  }

  public CnvStartRange.Builder toBuilder() {
    return new CnvStartRange.Builder().start(getStart()).end(getEnd());
  }
}
