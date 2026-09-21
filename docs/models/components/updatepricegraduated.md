# UpdatePriceGraduated

Graduated tiered price.

Each tier's rate applies only to the usage within that tier. Pricing can change
as cumulative usage crosses tier boundaries.

When UnitConfig is present on the containing resource, tier boundaries
(up_to_amount) are expressed in converted billing units.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Type`                                                                                     | [components.UpdatePriceGraduatedType](../../models/components/updatepricegraduatedtype.md) | :heavy_check_mark:                                                                         | The type of the price.                                                                     |
| `Tiers`                                                                                    | [][components.UpdatePriceTier](../../models/components/updatepricetier.md)                 | :heavy_check_mark:                                                                         | The tiers of the graduated price. At least one tier is required.                           |