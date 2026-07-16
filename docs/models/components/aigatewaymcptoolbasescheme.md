# AIGatewayMCPToolBaseScheme

The scheme of the exported API. By default, Kong will extract the scheme from API configuration. If the configured scheme is not expected, this field can be used to override it.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayMCPToolBaseSchemeHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayMCPToolBaseScheme("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `AIGatewayMCPToolBaseSchemeHTTP`  | http                              |
| `AIGatewayMCPToolBaseSchemeHTTPS` | https                             |