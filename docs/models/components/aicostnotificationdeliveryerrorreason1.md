# AICostNotificationDeliveryErrorReason1

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostNotificationDeliveryErrorReason1SlackNoService

// Open enum: custom values can be created with a direct type cast
custom := components.AICostNotificationDeliveryErrorReason1("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `AICostNotificationDeliveryErrorReason1SlackNoService`         | slack_no_service                                               |
| `AICostNotificationDeliveryErrorReason1SlackChannelNotFound`   | slack_channel_not_found                                        |
| `AICostNotificationDeliveryErrorReason1SlackChannelIsArchived` | slack_channel_is_archived                                      |
| `AICostNotificationDeliveryErrorReason1SlackActionProhibited`  | slack_action_prohibited                                        |
| `AICostNotificationDeliveryErrorReason1SlackInvalidPayload`    | slack_invalid_payload                                          |
| `AICostNotificationDeliveryErrorReason1SlackUserNotFound`      | slack_user_not_found                                           |
| `AICostNotificationDeliveryErrorReason1SlackTeamDisabled`      | slack_team_disabled                                            |
| `AICostNotificationDeliveryErrorReason1SlackInvalidToken`      | slack_invalid_token                                            |
| `AICostNotificationDeliveryErrorReason1SlackRedirected`        | slack_redirected                                               |
| `AICostNotificationDeliveryErrorReason1SlackInvalidTarget`     | slack_invalid_target                                           |
| `AICostNotificationDeliveryErrorReason1EmailInvalidTarget`     | email_invalid_target                                           |
| `AICostNotificationDeliveryErrorReason1MailboxPublishFailed`   | mailbox_publish_failed                                         |
| `AICostNotificationDeliveryErrorReason1EmailRenderFailed`      | email_render_failed                                            |
| `AICostNotificationDeliveryErrorReason1TransportError`         | transport_error                                                |
| `AICostNotificationDeliveryErrorReason1AttemptsExhausted`      | attempts_exhausted                                             |
| `AICostNotificationDeliveryErrorReason1WindowEnded`            | window_ended                                                   |