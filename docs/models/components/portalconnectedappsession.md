# PortalConnectedAppSession

An active Portal Connected App session (refresh token)


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | `string`                                                      | :heavy_check_mark:                                            | The ID of the session                                         |                                                               |
| `ClientID`                                                    | `string`                                                      | :heavy_check_mark:                                            | The client ID of the connected app                            | abc123-client                                                 |
| `ClientName`                                                  | `string`                                                      | :heavy_check_mark:                                            | The name of the connected app                                 | My Connected App                                              |
| `GrantedScopeNames`                                           | []`string`                                                    | :heavy_check_mark:                                            | The scopes that the session has been granted                  | [<br/>"konnect:read",<br/>"konnect:write"<br/>]               |
| `CreatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `ExpiresAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of the expiration date   | 2022-11-04T20:10:06.927Z                                      |
| `DeveloperID`                                                 | `string`                                                      | :heavy_check_mark:                                            | The ID of the developer who owns the session                  |                                                               |