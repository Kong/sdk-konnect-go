// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CursorMetaWithSizeAndTotal struct {
	// URI to the next page
	Next *string `json:"next"`
	// Requested page size
	Size float64 `json:"size"`
	// Total number of objects in the collection; will only be present on the first page
	Total *float64 `json:"total"`
}

func (o *CursorMetaWithSizeAndTotal) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *CursorMetaWithSizeAndTotal) GetSize() float64 {
	if o == nil {
		return 0.0
	}
	return o.Size
}

func (o *CursorMetaWithSizeAndTotal) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}
