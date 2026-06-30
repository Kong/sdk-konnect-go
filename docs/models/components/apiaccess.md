# APIAccess

Controls how data planes in a configuration are exposed. Supported values:
- `private` — data planes are accessible only within the VPC network; no public internet exposure
- `public` — data planes are accessible from the public internet
- `private+public` — equivalent to `public`; data planes are accessible from the public internet (default)


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIAccessPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.APIAccess("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `APIAccessPrivate`           | private                      |
| `APIAccessPublic`            | public                       |
| `APIAccessPrivatePlusPublic` | private+public               |