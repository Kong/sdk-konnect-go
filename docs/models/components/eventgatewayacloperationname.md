# EventGatewayACLOperationName

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EventGatewayACLOperationNameAll

// Open enum: custom values can be created with a direct type cast
custom := components.EventGatewayACLOperationName("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `EventGatewayACLOperationNameAll`             | all                                           |
| `EventGatewayACLOperationNameAlter`           | alter                                         |
| `EventGatewayACLOperationNameAlterConfigs`    | alter_configs                                 |
| `EventGatewayACLOperationNameCreate`          | create                                        |
| `EventGatewayACLOperationNameDelete`          | delete                                        |
| `EventGatewayACLOperationNameDescribe`        | describe                                      |
| `EventGatewayACLOperationNameDescribeConfigs` | describe_configs                              |
| `EventGatewayACLOperationNameIdempotentWrite` | idempotent_write                              |
| `EventGatewayACLOperationNameRead`            | read                                          |
| `EventGatewayACLOperationNameWrite`           | write                                         |