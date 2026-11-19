# Integrations

A tool whose own request headers the plugin reads and maps onto the standard cost tags.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IntegrationsClaudeCode

// Open enum: custom values can be created with a direct type cast
custom := components.Integrations("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `IntegrationsClaudeCode`          | claude-code                       |
| `IntegrationsCodex`               | codex                             |
| `IntegrationsOpencode`            | opencode                          |
| `IntegrationsCopilot`             | copilot                           |
| `IntegrationsHelicone`            | helicone                          |
| `IntegrationsPortkey`             | portkey                           |
| `IntegrationsCloudflareAiGateway` | cloudflare-ai-gateway             |