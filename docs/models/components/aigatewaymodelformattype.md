# AIGatewayModelFormatType

The request and response format this model supports. Set to a specific format to restrict the model to that format only, or to `*` to allow all formats that the targets support at the gateway's min_runtime_version. A format that needs a newer min_runtime_version is left out of `*`. The `*` format never includes passthrough.

**`passthrough` requires a minimum runtime version of `2.2`**.

**`typesafe` requires a minimum runtime version of `2.2`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelFormatTypeWildcard

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelFormatType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AIGatewayModelFormatTypeWildcard`    | *                                     |
| `AIGatewayModelFormatTypeAnthropic`   | anthropic                             |
| `AIGatewayModelFormatTypeBedrock`     | bedrock                               |
| `AIGatewayModelFormatTypeCohere`      | cohere                                |
| `AIGatewayModelFormatTypeGemini`      | gemini                                |
| `AIGatewayModelFormatTypeHuggingface` | huggingface                           |
| `AIGatewayModelFormatTypeOpenai`      | openai                                |
| `AIGatewayModelFormatTypePassthrough` | passthrough                           |
| `AIGatewayModelFormatTypeTypesafe`    | typesafe                              |