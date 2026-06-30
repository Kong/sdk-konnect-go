# PlatformFilter

A filter to apply to a platform usage query.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Field`                                                                                | [components.PlatformFilterField](../../models/components/platformfilterfield.md)       | :heavy_check_mark:                                                                     | The field to filter.                                                                   |
| `Operator`                                                                             | [components.PlatformFilterOperator](../../models/components/platformfilteroperator.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `Value`                                                                                | []`string`                                                                             | :heavy_check_mark:                                                                     | The values to filter by.                                                               |