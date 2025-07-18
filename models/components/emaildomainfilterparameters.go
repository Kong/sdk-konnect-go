// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EmailDomainFilterParameters struct {
	// Filters on the given string field value by exact match inequality.
	Domain *StringFieldFilter `queryParam:"name=domain"`
	// Filters on the given string field value by exact match inequality.
	VerificationStatus *StringFieldFilter `queryParam:"name=verification_status"`
}

func (o *EmailDomainFilterParameters) GetDomain() *StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *EmailDomainFilterParameters) GetVerificationStatus() *StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.VerificationStatus
}
