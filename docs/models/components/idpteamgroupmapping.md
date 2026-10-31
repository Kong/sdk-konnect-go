# IdpTeamGroupMapping

A mapping between a Konnect team and an identity provider group.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ID`                                                              | `string`                                                          | :heavy_check_mark:                                                | Contains a unique identifier used for this resource.              | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                              |
| `TeamID`                                                          | `string`                                                          | :heavy_check_mark:                                                | The Konnect team ID.                                              | 6801e673-cc10-498a-94cd-4271de07a0d3                              |
| `Group`                                                           | `string`                                                          | :heavy_check_mark:                                                | The identity provider group name. Group names are case sensitive. | Tech Leads                                                        |
| `CreatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | An ISO-8601 timestamp representation of entity creation date.     | 2022-11-04T20:10:06.927Z                                          |
| `UpdatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | An ISO-8601 timestamp representation of entity update date.       | 2022-11-04T20:10:06.927Z                                          |