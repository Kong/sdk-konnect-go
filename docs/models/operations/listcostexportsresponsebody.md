# ListCostExportsResponseBody

A page of exports.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Data`                                                               | [][components.AICostExport](../../models/components/aicostexport.md) | :heavy_check_mark:                                                   | The exports on this page, in `sort` order.                           |
| `Meta`                                                               | [components.CursorMeta](../../models/components/cursormeta.md)       | :heavy_check_mark:                                                   | Pagination metadata.                                                 |