package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class PatientRecordSystemicTherapiesInner {
  public static final String JSON_PROPERTY_HISTORY = "history";

  private List<MtbSystemicTherapy> history;

  public PatientRecordSystemicTherapiesInner() {}

  public PatientRecordSystemicTherapiesInner history(List<MtbSystemicTherapy> history) {
    this.history = history;
    return this;
  }

  public PatientRecordSystemicTherapiesInner addHistoryItem(MtbSystemicTherapy historyItem) {
    if (this.history == null) {
      this.history = new ArrayList<>();
    }
    this.history.add(historyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public List<MtbSystemicTherapy> getHistory() {
    return history;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public void setHistory(List<MtbSystemicTherapy> history) {
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
    PatientRecordSystemicTherapiesInner patientRecordSystemicTherapiesInner =
        (PatientRecordSystemicTherapiesInner) o;
    return Objects.equals(this.history, patientRecordSystemicTherapiesInner.history);
  }

  @Override
  public int hashCode() {
    return Objects.hash(history);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PatientRecordSystemicTherapiesInner {\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PatientRecordSystemicTherapiesInner instance;

    public Builder() {
      this(new PatientRecordSystemicTherapiesInner());
    }

    protected Builder(PatientRecordSystemicTherapiesInner instance) {
      this.instance = instance;
    }

    public PatientRecordSystemicTherapiesInner.Builder history(List<MtbSystemicTherapy> history) {
      this.instance.history = history;
      return this;
    }

    public PatientRecordSystemicTherapiesInner build() {
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

  public static PatientRecordSystemicTherapiesInner.Builder builder() {
    return new PatientRecordSystemicTherapiesInner.Builder();
  }

  public PatientRecordSystemicTherapiesInner.Builder toBuilder() {
    return new PatientRecordSystemicTherapiesInner.Builder().history(getHistory());
  }
}
