# MV §64e MTB data model for use with DNPM:DIP

[![Run Tests](https://github.com/pcvolkmer/mv64e-mtb-model/actions/workflows/ci.yml/badge.svg)](https://github.com/pcvolkmer/mv64e-mtb-model/actions/workflows/ci.yml)
[![Check Compatibility](https://github.com/pcvolkmer/mv64e-mtb-model/actions/workflows/compatibility.yml/badge.svg)](https://github.com/pcvolkmer/mv64e-mtb-model/actions/workflows/compatibility.yml)

Serialization and deserialization of DNPM:DIP MTB DTOs for some programming languages.

This library contains auto-generated code for the DNPM:DIP MTB DTOs as used by the DNPM:DIP backends Scala code. It is a
replacement for mostly manual adapted code
in [https://github.com/dnpm-dip/mv64e-mtb-dto-*](https://github.com/dnpm-dip).

## Supported DNPM-Datenmodell 2.1 versions

This library is intended to be used with DNPM-Datenmodell 2.1 as used in DNPM:DIP 1.3.3. For older DNPM:DIP versions,
please use the corresponding library versions:

| DNPM Datenmodell     | DNPM:DIP      | Language(s) | Library                                                                        | Latest Version       |
|----------------------|---------------|-------------|--------------------------------------------------------------------------------|----------------------|
| DNPM Datenmodell 2.1 | <= 1.3.1      | Java        | [mv64e-mtb-dto-java](https://github.com/dnpm-dip/mv64e-mtb-dto-java)           | 0.3.0                |
| DNPM Datenmodell 2.1 | <= 1.3.1      | C#          | [mv64e-mtb-dto-dotnet](https://github.com/dnpm-dip/mv64e-mtb-dto-dotnet)       | 0.3.0                |
| DNPM Datenmodell 2.1 | <= 1.3.1      | Rust        | [mv64e-mtb-dto-rs](https://github.com/dnpm-dip/mv64e-mtb-dto-rs)               | 0.3.0                |
| DNPM Datenmodell 2.1 | <= 1.3.1      | Go          | [mv64e-mtb-dto-go](https://github.com/dnpm-dip/mv64e-mtb-dto-go)               | 0.3.0                |
| DNPM Datenmodell 2.1 | 1.3.2 - 1.3.4 | Java, Rust  | [mv64e-mtb-model](https://github.com/pcvolkmer/mv64e-mtb-model) (this library) | ~~0.5.0~~ (outdated) |
| DNPM Datenmodell 2.1 | 1.3.5         | Java, Rust  | [mv64e-mtb-model](https://github.com/pcvolkmer/mv64e-mtb-model) (this library) | 0.6.0                |

## Usage notices

JSON (de)serialization includes format conversion for patients birthdate and date of death as introduced in
https://github.com/dnpm-dip/backend-core/commit/97c44aa8bbd6ba4ac81824c5178db23fd08f9068

### Java

Since Java code is also intended to be used with Onkostar, the default Java version is limited to Java 11 and all
date-time objects do not use JSR 310 types but `java.util.Date`.
This may change in the future (Onkostar 3).

