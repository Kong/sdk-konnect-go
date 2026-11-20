# AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProvider

Auth providers to be used to authenticate to a Cloud Provider's Redis instance.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProviderAws

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProvider("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProviderAws`   | aws                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProviderAzure` | azure                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseConfigAuthProviderGcp`   | gcp                                                                     |