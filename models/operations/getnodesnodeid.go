// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetNodesNodeIDRequest struct {
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (g GetNodesNodeIDRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetNodesNodeIDRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetNodesNodeIDRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *GetNodesNodeIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetNodesNodeIDRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetNodesNodeIDRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type GetNodesNodeIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Example response
	GetNode *components.GetNode
}

func (o *GetNodesNodeIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNodesNodeIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNodesNodeIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetNodesNodeIDResponse) GetGetNode() *components.GetNode {
	if o == nil {
		return nil
	}
	return o.GetNode
}
