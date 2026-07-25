package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class HistologyReportResults {
  public static final String JSON_PROPERTY_TUMOR_MORPHOLOGY = "tumorMorphology";

  private HistologyReportResultsTumorMorphology tumorMorphology;

  public static final String JSON_PROPERTY_TUMOR_CELL_CONTENT = "tumorCellContent";

  private @Nullable TumorCellContent tumorCellContent;

  public HistologyReportResults() {}

  public HistologyReportResults tumorMorphology(
      HistologyReportResultsTumorMorphology tumorMorphology) {
    this.tumorMorphology = tumorMorphology;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_MORPHOLOGY, required = true)
  public HistologyReportResultsTumorMorphology getTumorMorphology() {
    return tumorMorphology;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_MORPHOLOGY, required = true)
  public void setTumorMorphology(HistologyReportResultsTumorMorphology tumorMorphology) {
    this.tumorMorphology = tumorMorphology;
  }

  public HistologyReportResults tumorCellContent(@Nullable TumorCellContent tumorCellContent) {
    this.tumorCellContent = tumorCellContent;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_CELL_CONTENT, required = false)
  public @Nullable TumorCellContent getTumorCellContent() {
    return tumorCellContent;
  }

  @JsonProperty(value = JSON_PROPERTY_TUMOR_CELL_CONTENT, required = false)
  public void setTumorCellContent(@Nullable TumorCellContent tumorCellContent) {
    this.tumorCellContent = tumorCellContent;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HistologyReportResults histologyReportResults = (HistologyReportResults) o;
    return Objects.equals(this.tumorMorphology, histologyReportResults.tumorMorphology)
        && Objects.equals(this.tumorCellContent, histologyReportResults.tumorCellContent);
  }

  @Override
  public int hashCode() {
    return Objects.hash(tumorMorphology, tumorCellContent);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HistologyReportResults {\n");
    sb.append("    tumorMorphology: ").append(toIndentedString(tumorMorphology)).append("\n");
    sb.append("    tumorCellContent: ").append(toIndentedString(tumorCellContent)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private HistologyReportResults instance;

    public Builder() {
      this(new HistologyReportResults());
    }

    protected Builder(HistologyReportResults instance) {
      this.instance = instance;
    }

    public HistologyReportResults.Builder tumorMorphology(
        HistologyReportResultsTumorMorphology tumorMorphology) {
      this.instance.tumorMorphology = tumorMorphology;
      return this;
    }

    public HistologyReportResults.Builder tumorCellContent(TumorCellContent tumorCellContent) {
      this.instance.tumorCellContent = tumorCellContent;
      return this;
    }

    public HistologyReportResults build() {
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

  public static HistologyReportResults.Builder builder() {
    return new HistologyReportResults.Builder();
  }

  public HistologyReportResults.Builder toBuilder() {
    return new HistologyReportResults.Builder()
        .tumorMorphology(getTumorMorphology())
        .tumorCellContent(getTumorCellContent());
  }
}
