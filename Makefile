SHELL = bash
.SHELLFLAGS = -ec -o pipefail

OPENAPI_FILE = openapi.yaml
SPEAKEASY_DIR = .speakeasy
KUBEBUILDER_GENERATE_CODE_MARKER = +kubebuilder:object:generate=true

SED=sed
ifeq (Darwin,$(shell uname -s))
	SED=gsed
endif

.PHONY: generate.deepcopy
generate.deepcopy: deepcopy
# Generate deepcopy/runtime.Object implementations for a particular file
	controller-gen object paths=./models/components/

# TODO: add versioning
.PHONY: deepcopy
deepcopy:
	go install -v sigs.k8s.io/controller-tools/cmd/controller-gen@latest

# NOTE: We call speakeasy twice to generate DeepCopy() for the types that need it.
#       The generation is called twice because we want to generate the zz_generated.deepcopy.go
#       file for the types that need DeepCopy() to be generated but to not include
#       the related code markers in the generated sdk source code or docs.
# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	yq --inplace '.components.schemas.CreateControlPlaneRequest.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	yq --inplace '.components.schemas.CreateServiceWithoutParents.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	yq --inplace '.components.schemas.CreateRouteWithoutParents.description += "\n$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)

	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
	git checkout -- $(SPEAKEASY_DIR)/gen.lock $(SPEAKEASY_DIR)/gen.yaml sdk.go

	$(SED) -i 's#\(type CreateRouteWithoutParentsDestinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' models/components/createroutewithoutparents.go
	$(SED) -i 's#\(type CreateRouteDestinations struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' models/components/createroute.go
	$(SED) -i 's#\(type CreateRouteWithoutParentsSources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' models/components/createroutewithoutparents.go
	$(SED) -i 's#\(type CreateRouteSources struct\)#// $(KUBEBUILDER_GENERATE_CODE_MARKER)\n\1#g' models/components/createroute.go

	$(MAKE) generate.deepcopy
	git checkout -- $(OPENAPI_FILE) \
		$(shell git ls-files models/components/create*.go) \
		$(shell git ls-files docs/models/components/create*.md)
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
