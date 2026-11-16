# AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProvider

Auth providers to be used to authenticate to a Cloud Provider's Redis instance.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProviderAws

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProvider("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProviderAws`   | aws                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProviderAzure` | azure                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseConfigAuthProviderGcp`   | gcp                                                                 |