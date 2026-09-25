# AISettingsFeatures

AI features configuration. When top-level `enabled` is false, every feature toggle here is automatically reset to false.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `PortalAgent`                                                                                    | [*components.PortalAgent](../../models/components/portalagent.md)                                | :heavy_minus_sign:                                                                               | Portal Agent config                                                                              |
| `AiSearch`                                                                                       | [*components.AiSearch](../../models/components/aisearch.md)                                      | :heavy_minus_sign:                                                                               | AI Search config                                                                                 |
| `McpServer`                                                                                      | [components.AISettingsFeaturesMcpServer](../../models/components/aisettingsfeaturesmcpserver.md) | :heavy_check_mark:                                                                               | AI Features config                                                                               |