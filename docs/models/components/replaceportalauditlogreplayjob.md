# ReplacePortalAuditLogReplayJob

The request schema to replace a portal audit log replay job.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `StartAt`                                                                                                   | [time.Time](https://pkg.go.dev/time#Time)                                                                   | :heavy_check_mark:                                                                                          | The start of a date-time range in RFC3339 format e.g. 2017-07-21T17:32:28Z.<br/>Must be within the last 7 days. |
| `EndAt`                                                                                                     | [time.Time](https://pkg.go.dev/time#Time)                                                                   | :heavy_check_mark:                                                                                          | The end of a date-time range in RFC3339 format e.g. 2017-07-21T17:32:28Z.<br/>Must be within the last 7 days. |