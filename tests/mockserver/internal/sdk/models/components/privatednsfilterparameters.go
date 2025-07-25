// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PrivateDNSFilterParameters struct {
	Name  *CloudGatewaysStringFieldFilterOverride `queryParam:"name=name"`
	State *PrivateDNSStateFieldFilter             `queryParam:"name=state"`
}

func (o *PrivateDNSFilterParameters) GetName() *CloudGatewaysStringFieldFilterOverride {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PrivateDNSFilterParameters) GetState() *PrivateDNSStateFieldFilter {
	if o == nil {
		return nil
	}
	return o.State
}
