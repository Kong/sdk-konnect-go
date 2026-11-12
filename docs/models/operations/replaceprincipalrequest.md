# ReplacePrincipalRequest


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `DirectoryID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the directory                                                                |
| `PrincipalID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | Unique identifier of the principal (UUID format)                                   |
| `ReplacePrincipalBody`                                                             | [components.ReplacePrincipalBody](../../models/components/replaceprincipalbody.md) | :heavy_check_mark:                                                                 | Principal to be replaced                                                           |