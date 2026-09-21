# BillingInvoiceStandardLineServicePeriod

The service period covered by this invoice, spanning the earliest line start to
the latest line end across all of its lines.

For an invoice with no lines the period is empty, which means `from` will be
equal to `to`.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `From`                                                          | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | The start of the period.<br/><br/>The period is inclusive at the start. | 2023-01-01T01:01:01.001Z                                        |
| `To`                                                            | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | The end of the period.<br/><br/>The period is exclusive at the end. | 2023-01-01T01:01:01.001Z                                        |