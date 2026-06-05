# CreateAIGatewayConsumerRequestType

The type of the consumer.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateAIGatewayConsumerRequestTypeAPIKey

// Open enum: custom values can be created with a direct type cast
custom := components.CreateAIGatewayConsumerRequestType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateAIGatewayConsumerRequestTypeAPIKey` | api-key                                    |
| `CreateAIGatewayConsumerRequestTypeOauth`  | oauth                                      |