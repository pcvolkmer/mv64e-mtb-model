package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class MtbDiagnosisType {
  public static final String JSON_PROPERTY_HISTORY = "history";

  private List<MtbDiagnosisTypeHistoryInner> history;

  public MtbDiagnosisType() {}

  public MtbDiagnosisType history(List<MtbDiagnosisTypeHistoryInner> history) {
    this.history = history;
    return this;
  }

  public MtbDiagnosisType addHistoryItem(MtbDiagnosisTypeHistoryInner historyItem) {
    if (this.history == null) {
      this.history = new ArrayList<>();
    }
    this.history.add(historyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public List<MtbDiagnosisTypeHistoryInner> getHistory() {
    return history;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public void setHistory(List<MtbDiagnosisTypeHistoryInner> history) {
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
    MtbDiagnosisType mtbDiagnosisType = (MtbDiagnosisType) o;
    return Objects.equals(this.history, mtbDiagnosisType.history);
  }

  @Override
  public int hashCode() {
    return Objects.hash(history);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MtbDiagnosisType {\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MtbDiagnosisType instance;

    public Builder() {
      this(new MtbDiagnosisType());
    }

    protected Builder(MtbDiagnosisType instance) {
      this.instance = instance;
    }

    public MtbDiagnosisType.Builder history(List<MtbDiagnosisTypeHistoryInner> history) {
      this.instance.history = history;
      return this;
    }

    public MtbDiagnosisType build() {
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

  public static MtbDiagnosisType.Builder builder() {
    return new MtbDiagnosisType.Builder();
  }

  public MtbDiagnosisType.Builder toBuilder() {
    return new MtbDiagnosisType.Builder().history(getHistory());
  }
}
