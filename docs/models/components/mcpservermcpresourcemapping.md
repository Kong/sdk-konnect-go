# MCPServerMCPResourceMapping

A response to creating an MCP server MCP resource mapping.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ID`                                                           | `string`                                                       | :heavy_check_mark:                                             | The unique identifier for the MCP server MCP resource mapping. | edf35010-61d0-4fc4-a3b2-8095b02ec3fe                           |
| `McpServerID`                                                  | `string`                                                       | :heavy_check_mark:                                             | The ID of the MCP server.                                      | d5756e2a-75e4-470b-9a48-326eb9303fd6                           |
| `McpResourceID`                                                | `string`                                                       | :heavy_check_mark:                                             | The ID of the MCP resource.                                    | a60eceeb-dd7c-4884-8b26-806489c2f704                           |
| `CreatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity creation date.  | 2022-11-04T20:10:06.927Z                                       |
| `UpdatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity update date.    | 2022-11-04T20:10:06.927Z                                       |