# BillingAppCustomerDataStripe

Used if the customer has a linked Stripe app.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `CustomerID`                                        | `*string`                                           | :heavy_minus_sign:                                  | The Stripe customer ID used.                        | cus_1234567890                                      |
| `DefaultPaymentMethodID`                            | `*string`                                           | :heavy_minus_sign:                                  | The Stripe default payment method ID.               | pm_1234567890                                       |
| `Labels`                                            | map[string]`string`                                 | :heavy_minus_sign:                                  | Labels for this Stripe integration on the customer. | {<br/>"env": "test"<br/>}                           |