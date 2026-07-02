# AIGatewayMCPACLs

Access control rules for MCP resources. Configure `allow`, `deny`, or both.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Allow`                                            | []`string`                                         | :heavy_minus_sign:                                 | List of consumer groups that are permitted access. |
| `Deny`                                             | []`string`                                         | :heavy_minus_sign:                                 | List of consumer groups that are denied access.    |