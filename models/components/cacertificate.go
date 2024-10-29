// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CACertificate - A CA certificate object represents a trusted CA. These objects are used by Kong to verify the validity of a client or server certificate.
type CACertificate struct {
	// PEM-encoded public certificate of the CA.
	Cert string `json:"cert"`
	// SHA256 hex digest of the public certificate. This field is read-only and it cannot be set by the caller, the value is automatically computed.
	CertDigest *string `json:"cert_digest,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	// An optional set of strings associated with the Certificate for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *CACertificate) GetCert() string {
	if o == nil {
		return ""
	}
	return o.Cert
}

func (o *CACertificate) GetCertDigest() *string {
	if o == nil {
		return nil
	}
	return o.CertDigest
}

func (o *CACertificate) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CACertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CACertificate) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CACertificate) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
