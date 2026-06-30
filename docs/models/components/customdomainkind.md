# CustomDomainKind

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Kind of the custom domain based on Cloud Gateway deployment.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomDomainKindDedicatedV0

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDomainKind("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CustomDomainKindDedicatedV0`  | dedicated.v0                   |
| `CustomDomainKindServerlessV1` | serverless.v1                  |