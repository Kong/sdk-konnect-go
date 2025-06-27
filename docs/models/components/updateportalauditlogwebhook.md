# UpdatePortalAuditLogWebhook

The request schema to modify an portal audit log webhook.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Enabled`                                                           | **bool*                                                             | :heavy_minus_sign:                                                  | Indicates if the data should be sent to the configured destination. | true                                                                |
| `AuditLogDestinationID`                                             | **string*                                                           | :heavy_minus_sign:                                                  | ID of the audit log destination.                                    |                                                                     |