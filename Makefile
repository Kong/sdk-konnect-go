.PHONY: generate.deepcopy
generate.deepcopy: deepcopy
# Generate deepcopy/runtime.Object implementations for a particular file
	controller-gen object paths=./models/components/

# TODO: add versioning
.PHONY: deepcopy
deepcopy:
	go install -v sigs.k8s.io/controller-tools/cmd/controller-gen@latest

OPENAPI_FILE = openapi.yaml
SPEAKEASY_DIR = .speakeasy
KUBEBUILDER_GENERATE_CODE_MARKER = \n+kubebuilder:object:generate=true\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	yq --inplace '.components.schemas.CreateControlPlaneRequest.description += "$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	yq --inplace '.components.schemas.CreateServiceWithoutParents.description += "$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
	git checkout -- $(SPEAKEASY_DIR)/gen.lock $(SPEAKEASY_DIR)/gen.yaml sdk.go
	$(MAKE) generate.deepcopy
	git checkout -- $(OPENAPI_FILE) \
		models/components/create*.go \
		docs/models/components/create*.md
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
