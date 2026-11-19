# AICostBudgetScopeField

A usage field that a budget scope can filter on, and that `each` can split a budget by:

- `subject`: the subject reference, such as `principal/{uuid}`. Equal identifiers from different identity systems stay distinct.
- `subject_id`: the subject's identifier without its type, which pools equal identifiers across identity systems.
- `subject_type`: the subject's identity system: `principal`, `consumer`, or `oidc`.
- `group`: the usage's group references. Usage can carry several groups. `eq` and `in` match when any of them matches. `neq` and `nin` match only usage with at least one group, and exclude it when any of its groups matches.
- `team`: the usage's Kong Identity team references, with the membership rules of `group`.
- `model`: the model that the provider reports as served.
- `provider`: the provider that served the invocation.
- `request_mode`: how the invocation was served: `stream`, `oneshot`, or `realtime`.

A scope can also use `tags.<key>` for the value of one tag.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostBudgetScopeFieldSubjectType

// Open enum: custom values can be created with a direct type cast
custom := components.AICostBudgetScopeField("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AICostBudgetScopeFieldSubjectType` | subject_type                        |
| `AICostBudgetScopeFieldSubject`     | subject                             |
| `AICostBudgetScopeFieldSubjectID`   | subject_id                          |
| `AICostBudgetScopeFieldGroup`       | group                               |
| `AICostBudgetScopeFieldTeam`        | team                                |
| `AICostBudgetScopeFieldModel`       | model                               |
| `AICostBudgetScopeFieldProvider`    | provider                            |
| `AICostBudgetScopeFieldRequestMode` | request_mode                        |