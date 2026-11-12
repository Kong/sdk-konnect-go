# AISettingsRequestFeatures

AI Features configuration. Only nullable when top-level `enabled` is false. When top-level `enabled` is false, every feature toggle here is automatically reset to false.


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `McpServer`                                                                                                     | [*components.AISettingsRequestFeaturesMcpServer](../../models/components/aisettingsrequestfeaturesmcpserver.md) | :heavy_minus_sign:                                                                                              | MCP Server configuration                                                                                        |