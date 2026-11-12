# BillingEntitlementValueResultValue

Only available for metered entitlements. The balance details of the entitlement
at the evaluation time. Requires the `value` expand.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Balance`                                                                             | `string`                                                                              | :heavy_check_mark:                                                                    | The remaining balance of the entitlement in the usage period at the evaluation<br/>time. |
| `GrantBalances`                                                                       | map[string]`string`                                                                   | :heavy_check_mark:                                                                    | The remaining balance of each grant, keyed by grant ID.                               |
| `Overage`                                                                             | `string`                                                                              | :heavy_check_mark:                                                                    | The usage exceeding the available balance in the usage period at the evaluation<br/>time. |
| `TotalAvailableGrantAmount`                                                           | `string`                                                                              | :heavy_check_mark:                                                                    | The total amount granted and available to the entitlement at the evaluation<br/>time. |
| `Usage`                                                                               | `string`                                                                              | :heavy_check_mark:                                                                    | The usage recorded in the usage period at the evaluation time.                        |