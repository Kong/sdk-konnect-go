# APIKeyCredentialListItemStatus

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIKeyCredentialListItemStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.APIKeyCredentialListItemStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `APIKeyCredentialListItemStatusActive`  | active                                  |
| `APIKeyCredentialListItemStatusExpired` | expired                                 |
| `APIKeyCredentialListItemStatusRevoked` | revoked                                 |