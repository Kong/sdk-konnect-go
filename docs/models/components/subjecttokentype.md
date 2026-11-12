# SubjectTokenType

The kind of bearer token the caller presents, which the Token Vault exchanges for the
upstream credential. Omit it unless the request shape makes the token type ambiguous: it is
otherwise resolved from the request, and failing that from the token's shape (a JWS is an
`access_token`, anything else an `api_key`).


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubjectTokenTypeAccessToken

// Open enum: custom values can be created with a direct type cast
custom := components.SubjectTokenType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SubjectTokenTypeAccessToken` | access_token                  |
| `SubjectTokenTypeAPIKey`      | api_key                       |