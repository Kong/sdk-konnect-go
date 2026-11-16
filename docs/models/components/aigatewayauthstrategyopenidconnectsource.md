# AIGatewayAuthStrategyOpenIDConnectSource

Where to obtain the actor token. `header` reads it from a request header. `config` uses a fixed value. `none` disables actor token support.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectSourceConfig

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectSource("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `AIGatewayAuthStrategyOpenIDConnectSourceConfig` | config                                           |
| `AIGatewayAuthStrategyOpenIDConnectSourceHeader` | header                                           |
| `AIGatewayAuthStrategyOpenIDConnectSourceNone`   | none                                             |