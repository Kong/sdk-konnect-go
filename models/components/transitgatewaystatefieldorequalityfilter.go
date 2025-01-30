// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TransitGatewayStateFieldOrEqualityFilter - Filter transit-gateway state by determining if the value is equal to any in a set of values, where the set is a
// comma-delimited list.
type TransitGatewayStateFieldOrEqualityFilter struct {
	Oeq string `queryParam:"name=oeq"`
}

func (o *TransitGatewayStateFieldOrEqualityFilter) GetOeq() string {
	if o == nil {
		return ""
	}
	return o.Oeq
}
