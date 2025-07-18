// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type APIVersionFilterParameters struct {
	// Filters on the given string field value by exact match inequality.
	Type *StringFieldFilter `queryParam:"name=type"`
	// Filters on the given string field value by exact match inequality.
	Version *StringFieldFilter `queryParam:"name=version"`
}

func (o *APIVersionFilterParameters) GetType() *StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *APIVersionFilterParameters) GetVersion() *StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Version
}
