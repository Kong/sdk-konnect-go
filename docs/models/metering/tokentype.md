# TokenType

Static token type value. Use this when the feature tracks a single token type
(e.g., only input tokens). `request` is an alias for `input`, `response` is an
alias for `output`. Mutually exclusive with `token_type_property`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.TokenTypeInput

// Open enum: custom values can be created with a direct type cast
custom := metering.TokenType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `TokenTypeInput`      | input                 |
| `TokenTypeOutput`     | output                |
| `TokenTypeCacheRead`  | cache_read            |
| `TokenTypeCacheWrite` | cache_write           |
| `TokenTypeReasoning`  | reasoning             |
| `TokenTypeRequest`    | request               |
| `TokenTypeResponse`   | response              |