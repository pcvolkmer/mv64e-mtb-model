package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.util.Date;
import java.util.Objects;
import org.jspecify.annotations.Nullable;

public class Patient {
  public static final String JSON_PROPERTY_DATE_OF_DEATH = "dateOfDeath";

  private @Nullable Date dateOfDeath;

  public static final String JSON_PROPERTY_AGE = "age";

  private @Nullable Age age;

  public static final String JSON_PROPERTY_HEALTH_INSURANCE = "healthInsurance";

  private PatientHealthInsurance healthInsurance;

  public static final String JSON_PROPERTY_ADDRESS = "address";

  private @Nullable PatientAddress address;

  public static final String JSON_PROPERTY_VITAL_STATUS = "vitalStatus";

  private @Nullable VitalStatusCoding vitalStatus;

  public static final String JSON_PROPERTY_BIRTH_DATE = "birthDate";

  private Date birthDate;

  public static final String JSON_PROPERTY_ID = "id";

  private String id;

  public static final String JSON_PROPERTY_GENDER = "gender";

  private GenderCoding gender;

  public static final String JSON_PROPERTY_MANAGING_SITE = "managingSite";

  private @Nullable Coding managingSite;

  public Patient() {}

  public Patient dateOfDeath(@Nullable Date dateOfDeath) {
    this.dateOfDeath = dateOfDeath;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE_OF_DEATH, required = false)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM")
  public @Nullable Date getDateOfDeath() {
    return dateOfDeath;
  }

  @JsonProperty(value = JSON_PROPERTY_DATE_OF_DEATH, required = false)
  @JsonDeserialize(using = YearMonthDeserializerWithFallback.class)
  public void setDateOfDeath(@Nullable Date dateOfDeath) {
    this.dateOfDeath = dateOfDeath;
  }

  public Patient age(@Nullable Age age) {
    this.age = age;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_AGE, required = false)
  public @Nullable Age getAge() {
    return age;
  }

  @JsonProperty(value = JSON_PROPERTY_AGE, required = false)
  public void setAge(@Nullable Age age) {
    this.age = age;
  }

  public Patient healthInsurance(PatientHealthInsurance healthInsurance) {
    this.healthInsurance = healthInsurance;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_HEALTH_INSURANCE, required = true)
  public PatientHealthInsurance getHealthInsurance() {
    return healthInsurance;
  }

  @JsonProperty(value = JSON_PROPERTY_HEALTH_INSURANCE, required = true)
  public void setHealthInsurance(PatientHealthInsurance healthInsurance) {
    this.healthInsurance = healthInsurance;
  }

  public Patient address(@Nullable PatientAddress address) {
    this.address = address;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ADDRESS, required = false)
  public @Nullable PatientAddress getAddress() {
    return address;
  }

  @JsonProperty(value = JSON_PROPERTY_ADDRESS, required = false)
  public void setAddress(@Nullable PatientAddress address) {
    this.address = address;
  }

  public Patient vitalStatus(@Nullable VitalStatusCoding vitalStatus) {
    this.vitalStatus = vitalStatus;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_VITAL_STATUS, required = false)
  public @Nullable VitalStatusCoding getVitalStatus() {
    return vitalStatus;
  }

  @JsonProperty(value = JSON_PROPERTY_VITAL_STATUS, required = false)
  public void setVitalStatus(@Nullable VitalStatusCoding vitalStatus) {
    this.vitalStatus = vitalStatus;
  }

  public Patient birthDate(Date birthDate) {
    this.birthDate = birthDate;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_BIRTH_DATE, required = true)
  @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM")
  public Date getBirthDate() {
    return birthDate;
  }

  @JsonProperty(value = JSON_PROPERTY_BIRTH_DATE, required = true)
  @JsonDeserialize(using = YearMonthDeserializerWithFallback.class)
  public void setBirthDate(Date birthDate) {
    this.birthDate = birthDate;
  }

  public Patient id(String id) {
    this.id = id;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public String getId() {
    return id;
  }

  @JsonProperty(value = JSON_PROPERTY_ID, required = true)
  public void setId(String id) {
    this.id = id;
  }

  public Patient gender(GenderCoding gender) {
    this.gender = gender;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_GENDER, required = true)
  public GenderCoding getGender() {
    return gender;
  }

  @JsonProperty(value = JSON_PROPERTY_GENDER, required = true)
  public void setGender(GenderCoding gender) {
    this.gender = gender;
  }

  public Patient managingSite(@Nullable Coding managingSite) {
    this.managingSite = managingSite;
    return this;
  }

  @JsonProperty(value = JSON_PROPERTY_MANAGING_SITE, required = false)
  public @Nullable Coding getManagingSite() {
    return managingSite;
  }

  @JsonProperty(value = JSON_PROPERTY_MANAGING_SITE, required = false)
  public void setManagingSite(@Nullable Coding managingSite) {
    this.managingSite = managingSite;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Patient patient = (Patient) o;
    return Objects.equals(this.dateOfDeath, patient.dateOfDeath)
        && Objects.equals(this.age, patient.age)
        && Objects.equals(this.healthInsurance, patient.healthInsurance)
        && Objects.equals(this.address, patient.address)
        && Objects.equals(this.vitalStatus, patient.vitalStatus)
        && Objects.equals(this.birthDate, patient.birthDate)
        && Objects.equals(this.id, patient.id)
        && Objects.equals(this.gender, patient.gender)
        && Objects.equals(this.managingSite, patient.managingSite);
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        dateOfDeath,
        age,
        healthInsurance,
        address,
        vitalStatus,
        birthDate,
        id,
        gender,
        managingSite);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Patient {\n");
    sb.append("    dateOfDeath: ").append(toIndentedString(dateOfDeath)).append("\n");
    sb.append("    age: ").append(toIndentedString(age)).append("\n");
    sb.append("    healthInsurance: ").append(toIndentedString(healthInsurance)).append("\n");
    sb.append("    address: ").append(toIndentedString(address)).append("\n");
    sb.append("    vitalStatus: ").append(toIndentedString(vitalStatus)).append("\n");
    sb.append("    birthDate: ").append(toIndentedString(birthDate)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    gender: ").append(toIndentedString(gender)).append("\n");
    sb.append("    managingSite: ").append(toIndentedString(managingSite)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  private String toIndentedString(Object o) {
    return o == null ? "null" : o.toString().replace("\n", "\n    ");
  }

  public static class Builder {

    private Patient instance;

    public Builder() {
      this(new Patient());
    }

    protected Builder(Patient instance) {
      this.instance = instance;
    }

    public Patient.Builder dateOfDeath(Date dateOfDeath) {
      this.instance.dateOfDeath = dateOfDeath;
      return this;
    }

    public Patient.Builder age(Age age) {
      this.instance.age = age;
      return this;
    }

    public Patient.Builder healthInsurance(PatientHealthInsurance healthInsurance) {
      this.instance.healthInsurance = healthInsurance;
      return this;
    }

    public Patient.Builder address(PatientAddress address) {
      this.instance.address = address;
      return this;
    }

    public Patient.Builder vitalStatus(VitalStatusCoding vitalStatus) {
      this.instance.vitalStatus = vitalStatus;
      return this;
    }

    public Patient.Builder birthDate(Date birthDate) {
      this.instance.birthDate = birthDate;
      return this;
    }

    public Patient.Builder id(String id) {
      this.instance.id = id;
      return this;
    }

    public Patient.Builder gender(GenderCoding gender) {
      this.instance.gender = gender;
      return this;
    }

    public Patient.Builder managingSite(Coding managingSite) {
      this.instance.managingSite = managingSite;
      return this;
    }

    public Patient build() {
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

  public static Patient.Builder builder() {
    return new Patient.Builder();
  }

  public Patient.Builder toBuilder() {
    return new Patient.Builder()
        .dateOfDeath(getDateOfDeath())
        .age(getAge())
        .healthInsurance(getHealthInsurance())
        .address(getAddress())
        .vitalStatus(getVitalStatus())
        .birthDate(getBirthDate())
        .id(getId())
        .gender(getGender())
        .managingSite(getManagingSite());
  }
}
