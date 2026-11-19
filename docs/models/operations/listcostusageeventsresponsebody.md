# ListCostUsageEventsResponseBody

A page of usage events.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `Data`                                                                                                 | [][components.AICostUsageEvent](../../models/components/aicostusageevent.md)                           | :heavy_check_mark:                                                                                     | The usage events on this page, in `sort` order. Each has the `ai.usage.v1` form that producers report. |
| `Meta`                                                                                                 | [components.CursorMeta](../../models/components/cursormeta.md)                                         | :heavy_check_mark:                                                                                     | Pagination metadata.                                                                                   |