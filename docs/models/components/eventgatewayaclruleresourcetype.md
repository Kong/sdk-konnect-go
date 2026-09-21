# EventGatewayACLRuleResourceType

This rule applies to access only for type of resource

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayACLRuleResourceTypeTopic

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayACLRuleResourceType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `EventGatewayACLRuleResourceTypeTopic`           | topic                                            |
| `EventGatewayACLRuleResourceTypeGroup`           | group                                            |
| `EventGatewayACLRuleResourceTypeTransactionalID` | transactional_id                                 |
| `EventGatewayACLRuleResourceTypeCluster`         | cluster                                          |