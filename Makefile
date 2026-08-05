ifndef VERBOSE
.SILENT:
endif

# Code generation

.PHONY: generate-java
generate-java:
	./scripts/generate.sh java

.PHONY: generate-rust
generate-rust:
	./scripts/generate.sh rust

# Source-Code update

.PHONY: update-java
update-java: generate-java
	echo "Copying generated source code ..."
	cp -R ./generated/java/src/main/java/dev/pcvolkmer/mv64e/model/* ./java/src/main/java/dev/pcvolkmer/mv64e/model
	echo "Formatting ..."
	cd java && ./gradlew spotlessApply
	echo
	echo "WARNING: Keep track of manual changes!"

.PHONY: update-rust
update-rust: generate-rust
	echo "Copying generated source code ..."
	cp -R ./generated/rust/src/models/* ./rust/src/models
	echo "Formatting ..."
	cargo fmt
	echo
	echo "WARNING: Keep track of manual changes!"

