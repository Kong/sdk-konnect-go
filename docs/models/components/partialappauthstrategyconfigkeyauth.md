# PartialAppAuthStrategyConfigKeyAuth

Key Auth configuration for updating an Application Auth Strategy.
The ttl field can be set to null to unset the Time-To-Live.



## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `KeyNames`                                                                                                              | []*string*                                                                                                              | :heavy_minus_sign:                                                                                                      | The names of the headers containing the API key. You can specify multiple header names.                                 |
| `TTL`                                                                                                                   | [*components.PartialAppAuthStrategyConfigKeyAuthTTL](../../models/components/partialappauthstrategyconfigkeyauthttl.md) | :heavy_minus_sign:                                                                                                      | Default maximum Time-To-Live for keys created under this strategy. Set to null to unset.                                |