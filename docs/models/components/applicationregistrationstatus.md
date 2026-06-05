# ApplicationRegistrationStatus

The status of an application registration request. Each registration is linked to a single API, and application credentials will not grant access to the API until the registration is approved.
Pending, revoked, and rejected registrations will not provide access to the API.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ApplicationRegistrationStatusApproved

// Open enum: custom values can be created with a direct type cast
custom := components.ApplicationRegistrationStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ApplicationRegistrationStatusApproved` | approved                                |
| `ApplicationRegistrationStatusPending`  | pending                                 |
| `ApplicationRegistrationStatusRevoked`  | revoked                                 |
| `ApplicationRegistrationStatusRejected` | rejected                                |