# VirtualClusterAuthenticationSaslScramAlgorithm

The algorithm used for SASL/SCRAM authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VirtualClusterAuthenticationSaslScramAlgorithmSha256

// Open enum: custom values can be created with a direct type cast
custom := components.VirtualClusterAuthenticationSaslScramAlgorithm("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `VirtualClusterAuthenticationSaslScramAlgorithmSha256` | sha256                                                 |
| `VirtualClusterAuthenticationSaslScramAlgorithmSha512` | sha512                                                 |