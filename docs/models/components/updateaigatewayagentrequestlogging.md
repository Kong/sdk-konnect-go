# UpdateAIGatewayAgentRequestLogging

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for AI Gateway logging.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Payloads`                                                                                                  | `*bool`                                                                                                     | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |                                                                                                             |
| `MaxPayloadSize`                                                                                            | `*int64`                                                                                                    | :heavy_minus_sign:                                                                                          | Maximum size in bytes for logged request/response payloads. Payloads exceeding this size will be truncated. | 524288                                                                                                      |