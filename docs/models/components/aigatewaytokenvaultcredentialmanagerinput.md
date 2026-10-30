# AIGatewayTokenVaultCredentialManagerInput

Serves the self-service enrollment UI through which end users authorize the upstream providers
this auth strategy authenticates against. Omit `token_vault` entirely to serve no such UI.

Requires a `token-vault` Datastore reference in `datastores`; enabling the credential manager
without one is rejected.

**Requires a minimum runtime version of `2.3`**.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `EnableCredentialManager`                                                | `*bool`                                                                  | :heavy_minus_sign:                                                       | Whether the self-service enrollment UI is served for this auth strategy. |