# MCPCapabilityControlsDeny

The capabilities of the mapped MCP server that are denied.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Tools`                                    | []`string`                                 | :heavy_check_mark:                         | The names of the tools that are denied.    |
| `Resources`                                | []`string`                                 | :heavy_check_mark:                         | The URIs of the resources that are denied. |
| `Prompts`                                  | []`string`                                 | :heavy_check_mark:                         | The names of the prompts that are denied.  |