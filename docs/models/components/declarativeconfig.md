# DeclarativeConfig

Declarative Config Response


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Config`                                                      | `*string`                                                     | :heavy_minus_sign:                                            | The configuration as string. Must be valid yaml or json.      |                                                               |
| `Version`                                                     | `*string`                                                     | :heavy_minus_sign:                                            | An uuid identifier (v7) of the version of the config payload. | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                          |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `UpdatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | An ISO-8601 timestamp representation of entity update date.   | 2022-11-04T20:10:06.927Z                                      |