# Ssl


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Enabled`                                                 | `*bool`                                                   | :heavy_minus_sign:                                        | whether to use ssl for the pgvector database              |
| `Cert`                                                    | `*string`                                                 | :heavy_minus_sign:                                        | the path of ssl cert to use for the pgvector database     |
| `CertKey`                                                 | `*string`                                                 | :heavy_minus_sign:                                        | the path of ssl cert key to use for the pgvector database |
| `Required`                                                | `*bool`                                                   | :heavy_minus_sign:                                        | whether ssl is required for the pgvector database         |
| `Verify`                                                  | `*bool`                                                   | :heavy_minus_sign:                                        | whether to verify ssl for the pgvector database           |
| `Version`                                                 | [*components.Version](../../models/components/version.md) | :heavy_minus_sign:                                        | the ssl version to use for the pgvector database          |