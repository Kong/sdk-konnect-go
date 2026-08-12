# AIGatewayModelProviderAzureService

Selects the Azure backend for this provider instance. Use `azure-openai`
for Azure OpenAI deployments or `azure-foundry` for Azure AI Foundry.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelProviderAzureServiceAzureOpenai

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelProviderAzureService("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `AIGatewayModelProviderAzureServiceAzureOpenai`  | azure-openai                                     |
| `AIGatewayModelProviderAzureServiceAzureFoundry` | azure-foundry                                    |