# ResourceQuotaQualifier

Enumeration of resources available for quota enforcement.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ResourceQuotaQualifierCountProviderAccountsPerProvider

// Open enum: custom values can be created with a direct type cast
custom := components.ResourceQuotaQualifier("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `ResourceQuotaQualifierCountProviderAccountsPerProvider`  | count/provider-accounts.per-provider                      |
| `ResourceQuotaQualifierCountNetworksNotOffline`           | count/networks.not-offline                                |
| `ResourceQuotaQualifierCountDataPlanesEstimate`           | count/data-planes-estimate                                |
| `ResourceQuotaQualifierCountServerlessDataPlanesEstimate` | count/serverless-data-planes-estimate                     |