# UpstreamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpstreamTypeGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.UpstreamType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `UpstreamTypeGrpc`  | grpc                |
| `UpstreamTypeGrpcs` | grpcs               |
| `UpstreamTypeHTTP`  | http                |
| `UpstreamTypeHTTPS` | https               |
| `UpstreamTypeTCP`   | tcp                 |