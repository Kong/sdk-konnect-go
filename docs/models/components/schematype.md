# SchemaType

The format of the message.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SchemaTypeAvro

// Open enum: custom values can be created with a direct type cast
custom := components.SchemaType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `SchemaTypeAvro` | avro             |
| `SchemaTypeJSON` | json             |