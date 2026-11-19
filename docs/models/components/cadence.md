# Cadence

The length of each period, which sets the range of each export.

A replacement can't change it.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CadenceDaily

// Open enum: custom values can be created with a direct type cast
custom := components.Cadence("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `CadenceDaily`   | daily            |
| `CadenceWeekly`  | weekly           |
| `CadenceMonthly` | monthly          |