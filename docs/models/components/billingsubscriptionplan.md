# BillingSubscriptionPlan

The plan the subscription was created from, if any. Includes the plan key and
version so clients can resolve the exact plan revision.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ID`                                               | `string`                                           | :heavy_check_mark:                                 | The plan ID (exact revision).                      | 01G65Z755AFWAKHE12NY0CQ9FH                         |
| `Key`                                              | `string`                                           | :heavy_check_mark:                                 | The plan key. References the plan across versions. | resource_key                                       |
| `Version`                                          | `int64`                                            | :heavy_check_mark:                                 | The plan version.                                  |                                                    |