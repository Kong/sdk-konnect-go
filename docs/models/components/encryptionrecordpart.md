# EncryptionRecordPart

* key - encrypt the record key
* value - encrypt the record value


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EncryptionRecordPartKey

// Open enum: custom values can be created with a direct type cast
custom := components.EncryptionRecordPart("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `EncryptionRecordPartKey`   | key                         |
| `EncryptionRecordPartValue` | value                       |