# CustomFormStatus

Whether the form is live on the portal. `unpublished` forms are not served from the portal-client form-fetch endpoint; `developer_registration` falls back to a default schema in that case.


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