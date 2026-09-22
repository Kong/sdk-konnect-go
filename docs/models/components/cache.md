# Cache

Cache hints Kong emits on the cacheable operations it serves. Only clients on a protocol
revision that defines them receive them.

**Requires a minimum runtime version of `2.1`**.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ToolsList`                                                                                       | [*components.AIGatewayMCPServerCacheHint](../../models/components/aigatewaymcpservercachehint.md) | :heavy_minus_sign:                                                                                | A cache hint Kong emits on a cacheable operation it serves.                                       |
| `Discover`                                                                                        | [*components.AIGatewayMCPServerCacheHint](../../models/components/aigatewaymcpservercachehint.md) | :heavy_minus_sign:                                                                                | A cache hint Kong emits on a cacheable operation it serves.                                       |