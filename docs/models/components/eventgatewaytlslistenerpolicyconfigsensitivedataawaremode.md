# EventGatewayTLSListenerPolicyConfigSensitiveDataAwareMode

* required - Reject TLS connections without a valid client certificate.
* requested - Request a client certificate during the TLS handshake, but allow connections without one (falls back to other configured authentication methods). If a certificate is presented but cannot be verified, the connection is closed.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayTLSListenerPolicyConfigSensitiveDataAwareModeRequired

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayTLSListenerPolicyConfigSensitiveDataAwareMode("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `EventGatewayTLSListenerPolicyConfigSensitiveDataAwareModeRequired`  | required                                                             |
| `EventGatewayTLSListenerPolicyConfigSensitiveDataAwareModeRequested` | requested                                                            |