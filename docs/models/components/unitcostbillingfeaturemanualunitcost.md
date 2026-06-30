# UnitCostBillingFeatureManualUnitCost

A fixed per-unit cost amount.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | [components.UnitCostType](../../models/components/unitcosttype.md) | :heavy_check_mark:                                                 | The type discriminator for manual unit cost.                       |
| `Amount`                                                           | `string`                                                           | :heavy_check_mark:                                                 | Fixed per-unit cost amount in USD.                                 |