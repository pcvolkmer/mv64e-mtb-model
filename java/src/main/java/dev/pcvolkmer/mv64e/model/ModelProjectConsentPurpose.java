package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

public enum ModelProjectConsentPurpose {
  CASE_IDENTIFICATION("case-identification"),

  REIDENTIFICATION("reidentification"),

  SEQUENCING("sequencing");

  private String value;

  ModelProjectConsentPurpose(String value) {
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
  public static ModelProjectConsentPurpose fromValue(String value) {
    for (ModelProjectConsentPurpose b : ModelProjectConsentPurpose.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
