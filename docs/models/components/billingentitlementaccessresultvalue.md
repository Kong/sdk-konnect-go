# BillingEntitlementAccessResultValue

Only available for metered entitlements. The current balance details of the
entitlement. Requires the `value` expand.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Balance`                                                              | `string`                                                               | :heavy_check_mark:                                                     | The remaining balance of the entitlement in the current usage period.  |
| `GrantBalances`                                                        | map[string]`string`                                                    | :heavy_check_mark:                                                     | The remaining balance of each grant, keyed by grant ID.                |
| `Overage`                                                              | `string`                                                               | :heavy_check_mark:                                                     | The usage exceeding the available balance in the current usage period. |
| `TotalAvailableGrantAmount`                                            | `string`                                                               | :heavy_check_mark:                                                     | The total amount granted and currently available to the entitlement.   |
| `Usage`                                                                | `string`                                                               | :heavy_check_mark:                                                     | The usage recorded in the current usage period.                        |