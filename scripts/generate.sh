#!/bin/sh

echo "Fetching schema"
curl 'https://dnpm-dip.net/api/mtb/etl/patient-record/schema' > ./schema.json

# Draft 2020-12
SCHEMA_MOD=$(cat schema.json | \
  sed 's/\"#/\"#\/$defs\//' | \
  sed 's/\$defs/components\/schemas/' | \
  sed 's/\$anchor/\$id/' | \

  sed -r 's/Coding([A-Za-z_]*)/\1Coding/' | \
  sed 's/_//g' | \
  sed 's/ATC/Atc/g' | \
  sed 's/BRCA/Brca/g' | \
  sed 's/SNV/Snv/g' | \
  sed 's/CNV/Cnv/g' | \
  sed 's/DNA/Dna/g' | \
  sed 's/ECOG/Ecog/g' | \
  sed 's/HRD/Hrd/g' | \
  sed 's/ICScore/IcScore/g' | \
  sed 's/IHCReport/IhcReport/g' | \
  sed 's/MSI/Msi/g' | \
  sed 's/MTB/Mtb/g' | \
  sed 's/MVH/Mvh/g' | \
  sed 's/NGS/Ngs/g' | \
  sed 's/RECISTCoding/RecistCoding/g' | \
  sed 's/RNA/Rna/g' | \
  sed 's/TCScore/TcScore/g' | \
  sed 's/TMB/Tmb/g')

SCHEMAS=$(echo "$SCHEMA_MOD" | jq -r '."components/schemas"')

PATIENT_RECORD=$(echo "$SCHEMA_MOD" | jq '."properties"' )

RESULT=$(cat <<EOF
{
  "openapi": "3.1.0",
  "info": {
    "title": "DNPM MTB Patient Record API",
    "version": "1.3.0",
    "description": "OpenAPI representation of the DNPM DIP MTB Patient Record schema."
  },
  "paths": {
    "/api/mtb/etl/patient-record": {
      "post": {
        "summary": "Submit MTB patient record",
        "operationId": "submitPatientRecord",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "\$ref": "#/components/schemas/PatientRecord"
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas":
      $SCHEMAS YYY
      "PatientRecord": {
        "properties": $PATIENT_RECORD
      }
    }
  }
}
EOF
)

echo $RESULT | sed 's/} YYY/,/' | jq > openapi.json

rm -rf ./generated

case $1 in
  ### Rust
  rust)
    echo "Generate Rust code"
    mkdir -p ./generated/rust

    docker run --rm --user 1000 \
      -v ./openapi.json:/local/openapi.json \
      -v ./generated/rust/:/local/out/ \
      openapitools/openapi-generator-cli generate \
        --skip-validate-spec \
        -i /local/openapi.json \
        -g rust \
        -o /local/out \
        --additional-properties=packageName=mv64e-model,avoidBoxedModels=true,useChrono=false

    # cleanup rust code
    find ./generated/rust -name "*.rs" \
      -exec sed -i '/\/\*/,/\*\//d' {} \;

    # deny unknown fields
    find ./generated/rust -name "*.rs" \
      -exec sed -i 's/Deserialize)]/Deserialize)]\n#[serde(deny_unknown_fields)]/' {} \;
    ;;

  ### Java
  java)
    echo "Generate Java code"
    mkdir -p ./generated/java

    docker run --rm --user 1000 \
      -v ./openapi.json:/local/openapi.json \
      -v ./generated/java/:/local/out/ \
      openapitools/openapi-generator-cli generate \
        --skip-validate-spec \
        -i /local/openapi.json \
        -g java \
        -o /local/out \
        --library resttemplate \
        --global-property models \
        --additional-properties=modelPackage=dev.pcvolkmer.mv64e.model,useJakartaEe=true,uniqueItems=true,useJspecify=true,serializationLibrary=jackson,dateLibrary=legacy,supportUrlQuery=false,generateBuilders=true,containerDefaultToNull=true,sortModelPropertiesByRequiredFlag=false,sortParamsByRequiredFlag=false,openApiNullable=false


    # cleanup java code
    find ./generated/java -name "*.java" \
      -exec sed -i '/\/\*/,/\*\//d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/@JsonPropertyOrder/,/})/d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/@JsonTypeName/,/)/d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/import com.fasterxml.jackson.annotation.JsonPropertyOrder;/d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/import com.fasterxml.jackson.annotation.JsonTypeName;/d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/@jakarta.annotation.Generated/d' {} \;

    # No JsonInclude-Include for serialized null values
    find ./generated/java -name "*.java" \
      -exec sed -i '/import com.fasterxml.jackson.annotation.JsonInclude;/d' {} \;

    find ./generated/java -name "*.java" \
      -exec sed -i '/@JsonInclude/d' {} \;
    ;;

  *)
    echo "No code generation for $1 available"
esac