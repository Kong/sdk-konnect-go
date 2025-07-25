// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreatePortalCustomDomainRequest - Create a portal custom domain.
type CreatePortalCustomDomainRequest struct {
	Hostname string                      `json:"hostname"`
	Enabled  bool                        `json:"enabled"`
	Ssl      CreatePortalCustomDomainSSL `json:"ssl"`
}

func (o *CreatePortalCustomDomainRequest) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *CreatePortalCustomDomainRequest) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreatePortalCustomDomainRequest) GetSsl() CreatePortalCustomDomainSSL {
	if o == nil {
		return CreatePortalCustomDomainSSL{}
	}
	return o.Ssl
}
