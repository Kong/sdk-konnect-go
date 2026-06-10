# UpsertAppCustomerDataRequest

AppCustomerData upsert request.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Stripe`                                                                  | [*metering.Stripe](../../models/metering/stripe.md)                       | :heavy_minus_sign:                                                        | Used if the customer has a linked Stripe app.                             |
| `ExternalInvoicing`                                                       | [*metering.ExternalInvoicing](../../models/metering/externalinvoicing.md) | :heavy_minus_sign:                                                        | Used if the customer has a linked external invoicing app.                 |