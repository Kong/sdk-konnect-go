# AICostObservedSubjectType

The type of a subject reference, the part before the `/`:

- `principal`: a Kong Identity principal.
- `consumer`: a gateway consumer or an AI Gateway consumer.
- `oidc`: an OIDC subject, when only an OIDC credential identifies the caller.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostObservedSubjectTypePrincipal

// Open enum: custom values can be created with a direct type cast
custom := components.AICostObservedSubjectType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AICostObservedSubjectTypePrincipal` | principal                            |
| `AICostObservedSubjectTypeConsumer`  | consumer                             |
| `AICostObservedSubjectTypeOidc`      | oidc                                 |