package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class MtbDiagnosisStaging {
  public static final String JSON_PROPERTY_HISTORY = "history";

  private List<TumorStaging> history;

  public MtbDiagnosisStaging() {}

  public MtbDiagnosisStaging history(List<TumorStaging> history) {
    this.history = history;
    return this;
  }

  public MtbDiagnosisStaging addHistoryItem(TumorStaging historyItem) {
    if (this.history == null) {
      this.history = new ArrayList<>();
    }
    this.history.add(historyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public List<TumorStaging> getHistory() {
    return history;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public void setHistory(List<TumorStaging> history) {
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
    MtbDiagnosisStaging mtbDiagnosisStaging = (MtbDiagnosisStaging) o;
    return Objects.equals(this.history, mtbDiagnosisStaging.history);
  }

  @Override
  public int hashCode() {
    return Objects.hash(history);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbDiagnosisStaging {\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbDiagnosisStaging instance;

    public Builder() {
      this(new MtbDiagnosisStaging());
    }

    protected Builder(MtbDiagnosisStaging instance) {
      this.instance = instance;
    }

    public MtbDiagnosisStaging.Builder history(List<TumorStaging> history) {
      this.instance.history = history;
      return this;
    }

    public MtbDiagnosisStaging build() {
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

  public static MtbDiagnosisStaging.Builder builder() {
    return new MtbDiagnosisStaging.Builder();
  }

  public MtbDiagnosisStaging.Builder toBuilder() {
    return new MtbDiagnosisStaging.Builder().history(getHistory());
  }
}
