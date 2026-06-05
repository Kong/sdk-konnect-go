# MeterQueryResult

Meter query result.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `From`                                                                 | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | The start of the period the usage is queried from.                     | 2023-01-01T01:01:01.001Z                                               |
| `To`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | The end of the period the usage is queried to.                         | 2023-01-01T01:01:01.001Z                                               |
| `Data`                                                                 | [][components.MeterQueryRow](../../models/components/meterqueryrow.md) | :heavy_check_mark:                                                     | The usage data. If no data is available, an empty array is returned.   |                                                                        |