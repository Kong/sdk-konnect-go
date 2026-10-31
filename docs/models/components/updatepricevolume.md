# UpdatePriceVolume

Volume tiered price.

The maximum quantity within a period determines the per-unit price for all units
in that period.

When UnitConfig is present on the containing resource, tier boundaries
(up_to_amount) are expressed in converted billing units.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Type`                                                                               | [components.UpdatePriceVolumeType](../../models/components/updatepricevolumetype.md) | :heavy_check_mark:                                                                   | The type of the price.                                                               |
| `Tiers`                                                                              | [][components.UpdatePriceTier](../../models/components/updatepricetier.md)           | :heavy_check_mark:                                                                   | The tiers of the volume price. At least one tier is required.                        |