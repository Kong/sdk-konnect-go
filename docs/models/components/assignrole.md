# AssignRole

An assigned role is a role that has been assigned to a user or team.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `RoleName`                                                              | [*components.RoleName](../../models/components/rolename.md)             | :heavy_minus_sign:                                                      | The desired role.                                                       | Viewer                                                                  |
| `EntityID`                                                              | **string*                                                               | :heavy_minus_sign:                                                      | The ID of the entity.                                                   | e67490ce-44dc-4cbd-b65e-b52c746fc26a                                    |
| `EntityTypeName`                                                        | [*components.EntityTypeName](../../models/components/entitytypename.md) | :heavy_minus_sign:                                                      | The type of entity.                                                     | Control Planes                                                          |
| `EntityRegion`                                                          | [*components.EntityRegion](../../models/components/entityregion.md)     | :heavy_minus_sign:                                                      | N/A                                                                     | eu                                                                      |