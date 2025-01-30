# ConfigurationDataPlaneGroupAutoscaleStatic

Object that describes the static autoscaling strategy.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Kind`                                                                     | [components.Kind](../../models/components/kind.md)                         | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `InstanceType`                                                             | [components.InstanceTypeName](../../models/components/instancetypename.md) | :heavy_check_mark:                                                         | Instance type name to indicate capacity.                                   |                                                                            |
| `RequestedInstances`                                                       | *int64*                                                                    | :heavy_check_mark:                                                         | Number of data-planes the deployment target will contain.                  | 3                                                                          |