# Apps

The applications used by this billing profile.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Tax`                                                                            | [metering.Tax](../../models/metering/tax.md)                                     | :heavy_check_mark:                                                               | The tax app used for this workflow.                                              |
| `Invoicing`                                                                      | [metering.Invoicing](../../models/metering/invoicing.md)                         | :heavy_check_mark:                                                               | The invoicing app used for this workflow.                                        |
| `Payment`                                                                        | [metering.BillingProfilePayment](../../models/metering/billingprofilepayment.md) | :heavy_check_mark:                                                               | The payment app used for this workflow.                                          |