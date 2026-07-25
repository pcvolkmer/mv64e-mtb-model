package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class MvhMetadata {
  public static final String JSON_PROPERTY_MODEL_PROJECT_CONSENT = "modelProjectConsent";

  private MvhMetadataModelProjectConsent modelProjectConsent;

  public static final String JSON_PROPERTY_RESEARCH_CONSENTS = "researchConsents";

  private @Nullable List<Map<String, Object>> researchConsents;

  public static final String JSON_PROPERTY_EPISODE_OF_CARE = "episodeOfCare";

  private @Nullable Reference episodeOfCare;

  public static final String JSON_PROPERTY_REASON_RESEARCH_CONSENT_MISSING =
      "reasonResearchConsentMissing";

  private @Nullable BroadConsentReasonMissing reasonResearchConsentMissing;

  public static final String JSON_PROPERTY_TRANSFER_T_A_N = "transferTAN";

  private String transferTAN;

  public static final String JSON_PROPERTY_TYPE = "type";

  private MvhSubmissionType type;

  public MvhMetadata() {}

  public MvhMetadata modelProjectConsent(MvhMetadataModelProjectConsent modelProjectConsent) {
    this.modelProjectConsent = modelProjectConsent;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MODEL_PROJECT_CONSENT, required = true)
  public MvhMetadataModelProjectConsent getModelProjectConsent() {
    return modelProjectConsent;
  }

  @JsonProperty(value = JSON_PROPERTY_MODEL_PROJECT_CONSENT, required = true)
  public void setModelProjectConsent(MvhMetadataModelProjectConsent modelProjectConsent) {
    this.modelProjectConsent = modelProjectConsent;
  }

  public MvhMetadata researchConsents(@Nullable List<Map<String, Object>> researchConsents) {
    this.researchConsents = researchConsents;
    return this;
  }

  public MvhMetadata addResearchConsentsItem(Map<String, Object> researchConsentsItem) {
    if (this.researchConsents == null) {
      this.researchConsents = new ArrayList<>();
    }
    this.researchConsents.add(researchConsentsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RESEARCH_CONSENTS, required = false)
  public @Nullable List<Map<String, Object>> getResearchConsents() {
    return researchConsents;
  }

  @JsonProperty(value = JSON_PROPERTY_RESEARCH_CONSENTS, required = false)
  public void setResearchConsents(@Nullable List<Map<String, Object>> researchConsents) {
    this.researchConsents = researchConsents;
  }

  public MvhMetadata episodeOfCare(@Nullable Reference episodeOfCare) {
    this.episodeOfCare = episodeOfCare;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EPISODE_OF_CARE, required = false)
  public @Nullable Reference getEpisodeOfCare() {
    return episodeOfCare;
  }

  @JsonProperty(value = JSON_PROPERTY_EPISODE_OF_CARE, required = false)
  public void setEpisodeOfCare(@Nullable Reference episodeOfCare) {
    this.episodeOfCare = episodeOfCare;
  }

  public MvhMetadata reasonResearchConsentMissing(
      @Nullable BroadConsentReasonMissing reasonResearchConsentMissing) {
    this.reasonResearchConsentMissing = reasonResearchConsentMissing;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON_RESEARCH_CONSENT_MISSING, required = false)
  public @Nullable BroadConsentReasonMissing getReasonResearchConsentMissing() {
    return reasonResearchConsentMissing;
  }

  @JsonProperty(value = JSON_PROPERTY_REASON_RESEARCH_CONSENT_MISSING, required = false)
  public void setReasonResearchConsentMissing(
      @Nullable BroadConsentReasonMissing reasonResearchConsentMissing) {
    this.reasonResearchConsentMissing = reasonResearchConsentMissing;
  }

  public MvhMetadata transferTAN(String transferTAN) {
    this.transferTAN = transferTAN;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSFER_T_A_N, required = true)
  public String getTransferTAN() {
    return transferTAN;
  }

  @JsonProperty(value = JSON_PROPERTY_TRANSFER_T_A_N, required = true)
  public void setTransferTAN(String transferTAN) {
    this.transferTAN = transferTAN;
  }

  public MvhMetadata type(MvhSubmissionType type) {
    this.type = type;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public MvhSubmissionType getType() {
    return type;
  }

  @JsonProperty(value = JSON_PROPERTY_TYPE, required = true)
  public void setType(MvhSubmissionType type) {
    this.type = type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MvhMetadata mvhMetadata = (MvhMetadata) o;
    return Objects.equals(this.modelProjectConsent, mvhMetadata.modelProjectConsent)
        && Objects.equals(this.researchConsents, mvhMetadata.researchConsents)
        && Objects.equals(this.episodeOfCare, mvhMetadata.episodeOfCare)
        && Objects.equals(
            this.reasonResearchConsentMissing, mvhMetadata.reasonResearchConsentMissing)
        && Objects.equals(this.transferTAN, mvhMetadata.transferTAN)
        && Objects.equals(this.type, mvhMetadata.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        modelProjectConsent,
        researchConsents,
        episodeOfCare,
        reasonResearchConsentMissing,
        transferTAN,
        type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MvhMetadata {\n");
    sb.append("    modelProjectConsent: ")
        .append(toIndentedString(modelProjectConsent))
        .append("\n");
    sb.append("    researchConsents: ").append(toIndentedString(researchConsents)).append("\n");
    sb.append("    episodeOfCare: ").append(toIndentedString(episodeOfCare)).append("\n");
    sb.append("    reasonResearchConsentMissing: ")
        .append(toIndentedString(reasonResearchConsentMissing))
        .append("\n");
    sb.append("    transferTAN: ").append(toIndentedString(transferTAN)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private MvhMetadata instance;

    public Builder() {
      this(new MvhMetadata());
    }

    protected Builder(MvhMetadata instance) {
      this.instance = instance;
    }

    public MvhMetadata.Builder modelProjectConsent(
        MvhMetadataModelProjectConsent modelProjectConsent) {
      this.instance.modelProjectConsent = modelProjectConsent;
      return this;
    }

    public MvhMetadata.Builder researchConsents(List<Map<String, Object>> researchConsents) {
      this.instance.researchConsents = researchConsents;
      return this;
    }

    public MvhMetadata.Builder episodeOfCare(Reference episodeOfCare) {
      this.instance.episodeOfCare = episodeOfCare;
      return this;
    }

    public MvhMetadata.Builder reasonResearchConsentMissing(
        BroadConsentReasonMissing reasonResearchConsentMissing) {
      this.instance.reasonResearchConsentMissing = reasonResearchConsentMissing;
      return this;
    }

    public MvhMetadata.Builder transferTAN(String transferTAN) {
      this.instance.transferTAN = transferTAN;
      return this;
    }

    public MvhMetadata.Builder type(MvhSubmissionType type) {
      this.instance.type = type;
      return this;
    }

    public MvhMetadata build() {
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

  public static MvhMetadata.Builder builder() {
    return new MvhMetadata.Builder();
  }

  public MvhMetadata.Builder toBuilder() {
    return new MvhMetadata.Builder()
        .modelProjectConsent(getModelProjectConsent())
        .researchConsents(getResearchConsents())
        .episodeOfCare(getEpisodeOfCare())
        .reasonResearchConsentMissing(getReasonResearchConsentMissing())
        .transferTAN(getTransferTAN())
        .type(getType());
  }
}
