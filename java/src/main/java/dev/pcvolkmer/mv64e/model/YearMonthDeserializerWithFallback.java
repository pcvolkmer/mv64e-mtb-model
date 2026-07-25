package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import java.io.IOException;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class YearMonthDeserializerWithFallback extends JsonDeserializer<Date> {

  @Override
  public Date deserialize(JsonParser jsonParser, DeserializationContext deserializationContext)
      throws IOException {
    final var dateFormat = new SimpleDateFormat("yyyy-MM-dd");
    final var yearMonthFormat = new SimpleDateFormat("yyyy-MM");

    var str = jsonParser.getText();

    try {
      return dateFormat.parse(str);
    } catch (ParseException e) {
      try {
        return yearMonthFormat.parse(str);
      } catch (ParseException e1) {
        // left empty
      }
    }

    throw new IOException("Invalid date format for Patient birtDate or dataOfDeath");
  }
}
