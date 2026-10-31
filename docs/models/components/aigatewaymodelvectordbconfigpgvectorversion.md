# AIGatewayModelVectorDBConfigPgVectorVersion

the ssl version to use for the pgvector database

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelVectorDBConfigPgVectorVersionAny

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelVectorDBConfigPgVectorVersion("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AIGatewayModelVectorDBConfigPgVectorVersionAny`    | any                                                 |
| `AIGatewayModelVectorDBConfigPgVectorVersionTlsv12` | tlsv1_2                                             |
| `AIGatewayModelVectorDBConfigPgVectorVersionTlsv13` | tlsv1_3                                             |