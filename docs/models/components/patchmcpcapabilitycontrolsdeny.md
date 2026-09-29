# PatchMCPCapabilityControlsDeny

The capabilities of the mapped MCP server to deny. Omitted keys preserve their current value, an empty array clears the capability, and a non-empty array replaces.



## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Tools`                                    | []`string`                                 | :heavy_minus_sign:                         | The names of the tools that are denied.    |
| `Resources`                                | []`string`                                 | :heavy_minus_sign:                         | The URIs of the resources that are denied. |
| `Prompts`                                  | []`string`                                 | :heavy_minus_sign:                         | The names of the prompts that are denied.  |