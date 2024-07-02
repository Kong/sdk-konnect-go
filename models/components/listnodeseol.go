// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ListNodesEolItems struct {
	NodeID      *string `json:"node_id,omitempty"`
	NodeVersion *string `json:"node_version,omitempty"`
	Message     *string `json:"message,omitempty"`
}

func (o *ListNodesEolItems) GetNodeID() *string {
	if o == nil {
		return nil
	}
	return o.NodeID
}

func (o *ListNodesEolItems) GetNodeVersion() *string {
	if o == nil {
		return nil
	}
	return o.NodeVersion
}

func (o *ListNodesEolItems) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type ListNodesEolPage struct {
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o *ListNodesEolPage) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}

type ListNodesEol struct {
	Items []ListNodesEolItems `json:"items,omitempty"`
	Page  *ListNodesEolPage   `json:"page,omitempty"`
}

func (o *ListNodesEol) GetItems() []ListNodesEolItems {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ListNodesEol) GetPage() *ListNodesEolPage {
	if o == nil {
		return nil
	}
	return o.Page
}