# PatchAISettingsRequestFeatures

AI Features configuration. Only nullable when top-level `enabled` is false


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `PortalAgent`                                                                                                 | [*components.PatchAISettingsRequestPortalAgent](../../models/components/patchaisettingsrequestportalagent.md) | :heavy_minus_sign:                                                                                            | Portal Agent configuration                                                                                    |
| `McpServer`                                                                                                   | [components.PatchAISettingsRequestMcpServer](../../models/components/patchaisettingsrequestmcpserver.md)      | :heavy_check_mark:                                                                                            | MCP Server configuration                                                                                      |