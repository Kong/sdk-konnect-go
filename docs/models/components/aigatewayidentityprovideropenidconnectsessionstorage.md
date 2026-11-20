# AIGatewayIdentityProviderOpenIDConnectSessionStorage

The session storage for session data: - `cookie`: stores session data with the session cookie (the session cannot be invalidated or revoked without changing session secret, but is stateless, and doesn't require a database) - `memcache`: stores session data in memcached - `redis`: stores session data in Redis.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectSessionStorageCookie

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectSessionStorage("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectSessionStorageCookie`    | cookie                                                          |
| `AIGatewayIdentityProviderOpenIDConnectSessionStorageMemcache`  | memcache                                                        |
| `AIGatewayIdentityProviderOpenIDConnectSessionStorageMemcached` | memcached                                                       |
| `AIGatewayIdentityProviderOpenIDConnectSessionStorageRedis`     | redis                                                           |