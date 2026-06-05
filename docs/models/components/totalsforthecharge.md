# TotalsForTheCharge

Aggregated booked and realtime totals for the charge.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Booked`                                                                   | [components.Booked](../../models/components/booked.md)                     | :heavy_check_mark:                                                         | The amount of the charge already booked to the internal accounting system. |
| `Realtime`                                                                 | [*components.RealtimeTotals](../../models/components/realtimetotals.md)    | :heavy_minus_sign:                                                         | The realtime amount of the charge.<br/><br/>Requires the `realtime_usage` expand. |