# CustomerUpdate

Controls which customer fields can be updated by the checkout session.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Address`                                                                        | [*metering.Address](../../models/metering/address.md)                            | :heavy_minus_sign:                                                               | Whether to save the billing address to customer.address.<br/><br/>Defaults to "never". |
| `Name`                                                                           | [*metering.Name](../../models/metering/name.md)                                  | :heavy_minus_sign:                                                               | Whether to save the customer name to customer.name.<br/><br/>Defaults to "never". |
| `Shipping`                                                                       | [*metering.Shipping](../../models/metering/shipping.md)                          | :heavy_minus_sign:                                                               | Whether to save shipping information to customer.shipping.<br/><br/>Defaults to "never". |