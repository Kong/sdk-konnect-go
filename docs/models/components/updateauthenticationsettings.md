# UpdateAuthenticationSettings

The request schema to update an organization's authentication settings.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BasicAuthEnabled`                                         | **bool*                                                    | :heavy_minus_sign:                                         | The organization has basic auth enabled.                   | true                                                       |
| `OidcAuthEnabled`                                          | **bool*                                                    | :heavy_minus_sign:                                         | The organization has OIDC disabled.                        | false                                                      |
| `SamlAuthEnabled`                                          | **bool*                                                    | :heavy_minus_sign:                                         | The organization has SAML disabled.                        | false                                                      |
| `IdpMappingEnabled`                                        | **bool*                                                    | :heavy_minus_sign:                                         | Whether IdP groups determine the Konnect teams a user has. | true                                                       |
| `KonnectMappingEnabled`                                    | **bool*                                                    | :heavy_minus_sign:                                         | Whether a Konnect Identity Admin assigns teams to a user.  | false                                                      |