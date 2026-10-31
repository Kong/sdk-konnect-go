# AuthServerSigningAlgorithm

Algorithm used in the key signing process

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthServerSigningAlgorithmRs256

// Open enum: custom values can be created with a direct type cast
custom := components.AuthServerSigningAlgorithm("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `AuthServerSigningAlgorithmRs256` | RS256                             |
| `AuthServerSigningAlgorithmRs384` | RS384                             |
| `AuthServerSigningAlgorithmRs512` | RS512                             |
| `AuthServerSigningAlgorithmPs256` | PS256                             |
| `AuthServerSigningAlgorithmPs384` | PS384                             |
| `AuthServerSigningAlgorithmPs512` | PS512                             |