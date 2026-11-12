# ProviderName

Name of cloud provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProviderNameAws

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderName("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ProviderNameAws`   | aws                 |
| `ProviderNameAzure` | azure               |
| `ProviderNameGcp`   | gcp                 |