# Min

Minimum TLS version to use.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MinTlSv12

// Open enum: custom values can be created with a direct type cast
custom := components.Min("custom_value")
```


## Values

| Name        | Value       |
| ----------- | ----------- |
| `MinTlSv12` | TLSv1.2     |
| `MinTlSv13` | TLSv1.3     |