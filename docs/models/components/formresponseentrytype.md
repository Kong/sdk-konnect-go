# FormResponseEntryType

Field input type at submission time.

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