# Apps

The applications used by this billing profile.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Tax`                                                                                | [components.Tax](../../models/components/tax.md)                                     | :heavy_check_mark:                                                                   | The tax app used for this workflow.                                                  |
| `Invoicing`                                                                          | [components.Invoicing](../../models/components/invoicing.md)                         | :heavy_check_mark:                                                                   | The invoicing app used for this workflow.                                            |
| `Payment`                                                                            | [components.BillingProfilePayment](../../models/components/billingprofilepayment.md) | :heavy_check_mark:                                                                   | The payment app used for this workflow.                                              |