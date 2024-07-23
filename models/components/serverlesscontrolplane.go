// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ServerlessControlPlane struct {
	// ID of the serverless cloud gateway CP.
	ID string `json:"id"`
	// The prefix of the serverless cloud gateway CP.
	Prefix string `json:"prefix"`
	// The control plane region.
	Region CpRegion `json:"region"`
}

func (o *ServerlessControlPlane) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ServerlessControlPlane) GetPrefix() string {
	if o == nil {
		return ""
	}
	return o.Prefix
}

func (o *ServerlessControlPlane) GetRegion() CpRegion {
	if o == nil {
		return CpRegion("")
	}
	return o.Region
}
