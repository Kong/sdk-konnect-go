# SpecFormatSchema

The format of the returned `spec.content`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SpecFormatSchemaJSON

// Open enum: custom values can be created with a direct type cast
custom := components.SpecFormatSchema("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `SpecFormatSchemaJSON` | json                   |
| `SpecFormatSchemaYaml` | yaml                   |