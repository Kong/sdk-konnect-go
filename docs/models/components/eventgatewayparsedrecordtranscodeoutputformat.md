# EventGatewayParsedRecordTranscodeOutputFormat

The serialization format to convert the record value into.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayParsedRecordTranscodeOutputFormatJSON

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayParsedRecordTranscodeOutputFormat("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `EventGatewayParsedRecordTranscodeOutputFormatJSON` | json                                                |
| `EventGatewayParsedRecordTranscodeOutputFormatAvro` | avro                                                |