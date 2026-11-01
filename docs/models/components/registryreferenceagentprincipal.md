# RegistryReferenceAgentPrincipal

An agent's link to the Kong Identity principal that represents it — the stored reference plus its own id and timestamps. The registry does not manage the principal itself; clients manage it via Kong Identity.



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | `string`                                                           | :heavy_check_mark:                                                 | The principal link's own id.                                       | 423e4567-e89b-12d3-a456-426614174bbb                               |
| `PrincipalID`                                                      | `string`                                                           | :heavy_check_mark:                                                 | The linked Kong Identity principal's id, scoped to `directory_id`. | 523e4567-e89b-12d3-a456-426614174ccc                               |
| `DirectoryID`                                                      | `string`                                                           | :heavy_check_mark:                                                 | The id of the directory that hosts the principal.                  | 623e4567-e89b-12d3-a456-426614174ddd                               |
| `CreatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | When the link was created.                                         |                                                                    |
| `UpdatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | When the link was last updated.                                    |                                                                    |