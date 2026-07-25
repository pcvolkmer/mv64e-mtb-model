# MV §64e MTB datamodel for use with DNPM:DIP

Serialization and deserialization of DNPM:DIP MTB DTOs for some programming languages.

This library contains auto-generated code for the DNPM:DIP MTB DTOs as used by the DNPM:DIP backends Scala code. It is a
replacement for mostly manual adapted code
in [https://github.com/dnpm-dip/mv64e-mtb-dto-*](https://github.com/dnpm-dip).

## Usage notices

JSON (de)serialization includes format conversion for patients birthdate and date of death as introduced in
https://github.com/dnpm-dip/backend-core/commit/97c44aa8bbd6ba4ac81824c5178db23fd08f9068

### Java

Since java code is also intended to be used with Onkostar, the default Java version is limited to Java 11 and all
date-time objects do not use JSR 310 types but `java.util.Date`.

