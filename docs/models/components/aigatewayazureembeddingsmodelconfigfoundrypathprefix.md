# AIGatewayAzureEmbeddingsModelConfigFoundryPathPrefix

The API path prefix for the Azure AI Foundry endpoint, selecting the model's
API surface. `/openai/v1` targets the OpenAI-compatible surface; `/anthropic/v1`
targets the Anthropic surface. Applies when the Azure provider's `service` is
`azure-foundry`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAzureEmbeddingsModelConfigFoundryPathPrefixRootOpenaiV1

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAzureEmbeddingsModelConfigFoundryPathPrefix("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AIGatewayAzureEmbeddingsModelConfigFoundryPathPrefixRootOpenaiV1`    | /openai/v1                                                            |
| `AIGatewayAzureEmbeddingsModelConfigFoundryPathPrefixRootAnthropicV1` | /anthropic/v1                                                         |