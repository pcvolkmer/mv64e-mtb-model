package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.module.SimpleModule;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.TimeZone;

public class Converter {
  // Date-time helpers

  private static final String[] DATE_TIME_FORMATS = {
    "yyyy-MM-dd'T'HH:mm:ss.SX",
    "yyyy-MM-dd'T'HH:mm:ss.S",
    "yyyy-MM-dd'T'HH:mm:ssX",
    "yyyy-MM-dd'T'HH:mm:ss",
    "yyyy-MM-dd HH:mm:ss.SX",
    "yyyy-MM-dd HH:mm:ss.S",
    "yyyy-MM-dd HH:mm:ssX",
    "yyyy-MM-dd HH:mm:ss",
    "HH:mm:ss.SZ",
    "HH:mm:ss.S",
    "HH:mm:ssZ",
    "HH:mm:ss",
    "yyyy-MM-dd",
  };

  public static Date parseAllDateTimeString(String str) {
    for (String format : DATE_TIME_FORMATS) {
      try {
        return new SimpleDateFormat(format).parse(str);
      } catch (Exception ex) {
        // Ignored
      }
    }
    return null;
  }

  public static String serializeDateTime(Date datetime) {
    return new SimpleDateFormat("yyyy-MM-dd'T'hh:mm:ssZ").format(datetime);
  }

  public static String serializeDate(Date datetime) {
    return new SimpleDateFormat("yyyy-MM-dd").format(datetime);
  }

  public static String serializeTime(Date datetime) {
    return new SimpleDateFormat("hh:mm:ssZ").format(datetime);
  }

  // Serialize/deserialize helpers

  public static PatientRecord fromJsonString(String json) throws IOException {
    return getObjectReader().readValue(json);
  }

  public static String toJsonString(PatientRecord obj) throws JsonProcessingException {
    return getObjectWriter().writeValueAsString(obj);
  }

  private static ObjectReader reader;
  private static ObjectWriter writer;

  private static void instantiateMapper() {
    ObjectMapper mapper = new ObjectMapper();
    mapper.findAndRegisterModules();
    mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, true);
    mapper.configure(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS, false);
    mapper.setSerializationInclusion(JsonInclude.Include.NON_NULL);
    SimpleModule module = new SimpleModule();
    module.addDeserializer(
        Date.class,
        new JsonDeserializer<>() {
          @Override
          public Date deserialize(
              JsonParser jsonParser, DeserializationContext deserializationContext)
              throws IOException {
            String value = jsonParser.getText();
            return Converter.parseAllDateTimeString(value);
          }
        });
    module.addDeserializer(
        Date.class,
        new JsonDeserializer<>() {
          @Override
          public Date deserialize(
              JsonParser jsonParser, DeserializationContext deserializationContext)
              throws IOException {
            String value = jsonParser.getText();
            return Converter.parseAllDateTimeString(value);
          }
        });
    module.addDeserializer(
        Date.class,
        new JsonDeserializer<>() {
          @Override
          public Date deserialize(
              JsonParser jsonParser, DeserializationContext deserializationContext)
              throws IOException {
            String value = jsonParser.getText();
            return Converter.parseAllDateTimeString(value);
          }
        });
    mapper.registerModule(module);
    mapper.setTimeZone(TimeZone.getDefault());
    reader = mapper.readerFor(PatientRecord.class);
    writer = mapper.writerFor(PatientRecord.class);
  }

  private static ObjectReader getObjectReader() {
    if (reader == null) instantiateMapper();
    return reader;
  }

  private static ObjectWriter getObjectWriter() {
    if (writer == null) instantiateMapper();
    return writer;
  }
}
