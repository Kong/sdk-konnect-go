// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateAPIVersionRequestSpec struct {
	// The raw content of your API spec, in json or yaml format (OpenAPI or AsyncAPI).
	//
	Content *string `json:"content,omitempty"`
}

func (o *CreateAPIVersionRequestSpec) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

type CreateAPIVersionRequest struct {
	// The version of the api.
	Version *string                      `json:"version,omitempty"`
	Spec    *CreateAPIVersionRequestSpec `json:"spec,omitempty"`
}

func (o *CreateAPIVersionRequest) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *CreateAPIVersionRequest) GetSpec() *CreateAPIVersionRequestSpec {
	if o == nil {
		return nil
	}
	return o.Spec
}
