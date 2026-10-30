# AIGatewayTokenVaultCredential

Where the caller's subject token is read from, and where the credential the Token Vault returns
is placed on the upstream request.

Both placements default to `Authorization: Bearer <credential>` in a header, though an
`outbound` placement left unset honours whatever the provider advertises instead. Set them for
providers that take the credential elsewhere, such as an `x-api-key` header carrying the raw
credential.



## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Inbound`                                                                                                               | [*components.AIGatewayTokenVaultCredentialPlacement](../../models/components/aigatewaytokenvaultcredentialplacement.md) | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     | {<br/>"location": "header",<br/>"name": "x-api-key",<br/>"prefix": ""<br/>}                                             |
| `Outbound`                                                                                                              | [*components.AIGatewayTokenVaultCredentialPlacement](../../models/components/aigatewaytokenvaultcredentialplacement.md) | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     | {<br/>"location": "header",<br/>"name": "x-api-key",<br/>"prefix": ""<br/>}                                             |