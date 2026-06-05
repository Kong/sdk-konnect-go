# CreateWorkspaceRequest

The request schema for the create a workspace request.


## Fields

| Field                             | Type                              | Required                          | Description                       | Example                           |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `Name`                            | `string`                          | :heavy_check_mark:                | The name of the workspace.        | team-1                            |
| `Comment`                         | `*string`                         | :heavy_minus_sign:                | The description of the workspace. | A test workspace for team 1       |
| `Description`                     | `*string`                         | :heavy_minus_sign:                | The description of the workspace. | A test workspace for team 1       |
| `ManagedBy`                       | map[string]`any`                  | :heavy_minus_sign:                | N/A                               |                                   |