# VirtualClusterAuthenticationValidate

Validation rules.


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `Audiences`                                                                                                          | [][components.VirtualClusterAuthenticationAudience](../../models/components/virtualclusterauthenticationaudience.md) | :heavy_minus_sign:                                                                                                   | List of expected audience values. One of them has to match the audience claim in the token.                          |
| `Issuer`                                                                                                             | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Expected token issuer in the token.                                                                                  |