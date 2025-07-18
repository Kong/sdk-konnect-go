// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type MoveAPIDocumentRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The document identifier related to the API
	DocumentID string `pathParam:"style=simple,explode=false,name=documentId"`
	// move document
	MoveDocumentRequestPayload components.MoveDocumentRequestPayload `request:"mediaType=application/json"`
}

func (o *MoveAPIDocumentRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *MoveAPIDocumentRequest) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}

func (o *MoveAPIDocumentRequest) GetMoveDocumentRequestPayload() components.MoveDocumentRequestPayload {
	if o == nil {
		return components.MoveDocumentRequestPayload{}
	}
	return o.MoveDocumentRequestPayload
}

type MoveAPIDocumentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *MoveAPIDocumentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
