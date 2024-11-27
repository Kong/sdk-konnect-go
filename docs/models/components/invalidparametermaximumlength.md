# InvalidParameterMaximumLength


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Field`                                                                                                      | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          | name                                                                                                         |
| `Rule`                                                                                                       | [components.InvalidParameterMaximumLengthRule](../../models/components/invalidparametermaximumlengthrule.md) | :heavy_check_mark:                                                                                           | invalid parameters rules                                                                                     |                                                                                                              |
| `Maximum`                                                                                                    | *int64*                                                                                                      | :heavy_check_mark:                                                                                           | N/A                                                                                                          | 8                                                                                                            |
| `Source`                                                                                                     | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | N/A                                                                                                          | body                                                                                                         |
| `Reason`                                                                                                     | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          | must not have more than 8 characters                                                                         |