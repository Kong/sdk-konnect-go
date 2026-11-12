# ApplicationOwnerType

The type of the owner of the application. Can be `developer` (individual developer) or `team`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ApplicationOwnerTypeDeveloper

// Open enum: custom values can be created with a direct type cast
custom := components.ApplicationOwnerType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ApplicationOwnerTypeDeveloper` | developer                       |
| `ApplicationOwnerTypeTeam`      | team                            |