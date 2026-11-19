# ListCostNotificationChannelsResponseBody

A page of notification channels.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Data`                                                                                         | [][components.AICostNotificationChannel](../../models/components/aicostnotificationchannel.md) | :heavy_check_mark:                                                                             | The notification channels on this page, in `sort` order.                                       |
| `Meta`                                                                                         | [components.CursorMeta](../../models/components/cursormeta.md)                                 | :heavy_check_mark:                                                                             | Pagination metadata.                                                                           |