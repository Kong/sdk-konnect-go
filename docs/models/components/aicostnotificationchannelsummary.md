# AICostNotificationChannelSummary

A short form of the notification channel a message goes to, with its `id`, `name`, and `type`.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ID`                                                                                    | `string`                                                                                | :heavy_check_mark:                                                                      | The `id` of the channel, which the notification channel operations take as `channelId`. |
| `Name`                                                                                  | `string`                                                                                | :heavy_check_mark:                                                                      | The channel's human-readable name.                                                      |
| `Type`                                                                                  | [components.Type](../../models/components/type.md)                                      | :heavy_check_mark:                                                                      | Whether the channel sends email or posts to Slack.                                      |