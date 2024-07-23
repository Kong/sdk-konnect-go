// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertACLWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the ACL to lookup
	ACLID string `pathParam:"style=simple,explode=false,name=ACLId"`
	// Description of the ACL
	ACLWithoutParents components.ACLWithoutParents `request:"mediaType=application/json"`
}

func (o *UpsertACLWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertACLWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *UpsertACLWithConsumerRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

func (o *UpsertACLWithConsumerRequest) GetACLWithoutParents() components.ACLWithoutParents {
	if o == nil {
		return components.ACLWithoutParents{}
	}
	return o.ACLWithoutParents
}

type UpsertACLWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted ACL
	ACL *components.ACL
}

func (o *UpsertACLWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertACLWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertACLWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertACLWithConsumerResponse) GetACL() *components.ACL {
	if o == nil {
		return nil
	}
	return o.ACL
}
