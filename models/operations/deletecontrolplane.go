// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type DeleteControlPlaneRequest struct {
	// The control plane ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteControlPlaneResponse struct {
}
