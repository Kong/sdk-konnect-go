# BillingPriceGraduated

Graduated tiered price.

Each tier's rate applies only to the usage within that tier. Pricing can change
as cumulative usage crosses tier boundaries.

When UnitConfig is present on the rate card, tier boundaries (up_to_amount) are
expressed in converted billing units.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Type`                                                                                       | [components.BillingPriceGraduatedType](../../models/components/billingpricegraduatedtype.md) | :heavy_check_mark:                                                                           | The type of the price.                                                                       |
| `Tiers`                                                                                      | [][components.BillingPriceTier](../../models/components/billingpricetier.md)                 | :heavy_check_mark:                                                                           | The tiers of the graduated price. At least one tier is required.                             |