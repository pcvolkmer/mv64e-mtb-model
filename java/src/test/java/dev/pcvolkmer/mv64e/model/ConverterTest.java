package dev.pcvolkmer.mv64e.model;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import java.io.IOException;
import java.util.List;
import java.util.regex.Pattern;
import org.junit.jupiter.api.Test;

class ConverterTest {

  @Test
  void shouldDeserializeJsonString() throws IOException {
    var resource = getClass().getClassLoader().getResource("mv64e-mtb-fake-patient.json");
    assertNotNull(Converter.fromJsonString(new String(resource.openStream().readAllBytes())));
  }

  @Test
  void shouldKeepPatientBirthdateFormatInYearMonth() throws IOException {
    var resource = getClass().getClassLoader().getResource("mv64e-mtb-fake-patient.json");
    var json = new String(resource.openStream().readAllBytes());
    // Fake birthdate to old format
    json =
        Pattern.compile("\"birthDate\"\\s*:\\s*\"\\d{4}-\\d{2}(-\\d{2})?\"")
            .matcher(json)
            .replaceAll("\"birthDate\":\"2025-03\"");

    var mtb = Converter.fromJsonString(json);

    var actual = Converter.toJsonString(mtb);
    // Expect new format
    assertThat(actual).contains("\"birthDate\":\"2025-03\"");
  }

  @Test
  void shouldConvertPatientBirthdateFormatToYearMonth() throws IOException {
    var resource = getClass().getClassLoader().getResource("mv64e-mtb-fake-patient.json");
    var json = new String(resource.openStream().readAllBytes());
    // Fake birthdate to old format
    json =
        Pattern.compile("\"birthDate\"\\s*:\\s*\"\\d{4}-\\d{2}(-\\d{2})?\"")
            .matcher(json)
            .replaceAll("\"birthDate\":\"2025-03-19\"");

    var mtb = Converter.fromJsonString(json);

    var actual = Converter.toJsonString(mtb);
    // Expect new format
    assertThat(actual).contains("\"birthDate\":\"2025-03\"");
  }

  @Test
  void shouldSerializeWithJsonNodeResearchConsentInserted() throws IOException {
    var mtbResource = getClass().getClassLoader().getResource("mv64e-mtb-fake-patient.json");
    var mtb = Converter.fromJsonString(new String(mtbResource.openStream().readAllBytes()));

    // Load Research Consent as JsonNode
    ObjectMapper objectMapper = new ObjectMapper();
    var bcResource =
        getClass().getClassLoader().getResource("fake_broadConsent_mii_response_permit.json");
    var jsonObject =
        objectMapper.readValue(
            new String(bcResource.openStream().readAllBytes()), ObjectNode.class);

    // Add Research Consent into Metadata
    var researchConsent = ResearchConsent.from(jsonObject);
    var metadata = new MvhMetadata();
    metadata.setResearchConsents(List.of(researchConsent));
    mtb.setMetadata(metadata);

    // Check conversion to JSON String
    var actualMtb = Converter.toJsonString(mtb);
    assertThat(actualMtb)
        .contains(
            "\"lastUpdated\":\"2025-08-15T11:13:59.143+02:00\"",
            "http://localhost:8080/ttp-fhir/fhir/gics/Consent/7d3456c2-79b1-11f0-ab27-6ed0ed82d0fd",
            "https://ths-greifswald.de/fhir/CodeSystem/gics/Policy/MII");
  }

  @Test
  void shouldConvertResearchConsentToJsonNode() throws IOException {
    // Load Research Consent as JsonNode
    ObjectMapper objectMapper = new ObjectMapper();
    var bcResource =
        getClass().getClassLoader().getResource("fake_broadConsent_mii_response_permit.json");
    var jsonObject =
        objectMapper.readValue(
            new String(bcResource.openStream().readAllBytes()), ObjectNode.class);
    var researchConsent = ResearchConsent.from(jsonObject);

    // Check Conversion keeps all elements
    var actualRc = researchConsent.asObjectNode();
    assertThat(actualRc).hasSize(4);
  }
}
