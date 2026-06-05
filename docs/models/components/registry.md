# Registry


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `ID`                                      | `string`                                  | :heavy_check_mark:                        | Unique registry identifier                | 123e4567-e89b-12d3-a456-426614174000      |
| `Name`                                    | `string`                                  | :heavy_check_mark:                        | Unique registry name                      | my-registry                               |
| `DisplayName`                             | `string`                                  | :heavy_check_mark:                        | Human-readable registry name              | My MCP Registry                           |
| `Description`                             | `*string`                                 | :heavy_check_mark:                        | Optional registry description             | Registry for internal MCP servers         |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Registry creation timestamp               | 2023-12-01T10:30:00Z                      |
| `UpdatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Registry last update timestamp            | 2023-12-01T11:00:00Z                      |