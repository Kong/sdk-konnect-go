# AIGatewayTokenVaultOutput

Resolves an upstream credential per request via Kong's Token Vault instead of sending a static
credential. Exchanged credentials are cached per node and, when `redis` is configured, shared
across the cluster. Callers must enroll with the configured Token Vault provider before the
upstream tools are exposed: until enrollment completes, the MCP Server serves virtual
`authenticate` and `check_authentication_status` tools that guide the caller through the
enrollment flow.

**Requires a minimum runtime version of `2.3`**.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Directory`                                                                                                             | `string`                                                                                                                | :heavy_check_mark:                                                                                                      | Directory name segment of the vault token endpoint.                                                                     |
| `Provider`                                                                                                              | `string`                                                                                                                | :heavy_check_mark:                                                                                                      | Name of the upstream credential provider registered in the Token Vault directory.                                       |
| `Redis`                                                                                                                 | [*components.AIGatewayRedisCloudConfigurationOutput](../../models/components/aigatewayrediscloudconfigurationoutput.md) | :heavy_minus_sign:                                                                                                      | Config for connecting to a Cloud Provider's Redis instance.                                                             |