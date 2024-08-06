// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListServerlessCloudGatewaysResponse struct {
	// returns the pagination information
	Meta *PaginatedMeta           `json:"meta,omitempty"`
	Data []ServerlessCloudGateway `json:"data,omitempty"`
}

func (o *ListServerlessCloudGatewaysResponse) GetMeta() *PaginatedMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *ListServerlessCloudGatewaysResponse) GetData() []ServerlessCloudGateway {
	if o == nil {
		return nil
	}
	return o.Data
}