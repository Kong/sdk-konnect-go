# TaxConfig

The tax config of the rate card.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Behavior`                                                                                                   | [*metering.BillingTaxBehavior](../../models/metering/billingtaxbehavior.md)                                  | :heavy_minus_sign:                                                                                           | Tax behavior.<br/><br/>This enum is used to specify whether tax is included in the price or excluded<br/>from the price. |
| `Code`                                                                                                       | [metering.TaxCodeReferenceItem](../../models/metering/taxcodereferenceitem.md)                               | :heavy_check_mark:                                                                                           | TaxCode reference.                                                                                           |