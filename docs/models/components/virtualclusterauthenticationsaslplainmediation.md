# VirtualClusterAuthenticationSaslPlainMediation

The mediation type for SASL/PLAIN authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VirtualClusterAuthenticationSaslPlainMediationPassthrough

// Open enum: custom values can be created with a direct type cast
custom := components.VirtualClusterAuthenticationSaslPlainMediation("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `VirtualClusterAuthenticationSaslPlainMediationPassthrough` | passthrough                                                 |
| `VirtualClusterAuthenticationSaslPlainMediationTerminate`   | terminate                                                   |