# AIGatewayMCPPassthroughTool

A tool exposed by an MCP Server in `passthrough-listener` mode.



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Access`                                                                               | [components.AIGatewayMCPToolAccess](../../models/components/aigatewaymcptoolaccess.md) | :heavy_check_mark:                                                                     | Access-control rules for a tool.                                                       |
| `Name`                                                                                 | `string`                                                                               | :heavy_check_mark:                                                                     | Tool identifier used to match remote MCP Server tools for ACL enforcement.             |