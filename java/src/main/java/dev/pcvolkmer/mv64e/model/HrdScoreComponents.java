package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.math.BigDecimal;
import java.util.Objects;

public class HrdScoreComponents {
  public static final String JSON_PROPERTY_LST = "lst";

  private BigDecimal lst;

  public static final String JSON_PROPERTY_LOH = "loh";

  private BigDecimal loh;

  public static final String JSON_PROPERTY_TAI = "tai";

  private BigDecimal tai;

  public HrdScoreComponents() {}

  public HrdScoreComponents lst(BigDecimal lst) {
    this.lst = lst;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LST, required = true)
  public BigDecimal getLst() {
    return lst;
  }

  @JsonProperty(value = JSON_PROPERTY_LST, required = true)
  public void setLst(BigDecimal lst) {
    this.lst = lst;
  }

  public HrdScoreComponents loh(BigDecimal loh) {
    this.loh = loh;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_LOH, required = true)
  public BigDecimal getLoh() {
    return loh;
  }

  @JsonProperty(value = JSON_PROPERTY_LOH, required = true)
  public void setLoh(BigDecimal loh) {
    this.loh = loh;
  }

  public HrdScoreComponents tai(BigDecimal tai) {
    this.tai = tai;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TAI, required = true)
  public BigDecimal getTai() {
    return tai;
  }

  @JsonProperty(value = JSON_PROPERTY_TAI, required = true)
  public void setTai(BigDecimal tai) {
    this.tai = tai;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HrdScoreComponents hrdScoreComponents = (HrdScoreComponents) o;
    return Objects.equals(this.lst, hrdScoreComponents.lst)
        && Objects.equals(this.loh, hrdScoreComponents.loh)
        && Objects.equals(this.tai, hrdScoreComponents.tai);
  }

  @Override
  public int hashCode() {
    return Objects.hash(lst, loh, tai);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HrdScoreComponents {\n");
    sb.append("    lst: ").append(toIndentedString(lst)).append("\n");
    sb.append("    loh: ").append(toIndentedString(loh)).append("\n");
    sb.append("    tai: ").append(toIndentedString(tai)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private HrdScoreComponents instance;

    public Builder() {
      this(new HrdScoreComponents());
    }

    protected Builder(HrdScoreComponents instance) {
      this.instance = instance;
    }

    public HrdScoreComponents.Builder lst(BigDecimal lst) {
      this.instance.lst = lst;
      return this;
    }

    public HrdScoreComponents.Builder loh(BigDecimal loh) {
      this.instance.loh = loh;
      return this;
    }

    public HrdScoreComponents.Builder tai(BigDecimal tai) {
      this.instance.tai = tai;
      return this;
    }

    public HrdScoreComponents build() {
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

  public static HrdScoreComponents.Builder builder() {
    return new HrdScoreComponents.Builder();
  }

  public HrdScoreComponents.Builder toBuilder() {
    return new HrdScoreComponents.Builder().lst(getLst()).loh(getLoh()).tai(getTai());
  }
}
