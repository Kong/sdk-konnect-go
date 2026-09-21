# FormResponseEntryType

The field's type at the time it was submitted.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.FormResponseEntryTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.FormResponseEntryType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `FormResponseEntryTypeText`     | text                            |
| `FormResponseEntryTypeEmail`    | email                           |
| `FormResponseEntryTypeNumber`   | number                          |
| `FormResponseEntryTypeTextarea` | textarea                        |
| `FormResponseEntryTypeSelect`   | select                          |
| `FormResponseEntryTypeCheckbox` | checkbox                        |