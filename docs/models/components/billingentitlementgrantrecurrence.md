# BillingEntitlementGrantRecurrence

The recurrence of the grant. When set, the grant amount is re-issued every
interval from the anchor, which defaults to `effective_at`. Absent for
non-recurring grants.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Anchor`                                            | [time.Time](https://pkg.go.dev/time#Time)           | :heavy_check_mark:                                  | A date-time anchor to base the recurring period on. | 2023-01-01T01:01:01.001Z                            |
| `Interval`                                          | `string`                                            | :heavy_check_mark:                                  | The interval duration in ISO 8601 format.           | P1M                                                 |