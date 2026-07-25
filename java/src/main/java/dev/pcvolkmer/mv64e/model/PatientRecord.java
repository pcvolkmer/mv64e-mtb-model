package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class PatientRecord {
  public static final String JSON_PROPERTY_DIAGNOSES = "diagnoses";

  private @Nullable List<MtbDiagnosis> diagnoses;

  public static final String JSON_PROPERTY_NGS_REPORTS = "ngsReports";

  private @Nullable List<SomaticNgsReport> ngsReports;

  public static final String JSON_PROPERTY_PERFORMANCE_STATUS = "performanceStatus";

  private @Nullable List<PerformanceStatus> performanceStatus;

  public static final String JSON_PROPERTY_GUIDELINE_PROCEDURES = "guidelineProcedures";

  private @Nullable List<OncoProcedure> guidelineProcedures;

  public static final String JSON_PROPERTY_FAMILY_MEMBER_HISTORIES = "familyMemberHistories";

  private @Nullable List<FamilyMemberHistory> familyMemberHistories;

  public static final String JSON_PROPERTY_CLAIM_RESPONSES = "claimResponses";

  private @Nullable List<ClaimResponse> claimResponses;

  public static final String JSON_PROPERTY_MSI_FINDINGS = "msiFindings";

  private @Nullable List<Msi> msiFindings;

  public static final String JSON_PROPERTY_CARE_PLANS = "carePlans";

  private @Nullable List<MtbCarePlan> carePlans;

  public static final String JSON_PROPERTY_SPECIMENS = "specimens";

  private @Nullable List<TumorSpecimen> specimens;

  public static final String JSON_PROPERTY_IHC_REPORTS = "ihcReports";

  private @Nullable List<IhcReport> ihcReports;

  public static final String JSON_PROPERTY_HISTOLOGY_REPORTS = "histologyReports";

  private @Nullable List<HistologyReport> histologyReports;

  public static final String JSON_PROPERTY_GUIDELINE_THERAPIES = "guidelineTherapies";

  private @Nullable List<MtbSystemicTherapy> guidelineTherapies;

  public static final String JSON_PROPERTY_EPISODES_OF_CARE = "episodesOfCare";

  private @Nullable List<MtbEpisodeOfCare> episodesOfCare;

  public static final String JSON_PROPERTY_PRIOR_DIAGNOSTIC_REPORTS = "priorDiagnosticReports";

  private @Nullable List<MolecularDiagnosticReport> priorDiagnosticReports;

  public static final String JSON_PROPERTY_SYSTEMIC_THERAPIES = "systemicTherapies";

  private @Nullable List<PatientRecordSystemicTherapiesInner> systemicTherapies;

  public static final String JSON_PROPERTY_FOLLOW_UPS = "followUps";

  private @Nullable List<FollowUp> followUps;

  public static final String JSON_PROPERTY_CLAIMS = "claims";

  private @Nullable List<Claim> claims;

  public static final String JSON_PROPERTY_METADATA = "metadata";

  private @Nullable MvhMetadata metadata;

  public static final String JSON_PROPERTY_RESPONSES = "responses";

  private @Nullable List<Response> responses;

  public static final String JSON_PROPERTY_PATIENT = "patient";

  private @Nullable Patient patient;

  public PatientRecord() {}

  public PatientRecord diagnoses(@Nullable List<MtbDiagnosis> diagnoses) {
    this.diagnoses = diagnoses;
    return this;
  }

  public PatientRecord addDiagnosesItem(MtbDiagnosis diagnosesItem) {
    if (this.diagnoses == null) {
      this.diagnoses = new ArrayList<>();
    }
    this.diagnoses.add(diagnosesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSES, required = false)
  public @Nullable List<MtbDiagnosis> getDiagnoses() {
    return diagnoses;
  }

  @JsonProperty(value = JSON_PROPERTY_DIAGNOSES, required = false)
  public void setDiagnoses(@Nullable List<MtbDiagnosis> diagnoses) {
    this.diagnoses = diagnoses;
  }

  public PatientRecord ngsReports(@Nullable List<SomaticNgsReport> ngsReports) {
    this.ngsReports = ngsReports;
    return this;
  }

  public PatientRecord addNgsReportsItem(SomaticNgsReport ngsReportsItem) {
    if (this.ngsReports == null) {
      this.ngsReports = new ArrayList<>();
    }
    this.ngsReports.add(ngsReportsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_NGS_REPORTS, required = false)
  public @Nullable List<SomaticNgsReport> getNgsReports() {
    return ngsReports;
  }

  @JsonProperty(value = JSON_PROPERTY_NGS_REPORTS, required = false)
  public void setNgsReports(@Nullable List<SomaticNgsReport> ngsReports) {
    this.ngsReports = ngsReports;
  }

  public PatientRecord performanceStatus(@Nullable List<PerformanceStatus> performanceStatus) {
    this.performanceStatus = performanceStatus;
    return this;
  }

  public PatientRecord addPerformanceStatusItem(PerformanceStatus performanceStatusItem) {
    if (this.performanceStatus == null) {
      this.performanceStatus = new ArrayList<>();
    }
    this.performanceStatus.add(performanceStatusItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PERFORMANCE_STATUS, required = false)
  public @Nullable List<PerformanceStatus> getPerformanceStatus() {
    return performanceStatus;
  }

  @JsonProperty(value = JSON_PROPERTY_PERFORMANCE_STATUS, required = false)
  public void setPerformanceStatus(@Nullable List<PerformanceStatus> performanceStatus) {
    this.performanceStatus = performanceStatus;
  }

  public PatientRecord guidelineProcedures(@Nullable List<OncoProcedure> guidelineProcedures) {
    this.guidelineProcedures = guidelineProcedures;
    return this;
  }

  public PatientRecord addGuidelineProceduresItem(OncoProcedure guidelineProceduresItem) {
    if (this.guidelineProcedures == null) {
      this.guidelineProcedures = new ArrayList<>();
    }
    this.guidelineProcedures.add(guidelineProceduresItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_PROCEDURES, required = false)
  public @Nullable List<OncoProcedure> getGuidelineProcedures() {
    return guidelineProcedures;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_PROCEDURES, required = false)
  public void setGuidelineProcedures(@Nullable List<OncoProcedure> guidelineProcedures) {
    this.guidelineProcedures = guidelineProcedures;
  }

  public PatientRecord familyMemberHistories(
      @Nullable List<FamilyMemberHistory> familyMemberHistories) {
    this.familyMemberHistories = familyMemberHistories;
    return this;
  }

  public PatientRecord addFamilyMemberHistoriesItem(FamilyMemberHistory familyMemberHistoriesItem) {
    if (this.familyMemberHistories == null) {
      this.familyMemberHistories = new ArrayList<>();
    }
    this.familyMemberHistories.add(familyMemberHistoriesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_FAMILY_MEMBER_HISTORIES, required = false)
  public @Nullable List<FamilyMemberHistory> getFamilyMemberHistories() {
    return familyMemberHistories;
  }

  @JsonProperty(value = JSON_PROPERTY_FAMILY_MEMBER_HISTORIES, required = false)
  public void setFamilyMemberHistories(@Nullable List<FamilyMemberHistory> familyMemberHistories) {
    this.familyMemberHistories = familyMemberHistories;
  }

  public PatientRecord claimResponses(@Nullable List<ClaimResponse> claimResponses) {
    this.claimResponses = claimResponses;
    return this;
  }

  public PatientRecord addClaimResponsesItem(ClaimResponse claimResponsesItem) {
    if (this.claimResponses == null) {
      this.claimResponses = new ArrayList<>();
    }
    this.claimResponses.add(claimResponsesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIM_RESPONSES, required = false)
  public @Nullable List<ClaimResponse> getClaimResponses() {
    return claimResponses;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIM_RESPONSES, required = false)
  public void setClaimResponses(@Nullable List<ClaimResponse> claimResponses) {
    this.claimResponses = claimResponses;
  }

  public PatientRecord msiFindings(@Nullable List<Msi> msiFindings) {
    this.msiFindings = msiFindings;
    return this;
  }

  public PatientRecord addMsiFindingsItem(Msi msiFindingsItem) {
    if (this.msiFindings == null) {
      this.msiFindings = new ArrayList<>();
    }
    this.msiFindings.add(msiFindingsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MSI_FINDINGS, required = false)
  public @Nullable List<Msi> getMsiFindings() {
    return msiFindings;
  }

  @JsonProperty(value = JSON_PROPERTY_MSI_FINDINGS, required = false)
  public void setMsiFindings(@Nullable List<Msi> msiFindings) {
    this.msiFindings = msiFindings;
  }

  public PatientRecord carePlans(@Nullable List<MtbCarePlan> carePlans) {
    this.carePlans = carePlans;
    return this;
  }

  public PatientRecord addCarePlansItem(MtbCarePlan carePlansItem) {
    if (this.carePlans == null) {
      this.carePlans = new ArrayList<>();
    }
    this.carePlans.add(carePlansItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CARE_PLANS, required = false)
  public @Nullable List<MtbCarePlan> getCarePlans() {
    return carePlans;
  }

  @JsonProperty(value = JSON_PROPERTY_CARE_PLANS, required = false)
  public void setCarePlans(@Nullable List<MtbCarePlan> carePlans) {
    this.carePlans = carePlans;
  }

  public PatientRecord specimens(@Nullable List<TumorSpecimen> specimens) {
    this.specimens = specimens;
    return this;
  }

  public PatientRecord addSpecimensItem(TumorSpecimen specimensItem) {
    if (this.specimens == null) {
      this.specimens = new ArrayList<>();
    }
    this.specimens.add(specimensItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMENS, required = false)
  public @Nullable List<TumorSpecimen> getSpecimens() {
    return specimens;
  }

  @JsonProperty(value = JSON_PROPERTY_SPECIMENS, required = false)
  public void setSpecimens(@Nullable List<TumorSpecimen> specimens) {
    this.specimens = specimens;
  }

  public PatientRecord ihcReports(@Nullable List<IhcReport> ihcReports) {
    this.ihcReports = ihcReports;
    return this;
  }

  public PatientRecord addIhcReportsItem(IhcReport ihcReportsItem) {
    if (this.ihcReports == null) {
      this.ihcReports = new ArrayList<>();
    }
    this.ihcReports.add(ihcReportsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_IHC_REPORTS, required = false)
  public @Nullable List<IhcReport> getIhcReports() {
    return ihcReports;
  }

  @JsonProperty(value = JSON_PROPERTY_IHC_REPORTS, required = false)
  public void setIhcReports(@Nullable List<IhcReport> ihcReports) {
    this.ihcReports = ihcReports;
  }

  public PatientRecord histologyReports(@Nullable List<HistologyReport> histologyReports) {
    this.histologyReports = histologyReports;
    return this;
  }

  public PatientRecord addHistologyReportsItem(HistologyReport histologyReportsItem) {
    if (this.histologyReports == null) {
      this.histologyReports = new ArrayList<>();
    }
    this.histologyReports.add(histologyReportsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY_REPORTS, required = false)
  public @Nullable List<HistologyReport> getHistologyReports() {
    return histologyReports;
  }

  @JsonProperty(value = JSON_PROPERTY_HISTOLOGY_REPORTS, required = false)
  public void setHistologyReports(@Nullable List<HistologyReport> histologyReports) {
    this.histologyReports = histologyReports;
  }

  public PatientRecord guidelineTherapies(@Nullable List<MtbSystemicTherapy> guidelineTherapies) {
    this.guidelineTherapies = guidelineTherapies;
    return this;
  }

  public PatientRecord addGuidelineTherapiesItem(MtbSystemicTherapy guidelineTherapiesItem) {
    if (this.guidelineTherapies == null) {
      this.guidelineTherapies = new ArrayList<>();
    }
    this.guidelineTherapies.add(guidelineTherapiesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_THERAPIES, required = false)
  public @Nullable List<MtbSystemicTherapy> getGuidelineTherapies() {
    return guidelineTherapies;
  }

  @JsonProperty(value = JSON_PROPERTY_GUIDELINE_THERAPIES, required = false)
  public void setGuidelineTherapies(@Nullable List<MtbSystemicTherapy> guidelineTherapies) {
    this.guidelineTherapies = guidelineTherapies;
  }

  public PatientRecord episodesOfCare(@Nullable List<MtbEpisodeOfCare> episodesOfCare) {
    this.episodesOfCare = episodesOfCare;
    return this;
  }

  public PatientRecord addEpisodesOfCareItem(MtbEpisodeOfCare episodesOfCareItem) {
    if (this.episodesOfCare == null) {
      this.episodesOfCare = new ArrayList<>();
    }
    this.episodesOfCare.add(episodesOfCareItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_EPISODES_OF_CARE, required = false)
  public @Nullable List<MtbEpisodeOfCare> getEpisodesOfCare() {
    return episodesOfCare;
  }

  @JsonProperty(value = JSON_PROPERTY_EPISODES_OF_CARE, required = false)
  public void setEpisodesOfCare(@Nullable List<MtbEpisodeOfCare> episodesOfCare) {
    this.episodesOfCare = episodesOfCare;
  }

  public PatientRecord priorDiagnosticReports(
      @Nullable List<MolecularDiagnosticReport> priorDiagnosticReports) {
    this.priorDiagnosticReports = priorDiagnosticReports;
    return this;
  }

  public PatientRecord addPriorDiagnosticReportsItem(
      MolecularDiagnosticReport priorDiagnosticReportsItem) {
    if (this.priorDiagnosticReports == null) {
      this.priorDiagnosticReports = new ArrayList<>();
    }
    this.priorDiagnosticReports.add(priorDiagnosticReportsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PRIOR_DIAGNOSTIC_REPORTS, required = false)
  public @Nullable List<MolecularDiagnosticReport> getPriorDiagnosticReports() {
    return priorDiagnosticReports;
  }

  @JsonProperty(value = JSON_PROPERTY_PRIOR_DIAGNOSTIC_REPORTS, required = false)
  public void setPriorDiagnosticReports(
      @Nullable List<MolecularDiagnosticReport> priorDiagnosticReports) {
    this.priorDiagnosticReports = priorDiagnosticReports;
  }

  public PatientRecord systemicTherapies(
      @Nullable List<PatientRecordSystemicTherapiesInner> systemicTherapies) {
    this.systemicTherapies = systemicTherapies;
    return this;
  }

  public PatientRecord addSystemicTherapiesItem(
      PatientRecordSystemicTherapiesInner systemicTherapiesItem) {
    if (this.systemicTherapies == null) {
      this.systemicTherapies = new ArrayList<>();
    }
    this.systemicTherapies.add(systemicTherapiesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEMIC_THERAPIES, required = false)
  public @Nullable List<PatientRecordSystemicTherapiesInner> getSystemicTherapies() {
    return systemicTherapies;
  }

  @JsonProperty(value = JSON_PROPERTY_SYSTEMIC_THERAPIES, required = false)
  public void setSystemicTherapies(
      @Nullable List<PatientRecordSystemicTherapiesInner> systemicTherapies) {
    this.systemicTherapies = systemicTherapies;
  }

  public PatientRecord followUps(@Nullable List<FollowUp> followUps) {
    this.followUps = followUps;
    return this;
  }

  public PatientRecord addFollowUpsItem(FollowUp followUpsItem) {
    if (this.followUps == null) {
      this.followUps = new ArrayList<>();
    }
    this.followUps.add(followUpsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_FOLLOW_UPS, required = false)
  public @Nullable List<FollowUp> getFollowUps() {
    return followUps;
  }

  @JsonProperty(value = JSON_PROPERTY_FOLLOW_UPS, required = false)
  public void setFollowUps(@Nullable List<FollowUp> followUps) {
    this.followUps = followUps;
  }

  public PatientRecord claims(@Nullable List<Claim> claims) {
    this.claims = claims;
    return this;
  }

  public PatientRecord addClaimsItem(Claim claimsItem) {
    if (this.claims == null) {
      this.claims = new ArrayList<>();
    }
    this.claims.add(claimsItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIMS, required = false)
  public @Nullable List<Claim> getClaims() {
    return claims;
  }

  @JsonProperty(value = JSON_PROPERTY_CLAIMS, required = false)
  public void setClaims(@Nullable List<Claim> claims) {
    this.claims = claims;
  }

  public PatientRecord metadata(@Nullable MvhMetadata metadata) {
    this.metadata = metadata;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_METADATA, required = false)
  public @Nullable MvhMetadata getMetadata() {
    return metadata;
  }

  @JsonProperty(value = JSON_PROPERTY_METADATA, required = false)
  public void setMetadata(@Nullable MvhMetadata metadata) {
    this.metadata = metadata;
  }

  public PatientRecord responses(@Nullable List<Response> responses) {
    this.responses = responses;
    return this;
  }

  public PatientRecord addResponsesItem(Response responsesItem) {
    if (this.responses == null) {
      this.responses = new ArrayList<>();
    }
    this.responses.add(responsesItem);
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_RESPONSES, required = false)
  public @Nullable List<Response> getResponses() {
    return responses;
  }

  @JsonProperty(value = JSON_PROPERTY_RESPONSES, required = false)
  public void setResponses(@Nullable List<Response> responses) {
    this.responses = responses;
  }

  public PatientRecord patient(@Nullable Patient patient) {
    this.patient = patient;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = false)
  public @Nullable Patient getPatient() {
    return patient;
  }

  @JsonProperty(value = JSON_PROPERTY_PATIENT, required = false)
  public void setPatient(@Nullable Patient patient) {
    this.patient = patient;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PatientRecord patientRecord = (PatientRecord) o;
    return Objects.equals(this.diagnoses, patientRecord.diagnoses)
        && Objects.equals(this.ngsReports, patientRecord.ngsReports)
        && Objects.equals(this.performanceStatus, patientRecord.performanceStatus)
        && Objects.equals(this.guidelineProcedures, patientRecord.guidelineProcedures)
        && Objects.equals(this.familyMemberHistories, patientRecord.familyMemberHistories)
        && Objects.equals(this.claimResponses, patientRecord.claimResponses)
        && Objects.equals(this.msiFindings, patientRecord.msiFindings)
        && Objects.equals(this.carePlans, patientRecord.carePlans)
        && Objects.equals(this.specimens, patientRecord.specimens)
        && Objects.equals(this.ihcReports, patientRecord.ihcReports)
        && Objects.equals(this.histologyReports, patientRecord.histologyReports)
        && Objects.equals(this.guidelineTherapies, patientRecord.guidelineTherapies)
        && Objects.equals(this.episodesOfCare, patientRecord.episodesOfCare)
        && Objects.equals(this.priorDiagnosticReports, patientRecord.priorDiagnosticReports)
        && Objects.equals(this.systemicTherapies, patientRecord.systemicTherapies)
        && Objects.equals(this.followUps, patientRecord.followUps)
        && Objects.equals(this.claims, patientRecord.claims)
        && Objects.equals(this.metadata, patientRecord.metadata)
        && Objects.equals(this.responses, patientRecord.responses)
        && Objects.equals(this.patient, patientRecord.patient);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        diagnoses,
        ngsReports,
        performanceStatus,
        guidelineProcedures,
        familyMemberHistories,
        claimResponses,
        msiFindings,
        carePlans,
        specimens,
        ihcReports,
        histologyReports,
        guidelineTherapies,
        episodesOfCare,
        priorDiagnosticReports,
        systemicTherapies,
        followUps,
        claims,
        metadata,
        responses,
        patient);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PatientRecord {\n");
    sb.append("    diagnoses: ").append(toIndentedString(diagnoses)).append("\n");
    sb.append("    ngsReports: ").append(toIndentedString(ngsReports)).append("\n");
    sb.append("    performanceStatus: ").append(toIndentedString(performanceStatus)).append("\n");
    sb.append("    guidelineProcedures: ")
        .append(toIndentedString(guidelineProcedures))
        .append("\n");
    sb.append("    familyMemberHistories: ")
        .append(toIndentedString(familyMemberHistories))
        .append("\n");
    sb.append("    claimResponses: ").append(toIndentedString(claimResponses)).append("\n");
    sb.append("    msiFindings: ").append(toIndentedString(msiFindings)).append("\n");
    sb.append("    carePlans: ").append(toIndentedString(carePlans)).append("\n");
    sb.append("    specimens: ").append(toIndentedString(specimens)).append("\n");
    sb.append("    ihcReports: ").append(toIndentedString(ihcReports)).append("\n");
    sb.append("    histologyReports: ").append(toIndentedString(histologyReports)).append("\n");
    sb.append("    guidelineTherapies: ").append(toIndentedString(guidelineTherapies)).append("\n");
    sb.append("    episodesOfCare: ").append(toIndentedString(episodesOfCare)).append("\n");
    sb.append("    priorDiagnosticReports: ")
        .append(toIndentedString(priorDiagnosticReports))
        .append("\n");
    sb.append("    systemicTherapies: ").append(toIndentedString(systemicTherapies)).append("\n");
    sb.append("    followUps: ").append(toIndentedString(followUps)).append("\n");
    sb.append("    claims: ").append(toIndentedString(claims)).append("\n");
    sb.append("    metadata: ").append(toIndentedString(metadata)).append("\n");
    sb.append("    responses: ").append(toIndentedString(responses)).append("\n");
    sb.append("    patient: ").append(toIndentedString(patient)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private PatientRecord instance;

    public Builder() {
      this(new PatientRecord());
    }

    protected Builder(PatientRecord instance) {
      this.instance = instance;
    }

    public PatientRecord.Builder diagnoses(List<MtbDiagnosis> diagnoses) {
      this.instance.diagnoses = diagnoses;
      return this;
    }

    public PatientRecord.Builder ngsReports(List<SomaticNgsReport> ngsReports) {
      this.instance.ngsReports = ngsReports;
      return this;
    }

    public PatientRecord.Builder performanceStatus(List<PerformanceStatus> performanceStatus) {
      this.instance.performanceStatus = performanceStatus;
      return this;
    }

    public PatientRecord.Builder guidelineProcedures(List<OncoProcedure> guidelineProcedures) {
      this.instance.guidelineProcedures = guidelineProcedures;
      return this;
    }

    public PatientRecord.Builder familyMemberHistories(
        List<FamilyMemberHistory> familyMemberHistories) {
      this.instance.familyMemberHistories = familyMemberHistories;
      return this;
    }

    public PatientRecord.Builder claimResponses(List<ClaimResponse> claimResponses) {
      this.instance.claimResponses = claimResponses;
      return this;
    }

    public PatientRecord.Builder msiFindings(List<Msi> msiFindings) {
      this.instance.msiFindings = msiFindings;
      return this;
    }

    public PatientRecord.Builder carePlans(List<MtbCarePlan> carePlans) {
      this.instance.carePlans = carePlans;
      return this;
    }

    public PatientRecord.Builder specimens(List<TumorSpecimen> specimens) {
      this.instance.specimens = specimens;
      return this;
    }

    public PatientRecord.Builder ihcReports(List<IhcReport> ihcReports) {
      this.instance.ihcReports = ihcReports;
      return this;
    }

    public PatientRecord.Builder histologyReports(List<HistologyReport> histologyReports) {
      this.instance.histologyReports = histologyReports;
      return this;
    }

    public PatientRecord.Builder guidelineTherapies(List<MtbSystemicTherapy> guidelineTherapies) {
      this.instance.guidelineTherapies = guidelineTherapies;
      return this;
    }

    public PatientRecord.Builder episodesOfCare(List<MtbEpisodeOfCare> episodesOfCare) {
      this.instance.episodesOfCare = episodesOfCare;
      return this;
    }

    public PatientRecord.Builder priorDiagnosticReports(
        List<MolecularDiagnosticReport> priorDiagnosticReports) {
      this.instance.priorDiagnosticReports = priorDiagnosticReports;
      return this;
    }

    public PatientRecord.Builder systemicTherapies(
        List<PatientRecordSystemicTherapiesInner> systemicTherapies) {
      this.instance.systemicTherapies = systemicTherapies;
      return this;
    }

    public PatientRecord.Builder followUps(List<FollowUp> followUps) {
      this.instance.followUps = followUps;
      return this;
    }

    public PatientRecord.Builder claims(List<Claim> claims) {
      this.instance.claims = claims;
      return this;
    }

    public PatientRecord.Builder metadata(MvhMetadata metadata) {
      this.instance.metadata = metadata;
      return this;
    }

    public PatientRecord.Builder responses(List<Response> responses) {
      this.instance.responses = responses;
      return this;
    }

    public PatientRecord.Builder patient(Patient patient) {
      this.instance.patient = patient;
      return this;
    }

    public PatientRecord build() {
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

  public static PatientRecord.Builder builder() {
    return new PatientRecord.Builder();
  }

  public PatientRecord.Builder toBuilder() {
    return new PatientRecord.Builder()
        .diagnoses(getDiagnoses())
        .ngsReports(getNgsReports())
        .performanceStatus(getPerformanceStatus())
        .guidelineProcedures(getGuidelineProcedures())
        .familyMemberHistories(getFamilyMemberHistories())
        .claimResponses(getClaimResponses())
        .msiFindings(getMsiFindings())
        .carePlans(getCarePlans())
        .specimens(getSpecimens())
        .ihcReports(getIhcReports())
        .histologyReports(getHistologyReports())
        .guidelineTherapies(getGuidelineTherapies())
        .episodesOfCare(getEpisodesOfCare())
        .priorDiagnosticReports(getPriorDiagnosticReports())
        .systemicTherapies(getSystemicTherapies())
        .followUps(getFollowUps())
        .claims(getClaims())
        .metadata(getMetadata())
        .responses(getResponses())
        .patient(getPatient());
  }
}
