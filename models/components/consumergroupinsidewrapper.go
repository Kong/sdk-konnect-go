// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ConsumerGroupInsideWrapper struct {
	ConsumerGroup *ConsumerGroup `json:"consumer_group,omitempty"`
}

func (o *ConsumerGroupInsideWrapper) GetConsumerGroup() *ConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}
