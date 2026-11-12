# ServicePeriod

The service period covered by this invoice.

For flat fee the service period can be empty which means `from` will be equals
to `to`. In other cases those fields will be filled with the actual service
period.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `From`                                                          | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | The start of the period.<br/><br/>The period is inclusive at the start. | 2023-01-01T01:01:01.001Z                                        |
| `To`                                                            | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | The end of the period.<br/><br/>The period is exclusive at the end. | 2023-01-01T01:01:01.001Z                                        |