# BillingChargeRealizationDetailedLineCreditApplied

A credit allocation applied to a charge realization detailed line.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Amount`                                                                    | `string`                                                                    | :heavy_check_mark:                                                          | The monetary amount credited.                                               |
| `Description`                                                               | `*string`                                                                   | :heavy_minus_sign:                                                          | Human-readable description of the credit allocation.                        |
| `CreditRealizationID`                                                       | `string`                                                                    | :heavy_check_mark:                                                          | The ID of the credit realization (allocation) this credit was applied from. |