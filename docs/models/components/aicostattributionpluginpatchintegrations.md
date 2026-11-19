# AICostAttributionPluginPatchIntegrations

A tool whose own request headers the plugin reads and maps onto the standard cost tags.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostAttributionPluginPatchIntegrationsClaudeCode

// Open enum: custom values can be created with a direct type cast
custom := components.AICostAttributionPluginPatchIntegrations("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AICostAttributionPluginPatchIntegrationsClaudeCode`          | claude-code                                                   |
| `AICostAttributionPluginPatchIntegrationsCodex`               | codex                                                         |
| `AICostAttributionPluginPatchIntegrationsOpencode`            | opencode                                                      |
| `AICostAttributionPluginPatchIntegrationsCopilot`             | copilot                                                       |
| `AICostAttributionPluginPatchIntegrationsHelicone`            | helicone                                                      |
| `AICostAttributionPluginPatchIntegrationsPortkey`             | portkey                                                       |
| `AICostAttributionPluginPatchIntegrationsCloudflareAiGateway` | cloudflare-ai-gateway                                         |