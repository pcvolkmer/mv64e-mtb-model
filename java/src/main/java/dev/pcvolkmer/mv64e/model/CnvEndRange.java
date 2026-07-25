package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class CnvEndRange {
  public static final String JSON_PROPERTY_START = "start";

  private BigDecimal start;

  public static final String JSON_PROPERTY_END = "end";

  private @Nullable BigDecimal end;

  public CnvEndRange() {}

  public CnvEndRange start(BigDecimal start) {
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

  public CnvEndRange end(@Nullable BigDecimal end) {
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
    CnvEndRange cnvEndRange = (CnvEndRange) o;
    return Objects.equals(this.start, cnvEndRange.start)
        && Objects.equals(this.end, cnvEndRange.end);
  }

  @Override
  public int hashCode() {
    return Objects.hash(start, end);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CnvEndRange {\n");
    sb.append("    start: ").append(toIndentedString(start)).append("\n");
    sb.append("    end: ").append(toIndentedString(end)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private CnvEndRange instance;

    public Builder() {
      this(new CnvEndRange());
    }

    protected Builder(CnvEndRange instance) {
      this.instance = instance;
    }

    public CnvEndRange.Builder start(BigDecimal start) {
      this.instance.start = start;
      return this;
    }

    public CnvEndRange.Builder end(BigDecimal end) {
      this.instance.end = end;
      return this;
    }

    public CnvEndRange build() {
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

  public static CnvEndRange.Builder builder() {
    return new CnvEndRange.Builder();
  }

  public CnvEndRange.Builder toBuilder() {
    return new CnvEndRange.Builder().start(getStart()).end(getEnd());
  }
}
