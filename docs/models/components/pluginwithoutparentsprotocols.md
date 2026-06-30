# PluginWithoutParentsProtocols

A string representing a protocol, such as HTTP or HTTPS.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PluginWithoutParentsProtocolsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.PluginWithoutParentsProtocols("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `PluginWithoutParentsProtocolsGrpc`           | grpc                                          |
| `PluginWithoutParentsProtocolsGrpcs`          | grpcs                                         |
| `PluginWithoutParentsProtocolsHTTP`           | http                                          |
| `PluginWithoutParentsProtocolsHTTPS`          | https                                         |
| `PluginWithoutParentsProtocolsTCP`            | tcp                                           |
| `PluginWithoutParentsProtocolsTLS`            | tls                                           |
| `PluginWithoutParentsProtocolsTLSPassthrough` | tls_passthrough                               |
| `PluginWithoutParentsProtocolsUDP`            | udp                                           |
| `PluginWithoutParentsProtocolsWs`             | ws                                            |
| `PluginWithoutParentsProtocolsWss`            | wss                                           |