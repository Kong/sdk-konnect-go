# Version

the ssl version to use for the pgvector database

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VersionAny

// Open enum: custom values can be created with a direct type cast
custom := components.Version("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `VersionAny`    | any             |
| `VersionTlsv12` | tlsv1_2         |
| `VersionTlsv13` | tlsv1_3         |