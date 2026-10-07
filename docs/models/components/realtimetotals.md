# RealtimeTotals

The realtime amount of the charge, i.e. the whole usage rated at the charge's
price for its full service period, ignoring what has already been booked to a
realization. This differs from the `usage` of a realization with type
`outstanding`, which only covers the quantity not yet booked.

Requires the `real_time_usage` expand.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Amount`                                                                      | `string`                                                                      | :heavy_check_mark:                                                            | The total value of the resource before taxes, discounts and commitments.      |
| `TaxesTotal`                                                                  | `string`                                                                      | :heavy_check_mark:                                                            | The total tax amount applied to the resource.                                 |
| `TaxesInclusiveTotal`                                                         | `string`                                                                      | :heavy_check_mark:                                                            | The total tax amount already included in the resource amount.                 |
| `TaxesExclusiveTotal`                                                         | `string`                                                                      | :heavy_check_mark:                                                            | The total tax amount added on top of the resource amount.                     |
| `ChargesTotal`                                                                | `string`                                                                      | :heavy_check_mark:                                                            | The total amount contributed by additional charges.                           |
| `DiscountsTotal`                                                              | `string`                                                                      | :heavy_check_mark:                                                            | The total amount deducted through discounts.                                  |
| `CreditsTotal`                                                                | `string`                                                                      | :heavy_check_mark:                                                            | The total amount deducted through credits before taxes are applied.           |
| `Total`                                                                       | `string`                                                                      | :heavy_check_mark:                                                            | The final total value of the resource after taxes, discounts and commitments. |