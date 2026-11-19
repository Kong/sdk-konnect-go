# ListCostNotificationEventsResponseBody

A page of threshold events.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Data`                                                                                                       | [][components.AICostNotificationThresholdEvent](../../models/components/aicostnotificationthresholdevent.md) | :heavy_check_mark:                                                                                           | The threshold events on this page, in `sort` order. Channel tests never appear in this list.                 |
| `Meta`                                                                                                       | [components.CursorMeta](../../models/components/cursormeta.md)                                               | :heavy_check_mark:                                                                                           | Pagination metadata.                                                                                         |