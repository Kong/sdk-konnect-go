# APIPackageOperationResponseSchemaDuration

The timeframe to ensure that an application does not exceed the request limit.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIPackageOperationResponseSchemaDurationSeconds

// Open enum: custom values can be created with a direct type cast
custom := components.APIPackageOperationResponseSchemaDuration("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `APIPackageOperationResponseSchemaDurationSeconds` | seconds                                            |
| `APIPackageOperationResponseSchemaDurationMinutes` | minutes                                            |
| `APIPackageOperationResponseSchemaDurationHours`   | hours                                              |