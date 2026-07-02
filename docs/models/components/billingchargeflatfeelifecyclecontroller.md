# BillingChargeFlatFeeLifecycleController

Indicates whether the charge lifecycle is controlled by OpenMeter or manually
overridden by the API user.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeFlatFeeLifecycleControllerSystem

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeFlatFeeLifecycleController("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BillingChargeFlatFeeLifecycleControllerSystem` | system                                          |
| `BillingChargeFlatFeeLifecycleControllerManual` | manual                                          |