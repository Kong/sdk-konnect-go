# TransitGatewayState

The current state of the Transit Gateway. Possible values:
- `created` - The attachment has been created but is not attached to transit gateway.
- `initializing` - The attachment is in the process of being initialized and is setting up necessary resources.
- `pending-acceptance` The attachment request is awaiting acceptance in customer VPC.
- `ready` - The transit gateway attachment is fully operational and can route traffic as configured.
- `terminating` - The attachment is in the process of being deleted and is no longer accepting new traffic.
- `terminated` - The attachment has been fully deleted and is no longer available.



## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `TransitGatewayStateCreated`           | created                                |
| `TransitGatewayStateInitializing`      | initializing                           |
| `TransitGatewayStatePendingAcceptance` | pending-acceptance                     |
| `TransitGatewayStateReady`             | ready                                  |
| `TransitGatewayStateTerminating`       | terminating                            |
| `TransitGatewayStateTerminated`        | terminated                             |