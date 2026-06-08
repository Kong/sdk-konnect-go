# BillingPriceTier

A price tier used in graduated and volume pricing.

At least one price component (flat_price or unit_price) must be set. When
UnitConfig is present on the rate card, up_to_amount is expressed in converted
billing units.


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `UpToAmount`                                                                                                           | `*string`                                                                                                              | :heavy_minus_sign:                                                                                                     | Up to and including this quantity will be contained in the tier. If undefined,<br/>the tier is open-ended (the last tier). |
| `FlatPrice`                                                                                                            | [*components.FlatPriceComponent](../../models/components/flatpricecomponent.md)                                        | :heavy_minus_sign:                                                                                                     | The flat price component of the tier. Charged once when the tier is entered.                                           |
| `UnitPrice`                                                                                                            | [*components.UnitPriceComponent](../../models/components/unitpricecomponent.md)                                        | :heavy_minus_sign:                                                                                                     | The unit price component of the tier. Charged per billing unit within the tier.                                        |