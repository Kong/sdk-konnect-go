# UpsertAppCustomerDataRequest

AppCustomerData upsert request.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Stripe`                                                                      | [*components.Stripe](../../models/components/stripe.md)                       | :heavy_minus_sign:                                                            | Used if the customer has a linked Stripe app.                                 |
| `ExternalInvoicing`                                                           | [*components.ExternalInvoicing](../../models/components/externalinvoicing.md) | :heavy_minus_sign:                                                            | Used if the customer has a linked external invoicing app.                     |