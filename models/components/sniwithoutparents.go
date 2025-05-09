// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// SNIWithoutParentsCertificate - The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object.
type SNIWithoutParentsCertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *SNIWithoutParentsCertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SNIWithoutParents - An SNI object represents a many-to-one mapping of hostnames to a certificate. That is, a certificate object can have many hostnames associated with it; when Kong receives an SSL request, it uses the SNI field in the Client Hello to lookup the certificate object based on the SNI associated with the certificate.
type SNIWithoutParents struct {
	// The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object.
	Certificate *SNIWithoutParentsCertificate `json:"certificate,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	// The SNI name to associate with the given certificate.
	Name string `json:"name"`
	// An optional set of strings associated with the SNIs for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *SNIWithoutParents) GetCertificate() *SNIWithoutParentsCertificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *SNIWithoutParents) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SNIWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SNIWithoutParents) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SNIWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SNIWithoutParents) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
