# SchemaRegistryConfluentConfigSensitiveDataAwareSchemaType

The format of the message.

**`any` requires a minimum runtime version of `1.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SchemaRegistryConfluentConfigSensitiveDataAwareSchemaTypeAvro

// Open enum: custom values can be created with a direct type cast
custom := components.SchemaRegistryConfluentConfigSensitiveDataAwareSchemaType("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `SchemaRegistryConfluentConfigSensitiveDataAwareSchemaTypeAvro` | avro                                                            |
| `SchemaRegistryConfluentConfigSensitiveDataAwareSchemaTypeJSON` | json                                                            |
| `SchemaRegistryConfluentConfigSensitiveDataAwareSchemaTypeAny`  | any                                                             |