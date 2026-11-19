# ListCostBudgetsResponseBody

A page of budgets.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Data`                                                               | [][components.AICostBudget](../../models/components/aicostbudget.md) | :heavy_check_mark:                                                   | The budgets on this page, in `sort` order.                           |
| `Meta`                                                               | [components.CursorMeta](../../models/components/cursormeta.md)       | :heavy_check_mark:                                                   | Pagination metadata.                                                 |