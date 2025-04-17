// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IDFieldFilter - Filters on the given comma separated string by one or more exact match.
type IDFieldFilter struct {
	Eq  *string `queryParam:"name=eq"`
	Neq *string `queryParam:"name=neq"`
	Oeq *string `queryParam:"name=oeq"`
}

func (o *IDFieldFilter) GetEq() *string {
	if o == nil {
		return nil
	}
	return o.Eq
}

func (o *IDFieldFilter) GetNeq() *string {
	if o == nil {
		return nil
	}
	return o.Neq
}

func (o *IDFieldFilter) GetOeq() *string {
	if o == nil {
		return nil
	}
	return o.Oeq
}
