package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class PatientRecordStudyInclusionsInner {
  public static final String JSON_PROPERTY_HISTORY = "history";

  private List<StudyInclusion> history;

  public PatientRecordStudyInclusionsInner() {}

  public PatientRecordStudyInclusionsInner history(List<StudyInclusion> history) {

    this.history = history;
    return this;
  }

  public PatientRecordStudyInclusionsInner addHistoryItem(StudyInclusion historyItem) {
    if (this.history == null) {
      this.history = new ArrayList<>();
    }
    this.history.add(historyItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public List<StudyInclusion> getHistory() {
    return history;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTORY, required = true)
  public void setHistory(List<StudyInclusion> history) {
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
    PatientRecordStudyInclusionsInner patientRecordStudyInclusionsInner =
        (PatientRecordStudyInclusionsInner) o;
    return Objects.equals(this.history, patientRecordStudyInclusionsInner.history);
  }

  @Override
  public int hashCode() {
    return Objects.hash(history);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PatientRecordStudyInclusionsInner {\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PatientRecordStudyInclusionsInner instance;

    public Builder() {
      this(new PatientRecordStudyInclusionsInner());
    }

    protected Builder(PatientRecordStudyInclusionsInner instance) {
      this.instance = instance;
    }

    public PatientRecordStudyInclusionsInner.Builder history(List<StudyInclusion> history) {
      this.instance.history = history;
      return this;
    }

    public PatientRecordStudyInclusionsInner build() {
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

  public static PatientRecordStudyInclusionsInner.Builder builder() {
    return new PatientRecordStudyInclusionsInner.Builder();
  }

  public PatientRecordStudyInclusionsInner.Builder toBuilder() {
    return new PatientRecordStudyInclusionsInner.Builder().history(getHistory());
  }
}
