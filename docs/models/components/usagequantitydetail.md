# UsageQuantityDetail

Usage quantity details for this line when UnitConfig is in effect.

Read-only; omitted for lines without unit conversion.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `RawQuantity`                                                                                     | `string`                                                                                          | :heavy_check_mark:                                                                                | The raw quantity as reported by the meter (native units).                                         |
| `InvoicedQuantity`                                                                                | `string`                                                                                          | :heavy_check_mark:                                                                                | The net billed quantity for this line in converted units, after rounding and any<br/>usage discounts. |
| `DisplayUnit`                                                                                     | `*string`                                                                                         | :heavy_minus_sign:                                                                                | The display unit label (e.g., "GB", "hours", "M tokens").                                         |