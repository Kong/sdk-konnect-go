# Algorithm

Algorithm used in the key signing process

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AlgorithmRs256

// Open enum: custom values can be created with a direct type cast
custom := components.Algorithm("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `AlgorithmRs256` | RS256            |
| `AlgorithmRs384` | RS384            |
| `AlgorithmRs512` | RS512            |
| `AlgorithmPs256` | PS256            |
| `AlgorithmPs384` | PS384            |
| `AlgorithmPs512` | PS512            |