# AIGatewayTargetMistralConfigFormat

The request format to use when communicating with the Mistral model.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayTargetMistralConfigFormatOllama

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayTargetMistralConfigFormat("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `AIGatewayTargetMistralConfigFormatOllama` | ollama                                     |
| `AIGatewayTargetMistralConfigFormatOpenai` | openai                                     |