# RouteJSONProtocols

A string representing a protocol, such as HTTP or HTTPS.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RouteJSONProtocolsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.RouteJSONProtocols("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `RouteJSONProtocolsGrpc`           | grpc                               |
| `RouteJSONProtocolsGrpcs`          | grpcs                              |
| `RouteJSONProtocolsHTTP`           | http                               |
| `RouteJSONProtocolsHTTPS`          | https                              |
| `RouteJSONProtocolsTCP`            | tcp                                |
| `RouteJSONProtocolsTLS`            | tls                                |
| `RouteJSONProtocolsTLSPassthrough` | tls_passthrough                    |
| `RouteJSONProtocolsUDP`            | udp                                |
| `RouteJSONProtocolsWs`             | ws                                 |
| `RouteJSONProtocolsWss`            | wss                                |