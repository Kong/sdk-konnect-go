# ListCostNotificationRuleFiringsResponseBody

A page of firings.


## Fields

| Field                                                                                                              | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `Data`                                                                                                             | [][operations.ListCostNotificationRuleFiringsData](../../models/operations/listcostnotificationrulefiringsdata.md) | :heavy_check_mark:                                                                                                 | The firings on this page, in `sort` order.                                                                         |
| `Meta`                                                                                                             | [components.CursorMeta](../../models/components/cursormeta.md)                                                     | :heavy_check_mark:                                                                                                 | Pagination metadata.                                                                                               |