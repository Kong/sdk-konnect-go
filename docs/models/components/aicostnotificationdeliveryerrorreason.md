# AICostNotificationDeliveryErrorReason

Why a message couldn't be delivered to a channel. The value is one of a fixed set of reasons, never free text:

- `slack_no_service`, `slack_channel_not_found`, `slack_channel_is_archived`, `slack_action_prohibited`, `slack_invalid_payload`, `slack_user_not_found`, `slack_team_disabled`, `slack_invalid_token`: Slack answered with an error status and a response body that begins with the reason's name without `slack_`, such as `channel_is_archived`.
- `slack_http_<status>`: Slack answered with the HTTP status `<status>`, a 3-digit code. A `429` or `5xx` status always gives this reason. Any other error status gives it only when the response doesn't begin with one of the errors above.
- `slack_redirected`: Slack answered with a redirect, which the service never follows.
- `slack_invalid_target`: the stored webhook URL can no longer be decrypted or validated.
- `email_invalid_target`: the stored recipient list is no longer valid.
- `email_render_failed`: the service couldn't render the email.
- `mailbox_publish_failed`: handing the email to Konnect's email service, the mailbox, failed for a reason other than running out of time.
- `transport_error`: the attempt ended without an answer. The request to Slack failed at the network level, a Slack or email attempt ran out of time, or sending failed unexpectedly.
- `attempts_exhausted`: every attempt ended in a temporary error, so the service stopped retrying.
- `window_ended`: the delivery was still waiting to be sent when its budget window ended. Only deliveries from before a one-time data migration carry this reason.

A delivery fails at once on a permanent error. A temporary error (`transport_error`, `mailbox_publish_failed`, `slack_http_429`, or `slack_http_<status>` with a `5xx` status) makes the service retry, so a delivery never ends with one of those reasons. It ends with `attempts_exhausted` instead. A channel test isn't retried, so its `error` is the reason its single attempt failed, temporary or not.


## Supported Types

### AICostNotificationDeliveryErrorReason1

```go
aiCostNotificationDeliveryErrorReason := components.CreateAICostNotificationDeliveryErrorReasonAICostNotificationDeliveryErrorReason1(components.AICostNotificationDeliveryErrorReason1{/* values here */})
```

### 

```go
aiCostNotificationDeliveryErrorReason := components.CreateAICostNotificationDeliveryErrorReasonStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostNotificationDeliveryErrorReason.Type {
	case components.AICostNotificationDeliveryErrorReasonTypeAICostNotificationDeliveryErrorReason1:
		// aiCostNotificationDeliveryErrorReason.AICostNotificationDeliveryErrorReason1 is populated
	case components.AICostNotificationDeliveryErrorReasonTypeStr:
		// aiCostNotificationDeliveryErrorReason.Str is populated
}
```
