# AIGatewayModelFormatType

The format type.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelFormatTypeAnthropic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelFormatType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AIGatewayModelFormatTypeAnthropic`   | anthropic                             |
| `AIGatewayModelFormatTypeBedrock`     | bedrock                               |
| `AIGatewayModelFormatTypeCohere`      | cohere                                |
| `AIGatewayModelFormatTypeGemini`      | gemini                                |
| `AIGatewayModelFormatTypeHuggingface` | huggingface                           |
| `AIGatewayModelFormatTypeOpenai`      | openai                                |
| `AIGatewayModelFormatTypeVertex`      | vertex                                |