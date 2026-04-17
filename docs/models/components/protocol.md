# Protocol

The protocol used to communicate with the upstream.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.Protocol("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ProtocolGrpc`           | grpc                     |
| `ProtocolGrpcs`          | grpcs                    |
| `ProtocolHTTP`           | http                     |
| `ProtocolHTTPS`          | https                    |
| `ProtocolTCP`            | tcp                      |
| `ProtocolTLS`            | tls                      |
| `ProtocolTLSPassthrough` | tls_passthrough          |
| `ProtocolUDP`            | udp                      |
| `ProtocolWs`             | ws                       |
| `ProtocolWss`            | wss                      |