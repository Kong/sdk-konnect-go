# ExternalReferences

External identifiers assigned to this invoice by third-party systems.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `InvoicingID`                                                                | `*string`                                                                    | :heavy_minus_sign:                                                           | The ID assigned by the external invoicing app (e.g. Stripe invoice ID).      |
| `PaymentID`                                                                  | `*string`                                                                    | :heavy_minus_sign:                                                           | The ID assigned by the external payment app (e.g. Stripe payment intent ID). |