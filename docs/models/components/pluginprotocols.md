# PluginProtocols

A string representing a protocol, such as HTTP or HTTPS.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PluginProtocolsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.PluginProtocols("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `PluginProtocolsGrpc`           | grpc                            |
| `PluginProtocolsGrpcs`          | grpcs                           |
| `PluginProtocolsHTTP`           | http                            |
| `PluginProtocolsHTTPS`          | https                           |
| `PluginProtocolsTCP`            | tcp                             |
| `PluginProtocolsTLS`            | tls                             |
| `PluginProtocolsTLSPassthrough` | tls_passthrough                 |
| `PluginProtocolsUDP`            | udp                             |
| `PluginProtocolsWs`             | ws                              |
| `PluginProtocolsWss`            | wss                             |