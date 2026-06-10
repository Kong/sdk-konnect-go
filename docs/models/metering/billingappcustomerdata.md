# BillingAppCustomerData

App customer data.


## Fields

| Field                                                                                                                 | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `Stripe`                                                                                                              | [*metering.BillingAppCustomerDataStripe](../../models/metering/billingappcustomerdatastripe.md)                       | :heavy_minus_sign:                                                                                                    | Used if the customer has a linked Stripe app.                                                                         |
| `ExternalInvoicing`                                                                                                   | [*metering.BillingAppCustomerDataExternalInvoicing](../../models/metering/billingappcustomerdataexternalinvoicing.md) | :heavy_minus_sign:                                                                                                    | Used if the customer has a linked external invoicing app.                                                             |