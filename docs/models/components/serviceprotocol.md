# ServiceProtocol

The protocol used to communicate with the upstream.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ServiceProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceProtocol("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ServiceProtocolGrpc`           | grpc                            |
| `ServiceProtocolGrpcs`          | grpcs                           |
| `ServiceProtocolHTTP`           | http                            |
| `ServiceProtocolHTTPS`          | https                           |
| `ServiceProtocolTCP`            | tcp                             |
| `ServiceProtocolTLS`            | tls                             |
| `ServiceProtocolTLSPassthrough` | tls_passthrough                 |
| `ServiceProtocolUDP`            | udp                             |
| `ServiceProtocolWs`             | ws                              |
| `ServiceProtocolWss`            | wss                             |