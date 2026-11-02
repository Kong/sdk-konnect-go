# BackendClusterAuthenticationSaslScramSensitiveDataAwareAlgorithm

The algorithm used for SASL/SCRAM authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BackendClusterAuthenticationSaslScramSensitiveDataAwareAlgorithmSha256

// Open enum: custom values can be created with a direct type cast
custom := components.BackendClusterAuthenticationSaslScramSensitiveDataAwareAlgorithm("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `BackendClusterAuthenticationSaslScramSensitiveDataAwareAlgorithmSha256` | sha256                                                                   |
| `BackendClusterAuthenticationSaslScramSensitiveDataAwareAlgorithmSha512` | sha512                                                                   |