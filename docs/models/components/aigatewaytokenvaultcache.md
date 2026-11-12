# AIGatewayTokenVaultCache

Optional best-effort cache for exchanged credentials, shared across the cluster. Losing it
only costs a re-exchange. Credentials are encrypted with the keys held by the Token Vault
Datastore before being cached.



## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Redis`                                                                                                                 | [*components.AIGatewayRedisCloudConfigurationOutput](../../models/components/aigatewayrediscloudconfigurationoutput.md) | :heavy_minus_sign:                                                                                                      | Config for connecting to a Cloud Provider's Redis instance.                                                             |