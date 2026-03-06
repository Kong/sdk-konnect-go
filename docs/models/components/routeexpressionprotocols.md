# RouteExpressionProtocols

A string representing a protocol, such as HTTP or HTTPS.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RouteExpressionProtocolsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.RouteExpressionProtocols("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RouteExpressionProtocolsGrpc`           | grpc                                     |
| `RouteExpressionProtocolsGrpcs`          | grpcs                                    |
| `RouteExpressionProtocolsHTTP`           | http                                     |
| `RouteExpressionProtocolsHTTPS`          | https                                    |
| `RouteExpressionProtocolsTCP`            | tcp                                      |
| `RouteExpressionProtocolsTLS`            | tls                                      |
| `RouteExpressionProtocolsTLSPassthrough` | tls_passthrough                          |
| `RouteExpressionProtocolsUDP`            | udp                                      |
| `RouteExpressionProtocolsWs`             | ws                                       |
| `RouteExpressionProtocolsWss`            | wss                                      |