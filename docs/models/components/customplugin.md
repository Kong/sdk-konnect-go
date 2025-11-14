# CustomPlugin


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreatedAt`                                                   | **int64*                                                      | :heavy_minus_sign:                                            | Unix epoch when the resource was created.                     |
| `Handler`                                                     | *string*                                                      | :heavy_check_mark:                                            | The handler for the given custom plugin.                      |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | A string representing a UUID (universally unique identifier). |
| `Name`                                                        | *string*                                                      | :heavy_check_mark:                                            | The name to associate with the given custom plugin.           |
| `Schema`                                                      | *string*                                                      | :heavy_check_mark:                                            | The schema for the given custom plugin.                       |
| `Tags`                                                        | []*string*                                                    | :heavy_minus_sign:                                            | A set of strings representing tags.                           |
| `UpdatedAt`                                                   | **int64*                                                      | :heavy_minus_sign:                                            | Unix epoch when the resource was last updated.                |