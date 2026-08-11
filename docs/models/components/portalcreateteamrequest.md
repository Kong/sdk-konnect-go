# PortalCreateTeamRequest

Details about a team to create.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Name`                                          | `string`                                        | :heavy_check_mark:                              | The name of the team.                           | IDM - Developers                                |
| `Description`                                   | `*string`                                       | :heavy_minus_sign:                              | The description of the team.                    | The Identity Management (IDM) team.             |
| `CanOwnApplications`                            | `*bool`                                         | :heavy_minus_sign:                              | Whether the team is allowed to own applications | true                                            |