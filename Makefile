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

# ------------------------------------------------------------------------------
# Code generation
# ------------------------------------------------------------------------------

SED=sed
ifeq (Darwin,$(shell uname -s))
	SED=gsed
endif

OPENAPI_FILE = openapi.yaml
SPEAKEASY_DIR = .speakeasy

# TODO: this works around the fact that speakeasy does not support fields that
# do not have "omitempty" set and are not required.
# Related slack thread: https://kongstrong.slack.com/archives/C05MLAA216D/p1730369303065339
.PHONY: _generate.omitempty
_generate.omitempty:
	$(SED) -i 's#Members \[\]Members `json:"members,omitempty"`#Members \[\]Members `json:"members"`#g' \
		models/components/groupmembership.go

.PHONY: generate.sdk
generate.sdk:
	speakeasy run --skip-versioning --output console
	$(MAKE) _generate.omitempty
	go mod tidy