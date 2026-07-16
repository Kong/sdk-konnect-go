# CreateApplicationRegistrationRequestStatus

The status of the application registration. It must be a valid status value for a registration creation.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateApplicationRegistrationRequestStatusApproved

// Open enum: custom values can be created with a direct type cast
custom := components.CreateApplicationRegistrationRequestStatus("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CreateApplicationRegistrationRequestStatusApproved` | approved                                             |
| `CreateApplicationRegistrationRequestStatusPending`  | pending                                              |