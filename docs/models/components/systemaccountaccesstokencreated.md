# SystemAccountAccessTokenCreated

A response including a single system account access token with the token.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | **string*                                                           | :heavy_minus_sign:                                                  | ID of the system account access token.                              |
| `Name`                                                              | **string*                                                           | :heavy_minus_sign:                                                  | Name of the system account access token.                            |
| `CreatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | Timestamp of when the system account access token was created.      |
| `UpdatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | Timestamp of when the system account access token was last updated. |
| `ExpiresAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | Timestamp of when the system account access token will expire.      |
| `LastUsedAt`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | Timestamp of when the system account access token was last used.    |
| `Token`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | The token of the system account access token.                       |