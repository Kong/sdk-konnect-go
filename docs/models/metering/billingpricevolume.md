# BillingPriceVolume

Volume tiered price.

The maximum quantity within a period determines the per-unit price for all units
in that period.

When UnitConfig is present on the rate card, tier boundaries (up_to_amount) are
expressed in converted billing units.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Type`                                                                             | [metering.BillingPriceVolumeType](../../models/metering/billingpricevolumetype.md) | :heavy_check_mark:                                                                 | The type of the price.                                                             |
| `Tiers`                                                                            | [][metering.BillingPriceTier](../../models/metering/billingpricetier.md)           | :heavy_check_mark:                                                                 | The tiers of the volume price. At least one tier is required.                      |