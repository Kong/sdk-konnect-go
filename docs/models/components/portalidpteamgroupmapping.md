# PortalIdpTeamGroupMapping


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `ID`                                      | `string`                                  | :heavy_check_mark:                        | The mapping ID.                           | a1b2c3d4-e5f6-4a8b-9c0d-1e2f3a4b5c6d      |
| `TeamID`                                  | `string`                                  | :heavy_check_mark:                        | The Konnect team ID.                      | 6801e673-cc10-498a-94cd-4271de07a0d3      |
| `Group`                                   | `string`                                  | :heavy_check_mark:                        | The IdP group name.                       | API Engineers                             |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Creation timestamp.                       | 2024-01-15T10:30:00Z                      |
| `UpdatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Last update timestamp.                    | 2024-01-15T10:30:00Z                      |