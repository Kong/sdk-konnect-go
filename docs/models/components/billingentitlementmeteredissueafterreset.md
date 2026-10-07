# BillingEntitlementMeteredIssueAfterReset

Usage granted automatically after each reset. Cannot be combined with `grants`.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Amount`                                                                               | `string`                                                                               | :heavy_check_mark:                                                                     | The amount granted after each reset, in the feature's unit.                            | 100                                                                                    |
| `Priority`                                                                             | `*int64`                                                                               | :heavy_minus_sign:                                                                     | The priority of the grant created after each reset. Lower values have higher<br/>priority. |                                                                                        |