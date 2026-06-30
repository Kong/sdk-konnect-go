# UpstreamHealthchecksType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpstreamHealthchecksTypeGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.UpstreamHealthchecksType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `UpstreamHealthchecksTypeGrpc`  | grpc                            |
| `UpstreamHealthchecksTypeGrpcs` | grpcs                           |
| `UpstreamHealthchecksTypeHTTP`  | http                            |
| `UpstreamHealthchecksTypeHTTPS` | https                           |
| `UpstreamHealthchecksTypeTCP`   | tcp                             |