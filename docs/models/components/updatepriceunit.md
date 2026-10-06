# UpdatePriceUnit

Unit price.

Charges a fixed rate per billing unit. When UnitConfig is present on the object,
billing units are the converted quantities (e.g. GB instead of bytes).


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Type`                                                                           | [components.UpdatePriceUnitType](../../models/components/updatepriceunittype.md) | :heavy_check_mark:                                                               | The type of the price.                                                           |
| `Amount`                                                                         | `string`                                                                         | :heavy_check_mark:                                                               | The amount of the unit price.                                                    |