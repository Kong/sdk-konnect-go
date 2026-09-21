# ControlPlaneAddOnOwnerType

Type of gateway that owns the add-on: `api` for an API Gateway or `ai` for an
AI Gateway. Defaults to `api` when omitted.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ControlPlaneAddOnOwnerTypeAPI

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneAddOnOwnerType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ControlPlaneAddOnOwnerTypeAPI` | api                             |
| `ControlPlaneAddOnOwnerTypeAi`  | ai                              |