SHELL = bash
.SHELLFLAGS = -ec -o pipefail

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
TOOLS_VERSIONS_FILE = $(PROJECT_DIR)/.tools_versions.yaml
export MISE_DATA_DIR = $(PROJECT_DIR)/bin/

MISE := $(shell which mise)
.PHONY: mise
mise:
	@mise -V >/dev/null || (echo "mise - https://github.com/jdx/mise - not found. Please install it." && exit 1)

# Do not store yq's version in .tools_versions.yaml as it is used to get tool versions.
# renovate: datasource=github-releases depName=mikefarah/yq
YQ_VERSION = 4.43.1
YQ = $(PROJECT_DIR)/bin/installs/yq/$(YQ_VERSION)/bin/yq
.PHONY: yq
yq: mise # Download yq locally if necessary.
	@$(MISE) plugin install --yes -q yq
	@$(MISE) install -q yq@$(YQ_VERSION)

CONTROLLER_GEN_VERSION = $(shell $(YQ) -r '.controller-tools' < $(TOOLS_VERSIONS_FILE))
CONTROLLER_GEN = $(PROJECT_DIR)/bin/installs/kube-controller-tools/$(CONTROLLER_GEN_VERSION)/bin/controller-gen
.PHONY: controller-gen
controller-gen: mise yq ## Download controller-gen locally if necessary.
	@$(MISE) plugin install --yes -q kube-controller-tools
	@$(MISE) install -q kube-controller-tools@$(CONTROLLER_GEN_VERSION)

# ------------------------------------------------------------------------------
# Code generation
# ------------------------------------------------------------------------------

SED=sed
ifeq (Darwin,$(shell uname -s))
	SED=gsed
endif

OPENAPI_FILE = openapi.yaml
SPEAKEASY_DIR = .speakeasy
KUBEBUILDER_GENERATE_CODE_MARKER = +kubebuilder:object:generate=true

.PHONY: generate.deepcopy
generate.deepcopy: controller-gen
# Generate deepcopy/runtime.Object implementations for a particular file
	$(CONTROLLER_GEN) object paths=./models/components/

# NOTE: We call speakeasy twice to generate DeepCopy() for the types that need it.
#       The generation is called twice because we want to generate the zz_generated.deepcopy.go
#       file for the types that need DeepCopy() to be generated but to not include
#       the related code markers in the generated sdk source code or docs.
# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	yq --inplace '.components.schemas.CreateControlPlaneRequest.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	yq --inplace '.components.schemas.CreateServiceWithoutParents.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	yq --inplace '.components.schemas.RouteWithoutParents.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)

	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
	git checkout -- $(SPEAKEASY_DIR)/gen.lock $(SPEAKEASY_DIR)/gen.yaml sdk.go

	$(SED) -i 's#\(type RouteWithoutParentsDestinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type Destinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type Sources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type RouteService struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type RouteWithoutParentsSources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type RouteInput struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go

	$(MAKE) generate.deepcopy
	git checkout -- $(OPENAPI_FILE) \
		$(shell git ls-files models/components/route*.go) \
		$(shell git ls-files docs/models/components/route*.md) \
		$(shell git ls-files models/components/create*.go) \
		$(shell git ls-files docs/models/components/create*.md)
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
