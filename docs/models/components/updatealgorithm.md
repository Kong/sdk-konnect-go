# UpdateAlgorithm

Algorithm used in the key signing process

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpdateAlgorithmRs256

// Open enum: custom values can be created with a direct type cast
custom := components.UpdateAlgorithm("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `UpdateAlgorithmRs256` | RS256                  |
| `UpdateAlgorithmRs384` | RS384                  |
| `UpdateAlgorithmRs512` | RS512                  |
| `UpdateAlgorithmPs256` | PS256                  |
| `UpdateAlgorithmPs384` | PS384                  |
| `UpdateAlgorithmPs512` | PS512                  |