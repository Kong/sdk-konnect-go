// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// StringFieldOEQFilter - Returns entities that exact match any of the comma-delimited phrases in the filter string.
type StringFieldOEQFilter struct {
	Oeq string `queryParam:"name=oeq"`
}

func (o *StringFieldOEQFilter) GetOeq() string {
	if o == nil {
		return ""
	}
	return o.Oeq
}