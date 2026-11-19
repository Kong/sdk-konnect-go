# AICostTimeWindow

The start and end of one budget window.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `StartsAt`                                                                            | [time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_check_mark:                                                                    | When the budget window starts, as an RFC 3339 timestamp in UTC.                       |
| `EndsAt`                                                                              | [time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_check_mark:                                                                    | When the budget window ends and the next one starts, as an RFC 3339 timestamp in UTC. |