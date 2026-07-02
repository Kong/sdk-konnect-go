# BillingInvoiceStandardLinePriceBillingPriceVolume

Volume tiered price.

The maximum quantity within a period determines the per-unit price for all units
in that period.

When UnitConfig is present on the rate card, tier boundaries (up_to_amount) are
expressed in converted billing units.


## Fields

| Field                                                                                                                                                | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                                                               | [components.BillingPriceVolumePriceBillingInvoiceStandardLineType](../../models/components/billingpricevolumepricebillinginvoicestandardlinetype.md) | :heavy_check_mark:                                                                                                                                   | The type of the price.                                                                                                                               |
| `Tiers`                                                                                                                                              | [][components.BillingPriceTier](../../models/components/billingpricetier.md)                                                                         | :heavy_check_mark:                                                                                                                                   | The tiers of the volume price. At least one tier is required.                                                                                        |