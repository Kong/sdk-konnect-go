# BulkPayload

Request body schema for bulk status update.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Ids`                                                                          | []*string*                                                                     | :heavy_check_mark:                                                             | N/A                                                                            |
| `Status`                                                                       | [components.NotificationStatus](../../models/components/notificationstatus.md) | :heavy_check_mark:                                                             | Status of the notification.                                                    |