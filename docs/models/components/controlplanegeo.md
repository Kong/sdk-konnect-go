# ControlPlaneGeo

Geographic region of the control plane. Supported values:
- `us` — United States
- `eu` — Europe
- `au` — Australia
- `me` — Middle East
- `in` — India
- `sg` — Singapore


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ControlPlaneGeoUs

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneGeo("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ControlPlaneGeoUs` | us                  |
| `ControlPlaneGeoEu` | eu                  |
| `ControlPlaneGeoAu` | au                  |
| `ControlPlaneGeoMe` | me                  |
| `ControlPlaneGeoIn` | in                  |
| `ControlPlaneGeoSg` | sg                  |