# KongTeamManager

The manager of a team, set through the team's `team_manager` custom identity. Always null for plain groups, and for teams without a manager.



## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ID`                                                                                | `string`                                                                            | :heavy_check_mark:                                                                  | Contains a unique identifier used for this resource.                                | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                |
| `DisplayName`                                                                       | `string`                                                                            | :heavy_check_mark:                                                                  | Display name of the manager.                                                        | josh                                                                                |
| `Email`                                                                             | `*string`                                                                           | :heavy_check_mark:                                                                  | Value of the manager's `email` custom identity, or null when the manager has none.<br/> | josh@acme.com                                                                       |