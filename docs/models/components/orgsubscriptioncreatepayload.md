# OrgSubscriptionCreatePayload

Defines which destinations receive an event and the regions and entities for which notifications are sent.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `DestinationIds`                                                                  | []`string`                                                                        | :heavy_check_mark:                                                                | IDs of 1 to 50 HTTP destinations that receive this event. Each ID must be unique. |
| `Regions`                                                                         | [][components.NotificationRegion](../../models/components/notificationregion.md)  | :heavy_check_mark:                                                                | Regions for which notifications are sent. Use `*` for all regions.                |
| `Entities`                                                                        | []`string`                                                                        | :heavy_check_mark:                                                                | Entity IDs for which notifications are sent. Use `*` for all entities.            |
| `Enabled`                                                                         | `bool`                                                                            | :heavy_check_mark:                                                                | Set to false to create the subscription in a paused state.                        |
| `Name`                                                                            | `string`                                                                          | :heavy_check_mark:                                                                | Name that identifies the subscription.                                            |