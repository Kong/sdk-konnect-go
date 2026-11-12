# PortalCustomDomainVerificationStatus

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PortalCustomDomainVerificationStatusVerified

// Open enum: custom values can be created with a direct type cast
custom := components.PortalCustomDomainVerificationStatus("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `PortalCustomDomainVerificationStatusVerified` | verified                                       |
| `PortalCustomDomainVerificationStatusPending`  | pending                                        |
| `PortalCustomDomainVerificationStatusError`    | error                                          |