# UpdatePrincipalRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `DirectoryID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | ID of the directory                                                              |
| `PrincipalID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | Unique identifier of the principal (UUID format)                                 |
| `UpdatePrincipalBody`                                                            | [components.UpdatePrincipalBody](../../models/components/updateprincipalbody.md) | :heavy_check_mark:                                                               | Principal to be updated                                                          |