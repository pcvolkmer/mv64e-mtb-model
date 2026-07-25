package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class ProteinExpression {
  public static final String JSON_PROPERTY_CPS_SCORE = "cpsScore";

  private @Nullable Integer cpsScore;

  public static final String JSON_PROPERTY_IC_SCORE = "icScore";

  private @Nullable ProteinExpressionIcScoreCoding icScore;

  public static final String JSON_PROPERTY_VALUE = "value";

  private ProteinExpressionResultCoding value;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private Reference patient;

  public static final String JSON_PROPERTY_PROTEIN = "protein";

  private Coding protein;

  public static final String JSON_PROPERTY_TC_SCORE = "tcScore";

  private @Nullable ProteinExpressionTcScoreCoding tcScore;

  public static final String JSON_PROPERTY_TPS_SCORE = "tpsScore";

  private @Nullable Integer tpsScore;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public ProteinExpression() {}

  public ProteinExpression cpsScore(@Nullable Integer cpsScore) {
    this.cpsScore = cpsScore;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CPS_SCORE, required = false)
  public @Nullable Integer getCpsScore() {
    return cpsScore;
  }

  @JsonProperty(value = JSON_PROPERTY_CPS_SCORE, required = false)
  public void setCpsScore(@Nullable Integer cpsScore) {
    this.cpsScore = cpsScore;
  }

  public ProteinExpression icScore(@Nullable ProteinExpressionIcScoreCoding icScore) {
    this.icScore = icScore;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_IC_SCORE, required = false)
  public @Nullable ProteinExpressionIcScoreCoding getIcScore() {
    return icScore;
  }

  @JsonProperty(value = JSON_PROPERTY_IC_SCORE, required = false)
  public void setIcScore(@Nullable ProteinExpressionIcScoreCoding icScore) {
    this.icScore = icScore;
  }

  public ProteinExpression value(ProteinExpressionResultCoding value) {
    this.value = value;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public ProteinExpressionResultCoding getValue() {
    return value;
  }

  @JsonProperty(value = JSON_PROPERTY_VALUE, required = true)
  public void setValue(ProteinExpressionResultCoding value) {
    this.value = value;
  }

  public ProteinExpression patient(Reference patient) {
    this.patient = patient;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  public Reference getPatient() {
    return patient;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = true)
  public void setPatient(Reference patient) {
    this.patient = patient;
  }

  public ProteinExpression protein(Coding protein) {
    this.protein = protein;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN, required = true)
  public Coding getProtein() {
    return protein;
  }

  @JsonProperty(value = JSON_PROPERTY_PROTEIN, required = true)
  public void setProtein(Coding protein) {
    this.protein = protein;
  }

  public ProteinExpression tcScore(@Nullable ProteinExpressionTcScoreCoding tcScore) {
    this.tcScore = tcScore;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TC_SCORE, required = false)
  public @Nullable ProteinExpressionTcScoreCoding getTcScore() {
    return tcScore;
  }

  @JsonProperty(value = JSON_PROPERTY_TC_SCORE, required = false)
  public void setTcScore(@Nullable ProteinExpressionTcScoreCoding tcScore) {
    this.tcScore = tcScore;
  }

  public ProteinExpression tpsScore(@Nullable Integer tpsScore) {
    this.tpsScore = tpsScore;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TPS_SCORE, required = false)
  public @Nullable Integer getTpsScore() {
    return tpsScore;
  }

  @JsonProperty(value = JSON_PROPERTY_TPS_SCORE, required = false)
  public void setTpsScore(@Nullable Integer tpsScore) {
    this.tpsScore = tpsScore;
  }

  public ProteinExpression id(String id) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProteinExpression proteinExpression = (ProteinExpression) o;
    return Objects.equals(this.cpsScore, proteinExpression.cpsScore)
        && Objects.equals(this.icScore, proteinExpression.icScore)
        && Objects.equals(this.value, proteinExpression.value)
        && Objects.equals(this.patient, proteinExpression.patient)
        && Objects.equals(this.protein, proteinExpression.protein)
        && Objects.equals(this.tcScore, proteinExpression.tcScore)
        && Objects.equals(this.tpsScore, proteinExpression.tpsScore)
        && Objects.equals(this.id, proteinExpression.id);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cpsScore, icScore, value, patient, protein, tcScore, tpsScore, id);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProteinExpression {\n");
    sb.append("    cpsScore: ").append(toIndentedString(cpsScore)).append("\n");
    sb.append("    icScore: ").append(toIndentedString(icScore)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("    protein: ").append(toIndentedString(protein)).append("\n");
    sb.append("    tcScore: ").append(toIndentedString(tcScore)).append("\n");
    sb.append("    tpsScore: ").append(toIndentedString(tpsScore)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private ProteinExpression instance;

    public Builder() {
      this(new ProteinExpression());
    }

    protected Builder(ProteinExpression instance) {
      this.instance = instance;
    }

    public ProteinExpression.Builder cpsScore(Integer cpsScore) {
      this.instance.cpsScore = cpsScore;
      return this;
    }

    public ProteinExpression.Builder icScore(ProteinExpressionIcScoreCoding icScore) {
      this.instance.icScore = icScore;
      return this;
    }

    public ProteinExpression.Builder value(ProteinExpressionResultCoding value) {
      this.instance.value = value;
      return this;
    }

    public ProteinExpression.Builder patient(Reference patient) {
      this.instance.patient = patient;
      return this;
    }

    public ProteinExpression.Builder protein(Coding protein) {
      this.instance.protein = protein;
      return this;
    }

    public ProteinExpression.Builder tcScore(ProteinExpressionTcScoreCoding tcScore) {
      this.instance.tcScore = tcScore;
      return this;
    }

    public ProteinExpression.Builder tpsScore(Integer tpsScore) {
      this.instance.tpsScore = tpsScore;
      return this;
    }

    public ProteinExpression.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public ProteinExpression build() {
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

  public static ProteinExpression.Builder builder() {
    return new ProteinExpression.Builder();
  }

  public ProteinExpression.Builder toBuilder() {
    return new ProteinExpression.Builder()
        .cpsScore(getCpsScore())
        .icScore(getIcScore())
        .value(getValue())
        .patient(getPatient())
        .protein(getProtein())
        .tcScore(getTcScore())
        .tpsScore(getTpsScore())
        .id(getId());
  }
}
