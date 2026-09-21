# AIGatewayConfigStoreSecret


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Key`                                                          | `string`                                                       | :heavy_check_mark:                                             | The unique key identifying the secret within the Config Store. | my-secret-key                                                  |
| `CreatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity creation date.  | 2022-11-04T20:10:06.927Z                                       |
| `UpdatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity update date.    | 2022-11-04T20:10:06.927Z                                       |