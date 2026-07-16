# Dimensions

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DimensionsAPI

// Open enum: custom values can be created with a direct type cast
custom := components.Dimensions("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `DimensionsAPI`                       | api                                   |
| `DimensionsAPIPackage`                | api_package                           |
| `DimensionsAPIProduct`                | api_product                           |
| `DimensionsAPIProductVersion`         | api_product_version                   |
| `DimensionsApplication`               | application                           |
| `DimensionsConsumer`                  | consumer                              |
| `DimensionsControlPlane`              | control_plane                         |
| `DimensionsControlPlaneGroup`         | control_plane_group                   |
| `DimensionsCountryCode`               | country_code                          |
| `DimensionsDataPlaneNode`             | data_plane_node                       |
| `DimensionsDataPlaneNodeVersion`      | data_plane_node_version               |
| `DimensionsGatewayService`            | gateway_service                       |
| `DimensionsPortal`                    | portal                                |
| `DimensionsPrincipal`                 | principal                             |
| `DimensionsRealm`                     | realm                                 |
| `DimensionsResponseSource`            | response_source                       |
| `DimensionsRoute`                     | route                                 |
| `DimensionsStatusCode`                | status_code                           |
| `DimensionsStatusCodeGrouped`         | status_code_grouped                   |
| `DimensionsTime`                      | time                                  |
| `DimensionsUpstreamStatusCode`        | upstream_status_code                  |
| `DimensionsUpstreamStatusCodeGrouped` | upstream_status_code_grouped          |