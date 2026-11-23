# AIGatewayMCPServerUpstreamConfigOutput

Configuration applied when proxying to the upstream service, including authentication.
Applies to `conversion-listener`, `conversion-only`, and `upstream-server` MCP Servers.

**Requires a minimum runtime version of `2.4`**.


## Fields

| Field                                                                                                                           | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `Auth`                                                                                                                          | [*components.AIGatewayMCPServerUpstreamConfigAuthOutput](../../models/components/aigatewaymcpserverupstreamconfigauthoutput.md) | :heavy_minus_sign:                                                                                                              | Authentication to use when proxying to the upstream service.                                                                    |