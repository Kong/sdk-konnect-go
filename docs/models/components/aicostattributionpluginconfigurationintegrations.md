# AICostAttributionPluginConfigurationIntegrations

A tool whose own request headers the plugin reads and maps onto the standard cost tags.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostAttributionPluginConfigurationIntegrationsClaudeCode

// Open enum: custom values can be created with a direct type cast
custom := components.AICostAttributionPluginConfigurationIntegrations("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AICostAttributionPluginConfigurationIntegrationsClaudeCode`          | claude-code                                                           |
| `AICostAttributionPluginConfigurationIntegrationsCodex`               | codex                                                                 |
| `AICostAttributionPluginConfigurationIntegrationsOpencode`            | opencode                                                              |
| `AICostAttributionPluginConfigurationIntegrationsCopilot`             | copilot                                                               |
| `AICostAttributionPluginConfigurationIntegrationsHelicone`            | helicone                                                              |
| `AICostAttributionPluginConfigurationIntegrationsPortkey`             | portkey                                                               |
| `AICostAttributionPluginConfigurationIntegrationsCloudflareAiGateway` | cloudflare-ai-gateway                                                 |