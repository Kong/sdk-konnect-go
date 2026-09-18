# ExternalInvoicing

Used if the customer has a linked external invoicing app.

Set to `null` to remove the external invoicing data of the customer; omit to
leave it unchanged.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Labels`                                                        | map[string]`string`                                             | :heavy_minus_sign:                                              | Labels for this external invoicing integration on the customer. | {<br/>"env": "test"<br/>}                                       |