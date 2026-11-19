# ListCostAttributionPluginsResponseBody

A page of attribution plugin installations.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Data`                                                                                     | [][components.AICostAttributionPlugin](../../models/components/aicostattributionplugin.md) | :heavy_check_mark:                                                                         | The attribution plugin installations on this page, in `sort` order.                        |
| `Meta`                                                                                     | [components.CursorMeta](../../models/components/cursormeta.md)                             | :heavy_check_mark:                                                                         | Pagination metadata.                                                                       |