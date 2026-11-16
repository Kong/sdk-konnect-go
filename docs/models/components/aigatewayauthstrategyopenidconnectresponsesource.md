# AIGatewayAuthStrategyOpenIDConnectResponseSource

Where to obtain the actor token. `header` reads it from a request header. `config` uses a fixed value. `none` disables actor token support.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseSourceConfig

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseSource("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseSourceConfig` | config                                                   |
| `AIGatewayAuthStrategyOpenIDConnectResponseSourceHeader` | header                                                   |
| `AIGatewayAuthStrategyOpenIDConnectResponseSourceNone`   | none                                                     |