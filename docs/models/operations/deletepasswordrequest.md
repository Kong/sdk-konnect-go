# DeletePasswordRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `DirectoryID`                                                | `string`                                                     | :heavy_check_mark:                                           | ID of the directory                                          |
| `PrincipalID`                                                | `string`                                                     | :heavy_check_mark:                                           | Unique identifier of the principal (UUID format)             |
| `BasicAuthID`                                                | `string`                                                     | :heavy_check_mark:                                           | Unique identifier of the basic auth credential (UUID format) |
| `PasswordID`                                                 | `string`                                                     | :heavy_check_mark:                                           | Unique identifier of the password (UUID format)              |