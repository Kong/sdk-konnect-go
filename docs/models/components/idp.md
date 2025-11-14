# IDP

The IdP object contains the configuration data for the OIDC authentication integration.

NOTE: The `openid` scope is required. Removing it could break the OIDC integration.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Issuer`                                                              | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   | https://myidp.com/oauth2                                              |
| `LoginPath`                                                           | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   | myapp                                                                 |
| `ClientID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   | YOUR_CLIENT_ID                                                        |
| `Scopes`                                                              | []*string*                                                            | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |
| `ClaimMappings`                                                       | [*components.ClaimMappings](../../models/components/claimmappings.md) | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |