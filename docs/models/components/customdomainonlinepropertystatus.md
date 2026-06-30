# CustomDomainOnlinePropertyStatus

Set of available statuses for the online properties of a custom domain.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomDomainOnlinePropertyStatusVerified

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDomainOnlinePropertyStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CustomDomainOnlinePropertyStatusVerified`   | verified                                     |
| `CustomDomainOnlinePropertyStatusUnverified` | unverified                                   |