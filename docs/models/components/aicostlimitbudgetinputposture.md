# AICostLimitBudgetInputPosture

How the budget responds when spend reaches its limit. With `advisory`, the default, requests continue past the limit.

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