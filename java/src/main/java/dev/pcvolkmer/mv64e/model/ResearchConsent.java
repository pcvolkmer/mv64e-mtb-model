package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import java.util.TreeMap;

public class ResearchConsent extends TreeMap<String, Object> {

  /**
   * Creates a research consent from ObjectNode. As a developer, you are responsible to use a FHIR
   * R4 Consent JSON-Objekt.
   *
   * @param researchConsents The FHIR R4 Consent JSON-Objekt
   * @return the research consent object
   */
  public static ResearchConsent from(ObjectNode researchConsents) {
    if (null == researchConsents) {
      throw new IllegalArgumentException("researchConsent cannot be null");
    }

    ObjectMapper mapper = new ObjectMapper();
    return mapper.convertValue(researchConsents, new TypeReference<>() {});
  }

  public ObjectNode asObjectNode() {
    ObjectMapper mapper = new ObjectMapper();
    return mapper.convertValue(this, ObjectNode.class);
  }
}
