# CustomDomainType

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Type of gateway the custom domain belongs to: `api` for an API Gateway or
`ai` for an AI Gateway. Applies only to dedicated custom domains. Defaults to `api`
when omitted.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomDomainTypeAPI

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDomainType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CustomDomainTypeAPI` | api                   |
| `CustomDomainTypeAi`  | ai                    |