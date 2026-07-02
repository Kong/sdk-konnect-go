# EventGatewayACLRuleAction

How to handle the request if the rule matches

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayACLRuleActionAllow

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayACLRuleAction("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `EventGatewayACLRuleActionAllow` | allow                            |
| `EventGatewayACLRuleActionDeny`  | deny                             |