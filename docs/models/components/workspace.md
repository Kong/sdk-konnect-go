# Workspace

Workspaces provide a way to segment Kong entities.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | `*string`                                                     | :heavy_minus_sign:                                            | The unique UUID for this resource.                            |                                                               |
| `Name`                                                        | `string`                                                      | :heavy_check_mark:                                            | N/A                                                           |                                                               |
| `Comment`                                                     | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |                                                               |
| `Description`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |                                                               |
| `ManagedBy`                                                   | map[string]`any`                                              | :heavy_minus_sign:                                            | N/A                                                           |                                                               |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `UpdatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | An ISO-8601 timestamp representation of entity update date.   | 2022-11-04T20:10:06.927Z                                      |