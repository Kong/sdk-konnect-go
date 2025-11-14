# PersonalAccessTokenCreateRequest

Request body schema for creating personal access tokens.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Name`                                                          | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |                                                                 |
| `ExpiresAt`                                                     | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | An ISO-8601 timestamp representation of entity expiration date. | 2022-11-04T20:10:06.927Z                                        |