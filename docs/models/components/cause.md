# Cause

Represents what caused the resource action to occur:
- `user-action`: user performed a resource action
- `suggestion-rule`: resource action performed based on the configuration of a suggestion rule


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CauseUserAction

// Open enum: custom values can be created with a direct type cast
custom := components.Cause("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CauseUserAction`     | user-action           |
| `CauseSuggestionRule` | suggestion-rule       |