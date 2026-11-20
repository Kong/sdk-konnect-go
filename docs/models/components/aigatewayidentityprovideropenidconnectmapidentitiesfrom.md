# AIGatewayIdentityProviderOpenIDConnectMapIdentitiesFrom

Which token's claims to use for consumer, consumer group, and principal mapping. Only takes effect when token exchange is configured. `exchanged_tokens` (default) uses the token(s) returned by the exchange or tokens derived from it, as today. `subject_token` uses the original, pre-exchange bearer token's claims instead.

**Requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectMapIdentitiesFromExchangedTokens

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectMapIdentitiesFrom("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AIGatewayIdentityProviderOpenIDConnectMapIdentitiesFromExchangedTokens` | exchanged_tokens                                                         |
| `AIGatewayIdentityProviderOpenIDConnectMapIdentitiesFromSubjectToken`    | subject_token                                                            |