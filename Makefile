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
KUBEBUILDER_GENERATE_CODE_MARKER = \n+kubebuilder:object:generate=true
VERSION_TMP_FILE = $(SPEAKEASY_DIR)/.version_temp

# NOTE: add more types that need to have DeepCopy() generated.
.PHONY: generate.sdk
generate.sdk:
	yq --inplace '.components.schemas.CreateControlPlaneRequest.description += "$(KUBEBUILDER_GENERATE_CODE_MARKER)"' $(OPENAPI_FILE)
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)
	git checkout -- $(SPEAKEASY_DIR)/gen.lock $(SPEAKEASY_DIR)/gen.yaml sdk.go
	$(MAKE) generate.deepcopy
	git checkout -- $(OPENAPI_FILE)
	speakeasy generate sdk --lang go --out . --schema ./$(OPENAPI_FILE)

# yq '.management.releaseVersion' $(SPEAKEASY_DIR)/gen.lock > $(VERSION_TMP_FILE)
# yq --inplace '.management.releaseVersion = load("${VERSION_TMP_FILE}")' $(speakeasy_dir)/gen.lock
# yq --inplace '.go.version = load("${VERSION_TMP_FILE}")' $(SPEAKEASY_DIR)/gen.yaml
# rm $(VERSION_TMP_FILE)