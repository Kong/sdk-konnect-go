# SchemaValidationType

How to validate the schema and parse the record.
* confluent_schema_registry - validates against confluent schema registry.
* json - simple JSON parsing without the schema.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SchemaValidationTypeConfluentSchemaRegistry

// Open enum: custom values can be created with a direct type cast
custom := components.SchemaValidationType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `SchemaValidationTypeConfluentSchemaRegistry` | confluent_schema_registry                     |
| `SchemaValidationTypeJSON`                    | json                                          |