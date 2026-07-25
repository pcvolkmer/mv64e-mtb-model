package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

public enum MvhSubmissionType {
  INITIAL("initial"),

  CORRECTION("correction"),

  TEST("test"),

  ADDITION("addition"),

  FOLLOWUP("followup");

  private String value;

  MvhSubmissionType(String value) {
    this.value = value;
  }

  @JsonValue
  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  @JsonCreator
  public static MvhSubmissionType fromValue(String value) {
    for (MvhSubmissionType b : MvhSubmissionType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
