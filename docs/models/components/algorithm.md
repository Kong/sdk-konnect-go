# Algorithm

The algorithm used for SASL/SCRAM authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AlgorithmSha256

// Open enum: custom values can be created with a direct type cast
custom := components.Algorithm("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `AlgorithmSha256` | sha256            |
| `AlgorithmSha512` | sha512            |