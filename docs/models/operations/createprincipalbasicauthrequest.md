# CreatePrincipalBasicAuthRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `DirectoryID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | ID of the directory                                                              |
| `PrincipalID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | Unique identifier of the principal (UUID format)                                 |
| `CreateBasicAuthBody`                                                            | [components.CreateBasicAuthBody](../../models/components/createbasicauthbody.md) | :heavy_check_mark:                                                               | Basic auth credential to be created                                              |