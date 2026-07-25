package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class MtbDiagnosisGrading {
  public static final String JSON_PROPERTY_HISTORY = "history";

  private List<TumorGrading> history;

  public MtbDiagnosisGrading() {}

  public MtbDiagnosisGrading history(List<TumorGrading> history) {
    this.history = history;
    return this;
  }

  public MtbDiagnosisGrading addHistoryItem(TumorGrading historyItem) {
    if (this.history == null) {
      this.history = new ArrayList<>();
    }
    this.history.add(historyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public List<TumorGrading> getHistory() {
    return history;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public void setHistory(List<TumorGrading> history) {
    this.history = history;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MtbDiagnosisGrading mtbDiagnosisGrading = (MtbDiagnosisGrading) o;
    return Objects.equals(this.history, mtbDiagnosisGrading.history);
  }

  @Override
  public int hashCode() {
    return Objects.hash(history);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbDiagnosisGrading {\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbDiagnosisGrading instance;

    public Builder() {
      this(new MtbDiagnosisGrading());
    }

    protected Builder(MtbDiagnosisGrading instance) {
      this.instance = instance;
    }

    public MtbDiagnosisGrading.Builder history(List<TumorGrading> history) {
      this.instance.history = history;
      return this;
    }

    public MtbDiagnosisGrading build() {
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

  public static MtbDiagnosisGrading.Builder builder() {
    return new MtbDiagnosisGrading.Builder();
  }

  public MtbDiagnosisGrading.Builder toBuilder() {
    return new MtbDiagnosisGrading.Builder().history(getHistory());
  }
}
