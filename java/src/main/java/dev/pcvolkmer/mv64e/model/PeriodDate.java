package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.Date;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class PeriodDate {
  public static final String JSON_PROPERTY_START = "start";

  private Date start;

  public static final String JSON_PROPERTY_END = "end";

  private @Nullable Date end;

  public PeriodDate() {}

  public PeriodDate start(Date start) {
    this.start = start;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_START, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getStart() {
    return start;
  }

  @JsonProperty(value = JSON_PROPERTY_START, required = true)
  public void setStart(Date start) {
    this.start = start;
  }

  public PeriodDate end(@Nullable Date end) {
    this.end = end;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_END, required = false)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public @Nullable Date getEnd() {
    return end;
  }

  @JsonProperty(value = JSON_PROPERTY_END, required = false)
  public void setEnd(@Nullable Date end) {
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
    PeriodDate periodDate = (PeriodDate) o;
    return Objects.equals(this.start, periodDate.start) && Objects.equals(this.end, periodDate.end);
  }

  @Override
  public int hashCode() {
    return Objects.hash(start, end);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PeriodDate {\n");
    sb.append("    start: ").append(toIndentedString(start)).append("\n");
    sb.append("    end: ").append(toIndentedString(end)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PeriodDate instance;

    public Builder() {
      this(new PeriodDate());
    }

    protected Builder(PeriodDate instance) {
      this.instance = instance;
    }

    public PeriodDate.Builder start(Date start) {
      this.instance.start = start;
      return this;
    }

    public PeriodDate.Builder end(Date end) {
      this.instance.end = end;
      return this;
    }

    public PeriodDate build() {
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

  public static PeriodDate.Builder builder() {
    return new PeriodDate.Builder();
  }

  public PeriodDate.Builder toBuilder() {
    return new PeriodDate.Builder().start(getStart()).end(getEnd());
  }
}
