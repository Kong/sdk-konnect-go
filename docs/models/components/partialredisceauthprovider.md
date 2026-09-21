# PartialRedisCeAuthProvider

Auth providers to be used to authenticate to a Cloud Provider's Redis instance.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisCeAuthProviderAws

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisCeAuthProvider("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PartialRedisCeAuthProviderAws`   | aws                               |
| `PartialRedisCeAuthProviderAzure` | azure                             |
| `PartialRedisCeAuthProviderGcp`   | gcp                               |