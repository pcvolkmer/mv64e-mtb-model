package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.ArrayList;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import org.jspecify.annotations.Nullable;

public class LevelOfEvidence {
  public static final String JSON_PROPERTY_GRADING = "grading";

  private LevelOfEvidenceGradingCoding grading;

  public static final String JSON_PROPERTY_ADDENDUMS = "addendums";

  private @Nullable Set<LevelOfEvidenceAddendumCoding> addendums;

  public static final String JSON_PROPERTY_PUBLICATIONS = "publications";

  private @Nullable List<PublicationReference> publications;

  public LevelOfEvidence() {}

  public LevelOfEvidence grading(LevelOfEvidenceGradingCoding grading) {
    this.grading = grading;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GRADING, required = true)
  public LevelOfEvidenceGradingCoding getGrading() {
    return grading;
  }

  @JsonProperty(value = JSON_PROPERTY_GRADING, required = true)
  public void setGrading(LevelOfEvidenceGradingCoding grading) {
    this.grading = grading;
  }

  public LevelOfEvidence addendums(@Nullable Set<LevelOfEvidenceAddendumCoding> addendums) {
    this.addendums = addendums;
    return this;
  }

  public LevelOfEvidence addAddendumsItem(LevelOfEvidenceAddendumCoding addendumsItem) {
    if (this.addendums == null) {
      this.addendums = new LinkedHashSet<>();
    }
    this.addendums.add(addendumsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ADDENDUMS, required = false)
  public @Nullable Set<LevelOfEvidenceAddendumCoding> getAddendums() {
    return addendums;
  }

  @JsonDeserialize(as = LinkedHashSet.class)
  @JsonProperty(value = JSON_PROPERTY_ADDENDUMS, required = false)
  public void setAddendums(@Nullable Set<LevelOfEvidenceAddendumCoding> addendums) {
    this.addendums = addendums;
  }

  public LevelOfEvidence publications(@Nullable List<PublicationReference> publications) {
    this.publications = publications;
    return this;
  }

  public LevelOfEvidence addPublicationsItem(PublicationReference publicationsItem) {
    if (this.publications == null) {
      this.publications = new ArrayList<>();
    }
    this.publications.add(publicationsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PUBLICATIONS, required = false)
  public @Nullable List<PublicationReference> getPublications() {
    return publications;
  }

  @JsonProperty(value = JSON_PROPERTY_PUBLICATIONS, required = false)
  public void setPublications(@Nullable List<PublicationReference> publications) {
    this.publications = publications;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LevelOfEvidence levelOfEvidence = (LevelOfEvidence) o;
    return Objects.equals(this.grading, levelOfEvidence.grading)
        && Objects.equals(this.addendums, levelOfEvidence.addendums)
        && Objects.equals(this.publications, levelOfEvidence.publications);
  }

  @Override
  public int hashCode() {
    return Objects.hash(grading, addendums, publications);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LevelOfEvidence {\n");
    sb.append("    grading: ").append(toIndentedString(grading)).append("\n");
    sb.append("    addendums: ").append(toIndentedString(addendums)).append("\n");
    sb.append("    publications: ").append(toIndentedString(publications)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private LevelOfEvidence instance;

    public Builder() {
      this(new LevelOfEvidence());
    }

    protected Builder(LevelOfEvidence instance) {
      this.instance = instance;
    }

    public LevelOfEvidence.Builder grading(LevelOfEvidenceGradingCoding grading) {
      this.instance.grading = grading;
      return this;
    }

    public LevelOfEvidence.Builder addendums(Set<LevelOfEvidenceAddendumCoding> addendums) {
      this.instance.addendums = addendums;
      return this;
    }

    public LevelOfEvidence.Builder publications(List<PublicationReference> publications) {
      this.instance.publications = publications;
      return this;
    }

    public LevelOfEvidence build() {
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

  public static LevelOfEvidence.Builder builder() {
    return new LevelOfEvidence.Builder();
  }

  public LevelOfEvidence.Builder toBuilder() {
    return new LevelOfEvidence.Builder()
        .grading(getGrading())
        .addendums(getAddendums())
        .publications(getPublications());
  }
}
