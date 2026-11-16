# DestinationSigningAlgo

Signing algorithm of the signing key.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DestinationSigningAlgoHmacSha256

// Open enum: custom values can be created with a direct type cast
custom := components.DestinationSigningAlgo("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `DestinationSigningAlgoHmacSha256` | hmac-sha256                        |
| `DestinationSigningAlgoEd25519`    | ed25519                            |