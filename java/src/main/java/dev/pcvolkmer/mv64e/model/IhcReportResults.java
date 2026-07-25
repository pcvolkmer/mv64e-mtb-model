package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class IhcReportResults {
  public static final String JSON_PROPERTY_PROTEIN_EXPRESSION = "proteinExpression";

  private List<ProteinExpression> proteinExpression;

  public static final String JSON_PROPERTY_MSI_MMR = "msiMmr";

  private List<ProteinExpression> msiMmr;

  public IhcReportResults() {}

  public IhcReportResults proteinExpression(List<ProteinExpression> proteinExpression) {
    this.proteinExpression = proteinExpression;
    return this;
  }

  public IhcReportResults addProteinExpressionItem(ProteinExpression proteinExpressionItem) {
    if (this.proteinExpression == null) {
      this.proteinExpression = new ArrayList<>();
    }
    this.proteinExpression.add(proteinExpressionItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN_EXPRESSION, required = true)
  public List<ProteinExpression> getProteinExpression() {
    return proteinExpression;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN_EXPRESSION, required = true)
  public void setProteinExpression(List<ProteinExpression> proteinExpression) {
    this.proteinExpression = proteinExpression;
  }

  public IhcReportResults msiMmr(List<ProteinExpression> msiMmr) {
    this.msiMmr = msiMmr;
    return this;
  }

  public IhcReportResults addMsiMmrItem(ProteinExpression msiMmrItem) {
    if (this.msiMmr == null) {
      this.msiMmr = new ArrayList<>();
    }
    this.msiMmr.add(msiMmrItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MSI_MMR, required = true)
  public List<ProteinExpression> getMsiMmr() {
    return msiMmr;
  }

  @JsonProperty(value = JSON_PROPERTY_MSI_MMR, required = true)
  public void setMsiMmr(List<ProteinExpression> msiMmr) {
    this.msiMmr = msiMmr;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IhcReportResults ihcReportResults = (IhcReportResults) o;
    return Objects.equals(this.proteinExpression, ihcReportResults.proteinExpression)
        && Objects.equals(this.msiMmr, ihcReportResults.msiMmr);
  }

  @Override
  public int hashCode() {
    return Objects.hash(proteinExpression, msiMmr);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IhcReportResults {\n");
    sb.append("    proteinExpression: ").append(toIndentedString(proteinExpression)).append("\n");
    sb.append("    msiMmr: ").append(toIndentedString(msiMmr)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private IhcReportResults instance;

    public Builder() {
      this(new IhcReportResults());
    }

    protected Builder(IhcReportResults instance) {
      this.instance = instance;
    }

    public IhcReportResults.Builder proteinExpression(List<ProteinExpression> proteinExpression) {
      this.instance.proteinExpression = proteinExpression;
      return this;
    }

    public IhcReportResults.Builder msiMmr(List<ProteinExpression> msiMmr) {
      this.instance.msiMmr = msiMmr;
      return this;
    }

    public IhcReportResults build() {
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

  public static IhcReportResults.Builder builder() {
    return new IhcReportResults.Builder();
  }

  public IhcReportResults.Builder toBuilder() {
    return new IhcReportResults.Builder()
        .proteinExpression(getProteinExpression())
        .msiMmr(getMsiMmr());
  }
}
