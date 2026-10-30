# Cache

Optional best-effort cache for exchanged credentials, shared across the cluster. Losing it
only costs a re-exchange. Credentials are encrypted with the keys held by the Token Vault
Datastore before being cached.



## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Redis`                                                                                                     | [*components.AIGatewayRedisCloudConfiguration](../../models/components/aigatewayrediscloudconfiguration.md) | :heavy_minus_sign:                                                                                          | Config for connecting to a Cloud Provider's Redis instance.                                                 |