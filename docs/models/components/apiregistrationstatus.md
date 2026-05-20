# APIRegistrationStatus

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIRegistrationStatusApproved

// Open enum: custom values can be created with a direct type cast
custom := components.APIRegistrationStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `APIRegistrationStatusApproved` | approved                        |
| `APIRegistrationStatusPending`  | pending                         |
| `APIRegistrationStatusRejected` | rejected                        |
| `APIRegistrationStatusRevoked`  | revoked                         |