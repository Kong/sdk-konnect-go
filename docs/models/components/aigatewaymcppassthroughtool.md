# AIGatewayMCPPassthroughTool

**Pre-release Feature**
This feature is currently in beta and is subject to change.

A tool exposed by an MCP Server in `passthrough-listener` mode.


## Fields

| Field                                                                                                                 | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `Access`                                                                                                              | [components.AIGatewayMCPToolAccess](../../models/components/aigatewaymcptoolaccess.md)                                | :heavy_check_mark:                                                                                                    | **Pre-release Feature**<br/>This feature is currently in beta and is subject to change.<br/><br/>Access-control rules for a tool. |
| `Name`                                                                                                                | `string`                                                                                                              | :heavy_check_mark:                                                                                                    | Tool identifier used to match remote MCP Server tools for ACL enforcement.                                            |