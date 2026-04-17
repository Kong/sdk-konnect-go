# Subfield

Specify which subfield of the `url` custom field value to use for selection.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubfieldName

// Open enum: custom values can be created with a direct type cast
custom := components.Subfield("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `SubfieldName` | name           |
| `SubfieldLink` | link           |