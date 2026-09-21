# AuthProvider

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthProviderAws

// Open enum: custom values can be created with a direct type cast
custom := components.AuthProvider("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `AuthProviderAws`   | aws                 |
| `AuthProviderGcp`   | gcp                 |
| `AuthProviderAzure` | azure               |