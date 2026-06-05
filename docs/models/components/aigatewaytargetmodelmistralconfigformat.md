# AIGatewayTargetModelMistralConfigFormat

The request format to use when communicating with the Mistral model.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayTargetModelMistralConfigFormatOllama

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayTargetModelMistralConfigFormat("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AIGatewayTargetModelMistralConfigFormatOllama` | ollama                                          |
| `AIGatewayTargetModelMistralConfigFormatOpenai` | openai                                          |