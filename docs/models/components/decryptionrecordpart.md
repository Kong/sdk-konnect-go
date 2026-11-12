# DecryptionRecordPart

* key - decrypt the record key
* value - decrypt the record value


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DecryptionRecordPartKey

// Open enum: custom values can be created with a direct type cast
custom := components.DecryptionRecordPart("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `DecryptionRecordPartKey`   | key                         |
| `DecryptionRecordPartValue` | value                       |