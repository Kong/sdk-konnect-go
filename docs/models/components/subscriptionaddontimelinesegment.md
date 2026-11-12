# SubscriptionAddonTimelineSegment

A subscription add-on event.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ActiveFrom`                                                               | [time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_check_mark:                                                         | An ISO-8601 timestamp representation of the cadence start of the resource. | 2023-01-01T01:01:01.001Z                                                   |
| `ActiveTo`                                                                 | [*time.Time](https://pkg.go.dev/time#Time)                                 | :heavy_minus_sign:                                                         | An ISO-8601 timestamp representation of the cadence end of the resource.   | 2023-01-01T01:01:01.001Z                                                   |
| `Quantity`                                                                 | `int64`                                                                    | :heavy_check_mark:                                                         | The quantity of the add-on for the given period.                           | 1                                                                          |