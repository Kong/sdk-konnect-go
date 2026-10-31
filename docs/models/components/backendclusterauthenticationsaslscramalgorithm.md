# BackendClusterAuthenticationSaslScramAlgorithm

The algorithm used for SASL/SCRAM authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BackendClusterAuthenticationSaslScramAlgorithmSha256

// Open enum: custom values can be created with a direct type cast
custom := components.BackendClusterAuthenticationSaslScramAlgorithm("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `BackendClusterAuthenticationSaslScramAlgorithmSha256` | sha256                                                 |
| `BackendClusterAuthenticationSaslScramAlgorithmSha512` | sha512                                                 |