# UnitCostBillingFeatureManualUnitCost

A fixed per-unit cost amount.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Type`                                                         | [metering.UnitCostType](../../models/metering/unitcosttype.md) | :heavy_check_mark:                                             | The type discriminator for manual unit cost.                   |
| `Amount`                                                       | `string`                                                       | :heavy_check_mark:                                             | Fixed per-unit cost amount in USD.                             |