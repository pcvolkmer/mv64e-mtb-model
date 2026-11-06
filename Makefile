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

.PHONY: generate-go
generate-go:
	./scripts/generate.sh go

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

.PHONY: update-go
update-go: generate-go
	echo "Copying generated source code ..."
	cp ./generated/go/go.* ./go/
	cp ./generated/go/utils.go ./generated/go/model_*.go ./go/
	$(MAKE) rename-go
	echo "Formatting ..."
	cd go && go fmt &>/dev/null
	echo
	echo "WARNING: Keep track of manual changes!"

.PHONY: rename-go
rename-go:
	find ./go -type f -name 'model_*.go' ! -name 'model_model_*' | while read -r file; do \
		dir=$$(dirname "$$file"); \
		name=$$(basename "$$file"); \
		mv "$$file" "$$dir/$${name#model_}"; \
	done
	find ./go -type f -name 'model_model_*.go' | while read -r file; do \
		dir=$$(dirname "$$file"); \
		name=$$(basename "$$file"); \
		mv "$$file" "$$dir/$${name#model_}"; \
	done

