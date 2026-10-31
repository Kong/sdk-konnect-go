# DeveloperStatus

The status of a developer in a portal. Approved developers can log in, create applications, and view and register for products they have access to. Pending, revoked, and rejected developers cannot login or view any non-public portal information, or create or modify applications or registrations.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeveloperStatusApproved

// Open enum: custom values can be created with a direct type cast
custom := components.DeveloperStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DeveloperStatusApproved` | approved                  |
| `DeveloperStatusPending`  | pending                   |
| `DeveloperStatusRevoked`  | revoked                   |
| `DeveloperStatusRejected` | rejected                  |