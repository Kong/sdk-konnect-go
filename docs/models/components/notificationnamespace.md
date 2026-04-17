# NotificationNamespace

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NotificationNamespacePlanAndUsage

// Open enum: custom values can be created with a direct type cast
custom := components.NotificationNamespace("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `NotificationNamespacePlanAndUsage`   | plan-and-usage                        |
| `NotificationNamespaceOrganization`   | organization                          |
| `NotificationNamespaceDevPortal`      | dev-portal                            |
| `NotificationNamespaceCloudGateways`  | cloud-gateways                        |
| `NotificationNamespaceGatewayManager` | gateway-manager                       |