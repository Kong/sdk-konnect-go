// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type GetControlPlaneRequest struct {
	// The control plane ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
