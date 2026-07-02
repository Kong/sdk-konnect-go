# PortalMCPServerResponse

A portal MCP server belonging to a portal.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | `string`                                                      | :heavy_check_mark:                                            | Contains a unique identifier used for this resource.          | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                          |
| `Name`                                                        | `string`                                                      | :heavy_check_mark:                                            | N/A                                                           |                                                               |
| `Enabled`                                                     | `bool`                                                        | :heavy_check_mark:                                            | Is the MCP server enabled for the portal?                     |                                                               |
| `CreatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `UpdatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity update date.   | 2022-11-04T20:10:06.927Z                                      |