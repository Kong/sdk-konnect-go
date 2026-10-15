# ServiceRequest2Protocol

The protocol used to communicate with the upstream.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ServiceRequest2ProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceRequest2Protocol("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ServiceRequest2ProtocolGrpc`           | grpc                                    |
| `ServiceRequest2ProtocolGrpcs`          | grpcs                                   |
| `ServiceRequest2ProtocolHTTP`           | http                                    |
| `ServiceRequest2ProtocolHTTPS`          | https                                   |
| `ServiceRequest2ProtocolTCP`            | tcp                                     |
| `ServiceRequest2ProtocolTLS`            | tls                                     |
| `ServiceRequest2ProtocolTLSPassthrough` | tls_passthrough                         |
| `ServiceRequest2ProtocolUDP`            | udp                                     |
| `ServiceRequest2ProtocolWs`             | ws                                      |
| `ServiceRequest2ProtocolWss`            | wss                                     |