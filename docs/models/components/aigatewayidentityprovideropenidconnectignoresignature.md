# AIGatewayIdentityProviderOpenIDConnectIgnoreSignature

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectIgnoreSignature("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureAuthorizationCode` | authorization_code                                                       |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureClientCredentials` | client_credentials                                                       |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureIntrospection`     | introspection                                                            |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignaturePassword`          | password                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureRefreshToken`      | refresh_token                                                            |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureSession`           | session                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectIgnoreSignatureUserinfo`          | userinfo                                                                 |