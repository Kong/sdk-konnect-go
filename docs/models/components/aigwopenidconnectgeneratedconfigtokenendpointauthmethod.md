# AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                              |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodClientSecretPost`        | client_secret_post                                                               |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodNone`                    | none                                                                             |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                  |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                      |
| `AIGWOpenIDConnectGeneratedConfigTokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                  |