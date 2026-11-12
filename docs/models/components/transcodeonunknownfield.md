# TranscodeOnUnknownField

Describes how to handle JSON fields that are not defined in the target schema.
* `fail` - the transcoding fails. The policy `failure_mode` determines how the failure is handled.
* `drop` - the unknown fields are removed from the transcoded value.

**Requires a minimum runtime version of `1.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TranscodeOnUnknownFieldFail

// Open enum: custom values can be created with a direct type cast
custom := components.TranscodeOnUnknownField("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `TranscodeOnUnknownFieldFail` | fail                          |
| `TranscodeOnUnknownFieldDrop` | drop                          |