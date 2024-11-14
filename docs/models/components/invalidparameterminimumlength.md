# InvalidParameterMinimumLength


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Field`                                            | *string*                                           | :heavy_check_mark:                                 | N/A                                                | name                                               |
| `Rule`                                             | [components.Rule](../../models/components/rule.md) | :heavy_check_mark:                                 | invalid parameters rules                           |                                                    |
| `Minimum`                                          | *int64*                                            | :heavy_check_mark:                                 | N/A                                                | 8                                                  |
| `Source`                                           | **string*                                          | :heavy_minus_sign:                                 | N/A                                                | body                                               |
| `Reason`                                           | *string*                                           | :heavy_check_mark:                                 | N/A                                                | must have at least 8 characters                    |