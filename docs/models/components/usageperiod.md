# UsagePeriod

The usage period of the entitlement. The balance resets at the start of every
period. The anchor defaults to the entitlement creation time.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Interval`                                          | `string`                                            | :heavy_check_mark:                                  | The interval duration in ISO 8601 format.           | P1M                                                 |
| `Anchor`                                            | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | A date-time anchor to base the recurring period on. | 2023-01-01T01:01:01.001Z                            |