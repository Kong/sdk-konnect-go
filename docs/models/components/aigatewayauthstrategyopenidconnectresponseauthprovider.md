# AIGatewayAuthStrategyOpenIDConnectResponseAuthProvider

Auth providers to be used to authenticate to a Cloud Provider's Redis instance.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseAuthProviderAws

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseAuthProvider("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthProviderAws`   | aws                                                           |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthProviderAzure` | azure                                                         |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthProviderGcp`   | gcp                                                           |