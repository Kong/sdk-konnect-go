# EventGatewayRequestRuleAction

What to do when a rule evaluates to `false`.

`reject` fails the request with the `POLICY_VIOLATION` error code.
`passthrough` lets the request continue, but logs the violation in the same way as `reject`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayRequestRuleActionReject

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayRequestRuleAction("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `EventGatewayRequestRuleActionReject`      | reject                                     |
| `EventGatewayRequestRuleActionPassthrough` | passthrough                                |