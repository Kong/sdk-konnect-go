// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// CustomDomainStateMetadata - Metadata describing the backing state of the custom domain and why it may be in an erroneous state.
type CustomDomainStateMetadata struct {
	// Reported status of the custom domain from backing infrastructure.
	ReportedStatus *string `json:"reported_status,omitempty"`
	// Reason why the custom domain may be in an erroneous state, reported from backing infrastructure.
	//
	Reason *string `json:"reason,omitempty"`
}

func (o *CustomDomainStateMetadata) GetReportedStatus() *string {
	if o == nil {
		return nil
	}
	return o.ReportedStatus
}

func (o *CustomDomainStateMetadata) GetReason() *string {
	if o == nil {
		return nil
	}
	return o.Reason
}

// CustomDomain - Object containing information about a custom domain for a control-plane.
type CustomDomain struct {
	ID             string `json:"id"`
	ControlPlaneID string `json:"control_plane_id"`
	// Set of control-plane geos supported for deploying cloud-gateways configurations.
	ControlPlaneGeo ControlPlaneGeo `json:"control_plane_geo"`
	// Domain name of the custom domain.
	Domain string `json:"domain"`
	// Certificate ID for the certificate representing this domain and stored on data-planes for this
	// control-plane. Can be retrieved via the control-planes API for this custom domain's control-plane.
	//
	CertificateID *string `json:"certificate_id,omitempty"`
	// Server Name Indication ID for this domain and stored on data-planes for this control-plane. Can be retrieved
	// via the control-planes API for this custom domain's control-plane.
	//
	SniID *string `json:"sni_id,omitempty"`
	// State of the custom domain.
	State CustomDomainState `json:"state"`
	// Metadata describing the backing state of the custom domain and why it may be in an erroneous state.
	//
	StateMetadata CustomDomainStateMetadata `json:"state_metadata"`
	// Monotonically-increasing version count of the custom domain, to indicate the order of updates to the custom
	// domain.
	//
	EntityVersion int64 `json:"entity_version"`
	// An RFC-3339 timestamp representation of custom domain creation date.
	CreatedAt time.Time `json:"created_at"`
	// An RFC-3339 timestamp representation of custom domain update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (c CustomDomain) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomDomain) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomDomain) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CustomDomain) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CustomDomain) GetControlPlaneGeo() ControlPlaneGeo {
	if o == nil {
		return ControlPlaneGeo("")
	}
	return o.ControlPlaneGeo
}

func (o *CustomDomain) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *CustomDomain) GetCertificateID() *string {
	if o == nil {
		return nil
	}
	return o.CertificateID
}

func (o *CustomDomain) GetSniID() *string {
	if o == nil {
		return nil
	}
	return o.SniID
}

func (o *CustomDomain) GetState() CustomDomainState {
	if o == nil {
		return CustomDomainState("")
	}
	return o.State
}

func (o *CustomDomain) GetStateMetadata() CustomDomainStateMetadata {
	if o == nil {
		return CustomDomainStateMetadata{}
	}
	return o.StateMetadata
}

func (o *CustomDomain) GetEntityVersion() int64 {
	if o == nil {
		return 0
	}
	return o.EntityVersion
}

func (o *CustomDomain) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CustomDomain) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
