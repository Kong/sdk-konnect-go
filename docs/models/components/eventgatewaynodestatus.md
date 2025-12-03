# EventGatewayNodeStatus

The status of an event gateway node.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ConfigAppliedAt`                                            | [time.Time](https://pkg.go.dev/time#Time)                    | :heavy_check_mark:                                           | The time the node succeeds in applying the configuration.    |                                                              |
| `ConfigVersion`                                              | *string*                                                     | :heavy_check_mark:                                           | The version number of the configuration applied by the node. | v1.123                                                       |