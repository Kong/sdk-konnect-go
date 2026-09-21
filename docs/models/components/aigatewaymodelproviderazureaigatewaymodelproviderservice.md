# AIGatewayModelProviderAzureAIGatewayModelProviderService

Selects the Azure backend for this provider instance. Use `azure-openai`
for Azure OpenAI deployments or `azure-foundry` for Azure AI Foundry.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelProviderAzureAIGatewayModelProviderServiceAzureOpenai

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelProviderAzureAIGatewayModelProviderService("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AIGatewayModelProviderAzureAIGatewayModelProviderServiceAzureOpenai`  | azure-openai                                                           |
| `AIGatewayModelProviderAzureAIGatewayModelProviderServiceAzureFoundry` | azure-foundry                                                          |