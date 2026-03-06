# Protocols

A string representing a protocol, such as HTTP or HTTPS.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProtocolsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.Protocols("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ProtocolsGrpc`           | grpc                      |
| `ProtocolsGrpcs`          | grpcs                     |
| `ProtocolsHTTP`           | http                      |
| `ProtocolsHTTPS`          | https                     |
| `ProtocolsTCP`            | tcp                       |
| `ProtocolsTLS`            | tls                       |
| `ProtocolsTLSPassthrough` | tls_passthrough           |
| `ProtocolsUDP`            | udp                       |
| `ProtocolsWs`             | ws                        |
| `ProtocolsWss`            | wss                       |