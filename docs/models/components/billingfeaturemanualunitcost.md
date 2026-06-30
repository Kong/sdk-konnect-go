# BillingFeatureManualUnitCost

A fixed per-unit cost amount.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                     | [components.BillingFeatureManualUnitCostType](../../models/components/billingfeaturemanualunitcosttype.md) | :heavy_check_mark:                                                                                         | The type discriminator for manual unit cost.                                                               |
| `Amount`                                                                                                   | `string`                                                                                                   | :heavy_check_mark:                                                                                         | Fixed per-unit cost amount in USD.                                                                         |