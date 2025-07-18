// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// MoveDocumentRequestPayload - move document request payload
type MoveDocumentRequestPayload struct {
	// parent document id
	ParentDocumentID *string `json:"parent_document_id,omitempty"`
	// index of the document in the parent document's children
	Index *int64 `json:"index,omitempty"`
}

func (o *MoveDocumentRequestPayload) GetParentDocumentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentDocumentID
}

func (o *MoveDocumentRequestPayload) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}
