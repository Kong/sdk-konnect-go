# ServiceRequestProtocol

The protocol used to communicate with the upstream.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ServiceRequestProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceRequestProtocol("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ServiceRequestProtocolGrpc`           | grpc                                   |
| `ServiceRequestProtocolGrpcs`          | grpcs                                  |
| `ServiceRequestProtocolHTTP`           | http                                   |
| `ServiceRequestProtocolHTTPS`          | https                                  |
| `ServiceRequestProtocolTCP`            | tcp                                    |
| `ServiceRequestProtocolTLS`            | tls                                    |
| `ServiceRequestProtocolTLSPassthrough` | tls_passthrough                        |
| `ServiceRequestProtocolUDP`            | udp                                    |
| `ServiceRequestProtocolWs`             | ws                                     |
| `ServiceRequestProtocolWss`            | wss                                    |