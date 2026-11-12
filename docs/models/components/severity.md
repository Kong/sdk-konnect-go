# Severity

Severity of the validation issue.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SeverityCritical

// Open enum: custom values can be created with a direct type cast
custom := components.Severity("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SeverityCritical` | critical           |
| `SeverityWarning`  | warning            |