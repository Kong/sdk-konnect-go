# Token


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | `string`                                                      | :heavy_check_mark:                                            | The ID of the token                                           |                                                               |
| `GrantedScopeNames`                                           | []`string`                                                    | :heavy_check_mark:                                            | The scopes that the token has been granted                    |                                                               |
| `CreatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `ExpiresAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of the expiration date   | 2022-11-04T20:10:06.927Z                                      |