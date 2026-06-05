# CreateSystemAccountsAssignedRolesInternalRequest


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `AccountID`                                                                               | `string`                                                                                  | :heavy_check_mark:                                                                        | ID of the system account.                                                                 |
| `AssignParameterizedRole`                                                                 | [*components.AssignParameterizedRole](../../models/components/assignparameterizedrole.md) | :heavy_minus_sign:                                                                        | The request schema for assigning a role.                                                  |