# SimpleSchemaDefinition

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SimpleSchemaDefinitionString

// Open enum: custom values can be created with a direct type cast
custom := components.SimpleSchemaDefinition("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `SimpleSchemaDefinitionString`  | string                          |
| `SimpleSchemaDefinitionNumber`  | number                          |
| `SimpleSchemaDefinitionBoolean` | boolean                         |