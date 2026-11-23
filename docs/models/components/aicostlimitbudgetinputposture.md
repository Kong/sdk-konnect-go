# AICostLimitBudgetInputPosture

What the budget does when spend reaches its limit. `advisory`, the default, changes nothing at the limit. `enforce` belongs to budget blocking, which is in development: a request that sets it fails with `400` unless the organization has budget blocking turned on or a replacement keeps a budget's stored `enforce`. No request is refused because of `posture`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostLimitBudgetInputPostureAdvisory

// Open enum: custom values can be created with a direct type cast
custom := components.AICostLimitBudgetInputPosture("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AICostLimitBudgetInputPostureAdvisory` | advisory                                |
| `AICostLimitBudgetInputPostureEnforce`  | enforce                                 |