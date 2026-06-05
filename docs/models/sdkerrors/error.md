# Error

A single error code from the following enum.

  invalid_request
     The request is missing a required parameter, includes an unsupported parameter value (other than grant type),
     repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the
     client, or is otherwise malformed.

  invalid_client
     Client authentication failed (e.g., unknown client, no client authentication included, or unsupported
     authentication method).

  invalid_grant
     The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is
     invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to
     another client.

  unauthorized_client
     The authenticated client is not authorized to use this authorization grant type.

  unsupported_grant_type
     The authorization grant type is not supported by the authorization server.

  invalid_scope
     The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
)

value := sdkerrors.ErrorInvalidRequest

// Open enum: custom values can be created with a direct type cast
custom := sdkerrors.Error("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ErrorInvalidRequest`       | invalid_request             |
| `ErrorInvalidClient`        | invalid_client              |
| `ErrorInvalidGrant`         | invalid_grant               |
| `ErrorUnauthorizedClient`   | unauthorized_client         |
| `ErrorUnsupportedGrantType` | unsupported_grant_type      |
| `ErrorInvalidScope`         | invalid_scope               |