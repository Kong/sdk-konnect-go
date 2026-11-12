# SchemaType

The format of the message.

**`any` requires a minimum runtime version of `1.3`**.

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
| `SchemaTypeAny`  | any              |