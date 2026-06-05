# CreateMCPServerRequest


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Name`                                               | `string`                                             | :heavy_check_mark:                                   | The name of the MCP server.                          |
| `Description`                                        | `*string`                                            | :heavy_minus_sign:                                   | A description of the MCP server.                     |
| `Labels`                                             | map[string]`string`                                  | :heavy_minus_sign:                                   | Labels for the MCP server.                           |
| `ControlPlaneID`                                     | `*string`                                            | :heavy_minus_sign:                                   | The control plane ID associated with the MCP server. |
| `ResourceID`                                         | `*string`                                            | :heavy_minus_sign:                                   | The MCP resource ID associated with the MCP server.  |