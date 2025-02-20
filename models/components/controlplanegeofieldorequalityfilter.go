// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ControlPlaneGeoFieldOrEqualityFilter - Filter a control-plane geo by determining if the value is equal to any in a set of values, where the set is a
// comma-delimited list.
type ControlPlaneGeoFieldOrEqualityFilter struct {
	Oeq string `queryParam:"name=oeq"`
}

func (o *ControlPlaneGeoFieldOrEqualityFilter) GetOeq() string {
	if o == nil {
		return ""
	}
	return o.Oeq
}
