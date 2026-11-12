# Rule

invalid parameters rules

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RuleMinLength

// Open enum: custom values can be created with a direct type cast
custom := components.Rule("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `RuleMinLength`    | min_length         |
| `RuleMinDigits`    | min_digits         |
| `RuleMinLowercase` | min_lowercase      |
| `RuleMinUppercase` | min_uppercase      |
| `RuleMinSymbols`   | min_symbols        |
| `RuleMinItems`     | min_items          |
| `RuleMin`          | min                |