# DeclarativeConfigHTTPListenerProtocol

The protocol for the listener.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeclarativeConfigHTTPListenerProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.DeclarativeConfigHTTPListenerProtocol("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `DeclarativeConfigHTTPListenerProtocolHTTP`  | http                                         |
| `DeclarativeConfigHTTPListenerProtocolHTTPS` | https                                        |