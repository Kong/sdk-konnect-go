# UpdateIDPConfiguration


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `Issuer`                                                                                                          | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | N/A                                                                                                               | https://myidp.com/oauth2                                                                                          |
| `LoginPath`                                                                                                       | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | N/A                                                                                                               | myapp                                                                                                             |
| `ClientID`                                                                                                        | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | N/A                                                                                                               | YOUR_CLIENT_ID                                                                                                    |
| `ClientSecret`                                                                                                    | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | N/A                                                                                                               | YOUR_CLIENT_SECRET                                                                                                |
| `Scopes`                                                                                                          | []*string*                                                                                                        | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |                                                                                                                   |
| `ClaimMappings`                                                                                                   | [*components.UpdateIDPConfigurationClaimMappings](../../models/components/updateidpconfigurationclaimmappings.md) | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |                                                                                                                   |