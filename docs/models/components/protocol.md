# Protocol

The protocol to connect with.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.Protocol("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ProtocolHTTP`  | http            |
| `ProtocolHTTPS` | https           |