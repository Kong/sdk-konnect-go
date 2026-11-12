# CustomFormType

The kind of form. Determines which fields are required and where developers see the form on the portal.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomFormTypeDeveloperRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.CustomFormType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CustomFormTypeDeveloperRegistration` | developer_registration                |
| `CustomFormTypeAPIRegistration`       | api_registration                      |