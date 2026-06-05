# SslVersion

the ssl version to use for the pgvector database

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SslVersionAny

// Open enum: custom values can be created with a direct type cast
custom := components.SslVersion("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SslVersionAny`    | any                |
| `SslVersionTlsv12` | tlsv1_2            |
| `SslVersionTlsv13` | tlsv1_3            |