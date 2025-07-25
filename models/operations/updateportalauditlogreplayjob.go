// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpdatePortalAuditLogReplayJobRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// The request schema to replace a portal audit log replay job.
	ReplacePortalAuditLogReplayJob *components.ReplacePortalAuditLogReplayJob `request:"mediaType=application/json"`
}

func (o *UpdatePortalAuditLogReplayJobRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalAuditLogReplayJobRequest) GetReplacePortalAuditLogReplayJob() *components.ReplacePortalAuditLogReplayJob {
	if o == nil {
		return nil
	}
	return o.ReplacePortalAuditLogReplayJob
}

type UpdatePortalAuditLogReplayJobResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response from fetching or updating an portal audit log replay job
	PortalAuditLogReplayJob *components.PortalAuditLogReplayJob
}

func (o *UpdatePortalAuditLogReplayJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePortalAuditLogReplayJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePortalAuditLogReplayJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePortalAuditLogReplayJobResponse) GetPortalAuditLogReplayJob() *components.PortalAuditLogReplayJob {
	if o == nil {
		return nil
	}
	return o.PortalAuditLogReplayJob
}
