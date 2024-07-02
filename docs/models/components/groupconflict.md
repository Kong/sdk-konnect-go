# GroupConflict

The Group Conflict object contains information about a conflict in a control plane group.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ClusterID`                                                                           | **string*                                                                             | :heavy_minus_sign:                                                                    | The ID of a control plane member of a control plane group.                            | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                  |
| `Description`                                                                         | **string*                                                                             | :heavy_minus_sign:                                                                    | The description of the conflict.                                                      | conflicting entity found: ID=38d790ad-8b08-4ff5-a074-2e1e9e64d8bd, Name=foo           |
| `Resource`                                                                            | [*components.GroupConflictResource](../../models/components/groupconflictresource.md) | :heavy_minus_sign:                                                                    | A resource causing a conflict in a control plane group.                               |                                                                                       |