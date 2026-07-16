# AIGatewayConsumerType

The type of the consumer.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayConsumerTypeAPIKey

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayConsumerType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AIGatewayConsumerTypeAPIKey` | api-key                       |
| `AIGatewayConsumerTypeOauth`  | oauth                         |