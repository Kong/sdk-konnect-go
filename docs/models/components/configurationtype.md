# ConfigurationType

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Type of Cloud Gateway: `api` for an API Gateway or `ai` for an AI Gateway.
Applies only to dedicated Cloud Gateways. Defaults to `api` when omitted.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConfigurationTypeAPI

// Open enum: custom values can be created with a direct type cast
custom := components.ConfigurationType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ConfigurationTypeAPI` | api                    |
| `ConfigurationTypeAi`  | ai                     |