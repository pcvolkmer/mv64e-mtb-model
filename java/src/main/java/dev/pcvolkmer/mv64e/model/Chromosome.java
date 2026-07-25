package dev.pcvolkmer.mv64e.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

public enum Chromosome {
  CHR5("chr5"),

  CHR6("chr6"),

  CHR8("chr8"),

  CHR1("chr1"),

  CHR_X("chrX"),

  CHR12("chr12"),

  CHR11("chr11"),

  CHR17("chr17"),

  CHR21("chr21"),

  CHR19("chr19"),

  CHR14("chr14"),

  CHR9("chr9"),

  CHR4("chr4"),

  CHR_MT("chrMT"),

  CHR2("chr2"),

  CHR_Y("chrY"),

  CHR3("chr3"),

  CHR7("chr7"),

  CHR13("chr13"),

  CHR10("chr10"),

  CHR15("chr15"),

  CHR22("chr22"),

  CHR18("chr18"),

  CHR16("chr16"),

  CHR20("chr20");

  private String value;

  Chromosome(String value) {
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
  public static Chromosome fromValue(String value) {
    for (Chromosome b : Chromosome.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
