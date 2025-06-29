// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// HighVelocityProviderRegion - Region ID and human-readable name for a cloud provider region.
type HighVelocityProviderRegion struct {
	// Region ID for cloud provider region.
	Region string `json:"region"`
}

func (o *HighVelocityProviderRegion) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}
