# CustomFormType

The form's purpose. Determines built-in field requirements and where the form is consumed by the portal client.

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