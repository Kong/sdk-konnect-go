# SubscriptionAddonRateCardTaxConfig

The tax config of the rate card.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Behavior`                                                                                                   | [*components.BillingTaxBehavior](../../models/components/billingtaxbehavior.md)                              | :heavy_minus_sign:                                                                                           | Tax behavior.<br/><br/>This enum is used to specify whether tax is included in the price or excluded<br/>from the price. |
| `Code`                                                                                                       | [components.TaxCodeReference](../../models/components/taxcodereference.md)                                   | :heavy_check_mark:                                                                                           | TaxCode reference.                                                                                           |