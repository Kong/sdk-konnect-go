// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

type ListPluginWithRouteRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Route to lookup
	RouteID string `pathParam:"style=simple,explode=false,name=RouteId"`
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (l ListPluginWithRouteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPluginWithRouteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPluginWithRouteRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *ListPluginWithRouteRequest) GetRouteID() string {
	if o == nil {
		return ""
	}
	return o.RouteID
}

func (o *ListPluginWithRouteRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListPluginWithRouteRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListPluginWithRouteRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// ListPluginWithRouteResponseBody - A successful response listing Plugins
type ListPluginWithRouteResponseBody struct {
	Data []components.Plugin `json:"data,omitempty"`
	// URI to the next page (may be null)
	Next *string `json:"next,omitempty"`
	// Offset is used to paginate through the API. Provide this value to the next list operation to fetch the next page
	Offset *string `json:"offset,omitempty"`
}

func (o *ListPluginWithRouteResponseBody) GetData() []components.Plugin {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ListPluginWithRouteResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListPluginWithRouteResponseBody) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

type ListPluginWithRouteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A successful response listing Plugins
	Object *ListPluginWithRouteResponseBody
}

func (o *ListPluginWithRouteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListPluginWithRouteResponse) GetObject() *ListPluginWithRouteResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
