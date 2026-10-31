# CreateDeveloperRequestDeveloperStatus

When omitted, the portal's `auto_approve_developers` setting determines the initial status. Set this property to override the portal default with any valid developer status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateDeveloperRequestDeveloperStatusApproved

// Open enum: custom values can be created with a direct type cast
custom := components.CreateDeveloperRequestDeveloperStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `CreateDeveloperRequestDeveloperStatusApproved` | approved                                        |
| `CreateDeveloperRequestDeveloperStatusPending`  | pending                                         |
| `CreateDeveloperRequestDeveloperStatusRevoked`  | revoked                                         |
| `CreateDeveloperRequestDeveloperStatusRejected` | rejected                                        |