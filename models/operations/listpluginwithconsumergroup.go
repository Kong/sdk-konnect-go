// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListPluginWithConsumerGroupRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Consumer Group to lookup
	ConsumerGroupID string `pathParam:"style=simple,explode=false,name=ConsumerGroupId"`
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (l ListPluginWithConsumerGroupRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPluginWithConsumerGroupRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPluginWithConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *ListPluginWithConsumerGroupRequest) GetConsumerGroupID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerGroupID
}

func (o *ListPluginWithConsumerGroupRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListPluginWithConsumerGroupRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListPluginWithConsumerGroupRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// ListPluginWithConsumerGroupResponseBody - A successful response listing Plugins
type ListPluginWithConsumerGroupResponseBody struct {
	Data []components.Plugin `json:"data,omitempty"`
	// URI to the next page (may be null)
	Next *string `json:"next,omitempty"`
	// Offset is used to paginate through the API. Provide this value to the next list operation to fetch the next page
	Offset *string `json:"offset,omitempty"`
}

func (o *ListPluginWithConsumerGroupResponseBody) GetData() []components.Plugin {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ListPluginWithConsumerGroupResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListPluginWithConsumerGroupResponseBody) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

type ListPluginWithConsumerGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A successful response listing Plugins
	Object *ListPluginWithConsumerGroupResponseBody
}

func (o *ListPluginWithConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPluginWithConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPluginWithConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPluginWithConsumerGroupResponse) GetObject() *ListPluginWithConsumerGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
