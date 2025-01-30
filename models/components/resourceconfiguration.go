// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ResourceConfiguration - Object containing information about a resource configuration.
type ResourceConfiguration struct {
	ID string `json:"id"`
	// Enumeration of configuration qualifiers available for organization-wide configuration.
	Qualifier ResourceConfigurationQualifier `json:"qualifier"`
	// The value of this resource configuration.
	Value int64 `json:"value"`
}

func (o *ResourceConfiguration) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ResourceConfiguration) GetQualifier() ResourceConfigurationQualifier {
	if o == nil {
		return ResourceConfigurationQualifier("")
	}
	return o.Qualifier
}

func (o *ResourceConfiguration) GetValue() int64 {
	if o == nil {
		return 0
	}
	return o.Value
}
