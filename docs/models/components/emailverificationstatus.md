# EmailVerificationStatus

Specifies the status of the mail domain verification

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EmailVerificationStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.EmailVerificationStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `EmailVerificationStatusPending`          | pending                                   |
| `EmailVerificationStatusSuccess`          | success                                   |
| `EmailVerificationStatusFailed`           | failed                                    |
| `EmailVerificationStatusTemporaryFailure` | temporary_failure                         |
| `EmailVerificationStatusNotStarted`       | not_started                               |