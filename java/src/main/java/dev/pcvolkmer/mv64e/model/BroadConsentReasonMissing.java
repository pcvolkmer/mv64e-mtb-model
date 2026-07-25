package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

public enum BroadConsentReasonMissing {
  ORGANIZATIONAL_ISSUES("organizational-issues"),

  TECHNICAL_ISSUES("technical-issues"),

  PATIENT_INABILITY("patient-inability"),

  OTHER_PATIENT_REASON("other-patient-reason"),

  CONSENT_NOT_RETURNED("consent-not-returned"),

  PATIENT_REFUSAL("patient-refusal");

  private String value;

  BroadConsentReasonMissing(String value) {
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
  public static BroadConsentReasonMissing fromValue(String value) {
    for (BroadConsentReasonMissing b : BroadConsentReasonMissing.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
