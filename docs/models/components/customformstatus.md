# CustomFormStatus

Whether the form is visible to developers on the portal. `unpublished` forms aren't shown to developers; for the `developer_registration` form, a default form is shown instead.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomFormStatusPublished

// Open enum: custom values can be created with a direct type cast
custom := components.CustomFormStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CustomFormStatusPublished`   | published                     |
| `CustomFormStatusUnpublished` | unpublished                   |