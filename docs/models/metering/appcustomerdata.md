# AppCustomerData

App customer data.


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `Stripe`                                                                                                        | [*metering.BillingCustomerDataStripe](../../models/metering/billingcustomerdatastripe.md)                       | :heavy_minus_sign:                                                                                              | Used if the customer has a linked Stripe app.                                                                   |
| `ExternalInvoicing`                                                                                             | [*metering.BillingCustomerDataExternalInvoicing](../../models/metering/billingcustomerdataexternalinvoicing.md) | :heavy_minus_sign:                                                                                              | Used if the customer has a linked external invoicing app.                                                       |