# EventSubscription

Properties associated with an event subscription.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Regions`                                                                          | [][components.NotificationRegion](../../models/components/notificationregion.md)   | :heavy_check_mark:                                                                 | N/A                                                                                |
| `Entities`                                                                         | []*string*                                                                         | :heavy_check_mark:                                                                 | N/A                                                                                |
| `Channels`                                                                         | [][components.NotificationChannel](../../models/components/notificationchannel.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `Enabled`                                                                          | *bool*                                                                             | :heavy_check_mark:                                                                 | Enable/Disable complete subscription within an event.                              |