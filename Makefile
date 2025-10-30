SHELL = bash
.SHELLFLAGS = -ec -o pipefail

MODULE_NAME := github.com/Kong/sdk-konnect-go
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

MOCKERY_VERSION = $(shell $(YQ) -r '.mockery' < $(TOOLS_VERSIONS_FILE))
MOCKERY = $(PROJECT_DIR)/bin/installs/mockery/$(MOCKERY_VERSION)/bin/mockery
.PHONY: mockery
mockery: mise yq ## Download mockery locally if necessary.
	@$(MISE) plugin install --yes -q mockery https://github.com/cabify/asdf-mockery.git
	@$(MISE) install -q mockery@$(MOCKERY_VERSION)

IFACEMAKER_VERSION = $(shell $(YQ) -r '.ifacemaker' < $(TOOLS_VERSIONS_FILE))
IFACEMAKER = $(PROJECT_DIR)/bin/ifacemaker
.PHONY: ifacemaker
ifacemaker: yq
	GOBIN=$(PROJECT_DIR)/bin go install -v github.com/vburenin/ifacemaker@v$(IFACEMAKER_VERSION)

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


# TODO: this works around the fact that speakeasy does not support fields that
# do not have "omitempty" set and are not required.
# Related slack thread: https://kongstrong.slack.com/archives/C05MLAA216D/p1730369303065339
.PHONY: _generate.omitempty
_generate.omitempty:
	$(SED) -i 's#Members \[\]Members `json:"members,omitempty"`#Members \[\]Members `json:"members"`#g' \
		models/components/groupmembership.go

.PHONY: generate.deepcopy
generate.deepcopy: controller-gen
	$(SED) -i 's#\(type CreateControlPlaneRequest struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/createcontrolplanerequest.go
	$(SED) -i 's#\(type CreateNetworkRequest struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/createnetworkrequest.go
	$(SED) -i 's#\(type RouteWithoutParents struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type RouteWithoutParentsDestinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type Destinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type Sources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type Route struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type RouteService struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type RouteWithoutParentsSources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type RouteWithoutParentsService struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routewithoutparents.go
	$(SED) -i 's#\(type RouteInput struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/route.go
	$(SED) -i 's#\(type Destinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routejson.go
	$(SED) -i 's#\(type RouteJSON struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routejson.go
	$(SED) -i 's#\(type RouteJSONService struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routejson.go
	$(SED) -i 's#\(type Sources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routejson.go
	$(SED) -i 's#\(type RouteExpression struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routeexpression.go
	$(SED) -i 's#\(type RouteExpressionService struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/routeexpression.go
	$(SED) -i 's#\(type UpstreamClientCertificate struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type Healthchecks struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type Active struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type Passive struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type UpstreamHealthy struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type Healthy struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type UpstreamUnhealthy struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go
	$(SED) -i 's#\(type Unhealthy struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' \
		models/components/upstream.go

	$(CONTROLLER_GEN) object paths=./models/components/
	go mod tidy

# NOTE: This removes the +kubebuilder:object:generate=true markers after DeepCopy() generation.
# to prevent it from being committed to the codebase.
	git checkout -- $(OPENAPI_FILE) \
		$(shell git ls-files models/components/route*.go) \
		$(shell git ls-files docs/models/components/route*.md) \
		$(shell git ls-files models/components/create*.go) \
		$(shell git ls-files docs/models/components/create*.md) \
		$(shell git ls-files models/components/upstream*.go) \
		$(shell git ls-files docs/models/components/upstream*.md) \
		$(shell git ls-files docs/models/components/healthchecks*.md)

# NOTE: SDK generation consists of adding the kubebuilder code marker and generating
#       DeepCopy() for the types that need it and then using speakeasy to generate
#       the final sdk code.
# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	speakeasy run --skip-versioning --skip-testing --minimal --skip-upload-spec
	git add --update .
	$(MAKE) generate.deepcopy
	$(MAKE) _generate.omitempty
	go mod tidy

.PHONY: test
test: test.unit test.integration

.PHONY: test.unit
test.unit:
	go test -v -race $(GOTESTFLAGS) \
		./internal/...

.PHONY: test.integration
test.integration:
	KONNECT_TEST_RUN_ID=$(shell openssl rand -hex 8) \
		go test -v -race $(GOTESTFLAGS) \
		./test/integration/...

.PHONY: cleanup.konnect
cleanup.konnect:
	go run ./ci/konnect-cleanup

.PHONY: _generate.ifacemaker
_generate.ifacemaker:
	@$(eval LOWERCASE_STRUCT := $(shell echo $(STRUCT) | tr 'A-Z' 'a-z'))
	$(IFACEMAKER) \
		--file $(LOWERCASE_STRUCT).go \
		--struct $(STRUCT) \
		--iface $(STRUCT)SDK \
		--iface-comment "$(STRUCT)SDK is a generated interface." \
		--output $(LOWERCASE_STRUCT)_i.go \
		-p sdkkonnectgo

.PHONY: generate.interfaces
generate.interfaces: ifacemaker
# TODO: make this iterate over all structs that need mocks if necessary.
	$(MAKE) _generate.ifacemaker STRUCT=ControlPlanes
	$(MAKE) _generate.ifacemaker STRUCT=APIGatewayDataPlaneCertificates
	$(MAKE) _generate.ifacemaker STRUCT=APIGateways

# https://github.com/vektra/mockery/issues/803#issuecomment-2287198024
.PHONY: generate.mocks
generate.mocks: mockery
	GODEBUG=gotypesalias=0 $(MOCKERY)

.PHONY: verify.diff
verify.diff:
	@$(PROJECT_DIR)/scripts/verify-diff.sh $(PROJECT_DIR)
