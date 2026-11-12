# UpdateOIDCIdentityProviderClaimMappings

Defines the mappings between OpenID Connect (OIDC) claims and local claims used by your application for
authentication. Only the provided mappings are updated and provided values cannot be empty.



## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Name`                                                         | `*string`                                                      | :heavy_minus_sign:                                             | The claim mapping for the user's name.                         | name                                                           |
| `Email`                                                        | `*string`                                                      | :heavy_minus_sign:                                             | The claim mapping for the user's email address.                | email                                                          |
| `Groups`                                                       | `*string`                                                      | :heavy_minus_sign:                                             | The claim mapping for the user's group membership information. | groups                                                         |