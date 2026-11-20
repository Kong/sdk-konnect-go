# AIGatewayIdentityProviderOpenIDConnectResponseSource

Where to obtain the actor token. `header` reads it from a request header. `config` uses a fixed value. `none` disables actor token support.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseSourceConfig

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseSource("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AIGatewayIdentityProviderOpenIDConnectResponseSourceConfig` | config                                                       |
| `AIGatewayIdentityProviderOpenIDConnectResponseSourceHeader` | header                                                       |
| `AIGatewayIdentityProviderOpenIDConnectResponseSourceNone`   | none                                                         |