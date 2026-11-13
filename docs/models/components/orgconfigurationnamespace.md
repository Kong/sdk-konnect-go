# OrgConfigurationNamespace

Category that identifies the organization event's area of the product.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.OrgConfigurationNamespacePlanAndUsage

// Open enum: custom values can be created with a direct type cast
custom := components.OrgConfigurationNamespace("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `OrgConfigurationNamespacePlanAndUsage`           | plan-and-usage                                    |
| `OrgConfigurationNamespaceOrganization`           | organization                                      |
| `OrgConfigurationNamespaceDevPortal`              | dev-portal                                        |
| `OrgConfigurationNamespaceCloudGateways`          | cloud-gateways                                    |
| `OrgConfigurationNamespaceRegionalEventNamespace` | regional-event-namespace                          |
| `OrgConfigurationNamespaceGatewayManager`         | gateway-manager                                   |
| `OrgConfigurationNamespaceAdvancedAnalytics`      | advanced-analytics                                |
| `OrgConfigurationNamespaceAiManagement`           | ai-management                                     |