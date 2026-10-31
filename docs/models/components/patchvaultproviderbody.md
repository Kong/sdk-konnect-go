# PatchVaultProviderBody

Fields to update on a vault provider


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Name`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Enabled`                                                             | `*bool`                                                               | :heavy_minus_sign:                                                    | Defines if the provider is enabled or not                             |
| `ClientSecret`                                                        | `*string`                                                             | :heavy_minus_sign:                                                    | Replaces the stored OAuth client secret. Never returned in read APIs. |
| `Scopes`                                                              | []`string`                                                            | :heavy_minus_sign:                                                    | N/A                                                                   |
| `TokenEndpoint`                                                       | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `AuthorizationEndpoint`                                               | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `PkceEnabled`                                                         | `*bool`                                                               | :heavy_minus_sign:                                                    | N/A                                                                   |
| `RefreshBufferSeconds`                                                | `*int64`                                                              | :heavy_minus_sign:                                                    | N/A                                                                   |