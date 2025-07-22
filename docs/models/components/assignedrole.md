# AssignedRole

An assigned role is a role that has been assigned to a user or team.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | **string*                                                           | :heavy_minus_sign:                                                  | The ID of the role assignment.                                      | eaf7adf1-32c8-4bbf-b960-d1f8456afe67                                |
| `RoleName`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | Name of the role being assigned.                                    | Viewer                                                              |
| `EntityID`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | A RBAC entity ID.                                                   | 817d0422-45c9-4d88-8d64-45aef05c1ae7                                |
| `EntityTypeName`                                                    | **string*                                                           | :heavy_minus_sign:                                                  | Name of the entity type the role is being assigned to.              | Control Planes                                                      |
| `EntityRegion`                                                      | [*components.EntityRegion](../../models/components/entityregion.md) | :heavy_minus_sign:                                                  | N/A                                                                 | eu                                                                  |