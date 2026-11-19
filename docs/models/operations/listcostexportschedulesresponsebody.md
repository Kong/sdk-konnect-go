# ListCostExportSchedulesResponseBody

A page of export schedules.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Data`                                                                               | [][components.AICostExportSchedule](../../models/components/aicostexportschedule.md) | :heavy_check_mark:                                                                   | The export schedules on this page, in `sort` order.                                  |
| `Meta`                                                                               | [components.CursorMeta](../../models/components/cursormeta.md)                       | :heavy_check_mark:                                                                   | Pagination metadata.                                                                 |