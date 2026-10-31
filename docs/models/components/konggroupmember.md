# KongGroupMember

Identifies a single member (principal or group) of a group.



## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ID`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `Type`                                                                                           | [components.KongGroupMembershipType](../../models/components/konggroupmembershiptype.md)         | :heavy_check_mark:                                                                               | Classifies the kind of member referenced by a group membership entry.<br/>                       |
| `DisplayName`                                                                                    | `*string`                                                                                        | :heavy_minus_sign:                                                                               | Human-readable name of the member. For a principal, its `display_name`; for a group, its `name`. |