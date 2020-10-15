API_VERSION=v1.3.3

OPENAPI_GENERATOR_VERSION=4.3.1
JAVA=$(shell which java)
BINDIR=hack/bin

fmt:
	go fmt ./...

generate: update-version openapi-generator
	GO_POST_PROCESS_FILE="go fmt" \
	$(JAVA) -jar $(OPENAPI_GENERATOR_JAR) generate \
	--git-user-id TwoStone \
	--git-repo-id bitbucket-go-client \
  	--generator-name go-experimental \
  	--input-spec ../api/bitbucket-server-api.yaml \
  	--config hack/generator-config.yaml \
  	--api-package bitbucket \
  	--enable-post-process-file

validate: openapi-generator
	$(JAVA) -jar $(OPENAPI_GENERATOR_JAR) validate \
	--input-spec ../api/bitbucket-server-api.yaml \
	--recommend

build: generate fmt
	go mod download
	go build -v ./

version_regex := ^v([0-9]{1,}\.){2}[0-9]{1,}$

update-version:
ifeq ($(shell echo ${API_VERSION} | egrep "${version_regex}"),)
	@echo "No release version"
else
	@echo "Setting release version ${API_VERSION}"
	@sed -i '' 's/packageVersion:.*/packageVersion: $(API_VERSION:v%=%)/g' ./hack/generator-config.yaml
endif

openapi-generator:
ifeq (,$(wildcard $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar))
	mkdir -p $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)
	curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar \
		-o $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar
endif
OPENAPI_GENERATOR_JAR=$(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar

