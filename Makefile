SHELL = bash
.SHELLFLAGS = -ec -o pipefail

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
TOOLS_VERSIONS_FILE = $(PROJECT_DIR)/.tools_versions.yaml

MISE := $(shell which mise)
MISE_FILE := .mise/config.toml

.PHONY: mise
mise:
	@mise -V >/dev/null || (echo "mise - https://github.com/jdx/mise - not found. Please install it." && exit 1)

.PHONY: mise-install
mise-install: mise
	@$(MISE) install -q $(DEP_VER)

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')

# Do not store yq's version in .tools_versions.yaml as it is used to get tool versions.
# renovate: datasource=github-releases depName=mikefarah/yq
YQ_VERSION = 4.49.2
YQ = $(PROJECT_DIR)/bin/installs/github-mikefarah-yq/$(YQ_VERSION)/yq_$(OS)_$(ARCH)
.PHONY: yq
yq: mise
	MISE_DATA_DIR=$(PROJECT_DIR)/bin $(MAKE) mise-install DEP_VER=github:mikefarah/yq@$(YQ_VERSION)

CONTROLLER_GEN_VERSION = $(shell $(YQ) -r '.controller-tools' < $(TOOLS_VERSIONS_FILE))
CONTROLLER_GEN = $(PROJECT_DIR)/bin/installs/github-kubernetes-sigs-controller-tools/$(CONTROLLER_GEN_VERSION)/controller-gen
.PHONY: controller-gen
controller-gen: mise yq
	MISE_DATA_DIR=$(PROJECT_DIR)/bin $(MAKE) mise-install DEP_VER=github:kubernetes-sigs/controller-tools@$(CONTROLLER_GEN_VERSION)

MOCKERY_VERSION = $(shell $(YQ) -r '.mockery' < $(TOOLS_VERSIONS_FILE))
MOCKERY = $(PROJECT_DIR)/bin/installs/github-vektra-mockery/$(MOCKERY_VERSION)/mockery
.PHONY: mockery
mockery: mise yq
	MISE_DATA_DIR=$(PROJECT_DIR)/bin $(MAKE) mise-install DEP_VER=github:vektra/mockery@$(MOCKERY_VERSION)

IFACEMAKER_VERSION = $(shell $(YQ) -r '.ifacemaker' < $(TOOLS_VERSIONS_FILE))
IFACEMAKER = $(PROJECT_DIR)/bin/installs/go-github-com-vburenin-ifacemaker/$(IFACEMAKER_VERSION)/bin/ifacemaker
.PHONY: ifacemaker
ifacemaker: mise yq
	MISE_DATA_DIR=$(PROJECT_DIR)/bin $(MAKE) mise-install DEP_VER=go:github.com/vburenin/ifacemaker@$(IFACEMAKER_VERSION)

# NOTE: speakeasy is not placed in $(PROJECT_DIR)/bin as it is being used in
# GitHub workflows outside of this Makefile.
SPEAKEASY_VERSION = $(shell $(YQ) -p toml -o yaml '.tools["github:speakeasy-api/speakeasy"].version' < $(MISE_FILE))
.PHONY: speakeasy
speakeasy: mise yq
	$(MAKE) mise-install DEP_VER=github:speakeasy-api/speakeasy@$(SPEAKEASY_VERSION)

# ------------------------------------------------------------------------------
# Code generation
# ------------------------------------------------------------------------------

SED=sed
ifeq (Darwin,$(shell uname -s))
	SED=gsed
endif

OPENAPI_FILE = openapi.yaml
SPEAKEASY_DIR = .speakeasy
UPDATE_PORTAL_OVERLAY = $(SPEAKEASY_DIR)/overlays/update-portal-patch-defaults.yaml
UPDATE_PORTAL_AUDIT_LOG_WEBHOOK_OVERLAY = \
	$(SPEAKEASY_DIR)/overlays/update-portal-audit-log-webhook-patch-defaults.yaml
PATCH_CUSTOM_PORTAL_EMAIL_TEMPLATE_OVERLAY = \
	$(SPEAKEASY_DIR)/overlays/patch-custom-portal-email-template-defaults.yaml
UPDATE_PORTAL_IDENTITY_PROVIDER_OVERLAY = \
	$(SPEAKEASY_DIR)/overlays/update-portal-identity-provider-defaults.yaml
UPDATE_DCR_CONFIG_HTTP_OVERLAY = \
	$(SPEAKEASY_DIR)/overlays/update-dcr-config-http-defaults.yaml
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

.PHONY: validate.update-portal-overlay
validate.update-portal-overlay: speakeasy
	speakeasy overlay validate --overlay $(UPDATE_PORTAL_OVERLAY)
	speakeasy overlay apply --strict --schema $(OPENAPI_FILE) --overlay $(UPDATE_PORTAL_OVERLAY) --out /dev/null

.PHONY: validate.update-portal-audit-log-webhook-overlay
validate.update-portal-audit-log-webhook-overlay: speakeasy
	speakeasy overlay validate --overlay $(UPDATE_PORTAL_AUDIT_LOG_WEBHOOK_OVERLAY)
	speakeasy overlay apply --strict --schema $(OPENAPI_FILE) \
		--overlay $(UPDATE_PORTAL_AUDIT_LOG_WEBHOOK_OVERLAY) --out /dev/null

.PHONY: validate.patch-custom-portal-email-template-overlay
validate.patch-custom-portal-email-template-overlay: speakeasy
	speakeasy overlay validate --overlay $(PATCH_CUSTOM_PORTAL_EMAIL_TEMPLATE_OVERLAY)
	speakeasy overlay apply --strict --schema $(OPENAPI_FILE) \
		--overlay $(PATCH_CUSTOM_PORTAL_EMAIL_TEMPLATE_OVERLAY) --out /dev/null

.PHONY: validate.update-portal-identity-provider-overlay
validate.update-portal-identity-provider-overlay: speakeasy
	speakeasy overlay validate --overlay $(UPDATE_PORTAL_IDENTITY_PROVIDER_OVERLAY)
	speakeasy overlay apply --strict --schema $(OPENAPI_FILE) \
		--overlay $(UPDATE_PORTAL_IDENTITY_PROVIDER_OVERLAY) --out /dev/null

.PHONY: validate.update-dcr-config-http-overlay
validate.update-dcr-config-http-overlay: speakeasy
	speakeasy overlay validate --overlay $(UPDATE_DCR_CONFIG_HTTP_OVERLAY)
	speakeasy overlay apply --strict --schema $(OPENAPI_FILE) \
		--overlay $(UPDATE_DCR_CONFIG_HTTP_OVERLAY) --out /dev/null

.PHONY: generate.sdk.speakeasy
generate.sdk.speakeasy: validate.update-portal-overlay validate.update-portal-audit-log-webhook-overlay \
	validate.patch-custom-portal-email-template-overlay validate.update-portal-identity-provider-overlay \
	validate.update-dcr-config-http-overlay
	speakeasy run --skip-versioning --skip-testing --minimal --skip-upload-spec

# NOTE: SDK generation consists of adding the kubebuilder code marker and generating
#       DeepCopy() for the types that need it and then using speakeasy to generate
#       the final sdk code.
# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	$(MAKE) generate.sdk.speakeasy
	git add --update .
	$(MAKE) generate.deepcopy
	$(MAKE) _generate.omitempty
	go mod tidy

# NOTE: Use this when there are breaking changes in the generated SDK code
#       that require the interfaces and mocks to be regenerated.
.PHONY: generate.sdk.from-scratch
generate.sdk.from-scratch:
	$(MAKE) remove.field_tests
	$(MAKE) remove.interfaces
	$(MAKE) remove.mocks
	$(MAKE) generate.sdk
	$(MAKE) generate.interfaces
	$(MAKE) generate.mocks
	$(MAKE) generate.field_tests

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

TYPES_TO_MOCK := $(shell grep -B 20 'rootSDK.*\*SDK' *.go | grep 'type.*struct' | awk '{print $$2}' | sort -u)

.PHONY: remove.interfaces
remove.interfaces:
	@echo "Removing existing interfaces (to prevent breakage on breaking changes)..."
	@$(foreach s, $(TYPES_TO_MOCK), \
		rm -f $(shell echo $(s) | tr 'A-Z' 'a-z')_i.go; )

.PHONY: generate.interfaces
generate.interfaces: ifacemaker remove.interfaces
	@$(foreach s, $(TYPES_TO_MOCK), \
		$(MAKE) _generate.ifacemaker STRUCT=$(s) || exit 1;)

.PHONY: remove.mocks
remove.mocks:
	@echo "Removing existing mocks (to prevent breakage on breaking changes) and generating new ones..."
	rm -f test/mocks/zz_generated*.go

# https://github.com/vektra/mockery/issues/803#issuecomment-2287198024
.PHONY: generate.mocks
generate.mocks: mockery remove.mocks
	GODEBUG=gotypesalias=0 $(MOCKERY)

# TYPES_TO_TEST_FIELDS is a list of types in models/components/
# which have field tests generated for them.
# This prevents breaking changes to fields of these types from being merged.
# Note that this is a different list than TYPES_TO_MOCK as that one
# is a list of SDK types which have CRUD operations attached to them.
TYPES_TO_TEST_FIELDS := \
	ACL \
	BasicAuth \
	CACertificate \
	Certificate \
	Consumer \
	ConsumerGroup \
	CreateControlPlaneRequest \
	CreateNetworkRequest \
	HMACAuth \
	Jwt \
	Key \
	KeySet \
	Plugin \
	CreatePluginWithConsumerGroupRequest \
	UpsertPluginWithConsumerGroupRequest \
	CreatePluginWithConsumerRequest \
	UpsertPluginWithConsumerRequest \
	Route \
	Service \
	Target \
	Upstream \
	Vault

comma := ,
empty :=
space := $(empty) $(empty)
TYPES_TO_TEST_FIELDS_COMMA := $(subst $(space),$(comma),$(strip $(TYPES_TO_TEST_FIELDS)))

.PHONY: remove.field_tests
remove.field_tests:
	rm -rf test/fields/*

.PHONY: generate.field_tests
generate.field_tests: remove.field_tests
	go run ./scripts/gentests/ \
		-types=$(TYPES_TO_TEST_FIELDS_COMMA)

.PHONY: test.fields
test.fields:
	go test -v -race $(GOTESTFLAGS) \
		./test/fields/...

.PHONY: verify.diff
verify.diff:
	@$(PROJECT_DIR)/scripts/verify-diff.sh $(PROJECT_DIR)
