# SystemSuggestionRuleType

The type of action this system suggestion rule performs

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SystemSuggestionRuleTypeArchive

// Open enum: custom values can be created with a direct type cast
custom := components.SystemSuggestionRuleType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `SystemSuggestionRuleTypeArchive`     | archive                               |
| `SystemSuggestionRuleTypeMap`         | map                                   |
| `SystemSuggestionRuleTypeCreateOrMap` | create-or-map                         |