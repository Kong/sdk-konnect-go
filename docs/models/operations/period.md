# Period

The calendar month the forecast covers: the current month in `time_zone`.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `StartsAt`                                                                                   | [time.Time](https://pkg.go.dev/time#Time)                                                    | :heavy_check_mark:                                                                           | The start of the month, as an RFC 3339 timestamp in UTC.                                     |
| `EndsAt`                                                                                     | [time.Time](https://pkg.go.dev/time#Time)                                                    | :heavy_check_mark:                                                                           | The end of the month, as an RFC 3339 timestamp in UTC. The month ends just before this time. |
| `TimeZone`                                                                                   | `string`                                                                                     | :heavy_check_mark:                                                                           | The IANA time zone of the month and its days. `UTC` when the request omits `time_zone`.      |