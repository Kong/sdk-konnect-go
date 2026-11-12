# MeOrganizationState

State of the organization

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeOrganizationStateActive

// Open enum: custom values can be created with a direct type cast
custom := components.MeOrganizationState("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MeOrganizationStateActive`   | active                        |
| `MeOrganizationStateInactive` | inactive                      |
| `MeOrganizationStateDeleting` | deleting                      |
| `MeOrganizationStateDeleted`  | deleted                       |