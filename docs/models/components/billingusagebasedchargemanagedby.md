# BillingUsageBasedChargeManagedBy

The charge is managed by the following entity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingUsageBasedChargeManagedByManual

// Open enum: custom values can be created with a direct type cast
custom := components.BillingUsageBasedChargeManagedBy("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BillingUsageBasedChargeManagedByManual`       | manual                                         |
| `BillingUsageBasedChargeManagedBySystem`       | system                                         |
| `BillingUsageBasedChargeManagedBySubscription` | subscription                                   |