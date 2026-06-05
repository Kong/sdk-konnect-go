# BillingPriceUnit

Unit price.

Charges a fixed rate per billing unit. When UnitConfig is present on the rate
card, billing units are the converted quantities (e.g. GB instead of bytes).


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Type`                                                                             | [components.BillingPriceUnitType](../../models/components/billingpriceunittype.md) | :heavy_check_mark:                                                                 | The type of the price.                                                             |
| `Amount`                                                                           | `string`                                                                           | :heavy_check_mark:                                                                 | The amount of the unit price.                                                      |