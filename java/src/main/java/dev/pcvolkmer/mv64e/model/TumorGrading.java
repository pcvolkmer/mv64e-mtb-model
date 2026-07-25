package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;

public class TumorGrading {
  public static final String JSON_PROPERTY_DATE = "date";

  private Date date;

  public static final String JSON_PROPERTY_CODES = "codes";

  private List<Coding> codes;

  public TumorGrading() {}

  public TumorGrading date(Date date) {
    this.date = date;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
  public Date getDate() {
    return date;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDate(Date date) {
    this.date = date;
  }

  public TumorGrading codes(List<Coding> codes) {
    this.codes = codes;
    return this;
  }

  public TumorGrading addCodesItem(Coding codesItem) {
    if (this.codes == null) {
      this.codes = new ArrayList<>();
    }
    this.codes.add(codesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CODES, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public List<Coding> getCodes() {
    return codes;
  }

  @JsonProperty(value = JSON_PROPERTY_CODES, required = true)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setCodes(List<Coding> codes) {
    this.codes = codes;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TumorGrading tumorGrading = (TumorGrading) o;
    return Objects.equals(this.date, tumorGrading.date)
        && Objects.equals(this.codes, tumorGrading.codes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, codes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TumorGrading {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    codes: ").append(toIndentedString(codes)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private TumorGrading instance;

    public Builder() {
      this(new TumorGrading());
    }

    protected Builder(TumorGrading instance) {
      this.instance = instance;
    }

    public TumorGrading.Builder date(Date date) {
      this.instance.date = date;
      return this;
    }

    public TumorGrading.Builder codes(List<Coding> codes) {
      this.instance.codes = codes;
      return this;
    }

    public TumorGrading build() {
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

  public static TumorGrading.Builder builder() {
    return new TumorGrading.Builder();
  }

  public TumorGrading.Builder toBuilder() {
    return new TumorGrading.Builder().date(getDate()).codes(getCodes());
  }
}
